package store

import "database/sql"

const tableNameUser = "users"

type User struct {
	ID            int          `db:"id"`
	FirstName     string       `db:"first_name"`
	LastName      string       `db:"last_name"`
	Age           int          `db:"age"`
	RecordingDate sql.NullTime `db:"recording_date"`
}
