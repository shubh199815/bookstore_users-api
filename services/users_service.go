package services

import (
	"github.com/shubh199815/bookstore_users-api/domain/users"
	"github.com/shubh199815/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr){

	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil{
		return nil, err
	}
	return &user, nil

	// return &user, nil
	// return nil, &errors.RestErr{
	// 	Status: http.StatusInternalServerError,
	// }
}