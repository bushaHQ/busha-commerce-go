package busha_commerce_go

import "strings"

type ErrResponse struct {
	Errors Error `json:"error"`
}

type Error struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func (e ErrResponse) Error() string {
	return strings.ToLower(e.Errors.Message)
}
