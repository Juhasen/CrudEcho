package dto

type UserUpdateDTO struct {
	ID    string  `json:"id"`
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}
