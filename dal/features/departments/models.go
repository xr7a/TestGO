package departments

type DbDepartment struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	Phone string `db:"phone"`
}
