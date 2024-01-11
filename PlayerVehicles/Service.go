package PlayerVehicles

import (
	wgapi "github.com/Thenecromance/WarGamingAPI"
)

/* type Service struct {
	host        string `default:"https://api.worldoftanks.asia"`
	accountPath string `default:"wot/tanks"`
	RequestProcessor.IProcessor
	methodDesc map[string]string
	methodPath map[string]string
}

func newDescription() map[string]string {
	return map[string]string{
		"stats":        "Method returns overall statistics, Tank Company statistics, and clan statistics per each vehicle for each user.",
		"achievements": "Method returns list of achievements on all player's vehicles.\n\nAchievement properties define the achievements field values:\n\n1-4 for Mastery Badges and Stage Achievements (type: \"class\");\nmaximum value of Achievement series (type: \"series\");\nnumber of achievements earned from sections: Battle Hero, Epic Achievements, Group Achievements, Special Achievements, etc. (type: \"repeatable, single, custom\").",
		"mastery":      "The method returns percentiles of the distribution of average damage or experience values for each piece of equipment",
	}
}
func newPath() map[string]string {
	return map[string]string{
		"stats":        "stats",
		"achievements": "achievements",
		"mastery":      "mastery",
	}
}

func NewService(processor RequestProcessor.IProcessor) *Service {
	service := &Service{
		host:        "https://api.worldoftanks.asia",
		accountPath: "wot/tanks",
		methodDesc:  newDescription(),
		methodPath:  newPath(),
		IProcessor:  processor,
	}
	return service
}
*/

const (
	path = "wot/tanks"
)
const (
	stats        = "stats"
	achievements = "achievements"
	mastery      = "mastery"
)

var methodDesc = map[string]string{
	stats:        "Method returns overall statistics, Tank Company statistics, and clan statistics per each vehicle for each user.",
	achievements: "Method returns list of achievements on all player's vehicles.\n\nAchievement properties define the achievements field values:\n\n1-4 for Mastery Badges and Stage Achievements (type: \"class\");\nmaximum value of Achievement series (type: \"series\");\nnumber of achievements earned from sections: Battle Hero, Epic Achievements, Group Achievements, Special Achievements, etc. (type: \"repeatable, single, custom\").",
	mastery:      "The method returns percentiles of the distribution of average damage or experience values for each piece of equipment",
}

func buildPath(method string) string {
	return wgapi.Host + "/" + path + "/" + method
}
