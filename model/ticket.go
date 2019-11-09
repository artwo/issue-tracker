package model

import "errors"

type Ticket struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Status      Status `json:"status,omitempty"`
}

func (t *Ticket) Validate() []error {
	var errs []error
	if t.Title == "" {
		errs = append(errs, errors.New("title field of Ticket is empty or undefined"))
	}
	if t.Status == "" {
		t.Status = StatusNew
	} else if !t.Status.IsValid() {
		errs = append(errs, errors.New("status field of Ticket is unknown"))
	}
	return errs
}
