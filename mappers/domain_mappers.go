package mappers

import (
	"awesomeProject/dal/features/passports"
	"awesomeProject/dal/features/users"
	"awesomeProject/models"
)

func MapToDomainUser(source users.DbUser) models.User {
	return models.User{
		Name:         source.Name,
		Surname:      source.Surname,
		Phone:        source.Phone,
		CompanyId:    source.CompanyId,
		DepartmentId: source.DepartmentId,
	}
}

func MapToDomainSuccessCreateUser(source users.DbUser) models.SuccessCreateUser {
	return models.SuccessCreateUser{
		Id:           source.Id,
		Name:         source.Name,
		Surname:      source.Surname,
		Phone:        source.Phone,
		CompanyId:    source.CompanyId,
		DepartmentId: source.DepartmentId,
	}
}

func MapToDomainPassport(source passports.DbPassport) models.Passport {
	return models.Passport{
		PassportNumber: source.Number,
		PassportType:   source.Type,
	}
}
