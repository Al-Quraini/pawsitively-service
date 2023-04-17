package db

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// )

// // Operation provides all functions to execute db queries and transactions
// type Action struct {
// 	*Queries
// 	db *sql.DB
// }

// // NewStore creates a new store
// func NewStore(db *sql.DB) *Action {
// 	return &Action{
// 		db:      db,
// 		Queries: New(db),
// 	}
// }

// // execTx executes a function within a database transaction
// func (operation *Action) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := operation.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	q := New(tx)
// 	err = fn(q)
// 	if err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return fmt.Errorf("tx err: &v, rb err: %v", err, rbErr)
// 		}
// 		return err
// 	}
// 	return tx.Commit()
// }

// func (action *Action) LikeTx(ctx context.Context, post Post) error {
// 	err := action.execTx(ctx, func(q *Queries) error {

// 		user, err := action.GetUser(ctx, post.UserID)
// 		if err != nil {
// 			return err
// 		}
// 		like, err := action.CreateLike(ctx, CreateLikeParams{
// 			LikedPostID: post.ID,
// 			UserID:      user.ID,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})

// 	return err
// }
