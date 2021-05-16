// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createAdminStmt, err = db.PrepareContext(ctx, createAdmin); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAdmin: %w", err)
	}
	if q.createPostStmt, err = db.PrepareContext(ctx, createPost); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePost: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.getAllPostsStmt, err = db.PrepareContext(ctx, getAllPosts); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllPosts: %w", err)
	}
	if q.getPostByIDStmt, err = db.PrepareContext(ctx, getPostByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostByID: %w", err)
	}
	if q.getPostsByTagStmt, err = db.PrepareContext(ctx, getPostsByTag); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostsByTag: %w", err)
	}
	if q.listAllAdminStmt, err = db.PrepareContext(ctx, listAllAdmin); err != nil {
		return nil, fmt.Errorf("error preparing query ListAllAdmin: %w", err)
	}
	if q.loginStmt, err = db.PrepareContext(ctx, login); err != nil {
		return nil, fmt.Errorf("error preparing query Login: %w", err)
	}
	if q.readAllUsersStmt, err = db.PrepareContext(ctx, readAllUsers); err != nil {
		return nil, fmt.Errorf("error preparing query ReadAllUsers: %w", err)
	}
	if q.readUserByIDStmt, err = db.PrepareContext(ctx, readUserByID); err != nil {
		return nil, fmt.Errorf("error preparing query ReadUserByID: %w", err)
	}
	if q.readUsersByFaceStmt, err = db.PrepareContext(ctx, readUsersByFace); err != nil {
		return nil, fmt.Errorf("error preparing query ReadUsersByFace: %w", err)
	}
	if q.updateUserFlushStmt, err = db.PrepareContext(ctx, updateUserFlush); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserFlush: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createAdminStmt != nil {
		if cerr := q.createAdminStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createAdminStmt: %w", cerr)
		}
	}
	if q.createPostStmt != nil {
		if cerr := q.createPostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPostStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.getAllPostsStmt != nil {
		if cerr := q.getAllPostsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllPostsStmt: %w", cerr)
		}
	}
	if q.getPostByIDStmt != nil {
		if cerr := q.getPostByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostByIDStmt: %w", cerr)
		}
	}
	if q.getPostsByTagStmt != nil {
		if cerr := q.getPostsByTagStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostsByTagStmt: %w", cerr)
		}
	}
	if q.listAllAdminStmt != nil {
		if cerr := q.listAllAdminStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listAllAdminStmt: %w", cerr)
		}
	}
	if q.loginStmt != nil {
		if cerr := q.loginStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing loginStmt: %w", cerr)
		}
	}
	if q.readAllUsersStmt != nil {
		if cerr := q.readAllUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing readAllUsersStmt: %w", cerr)
		}
	}
	if q.readUserByIDStmt != nil {
		if cerr := q.readUserByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing readUserByIDStmt: %w", cerr)
		}
	}
	if q.readUsersByFaceStmt != nil {
		if cerr := q.readUsersByFaceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing readUsersByFaceStmt: %w", cerr)
		}
	}
	if q.updateUserFlushStmt != nil {
		if cerr := q.updateUserFlushStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserFlushStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                  DBTX
	tx                  *sql.Tx
	createAdminStmt     *sql.Stmt
	createPostStmt      *sql.Stmt
	createUserStmt      *sql.Stmt
	getAllPostsStmt     *sql.Stmt
	getPostByIDStmt     *sql.Stmt
	getPostsByTagStmt   *sql.Stmt
	listAllAdminStmt    *sql.Stmt
	loginStmt           *sql.Stmt
	readAllUsersStmt    *sql.Stmt
	readUserByIDStmt    *sql.Stmt
	readUsersByFaceStmt *sql.Stmt
	updateUserFlushStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                  tx,
		tx:                  tx,
		createAdminStmt:     q.createAdminStmt,
		createPostStmt:      q.createPostStmt,
		createUserStmt:      q.createUserStmt,
		getAllPostsStmt:     q.getAllPostsStmt,
		getPostByIDStmt:     q.getPostByIDStmt,
		getPostsByTagStmt:   q.getPostsByTagStmt,
		listAllAdminStmt:    q.listAllAdminStmt,
		loginStmt:           q.loginStmt,
		readAllUsersStmt:    q.readAllUsersStmt,
		readUserByIDStmt:    q.readUserByIDStmt,
		readUsersByFaceStmt: q.readUsersByFaceStmt,
		updateUserFlushStmt: q.updateUserFlushStmt,
	}
}
