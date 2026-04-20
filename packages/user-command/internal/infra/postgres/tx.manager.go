package postgres

// import (
// 	"context"
// 	"fmt"
// 	"shopify-user-command-module/internal/application/port"

// 	"github.com/jackc/pgx/v5"
// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// type txKey struct{}

// type TxManager struct {
// 	pool *pgxpool.Pool
// }

// func NewTxManager(pool *pgxpool.Pool) port.TxManager {
// 	return &TxManager{
// 		pool: pool,
// 	}
// }

// func (m *TxManager) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
// 	tx, err := m.pool.Begin(ctx)

// 	if err != nil {
// 		return fmt.Errorf("begin transaction: %w", err)
// 	}

// 	txCtx := context.WithValue(ctx, txKey{}, tx)

// 	if err := fn(txCtx); err != nil {
// 		_ = tx.Rollback(txCtx)
// 		return err
// 	}

// 	if err := tx.Commit(txCtx); err != nil {
// 		return fmt.Errorf("Commit transaction: %w", err)
// 	}
// 	return nil

// }

// func ExtractTx(ctx context.Context) pgx.Tx {
// 	tx, _ := ctx.Value(txKey{}).(pgx.Tx)
// 	return tx
// }
