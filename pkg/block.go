package block

import (
	"fmt"
	"github.com/txn2/txeh"
	"net/http"
)

func BlockSites(hosts *txeh.Hosts, sites []string) (blacklistConfiguredSites []string) {
	for _, s := range sites {
		blacklistConfiguredSites = append(blacklistConfiguredSites, s)
	}
	//TODO: Leave this for later, might not work
	//sites = state.FilterStateListedSites(sites)

	for _, site := range sites {
		if isTargetAlive(site) {
			blockSite(hosts, site)
		} else {
			//blacklistConfiguredSites, _ = state.RemoveUrlFromList(blacklistConfiguredSites, site)
		}
	}
	return blacklistConfiguredSites
}

func isTargetAlive(url string) bool {
	resp, err := http.Get(fmt.Sprintf("https://%s", url))
	if err != nil {
		return true
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 500 {
		return true
	} else {
		fmt.Printf("Status code for %s is %d, Ignoring...\n", url, resp.StatusCode)
		return false
	}
}

func blockSite(hosts *txeh.Hosts, site string) {
	// TODO integrate Viper
	// target := viper.GetString("app.blockRoute")
	target := "0.0.0.0"
	hosts.AddHost(target, site)
	hosts.Save()
}
