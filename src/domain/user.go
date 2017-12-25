package domain

import "database/sql/driver"

type Users []*User

type User struct {
	ID        UserID        `json:"id"`
	FirstName UserFirstName `json:"first_name"`
	LastName  UserLastName  `json:"last_name"`
	CreatedAt UserCreatedAt `json:"created_at"`
}

type (
	UserID        int
	UserFirstName string
	UserLastName  string
	UserCreatedAt = NullTime
)

func (user *User) FullName() string {
	return string(user.FirstName) + " " + string(user.LastName)
}

func (s *UserFirstName) Scan(value interface{}) (error) {
	*s = UserFirstName(value.([]byte))
	return nil
}

func (s UserFirstName) Value() (driver.Value, error) {
	return string(s), nil
}

func (s *UserLastName) Scan(value interface{}) (error) {
	*s = UserLastName(value.([]byte))
	return nil
}

func (s UserLastName) Value() (driver.Value, error) {
	return string(s), nil
}
