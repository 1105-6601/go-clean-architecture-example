package domain

import (
	"github.com/go-sql-driver/mysql"
	"encoding/json"
	"time"
)

type NullTime struct {
	mysql.NullTime
}

func NewNullTimeNow() NullTime {
	return NullTime{
		NullTime: mysql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
}

func NewNullTime(time time.Time) NullTime {
	return NullTime{
		NullTime: mysql.NullTime{
			Time:  time,
			Valid: true,
		},
	}
}

func (t NullTime) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.NullTime.Time)
	} else {
		return json.Marshal(nil)
	}
}

func (t *NullTime) UnmarshalJSON(data []byte) error {

	_t := time.Time{}

	err := json.Unmarshal(data, &_t)

	t.NullTime.Time  = _t
	t.NullTime.Valid = err == nil

	return err
}
