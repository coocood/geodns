package main

import (
	"strings"
	"time"

	"github.com/abh/go-metrics"
	"github.com/miekg/dns"
)

// ZoneOptions type
type ZoneOptions struct {
	Serial    int
	TTL       int
	MaxHosts  int
	Contact   string
	Targeting TargetOptions
}

// ZoneLogging type
type ZoneLogging struct {
	StatHat    bool
	StatHatAPI string
}

// Record type
type Record struct {
	RR     dns.RR
	Weight int
	Active bool
	Backup bool
}

// Records type
type Records []Record

// Len func
func (s Records) Len() int { return len(s) }

// Swap func
func (s Records) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// RecordsByWeight type
type RecordsByWeight struct{ Records }

// Less func
func (s RecordsByWeight) Less(i, j int) bool { return s.Records[i].Weight > s.Records[j].Weight }

// Label type
type Label struct {
	Label    string
	MaxHosts int
	TTL      int
	Records  map[uint16]Records
	Weight   map[uint16]int
}

type labels map[string]*Label

// ZoneMetrics type
type ZoneMetrics struct {
	Queries     metrics.Meter
	EdnsQueries metrics.Meter
	LabelStats  *ZoneLabelStats
	ClientStats *ZoneLabelStats
}

// Zone type
type Zone struct {
	Origin     string
	Labels     labels
	LabelCount int
	Options    ZoneOptions
	Logging    *ZoneLogging
	LastRead   time.Time
	Metrics    ZoneMetrics
}

type qTypes []uint16

// NewZone func
func NewZone(name string) *Zone {
	zone := new(Zone)
	zone.Labels = make(labels)
	zone.Origin = name
	zone.LabelCount = dns.CountLabel(zone.Origin)

	// defaults
	zone.Options.TTL = 120
	zone.Options.MaxHosts = 2
	zone.Options.Contact = "support.bitnames.com"
	zone.Options.Targeting = TargetGlobal + TargetCountry + TargetContinent

	return zone
}

// SetupMetrics func
func (z *Zone) SetupMetrics(old *Zone) {
	if old != nil &&
		old.Metrics.Queries != nil &&
		old.Metrics.EdnsQueries != nil &&
		old.Metrics.LabelStats != nil &&
		old.Metrics.ClientStats != nil {
		z.Metrics = old.Metrics
	} else {
		z.Metrics.Queries = metrics.NewMeter()
		z.Metrics.EdnsQueries = metrics.NewMeter()
		metrics.Register(z.Origin+" queries", z.Metrics.Queries)
		metrics.Register(z.Origin+" EDNS queries", z.Metrics.EdnsQueries)
		z.Metrics.LabelStats = NewZoneLabelStats(10000)
		z.Metrics.ClientStats = NewZoneLabelStats(10000)
	}
}

// Close func
func (z *Zone) Close() {
	metrics.Unregister(z.Origin + " queries")
	metrics.Unregister(z.Origin + " EDNS queries")
	z.Metrics.LabelStats.Close()
	z.Metrics.ClientStats.Close()
}

func (l *Label) firstRR(dnsType uint16) dns.RR {
	return l.Records[dnsType][0].RR
}

func (l *Label) pickerRR(dnsType uint16) dns.RR {
	return l.Picker(dnsType, 1)[0].RR
}

// AddLabel func
func (z *Zone) AddLabel(k string) *Label {
	k = strings.ToLower(k)
	z.Labels[k] = new(Label)
	label := z.Labels[k]
	label.Label = k
	label.TTL = z.Options.TTL
	label.MaxHosts = z.Options.MaxHosts

	label.Records = make(map[uint16]Records)
	label.Weight = make(map[uint16]int)

	return label
}

// SoaRR func
func (z *Zone) SoaRR() dns.RR {
	return z.Labels[""].firstRR(dns.TypeSOA)
}

// Find label "s" in country "cc" falling back to the appropriate
// continent and the global label name as needed. Looks for the
// first available qType at each targeting level. Return a Label
// and the qtype that was "found"
func (z *Zone) findLabels(s string, targets []string, qts qTypes) (*Label, uint16) {

	for _, target := range targets {

		var name string

		switch target {
		case "@":
			name = s
		default:
			if len(s) > 0 {
				name = s + "." + target
			} else {
				name = target
			}
		}

		if label, ok := z.Labels[name]; ok {

			for _, qtype := range qts {

				switch qtype {
				case dns.TypeANY:
					// short-circuit mostly to avoid subtle bugs later
					// to be correct we should run through all the selectors and
					// pick types not already picked
					return z.Labels[s], qtype
				case dns.TypeMF:
					if label.Records[dns.TypeMF] != nil {
						name = label.pickerRR(dns.TypeMF).(*dns.MF).Mf
						// TODO: need to avoid loops here somehow
						return z.findLabels(name, targets, qts)
					}
				default:
					// return the label if it has the right record
					if label.Records[qtype] != nil && len(label.Records[qtype]) > 0 {
						return label, qtype
					}
				}
			}
		}
	}

	return z.Labels[s], 0
}
