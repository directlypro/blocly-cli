package block

import "github.com/txn2/txeh"

func CleanBlock(hosts *txeh.Hosts, blacklistConfiguredSites []string) (cleanedTargets []string) {
	stateSites := blacklistConfiguredSites
	for _, stateSite := range stateSites {
		hosts.RemoveHost(stateSite)
		hosts.Save()

		cleanedTargets = append(cleanedTargets, stateSite)
	}
	return
}
