package StrongHold

import (
	wgapi "github.com/Thenecromance/WarGamingAPI"
)

const (
	path = "wot/stronghold"
)
const (
	claninfo            = "claninfo"
	clanreserves        = "clanreserves"
	activateclanreserve = "activateclanreserve"
)

var methodDesc = map[string]string{
	claninfo:            "Method returns general information and the battle statistics of clans in the Stronghold mode. Please note that information about the number of battles fought as well as the number of defeats and victories is updated once every 24 hours.",
	clanreserves:        "Method returns information about available Reserves and their current status.",
	activateclanreserve: "This method activates an available clan Reserve. A clan Reserve can be activated only by a clan member with the required permission.",
}

func buildPath(method string) string {
	return wgapi.Host + "/" + path + "/" + method
}
