package port

import "context"

type ESRepository interface {
	SyncData(ctx context.Context, index string, id string, data any) error
}
