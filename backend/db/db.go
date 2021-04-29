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
	if q.createStmt, err = db.PrepareContext(ctx, create); err != nil {
		return nil, fmt.Errorf("error preparing query Create: %w", err)
	}
	if q.createAdminStmt, err = db.PrepareContext(ctx, createAdmin); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAdmin: %w", err)
	}
	if q.createPostStmt, err = db.PrepareContext(ctx, createPost); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePost: %w", err)
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
	if q.loginStmt, err = db.PrepareContext(ctx, login); err != nil {
		return nil, fmt.Errorf("error preparing query Login: %w", err)
	}
	if q.readAllStmt, err = db.PrepareContext(ctx, readAll); err != nil {
		return nil, fmt.Errorf("error preparing query ReadAll: %w", err)
	}
	if q.readByIDStmt, err = db.PrepareContext(ctx, readByID); err != nil {
		return nil, fmt.Errorf("error preparing query ReadByID: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createStmt != nil {
		if cerr := q.createStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createStmt: %w", cerr)
		}
	}
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
	if q.loginStmt != nil {
		if cerr := q.loginStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing loginStmt: %w", cerr)
		}
	}
	if q.readAllStmt != nil {
		if cerr := q.readAllStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing readAllStmt: %w", cerr)
		}
	}
	if q.readByIDStmt != nil {
		if cerr := q.readByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing readByIDStmt: %w", cerr)
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
	db                DBTX
	tx                *sql.Tx
	createStmt        *sql.Stmt
	createAdminStmt   *sql.Stmt
	createPostStmt    *sql.Stmt
	getAllPostsStmt   *sql.Stmt
	getPostByIDStmt   *sql.Stmt
	getPostsByTagStmt *sql.Stmt
	loginStmt         *sql.Stmt
	readAllStmt       *sql.Stmt
	readByIDStmt      *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                tx,
		tx:                tx,
		createStmt:        q.createStmt,
		createAdminStmt:   q.createAdminStmt,
		createPostStmt:    q.createPostStmt,
		getAllPostsStmt:   q.getAllPostsStmt,
		getPostByIDStmt:   q.getPostByIDStmt,
		getPostsByTagStmt: q.getPostsByTagStmt,
		loginStmt:         q.loginStmt,
		readAllStmt:       q.readAllStmt,
		readByIDStmt:      q.readByIDStmt,
	}
}
