package pgrepo

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

const accountTable = "accounts"
const appointmentTable = "appointments"
const roleTable = "roles"

type Postgres struct {
	Pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Postgres {
	return &Postgres{Pool: pool}
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
