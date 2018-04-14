package models

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid"
)

func Now() time.Time {
	return time.Now().UTC().Truncate(time.Second)
}

func NewID() string {
	return ulid.MustNew(ulid.Timestamp(Now()), rand.Reader).String()
}

type Pagination struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}
