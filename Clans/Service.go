package Clans

import wgapi "github.com/Thenecromance/WGAPI"

// import (
// 	"encoding/json"

// 	wgapi "github.com/Thenecromance/WGAPI"
// 	"github.com/Thenecromance/WGAPI/RequestProcessor"
// 	structure "github.com/Thenecromance/WGAPI/Service/Clans/structure/Clans"
// )

// var (
// 	wgapi.Logger wgapi.Iwgapi.Logger
// )

// type timestamp = int64
// type Service struct {
// 	host        string `default:"https://api.worldoftanks.asia"`
// 	accountPath string `default:"wot/account"`
// 	RequestProcessor.IProcessor
// 	methodDesc map[string]string
// 	methodPath map[string]string
// }

// type clanResponse struct {
// 	RequestProcessor.ResponseBase
// 	Data []structure.Clans `json:"data"`
// }

// func (this *Service) Clans(name string) (ret []structure.Clans) {
// 	request, err := this.Request(this.host+"/"+this.accountPath+"/"+this.methodPath["clans"], map[string]string{"search": name, "language": "zh-cn"})
// 	if err != nil {
// 		wgapi.Logger.Error(err)
// 		return
// 	}

// 	//parse the json
// 	resp := &clanResponse{}
// 	jsonErr := json.Unmarshal(request, resp)
// 	if jsonErr != nil {
// 		wgapi.Logger.Error(jsonErr)
// 		return
// 	}
// 	return resp.Data
// }

// func NewService(processor RequestProcessor.IProcessor) *Service {
// 	service := &Service{
// 		host:        "https://api.worldoftanks.asia",
// 		accountPath: "wot/clans",
// 		methodDesc: map[string]string{
// 			"clans":       "Method searches through clans and sorts them in a specified order.",
// 			"clan_detail": "Method returns detailed clan information.",
// 			"glossary":    "Method returns information on clan entities.",
// 		},
// 		methodPath: map[string]string{
// 			"clans":       "list",
// 			"clan_detail": "info",
// 			"glossary":    "glossary",
// 		},
// 		IProcessor: processor,
// 	}
// 	return service
// }

const path = "wot/clans"
const (
	clans       = "list"
	clan_detail = "info"
	glossary    = "glossary"
)

var methodDesc = map[string]string{
	clans:       "Method searches through clans and sorts them in a specified order.",
	clan_detail: "Method returns detailed clan information.",
	glossary:    "Method returns information on clan entities.",
}

func buildPath(method string) string {
	return wgapi.Host + "/" + path + "/" + method
}
