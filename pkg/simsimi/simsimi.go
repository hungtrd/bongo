package simsimi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	url   = "https://api.simsimi.net/v2/"
	local = "vn"
)

type SimsimiResp struct {
	Methods      string `json:"methods"`
	Success      string `json:"success"`
	Notification string `json:"noti"`
	Location     string `json:"location"`
}

func SendMessage(msg string) (string, error) {
	reqURL := fmt.Sprintf("%s?text=%s&lc=%s", url, msg, local)
	resp, err := http.Get(reqURL)
	if err != nil {
		fmt.Println("Error: ", err)
		return "", err
	}
	var simResp SimsimiResp
	err = json.NewDecoder(resp.Body).Decode(&simResp)
	if err != nil {
		fmt.Println("Error: ", err)
		return "", err
	}

	return simResp.Success, nil
}
