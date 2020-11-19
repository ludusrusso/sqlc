// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package querytest

import (
	"context"
)

const countStar = `-- name: CountStar :one
SELECT count(*) FROM bar
`

func (q *Queries) CountStar(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countStar)
	var count int64
	err := row.Scan(&count)
	return count, err
}