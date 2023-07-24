package responses

import (
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/resources"
)

func NewUserResponse(user data.User) resources.UserResponse {
	return resources.UserResponse{
		Data: newUser(user),
	}
}

func NewUserListResponse(users []data.User) resources.UserListResponse {
	return resources.UserListResponse{
		Data: newUserList(users),
	}
}

func newUserList(users []data.User) []resources.User {
	var usersList = make([]resources.User, 0)
	for _, user := range users {
		usersList = append(usersList, newUser(user))
	}

	return usersList
}

func newUser(user data.User) resources.User {
	return resources.User{
		Key: resources.Key{
			ID:   user.Address,
			Type: resources.USER,
		},
		Attributes: resources.UserAttributes{
			Address:   user.Address,
			PublicKey: user.PublicKey,
		},
	}
}
