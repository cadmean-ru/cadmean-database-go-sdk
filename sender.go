package database

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type RequestSender struct {
	req    *ApiRequest
	client *Client
}

func (s *RequestSender) Send() (ApiResponse, error)   {
	var res ApiResponse

	js, err := json.Marshal(s.req)
	if err != nil {
		return res, err
	}

	resp, err := http.Post(s.client.ServerURL + "/api/v1/", "application/json", bytes.NewBuffer(js))
	if err != nil {
		return res, err
	}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return res, err
	}

	_ = resp.Body.Close()

	return res, nil
}


