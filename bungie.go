package cryptarch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BungieService struct {
	apiKey string
	client *http.Client
}

func NewBungieService(apiKey string) *BungieService {
	return &BungieService{apiKey: apiKey, client: &http.Client{}}
}

func (b BungieService) call(endPoint endpoint, result interface{}) error {
	request, err := http.NewRequest(endPoint.method, endPoint.build(), nil)
	if err != nil {
		return fmt.Errorf("Could not build path to endpoint: %v", err.Error())
	}
	request.Header.Add("X-API-Key", b.apiKey)
	response, err := b.client.Do(request)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	type bungieResponse struct {
		Response       json.RawMessage   `json:"Response,omitempty"`
		ErrorCode      int32             `json:"ErrorCode,omitempty"`
		ErrorStatus    string            `json:"ErrorStatus,omitempty"`
		Message        string            `json:"Message,omitempty"`
		MessageData    map[string]string `json:"MessageData,omitempty"`
		ThrottleStatus int32             `json:"ThrottleStatus,omitempty"`
	}
	var bungieResp bungieResponse
	if err = json.Unmarshal(body, &bungieResp); err != nil {
		return fmt.Errorf("Could not unmarshal bungie response: %v", err)
	}
	if bungieResp.ErrorCode != 1 && bungieResp.ErrorStatus != "Ok" {
		return fmt.Errorf("API Error: %d %v", bungieResp.ErrorCode, bungieResp.ErrorStatus)
	}
	fmt.Println(string(bungieResp.Response))
	if err = json.Unmarshal(bungieResp.Response, &result); err != nil {
		return fmt.Errorf("Could not unmarshal bungie response data: %v", err)
	}
	return nil
}
