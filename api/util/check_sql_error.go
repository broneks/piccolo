package util

import (
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgconn"
)

func CheckSqlError(err error) string {
	if err == nil {
		return ""
	}

	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return "unique-violation"
		default:
			return pgErr.Message
		}
	} else {
		log.Println("Unexpected non-sql error:", err.Error())
	}

	return ""
}
