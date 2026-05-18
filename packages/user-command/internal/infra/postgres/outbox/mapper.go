package outbox

import (
	"encoding/json"
	"user-command-module/db/repository"
	"user-command-module/internal/application/port"
	"user-shared-module/common"

	"github.com/jackc/pgx/v5/pgtype"
)

func toInfraOutbox(e *port.OutboxEvent) repository.SaveOutboxParams {
	payload, _ := json.Marshal(e.Payload)
	return repository.SaveOutboxParams{
		ID:             common.ToPgUUID(e.ID),
		AggregateID:    common.ToPgUUID(e.AggregateID),
		AggregateType:  e.AggregateType,
		EventType:      e.EventType,
		Payload:        payload,
		PartitionKey:   e.PartitionKey,
		IdempotencyKey: common.ToPgUUID(e.IdempotencyKey),
		CreatedAt:      common.ToPgTimeStampZ(&e.CreatedAt),
	}
}

func toInfraOutboxBatch(events []port.OutboxEvent) repository.SaveOutboxBatchParams {
	n := len(events)
	params := repository.SaveOutboxBatchParams{
		Ids:             make([]pgtype.UUID, n),
		AggregateIds:    make([]pgtype.UUID, n),
		AggregateTypes:  make([]string, n),
		EventTypes:      make([]string, n),
		Payloads:        make([][]byte, n),
		PartitionKeys:   make([]string, n),
		IdempotencyKeys: make([]pgtype.UUID, n),
		CreatedAts:      make([]pgtype.Timestamptz, n),
	}

	for i, e := range events {
		payload, err := json.Marshal(e.Payload)
		if err != nil {
			// Nếu lỗi, lưu object rỗng để không làm gãy Batch Insert
			payload = []byte("{}")
		}

		params.Ids[i] = common.ToPgUUID(e.ID)
		params.AggregateIds[i] = common.ToPgUUID(e.AggregateID)
		params.AggregateTypes[i] = e.AggregateType
		params.EventTypes[i] = e.EventType
		params.Payloads[i] = payload
		params.PartitionKeys[i] = e.PartitionKey
		params.IdempotencyKeys[i] = common.ToPgUUID(e.IdempotencyKey)
		params.CreatedAts[i] = common.ToPgTimeStampZ(&e.CreatedAt)
	}

	return params
}
