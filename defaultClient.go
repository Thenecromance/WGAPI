package wgapi

import (
	"errors"
	"io"
	"net/http"
)

type client struct {
}

func (this *client) Request(method string, url string, params map[string]string) ([]byte, error) {

	//// first try to get the request from local
	//resp, find := this.requestFromLocal(url, params)
	//if find {
	//	return resp, nil
	//}
	if params == nil {
		params = make(map[string]string)
	}

	params["application_id"] = token

	// if service could not find the info  on local storage,then start to request from remote
	resp, find := this.requestFromRemote(method, url, params)
	if find {
		//save the info to local machine
		return resp, nil
	}

	return nil, errors.New("Request Error!")

}

func (this *client) requestFromRemote(method string, url string, params map[string]string) ([]byte, bool) {
	url += "/?"
	for k, v := range params {
		if v != "" {
			url += "&" + k + "=" + v
		}
	}
	Logger.Debugf("Request URL: %s", url)
	req, _ := http.NewRequest(method, url, nil)

	//set up header
	{
		//req.Header.Add("Authorization", "Bearer "+this.Token)
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		Logger.Errorf("Request failed with error: %s", err.Error())
		return nil, false
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			Logger.Errorf("failed to close response stream: %s", err.Error())
		}
	}(response.Body)

	//read the response
	body, _ := io.ReadAll(response.Body)
	return body, true
}
