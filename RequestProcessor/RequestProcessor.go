package RequestProcessor

import (
	"errors"
	"io"
	"net/http"

	wgapi "github.com/Thenecromance/WarGamingAPI"
)

var (
	wgapi.Logger wgapi.Iwgapi.Logger
)

type IProcessor interface {
	Request(url string, params map[string]string) ([]byte, error)
}

type RequestProcessor struct {
	Token string `default:""`
	//d     map[string][]map[string]string
}

func (this *RequestProcessor) Request(url string, params map[string]string) ([]byte, error) {

	// first try to get the request from local
	resp, find := this.requestFromLocal(url, params)
	if find {
		return resp, nil
	}

	params["application_id"] = "c02d05ea55d29361b5e7ec9e51d3e60a"

	// if service could not find the info  on local storage,then start to request from remote
	resp, find = this.requestFromRemote(url, params)
	if find {

		return resp, nil
	}

	return nil, errors.New("Request Error!")

}

func (this *RequestProcessor) requestFromLocal(url string, params map[string]string) ([]byte, bool) {
	//request data from the local storage or from our database
	//wgapi.Logger.Warnf("requestFromLocal not implement yet!")
	return nil, false
}

func (this *RequestProcessor) requestFromRemote(url string, params map[string]string) ([]byte, bool) {
	url += "/?"
	for k, v := range params {
		if v != "" {
			url += "&" + k + "=" + v
		}
	}
	if wgapi.Logger != nil {
		wgapi.Logger.Debugf("Request URL: %s", url)
	}
	req, _ := http.NewRequest("GET", url, nil)

	//set up header
	{
		//req.Header.Add("Authorization", "Bearer "+this.Token)
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		wgapi.Logger.Errorf("Request Error: %s", err.Error())
		return nil, false
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			wgapi.Logger.Errorf("Request Error: %s", err.Error())
		}
	}(response.Body)

	//read the response
	body, _ := io.ReadAll(response.Body)
	return body, true
}
