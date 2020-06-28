package signup

import (
	"context"
	"time"
)

type VerifiedCodeRepository interface {
	Save(ctx context.Context, id string, code string, expireAt time.Time) (int64, error)
	Load(ctx context.Context, id string) (string, time.Time, error)
	Delete(ctx context.Context, id string) (int64, error)
}
