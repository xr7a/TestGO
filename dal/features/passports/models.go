package passports

type DbPassport struct {
	Id     int    `db:"id"`
	Type   string `db:"type"`
	Number string `db:"number"`
	UserId int    `db:"user_id"`
}
