package domain

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type State int32

const (
	AvailableState State = iota
	InUseState
	InactiveState
)

func ParseState(state string) (State, error) {
	switch state {
	case "available":
		return AvailableState, nil
	case "in-use":
		return InUseState, nil
	case "inactive":
		return InactiveState, nil
	default:
		return 0, fmt.Errorf("not a valid state: %s", state)
	}
}

func (s *State) UnmarshalJSON(data []byte) error {
	var stateStr string
	if err := json.Unmarshal(data, &stateStr); err != nil {
		return err
	}

	parsedState, err := ParseState(stateStr)
	if err != nil {
		return err
	}

	*s = parsedState
	return nil
}

type Device struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Brand        string `json:"brand"`
	State        State  `json:"state"`
	CreationTime string `json:"creation_time"`
}

type GetById struct {
	Id int `param:"id" validate:"required"`
}

type GetByBrand struct {
	Brand string `param:"brand" validate:"required"`
}

type GetByState struct {
	State State `param:"state" validate:"required"`
}

type Update struct {
	Id     int `param:"id" validate:"required"`
	Device `validate:"required"`
}

func (u *Update) Validate(fl validator.StructLevel) {
	if req, ok := fl.Current().Interface().(Update); ok {
		if req.CreationTime != "" {
			fl.ReportError(
				req.CreationTime,
				"CreationTime",
				"creation_time",
				"no_update",
				"")
		}
	}
}

type Patch struct {
	Id int `param:"id" validate:"required"`
	Device
}

type Delete struct {
	Id int `param:"id" validate:"required"`
}
