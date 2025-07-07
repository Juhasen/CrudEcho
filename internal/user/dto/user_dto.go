package dto

type UserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserUpdateDTO struct {
	ID    string  `json:"id"`
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}
