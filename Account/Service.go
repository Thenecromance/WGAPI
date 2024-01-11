package Account

import (
	wgapi "github.com/Thenecromance/WarGamingAPI"
)

const (
	path = "wot/account"
)

const (
	players       = "list"
	personal_data = "info"
	vehicles      = "tanks"
	achievement   = "achievements"
)

var methodDesc = map[string]string{
	players:       "Method returns partial list of players. The list is filtered by initial characters of user name and sorted alphabetically.",
	personal_data: "Method returns details on player's vehicles.",
	vehicles:      "Search account tanks by account id",
	achievement:   "Method returns players' achievement details.",
}

func buildPath(method string) string {
	return wgapi.Host + "/" + path + "/" + method
}

// type Service struct {
// 	host        string `default:"https://api.worldoftanks.asia"`
// 	accountPath string `default:"wot/account"`
// 	RequestProcessor.IProcessor
// 	methodDesc map[string]string
// 	methodPath map[string]string
// }

// func NewService(processor RequestProcessor.IProcessor) *Service {
// 	service := &Service{
// 		host:        "https://api.worldoftanks.asia",
// 		accountPath: "wot/account",
// 		methodDesc: map[string]string{
// 			"players":       "Method returns partial list of players. The list is filtered by initial characters of user name and sorted alphabetically.",
// 			"personal_data": "Method returns details on player's vehicles.",
// 			"vehicles":      "Search account tanks by account id",
// 			"achievement":   "Method returns players' achievement details.",
// 		},
// 		methodPath: map[string]string{
// 			"players":       "list",
// 			"personal_data": "info",
// 			"vehicles":      "tanks",
// 			"achievement":   "achievements",
// 		},
// 		IProcessor: processor,
// 	}
// 	return service
// }
