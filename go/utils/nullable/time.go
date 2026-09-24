package nullable

import (
	"database/sql"
	"time"
)

func ToSqlNullTime(time time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  time,
		Valid: !time.IsZero(),
	}
}
