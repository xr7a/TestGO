package mappers

import (
	"awesomeProject/dal/features/departments"
	"awesomeProject/dal/features/passports"
	"awesomeProject/dal/features/users"
	"awesomeProject/models"
)

func MapToDbUser(source models.UserPassport) users.DbUser {
	return users.DbUser{
		Name:         source.Name,
		Surname:      source.Surname,
		Phone:        source.Phone,
		CompanyId:    source.CompanyId,
		DepartmentId: source.DepartmentId,
	}
}

func MapToDbPassport(source models.UserPassport, userId int) passports.DbPassport {
	return passports.DbPassport{
		UserId: userId,
		Number: source.PassportNumber,
		Type:   source.PassportType,
	}
}

func MapToDbDepartment(source models.Department) departments.DbDepartment {
	return departments.DbDepartment{
		Name:  source.Name,
		Phone: source.Phone,
	}
}
