package models

import (
	"database/sql"
)

type Milliseconds float64

func (m Milliseconds) Seconds() float64 {
	return float64(m) / 1000
}

type NullMilliseconds sql.NullFloat64

func (m NullMilliseconds) Seconds() float64 {
	if m.Valid {
		return m.Float64 / 1000
	}
	return 0.0
}

