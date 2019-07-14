package utils

import (
	"encoding/json"

	"github.com/FlowerWrong/exchange/models"
	"gopkg.in/resty.v1"
)

// NextID return flake id form flake server
func NextID() (uint64, error) {
	resp, err := resty.R().Get("http://127.0.0.1:8090")
	if err != nil {
		return 0, err
	}
	flake := &models.Flake{}
	err = json.Unmarshal(resp.Body(), &flake)
	if err != nil {
		return 0, err
	}
	return flake.ID, nil
}
