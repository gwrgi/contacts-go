package types

import "time"

type Contact struct {
	Name    string               `json:"name"`
	Address Address              `json:"address"`
	Phone   string               `json:"phone"`
	Email   string               `json:"email"`
	Dates   map[string]time.Time `json:"dates"`
}
