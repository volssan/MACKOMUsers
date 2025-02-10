package core

import "time"

type UserStore interface {
	AddUser(user User) error
	GetUserList() ([]User, error)
	GetUserListByFilter(fromDate, toDate time.Time, minAge, maxAge int) ([]User, error)
}
