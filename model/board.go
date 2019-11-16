package model

import "errors"

type Board struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Tickets     []Ticket `json:"description,omitempty"`
}

func (b *Board) Validate() []error {
	var errs []error
	if b.Name == "" {
		errs = append(errs, errors.New("name field of Board is empty or undefined"))
	}
	return errs
}

func (b *Board) IsNull() bool {
	return b.ID == "" &&
		b.Name == "" &&
		len(b.Tickets) == 0
}
