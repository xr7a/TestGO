package users

type DbUser struct {
	Id           int    `db:"id"`
	Name         string `db:"name"`
	Surname      string `db:"surname"`
	Phone        string `db:"phone"`
	CompanyId    int    `db:"company_id"`
	DepartmentId int    `db:"department_id"`
}

type DbUserPassport struct {
	Id             int    `db:"id"`
	Name           string `db:"name"`
	Surname        string `db:"surname"`
	Phone          string `db:"phone"`
	CompanyId      int    `db:"company_id"`
	DepartmentId   int    `db:"department_id"`
	PassportType   string `db:"passport_type"`
	PassportNumber string `db:"passport_number"`
}

type DbUserFull struct {
	Id              int    `db:"id"`
	Name            string `db:"name"`
	Surname         string `db:"surname"`
	Phone           string `db:"phone"`
	CompanyId       int    `db:"company_id"`
	PassportType    string `db:"passport_type"`
	PassportNumber  string `db:"passport_number"`
	DepartmentName  string `db:"department_name"`
	DepartmentPhone string `db:"department_phone"`
}
