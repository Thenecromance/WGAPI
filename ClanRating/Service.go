package ClanRating

import (
	wgapi "github.com/Thenecromance/WarGamingAPI"
)

const (
	path = "clanratings/types"
)
const (
	types     = "types"
	dates     = "dates"
	clans     = "clans"
	neighbors = "neighbors"
	top       = "top"
)

var methodDesc = map[string]string{
	types:     "Method returns details on ratings types and categories.",
	dates:     "Method returns dates with available rating data.",
	clans:     "Method returns clan ratings by specified IDs.",
	neighbors: "Method returns list of adjacent positions in specified clan rating.",
	top:       "Method returns the list of top clans by specified parameters.",
}

func buildPath(method string) string {
	return wgapi.Host + "/" + path + "/" + method
}

/*
type Service struct {
	host        string `default:"https://api.worldoftanks.asia"`
	accountPath string `default:"clanratings/types"`
	RequestProcessor.IProcessor
	methodDesc map[string]string
	methodPath map[string]string
}

func newDescription() map[string]string {
	return map[string]string{
		"types":     "Method returns details on ratings types and categories.",
		"dates":     "Method returns dates with available rating data.",
		"clans":     "Method returns clan ratings by specified IDs.",
		"neighbors": "Method returns list of adjacent positions in specified clan rating.",
		"top":       "Method returns the list of top clans by specified parameters.",
	}
}
func newPath() map[string]string {
	return map[string]string{
		"types":     "types",
		"dates":     "dates",
		"clans":     "clans",
		"neighbors": "neighbors",
		"top":       "top",
	}
}

func NewService(processor RequestProcessor.IProcessor) *Service {
	service := &Service{
		host:        "https://api.worldoftanks.asia",
		accountPath: "wot/clanratings",
		methodDesc:  newDescription(),
		methodPath:  newPath(),
		IProcessor:  processor,
	}
	return service
}
*/
