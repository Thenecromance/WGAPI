package Tankopedia

import (
	wgapi "github.com/Thenecromance/WarGamingAPI"
)

// type Service struct {
// 	host        string `default:"https://api.worldoftanks.asia"`
// 	accountPath string `default:"wot/account"`
// 	RequestProcessor.IProcessor
// 	methodDesc map[string]string
// 	methodPath map[string]string
// }

// func newDescription() map[string]string {
// 	return map[string]string{
// 		"vehicles":         "Method returns list of available vehicles.",
// 		"vehicleprofile":   "Method returns vehicle configuration characteristics based on the specified module IDs.",
// 		"achievements":     "Method returns information about achievements.",
// 		"info":             "Method returns information about Tankopedia.",
// 		"arenas":           "Method returns information about maps.",
// 		"provisions":       "Method returns a list of available equipment and consumables.",
// 		"personalmissions": "Method returns details on Personal Missions on the basis of specified campaign IDs, operation IDs, mission branch and tag IDs.",
// 		"boosters":         "Method returns information about Personal Reserves.",
// 		"vehicleprofiles":  "Method returns vehicle configuration characteristics.",
// 		"modules":          "Method returns list of available modules that can be installed on vehicles, such as engines, turrets, etc. At least one input filter parameter (module ID, type) is required to be indicated.",
// 		"badges":           "Method returns list of available badges a player can gain in Ranked Battles.",
// 		"crewroles":        "Method returns full description of all crew qualifications.",
// 		"crewskills":       "Method returns full description of all crew skills.",
// 	}
// }
// func newPath() map[string]string {
// 	return map[string]string{
// 		"vehicles":         "vehicles",
// 		"vehicleprofile":   "vehicleprofile",
// 		"achievements":     "achievements",
// 		"info":             "info",
// 		"arenas":           "arenas",
// 		"provisions":       "provisions",
// 		"personalmissions": "personalmissions",
// 		"boosters":         "boosters",
// 		"vehicleprofiles":  "vehicleprofiles",
// 		"modules":          "modules",
// 		"badges":           "badges",
// 		"crewroles":        "crewroles",
// 		"crewskills":       "crewskills",
// 	}
// }
// func NewService(processor RequestProcessor.IProcessor) *Service {
// 	service := &Service{
// 		host:        "https://api.worldoftanks.asia",
// 		accountPath: "wot/encyclopedia",
// 		methodDesc:  newDescription(),
// 		methodPath:  newPath(),
// 		IProcessor:  processor,
// 	}
// 	return service
// }

/*
"vehicles":         "vehicles",
"vehicleprofile":   "vehicleprofile",
"achievements":     "achievements",
"info":             "info",
"arenas":           "arenas",
"provisions":       "provisions",
"personalmissions": "personalmissions",
"boosters":         "boosters",
"vehicleprofiles":  "vehicleprofiles",
"modules":          "modules",
"badges":           "badges",
"crewroles":        "crewroles",
"crewskills":       "crewskills",
*/

const (
	path = "wot/encyclopedia"
)
const (
	vehicles         = "vehicles"
	vehicleprofile   = "vehicleprofile"
	achievements     = "achievements"
	info             = "info"
	arenas           = "arenas"
	provisions       = "provisions"
	personalmissions = "personalmissions"
	boosters         = "boosters"
	vehicleprofiles  = "vehicleprofiles"
	modules          = "modules"
	badges           = "badges"
	crewroles        = "crewroles"
	crewskills       = "crewskills"
)

var methodDesc = map[string]string{
	vehicles:         "Method returns list of available vehicles.",
	vehicleprofile:   "Method returns vehicle configuration characteristics based on the specified module IDs.",
	achievements:     "Method returns information about achievements.",
	info:             "Method returns information about Tankopedia.",
	arenas:           "Method returns information about maps.",
	provisions:       "Method returns a list of available equipment and consumables.",
	personalmissions: "Method returns details on Personal Missions on the basis of specified campaign IDs, operation IDs, mission branch and tag IDs.",
	boosters:         "Method returns information about Personal Reserves.",
	vehicleprofiles:  "Method returns vehicle configuration characteristics.",
	modules:          "Method returns list of available modules that can be installed on vehicles, such as engines, turrets, etc. At least one input filter parameter (module ID, type) is required to be indicated.",
	badges:           "Method returns list of available badges a player can gain in Ranked Battles.",
	crewroles:        "Method returns full description of all crew qualifications.",
	crewskills:       "Method returns full description of all crew skills.",
}

func buildPath(method string) string {
	return wgapi.Host + "/" + path + "/" + method
}
