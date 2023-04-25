package db

import (
	"database/sql"
)

// Operation provides all functions to execute db queries and transactions
type Action interface {
	Querier
}

// SQLAction provides all functions to execute db queries and transactions
type SQLAction struct {
	db *sql.DB
	*Queries
}

// NewAction creates a new action
func NewAction(db *sql.DB) Action {
	return &SQLAction{
		db:      db,
		Queries: New(db),
	}
}

// // execTx executes a function within a database transaction
// func (operation *SQLAction) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := operation.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	q := New(tx)
// 	err = fn(q)
// 	if err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
// 		}
// 		return err
// 	}
// 	return tx.Commit()
// }
