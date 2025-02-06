package domain

import (
	"fmt"
)

type Error struct {
	Type   string `json:"type"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s - %v: %s", e.Type, e.Status, e.Detail)
}
