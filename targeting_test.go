package main

import (
	"net"

	. "gopkg.in/check.v1"
)

type TargetingSuite struct {
}

var _ = Suite(&TargetingSuite{})

func (s *TargetingSuite) SetUpSuite(c *C) {
	Config.GeoIP.Directory = "/usr/local/var/GeoIP"
}

func (s *TargetingSuite) TestTargetString(c *C) {
	var tgt TargetOptions
	tgt = TargetGlobal + TargetCountry + TargetContinent

	str := tgt.String()
	c.Check(str, Equals, "@ continent country")
}

func (s *TargetingSuite) TestTargetParse(c *C) {

	tgt, err := parseTargets("@ foo country")
	str := tgt.String()
	c.Check(str, Equals, "@ country")
	c.Check(err.Error(), Equals, "Unknown targeting option 'foo'")

	tgt, err = parseTargets("@ continent country asn")
	c.Assert(err, IsNil)
	str = tgt.String()
	c.Check(str, Equals, "@ continent country asn")
}
func (s *TargetingSuite) TestGetTargets(c *C) {

	ip := net.ParseIP("207.171.7.51")

	geoIP.setupGeoIPCity()
	geoIP.setupGeoIPCountry()
	geoIP.setupGeoIPASN()

	tgt, _ := parseTargets("@ continent country")
	targets, _ := tgt.GetTargets(ip)
	c.Check(targets, DeepEquals, []string{"us", "north-america", "@"})

	if geoIP.city == nil {
		c.Log("City GeoIP database requred for these tests")
		return
	}

	tgt, _ = parseTargets("@ continent country region ")
	targets, _ = tgt.GetTargets(ip)
	c.Check(targets, DeepEquals, []string{"us-ca", "us", "north-america", "@"})

	tgt, _ = parseTargets("@ continent regiongroup country region ")
	targets, _ = tgt.GetTargets(ip)
	c.Check(targets, DeepEquals, []string{"us-ca", "us-west", "us", "north-america", "@"})

	tgt, _ = parseTargets("@ continent regiongroup country region asn ip")
	targets, _ = tgt.GetTargets(ip)
	c.Check(targets, DeepEquals, []string{"207.171.7.51", "as53582", "us-ca", "us-west", "us", "north-america", "@"})
}
