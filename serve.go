package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/miekg/dns"
)

func getQuestionName(z *Zone, req *dns.Msg) string {
	lx := dns.SplitDomainName(req.Question[0].Name)
	ql := lx[0 : len(lx)-z.LabelCount]
	return strings.ToLower(strings.Join(ql, "."))
}

func query(req *dns.Msg, z *Zone, remoteAddr string) *dns.Msg {
	qtype := req.Question[0].Qtype

	logPrintf("[zone %s] incoming %s %s %d from %s\n", z.Origin, req.Question[0].Name,
		dns.TypeToString[qtype], req.MsgHdr.Id, remoteAddr)

	// Global meter
	qCounter.Mark(1)

	// Zone meter
	z.Metrics.Queries.Mark(1)

	logPrintln("Got request", req)

	label := getQuestionName(z, req)

	z.Metrics.LabelStats.Add(label)

	realIP, _, _ := net.SplitHostPort(remoteAddr)

	z.Metrics.ClientStats.Add(realIP)

	var ip net.IP // EDNS or real IP
	var edns *dns.EDNS0_SUBNET
	var optRR *dns.OPT

	for _, extra := range req.Extra {

		switch extra.(type) {
		case *dns.OPT:
			for _, o := range extra.(*dns.OPT).Option {
				optRR = extra.(*dns.OPT)
				switch e := o.(type) {
				case *dns.EDNS0_NSID:
					// do stuff with e.Nsid
				case *dns.EDNS0_SUBNET:
					z.Metrics.EdnsQueries.Mark(1)
					logPrintln("Got edns", e.Address, e.Family, e.SourceNetmask, e.SourceScope)
					if e.Address != nil {
						edns = e
						ip = e.Address
					}
				}
			}
		}
	}

	if len(ip) == 0 { // no edns subnet
		ip = net.ParseIP(realIP)
	}

	targets, netmask := z.Options.Targeting.GetTargets(ip)

	m := new(dns.Msg)
	m.SetReply(req)
	if e := m.IsEdns0(); e != nil {
		m.SetEdns0(4096, e.Do())
	}
	m.Authoritative = true

	// TODO: set scope to 0 if there are no alternate responses
	if edns != nil {
		if edns.Family != 0 {
			if netmask < 16 {
				netmask = 16
			}
			edns.SourceScope = uint8(netmask)
			m.Extra = append(m.Extra, optRR)
		}
	}

	labels, labelQtype := z.findLabels(label, targets, qTypes{dns.TypeMF, dns.TypeCNAME, qtype})
	if labelQtype == 0 {
		labelQtype = qtype
	}

	if labels == nil {

		firstLabel := (strings.Split(label, "."))[0]

		if firstLabel == "_status" {
			if qtype == dns.TypeANY || qtype == dns.TypeTXT {
				m.Answer = statusRR(label + "." + z.Origin + ".")
			} else {
				m.Ns = append(m.Ns, z.SoaRR())
			}
			m.Authoritative = true
			return m
		}

		if firstLabel == "_country" {
			if qtype == dns.TypeANY || qtype == dns.TypeTXT {
				h := dns.RR_Header{Ttl: 1, Class: dns.ClassINET, Rrtype: dns.TypeTXT}
				h.Name = label + "." + z.Origin + "."

				txt := []string{
					remoteAddr,
					ip.String(),
				}

				targets, netmask := z.Options.Targeting.GetTargets(ip)
				txt = append(txt, strings.Join(targets, " "))
				txt = append(txt, fmt.Sprintf("/%d", netmask), serverID, serverIP)

				m.Answer = []dns.RR{&dns.TXT{Hdr: h,
					Txt: txt,
				}}
			} else {
				m.Ns = append(m.Ns, z.SoaRR())
			}

			m.Authoritative = true
			return m
		}

		// return NXDOMAIN
		m.SetRcode(req, dns.RcodeNameError)
		m.Authoritative = true

		m.Ns = []dns.RR{z.SoaRR()}
		return m
	}

	if servers := labels.Picker(labelQtype, labels.MaxHosts); servers != nil {
		var rrs []dns.RR
		for _, record := range servers {
			rr := record.RR
			if rr.Header().Rrtype != qtype && qtype != dns.TypeA && qtype != dns.TypeANY {
				rrs = rrs[:0]
				break
			}
			rr.Header().Name = req.Question[0].Name
			rrs = append(rrs, rr)

			if rr.Header().Rrtype == dns.TypeCNAME {
				if rr.(*dns.CNAME).Target == req.Question[0].Name {
					rrs = rrs[:0]
					break
				}
				for _, r := range additionalAnswer(rr.(*dns.CNAME).Target, remoteAddr).Answer {
					rrs = append(rrs, r)
				}
			} else if rr.Header().Rrtype == dns.TypeNS {
				for _, r := range additionalAnswer(rr.(*dns.NS).Ns, remoteAddr).Answer {
					m.Extra = append(m.Extra, r)
				}
			} else if rr.Header().Rrtype == dns.TypeMX {
				for _, r := range additionalAnswer(rr.(*dns.MX).Mx, remoteAddr).Answer {
					m.Extra = append(m.Extra, r)
				}
			}

		}
		m.Answer = rrs
	}

	if len(m.Answer) == 0 {
		m.Ns = append(m.Ns, z.SoaRR())
	}

	return m
}

func additionalAnswer(target, remoteAddr string) *dns.Msg {
	req := new(dns.Msg)
	req.SetQuestion(target, dns.TypeA)
	tZone, ok := findZone(target)
	if ok {
		return query(req, tZone, remoteAddr)
	}

	return new(dns.Msg)
}

func serve(w dns.ResponseWriter, req *dns.Msg, z *Zone) {
	m := query(req, z, w.RemoteAddr().String())

	err := w.WriteMsg(m)
	if err != nil {
		// if Pack'ing fails the Write fails. Return SERVFAIL.
		log.Println("Error writing packet", m)
		dns.HandleFailed(w, req)
	}
}

func statusRR(label string) []dns.RR {
	h := dns.RR_Header{Ttl: 1, Class: dns.ClassINET, Rrtype: dns.TypeTXT}
	h.Name = label

	status := map[string]string{"v": VERSION, "id": serverID}

	hostname, err := os.Hostname()
	if err == nil {
		status["h"] = hostname
	}
	status["up"] = strconv.Itoa(int(time.Since(timeStarted).Seconds()))
	status["qs"] = strconv.FormatInt(qCounter.Count(), 10)
	status["qps1"] = fmt.Sprintf("%.4f", qCounter.Rate1())

	js, err := json.Marshal(status)

	return []dns.RR{&dns.TXT{Hdr: h, Txt: []string{string(js)}}}
}

func setupServerFunc(z *Zone) func(dns.ResponseWriter, *dns.Msg) {
	return func(w dns.ResponseWriter, r *dns.Msg) {
		serve(w, r, z)
	}
}

func listenAndServe(ip string) {

	prots := []string{"udp", "tcp"}

	for _, prot := range prots {
		go func(p string) {
			server := &dns.Server{Addr: ip, Net: p}

			log.Printf("Opening on %s %s", ip, p)
			if err := server.ListenAndServe(); err != nil {
				log.Fatalf("geodns: failed to setup %s %s: %s", ip, p, err)
			}
			log.Fatalf("geodns: ListenAndServe unexpectedly returned")
		}(prot)
	}

}
