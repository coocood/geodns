package main

import (
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/stathat/go"
)

func (zs *Zones) statHatPoster() {

	if len(Config.StatHat.APIKey) == 0 {
		return
	}

	stathatGroups := append(serverGroups, "total", serverID)
	suffix := strings.Join(stathatGroups, ",")

	lastCounts := map[string]int64{}
	lastEdnsCounts := map[string]int64{}

	for name, zone := range *zs {
		if zone.Logging.StatHat == true {
			lastCounts[name] = zone.Metrics.Queries.Count()
			lastEdnsCounts[name] = zone.Metrics.EdnsQueries.Count()
		}
	}

	for {
		time.Sleep(60 * time.Second)

		for name, zone := range *zs {

			count := zone.Metrics.Queries.Count()
			newCount := count - lastCounts[name]
			lastCounts[name] = count

			if zone.Logging != nil && zone.Logging.StatHat == true {

				APIKey := zone.Logging.StatHatAPI
				if len(APIKey) == 0 {
					APIKey = Config.StatHat.APIKey
				}
				if len(APIKey) == 0 {
					continue
				}
				stathat.PostEZCount("zone "+name+" queries~"+suffix, Config.StatHat.APIKey, int(newCount))

				ednsCount := zone.Metrics.EdnsQueries.Count()
				newEdnsCount := ednsCount - lastEdnsCounts[name]
				lastEdnsCounts[name] = ednsCount
				stathat.PostEZCount("zone "+name+" edns queries~"+suffix, Config.StatHat.APIKey, int(newEdnsCount))

			}
		}
	}
}

func statHatPoster() {

	lastQueryCount := qCounter.Count()
	stathatGroups := append(serverGroups, "total", serverID)
	suffix := strings.Join(stathatGroups, ",")
	// stathat.Verbose = true

	for {
		time.Sleep(60 * time.Second)

		if !Config.Flags.HasStatHat {
			log.Println("No stathat configuration")
			continue
		}

		log.Println("Posting to stathat")

		current := qCounter.Count()
		newQueries := current - lastQueryCount
		lastQueryCount = current

		stathat.PostEZCount("queries~"+suffix, Config.StatHat.APIKey, int(newQueries))
		stathat.PostEZValue("goroutines "+serverID, Config.StatHat.APIKey, float64(runtime.NumGoroutine()))

	}
}
