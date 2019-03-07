package commands_bob

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type BobResponse struct {
	Message string `json:"message"`
}

type API struct {
	Client  *http.Client
	BaseURL string
}

const statusEndpoint string = "/can-we-build-it"

func (api *API) RunningStatus()(BobResponse, error) {
	var bobResponse BobResponse
	uri := api.BaseURL + statusEndpoint

	res, e := api.Client.Get(uri)
	if e != nil {
		return BobResponse{}, e
	}
	defer res.Body.Close()

	body, e := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(body, &bobResponse)

	return bobResponse, err
}
