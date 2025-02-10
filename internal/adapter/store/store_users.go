package store

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"MACKOMUsers/internal/core"
)

func (x *Store) AddUser(user core.User) error {
	query, args, err := sq.
		Insert(tableNameUser).
		Columns(
			"first_name",
			"last_name",
			"age",
		).
		Values(
			user.Firstname,
			user.Lastname,
			user.Age,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "ToSql")
	}

	_, err = x.db.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "Exec")
	}

	return nil
}

func (x *Store) GetUserList() ([]core.User, error) {
	var userList []User

	query, arg, err := sq.Select(
		"first_name",
		"last_name",
		"age",
		"recording_date",
	).
		From(tableNameUser).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "ToSql")
	}

	err = x.db.Select(&userList, query, arg...)
	if err != nil {
		return nil, errors.Wrap(err, "Select")
	}

	return convertUserListFromDB(userList), nil
}

func (x *Store) GetUserListByFilter(fromDate, toDate time.Time, minAge, maxAge int) ([]core.User, error) {
	var userList []User

	builder := sq.Select(
		"first_name",
		"last_name",
		"age",
		"recording_date",
	).
		From(tableNameUser).
		PlaceholderFormat(sq.Dollar)
	if !fromDate.IsZero() {
		builder = builder.Where(
			sq.GtOrEq{
				"recording_date": sql.NullTime{
					Time:  fromDate,
					Valid: true,
				},
			},
		)
	}
	if !toDate.IsZero() {
		builder = builder.Where(
			sq.LtOrEq{
				"recording_date": sql.NullTime{
					Time:  toDate,
					Valid: true,
				},
			},
		)
	}
	if minAge > 0 {
		builder = builder.Where(sq.GtOrEq{"age": minAge})
	}
	if maxAge > 0 {
		builder = builder.Where(sq.LtOrEq{"age": maxAge})
	}

	query, arg, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "ToSql")
	}

	err = x.db.Select(&userList, query, arg...)
	if err != nil {
		return nil, errors.Wrap(err, "Select")
	}

	return convertUserListFromDB(userList), nil
}

func convertUserListFromDB(userList []User) []core.User {
	result := make([]core.User, 0, len(userList))

	for _, user := range userList {
		item := core.User{
			Firstname: user.FirstName,
			Lastname:  user.LastName,
			Age:       user.Age,
		}
		if user.RecordingDate.Valid {
			item.RecordingDate = user.RecordingDate.Time
		}

		result = append(result, item)
	}

	return result
}
