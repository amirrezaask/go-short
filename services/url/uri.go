package url

import (
	"math/rand"
	"time"

	"go-short/database"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}

func randomString(length int) string {
	return stringWithCharset(length, charset)
}

func newUri() string {
	uri := randomString(5)

	for database.ORM().Where("uri = ?", uri).RowsAffected > 0 {
		uri = randomString(5)
	}
	return uri
}
