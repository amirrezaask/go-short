package random

import (
	"go-short/services/database"
)

func Uri() string {
	uri := String(5)

	for database.ORM().Where("uri = ?", uri).RowsAffected > 0 {
		uri = String(5)
	}

	return uri
}
