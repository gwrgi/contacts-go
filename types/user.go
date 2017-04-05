package types

import (
	"container/list"
	"github.com/satori/go.uuid"
)

type User struct {
	Name     string    `json:"name"`
	Id       uuid.UUID `json:"uuid"`
	Contact  Contact   `json:"contact"`
	Contacts list.List `json:"contacts"`
}
