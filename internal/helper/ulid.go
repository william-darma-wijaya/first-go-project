package helper

import (
	"crypto/rand"
	"time"

	"go.rtnl.ai/ulid"
)

func GenerateULID() string {
	t := time.Now()
	id := ulid.MustNew(ulid.Timestamp(t), rand.Reader)
	return id.String()
}