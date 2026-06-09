package port

import "context"

type NestedParams struct {
	Index         string
	DocID         string
	NestedField   string
	NestedFieldID string
	Data          any
}

type ESRepository interface {
	SyncData(ctx context.Context, index string, docID string, data any) error
	SyncNestedData(ctx context.Context, param NestedParams) error
}
