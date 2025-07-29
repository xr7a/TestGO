package models

type SuccessCreateUser struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Surname      string   `json:"surname"`
	Phone        string   `json:"phone"`
	CompanyId    int      `json:"companyId"`
	DepartmentId int      `json:"departmentId"`
	Passport     Passport `json:"passport"`
}

type User struct {
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Phone        string `json:"phone"`
	CompanyId    int    `json:"companyId"`
	DepartmentId int    `json:"departmentId"`
}

type UserPassport struct {
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Phone          string `json:"phone"`
	CompanyId      int    `json:"companyId"`
	DepartmentId   int    `json:"departmentId"`
	PassportType   string `json:"passportType"`
	PassportNumber string `json:"passportNumber"`
}

type UserFull struct {
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Phone           string `json:"phone"`
	CompanyId       int    `json:"company_id"`
	PassportType    string `json:"passport_type"`
	PassportNumber  string `json:"passport_number"`
	DepartmentName  string `json:"department_name"`
	DepartmentPhone string `json:"department_phone"`
}
