package domain

import (
	"encoding/json"
	"fmt"
	"time"
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

func (s *State) UnmarshalParam(param string) error {
	parsedState, err := ParseState(param)
	if err != nil {
		return err
	}
	*s = parsedState
	return nil
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

func (s *State) String() string {
	switch *s {
	case AvailableState:
		return "available"
	case InUseState:
		return "in-use"
	case InactiveState:
		return "inactive"
	default:
		return "unknown"
	}
}

func (s *State) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

type Device struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Brand        string    `json:"brand"`
	State        State     `json:"state"`
	CreationTime time.Time `json:"creation_time"`
}

type GetById struct {
	Id int `param:"id" validate:"required"`
}

type GetByBrand struct {
	Brand string `param:"brand" validate:"required"`
}

type GetByState struct {
	State State `param:"state"`
}

type Update struct {
	Id           int       `param:"id" validate:"required"`
	Name         *string   `json:"name" validate:"required"`
	Brand        *string   `json:"brand" validate:"required"`
	State        *State    `json:"state" validate:"required"`
	CreationTime time.Time `json:"creation_time"`
}

type Patch struct {
	Id           int       `param:"id" validate:"required"`
	Name         *string   `json:"name,omitempty"`
	Brand        *string   `json:"brand,omitempty"`
	State        *State    `json:"state,omitempty"`
	CreationTime time.Time `json:"creation_time"`
}

type Delete struct {
	Id int `param:"id" validate:"required"`
}
