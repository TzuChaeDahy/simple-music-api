package models

type Error struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	Err         error  `json:"err,omitempty"`
}

func NewError(name string, description string, status int, err error) *Error {
	return &Error{
		Name:        name,
		Description: description,
		Status:      status,
		Err:         err,
	}
}
