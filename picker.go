package main

import (
	"math/rand"

	"github.com/miekg/dns"
)

// Picker func
func (label *Label) Picker(qtype uint16, max int) Records {

	if qtype == dns.TypeANY {
		result := make([]Record, 0)
		for rtype := range label.Records {

			rtypeRecords := label.Picker(rtype, max)

			tmpResult := make(Records, len(result)+len(rtypeRecords))

			copy(tmpResult, result)
			copy(tmpResult[len(result):], rtypeRecords)
			result = tmpResult
		}

		return result
	}

	if labelRR := label.Records[qtype]; labelRR != nil {

		var servers []Record
		tmpservers := make([]Record, len(labelRR))
		copy(tmpservers, labelRR)
		sum := label.Weight[qtype]
		backupcount := 0

		for i := range tmpservers {
			if !tmpservers[i].Active || tmpservers[i].Backup {
				sum -= tmpservers[i].Weight
			} else {
				servers = append(servers, tmpservers[i])
			}

			if tmpservers[i].Backup {
				backupcount++
			}
		}

		// not "balanced", just return all
		if sum == 0 {
			if backupcount > 0 {
				for i := range tmpservers {
					if tmpservers[i].Backup {
						sum += tmpservers[i].Weight
						servers = append(servers, tmpservers[i])
					}
				}
			} else {
				return labelRR
			}
		}

		rrCount := len(servers)
		if max > rrCount {
			max = rrCount
		}

		result := make([]Record, max)

		for si := 0; si < max; si++ {
			n := rand.Intn(sum + 1)
			s := 0

			for i := range servers {
				s += int(servers[i].Weight)
				if s >= n {
					sum -= servers[i].Weight
					result[si] = servers[i]

					// remove the server from the list
					servers = append(servers[:i], servers[i+1:]...)
					break
				}
			}
		}

		return result
	}
	return nil
}
