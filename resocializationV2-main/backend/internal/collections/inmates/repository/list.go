package repository

import (
	"context"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ListRepo struct{ pool *pgxpool.Pool }

func NewListRepo(pool *pgxpool.Pool) *ListRepo { return &ListRepo{pool: pool} }

// Execute lista presos com paginação filtrando por user_id.
// Se limit <= 0, usa 100. Se offset < 0, usa 0.
func (r *ListRepo) Execute(ctx context.Context, limit, offset int, userID int64) (error, context.Context, []entity.InmatesList) {
	if limit <= 0 || limit > 1000 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	rows, err := r.pool.Query(ctx, listInmatesSQL, userID, limit, offset)
	if err != nil {
		return err, ctx, nil
	}
	defer rows.Close()

	out := make([]entity.InmatesList, 0, limit)
	for rows.Next() {
		var (
			id            int32
			custody       string
			originID      int32
			destinationID int32
			// Origin city
			originCityID   int32
			originIBGECode string
			originName     string
			originUFCode   string
			// Destination city
			destCityID   int32
			destIBGECode string
			destName     string
			destUFCode   string
		)

		if err := rows.Scan(
			&id,
			&custody,
			&originID,
			&destinationID,
			&originCityID,
			&originIBGECode,
			&originName,
			&originUFCode,
			&destCityID,
			&destIBGECode,
			&destName,
			&destUFCode,
		); err != nil {
			return err, ctx, nil
		}

		out = append(out, entity.InmatesList{
			Id:            int(id),
			Custody:       custody,
			OriginID:      int(originID),
			DestinationID: int(destinationID),
			Origin: &entity.City{
				ID:       int(originCityID),
				IBGECode: originIBGECode,
				Name:     originName,
				UFCode:   originUFCode,
			},
			Destination: &entity.City{
				ID:       int(destCityID),
				IBGECode: destIBGECode,
				Name:     destName,
				UFCode:   destUFCode,
			},
		})
	}

	if err := rows.Err(); err != nil {
		return err, ctx, nil
	}

	return nil, ctx, out
}

var listInmatesSQL = `
	SELECT
	  i.id,
	  i.custody,
	  i.origin_id,
	  i.destination_id,
	  oc.id, oc.ibge_code, oc.name, oc.uf_code,
	  dc.id, dc.ibge_code, dc.name, dc.uf_code
	FROM resocialization.inmates AS i
	LEFT JOIN public.cities AS oc ON i.origin_id = oc.id
	LEFT JOIN public.cities AS dc ON i.destination_id = dc.id
	WHERE i.user_id = $1
	ORDER BY i.id DESC
	LIMIT $2 OFFSET $3;
`
