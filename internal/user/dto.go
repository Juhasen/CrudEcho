package user

import "RestCrud/internal/model"

type RequestDTO struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type ResponseDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func userToDTO(u *model.User) *ResponseDTO {
	if u == nil {
		return nil
	}
	return &ResponseDTO{
		Name:  u.Name,
		Email: u.Email,
	}
}

func dtoToUser(u *ResponseDTO) *model.User {
	if u == nil {
		return nil
	}
	return &model.User{
		Name:  u.Name,
		Email: u.Email,
	}
}
