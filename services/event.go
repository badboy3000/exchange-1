package services

import "encoding/json"

// Event ...
type Event struct {
	Name string          `json:"name"`
	Data json.RawMessage `json:"data"`
}
