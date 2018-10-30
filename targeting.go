package main

import (
	"fmt"
	"net"
	"strings"
)

// TargetOptions type
type TargetOptions int

const (
	// TargetGlobal const
	TargetGlobal = 1 << iota
	// TargetContinent const
	TargetContinent
	// TargetCountry const
	TargetCountry
	// TargetRegionGroup const
	TargetRegionGroup
	// TargetRegion const
	TargetRegion
	// TargetASN const
	TargetASN
	// TargetIP const
	TargetIP
)

// GetTargets func
func (t TargetOptions) GetTargets(ip net.IP) ([]string, int) {

	targets := make([]string, 0)

	var country, continent, region, regionGroup, asn string
	var netmask int

	if t&TargetASN > 0 {
		asn, netmask = geoIP.GetASN(ip)
	}

	if t&TargetRegion > 0 || t&TargetRegionGroup > 0 {
		country, continent, regionGroup, region, netmask = geoIP.GetCountryRegion(ip)

	} else if t&TargetCountry > 0 || t&TargetContinent > 0 {
		country, continent, netmask = geoIP.GetCountry(ip)
	}

	if t&TargetIP > 0 && len(ip) > 0 && ip.To4() != nil {
		targets = append(targets, ip.String())
	}

	if t&TargetASN > 0 && len(asn) > 0 {
		targets = append(targets, asn)
	}

	if t&TargetRegion > 0 && len(region) > 0 {
		targets = append(targets, region)
	}
	if t&TargetRegionGroup > 0 && len(regionGroup) > 0 {
		targets = append(targets, regionGroup)
	}

	if t&TargetCountry > 0 && len(country) > 0 {
		targets = append(targets, country)
	}

	if t&TargetContinent > 0 && len(continent) > 0 {
		targets = append(targets, continent)
	}

	if t&TargetGlobal > 0 {
		targets = append(targets, "@")
	}
	return targets, netmask
}

func (t TargetOptions) String() string {
	targets := make([]string, 0)
	if t&TargetGlobal > 0 {
		targets = append(targets, "@")
	}
	if t&TargetContinent > 0 {
		targets = append(targets, "continent")
	}
	if t&TargetCountry > 0 {
		targets = append(targets, "country")
	}
	if t&TargetRegionGroup > 0 {
		targets = append(targets, "regiongroup")
	}
	if t&TargetRegion > 0 {
		targets = append(targets, "region")
	}
	if t&TargetASN > 0 {
		targets = append(targets, "asn")
	}
	if t&TargetIP > 0 {
		targets = append(targets, "ip")
	}
	return strings.Join(targets, " ")
}

func parseTargets(v string) (tgt TargetOptions, err error) {
	targets := strings.Split(v, " ")
	for _, t := range targets {
		var x TargetOptions
		switch t {
		case "@":
			x = TargetGlobal
		case "country":
			x = TargetCountry
		case "continent":
			x = TargetContinent
		case "regiongroup":
			x = TargetRegionGroup
		case "region":
			x = TargetRegion
		case "asn":
			x = TargetASN
		case "ip":
			x = TargetIP
		default:
			err = fmt.Errorf("Unknown targeting option '%s'", t)
		}
		tgt = tgt | x
	}
	return
}
