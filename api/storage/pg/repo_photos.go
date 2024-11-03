package pg

import (
	"context"
	"fmt"
	"piccolo/api/model"

	"github.com/jackc/pgx/v5"
)

func (pg *postgres) GetPhotos() ([]model.Photo, error) {
	query := `select * from photos`

	rows, err := pg.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %v", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Photo])
}
