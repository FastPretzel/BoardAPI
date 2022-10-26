package queries

import "github.com/jackc/pgx/v5"

type Queries struct {
	conn *pgx.Conn
}

func NewQuery(conn *pgx.Conn) *Queries {
	return &Queries{
		conn: conn,
	}
}
