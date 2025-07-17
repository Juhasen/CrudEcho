package user

import (
	"RestCrud/internal/model"
	generated "RestCrud/openapi"
	"github.com/oapi-codegen/runtime/types"
)

func userToDTO(u *model.User) *generated.UserResponse {
	if u == nil {
		return nil
	}
	return &generated.UserResponse{
		Id:    u.ID,
		Name:  u.Name,
		Email: types.Email(u.Email),
	}
}

func dtoToUser(u *generated.UserResponse) *model.User {
	if u == nil {
		return nil
	}
	return &model.User{
		Name:  u.Name,
		Email: string(u.Email),
	}
}
