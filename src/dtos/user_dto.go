package dtos

import "go-boilerplate/src/models"

type GetUserByIDResponse struct {
	models.User
}

type CreateUserRequest struct {
	Username	string	`json:"username" validate:"required"`
	Password	string	`json:"password" validate:"required"`
	FullName	string	`json:"full_name"`
	PhoneNumber	string	`json:"phone_number"`
}

type UpdateUserParams struct {
	ID			uint	`param:"id" validate:"required"`
	FullName	string	`json:"full_name"`
	PhoneNumber	string	`json:"phone_number"`
}

type UserIDParams struct {
	ID uint `param:"id" validate:"required"`
}