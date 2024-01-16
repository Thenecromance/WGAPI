package wargaming

import (
	wgapi "github.com/Thenecromance/WGAPI"
)

// https://api.worldoftanks.asia/wgn/servers/info/?application_id=2d0477de2e50dcf7ec4cb28baf73e639
const (
	path = "wgn/servers"
)
const (
	info = "info"
)

var methodDesc = map[string]string{
	info: "Method returns the number of online players on the servers.",
}

func buildPath(method string) string {
	return wgapi.Host + "/" + path + "/" + method
}
