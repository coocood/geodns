package main

import (
	"log"
	"net"
	"strings"

	"github.com/semihalev/dns"
)

func findZone(name string) (*Zone, bool) {
	name = strings.ToLower(name)

	zonelist.RLock()
	defer zonelist.RUnlock()

	for off, end := 0, false; !end; off, end = dns.NextLabel(name, off) {
		if z, ok := zonelist.List[name[off:]]; ok {
			return z, ok
		}
	}

	return nil, false
}

func getInterfaces() []string {

	var inter []string
	uniq := make(map[string]bool)

	for _, host := range strings.Split(*flaginter, ",") {
		ip, port, err := net.SplitHostPort(host)
		if err != nil {
			switch {
			case strings.Contains(err.Error(), "missing port in address"):
				// 127.0.0.1
				ip = host
			case strings.Contains(err.Error(), "too many colons in address") &&
				// [a:b::c]
				strings.LastIndex(host, "]") == len(host)-1:
				ip = host[1 : len(host)-1]
				port = ""
			case strings.Contains(err.Error(), "too many colons in address"):
				// a:b::c
				ip = host
				port = ""
			default:
				log.Fatalf("Could not parse %s: %s\n", host, err)
			}
		}
		if len(port) == 0 {
			port = *flagport
		}
		host = net.JoinHostPort(ip, port)
		if uniq[host] {
			continue
		}
		uniq[host] = true

		if len(serverID) == 0 {
			serverID = ip
		}
		if len(serverIP) == 0 {
			serverIP = ip
		}
		inter = append(inter, host)

	}

	return inter
}
