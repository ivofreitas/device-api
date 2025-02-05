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
	return fmt.Sprintf("%s - %s: %s", e.Type, e.Status, e.Detail)
}
