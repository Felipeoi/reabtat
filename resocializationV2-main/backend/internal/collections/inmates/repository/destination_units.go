package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

func ReplaceDestinationUnits(ctx context.Context, tx pgx.Tx, inmateID int, unitIDs []int) error {
	if _, err := tx.Exec(ctx, `DELETE FROM resocialization.inmate_destination_units WHERE inmate_id = $1`, inmateID); err != nil {
		return err
	}

	for _, unitID := range unitIDs {
		if _, err := tx.Exec(ctx, `
			INSERT INTO resocialization.inmate_destination_units (inmate_id, prison_unit_id)
			VALUES ($1, $2)
		`, inmateID, unitID); err != nil {
			return err
		}
	}

	return nil
}

func LoadDestinationUnitsMap(ctx context.Context, pool *pgxpool.Pool, inmateIDs []int) (map[int][]entity.PrisonUnit, error) {
	out := make(map[int][]entity.PrisonUnit)
	if len(inmateIDs) == 0 {
		return out, nil
	}

	rows, err := pool.Query(ctx, `
		SELECT idu.inmate_id, pu.id, pu.name, pu.uf_code
		FROM resocialization.inmate_destination_units AS idu
		INNER JOIN public.prison_units AS pu ON pu.id = idu.prison_unit_id
		WHERE idu.inmate_id = ANY($1)
		ORDER BY idu.inmate_id, pu.name
	`, inmateIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var inmateID int32
		var unit entity.PrisonUnit
		if err := rows.Scan(&inmateID, &unit.ID, &unit.Name, &unit.UFCode); err != nil {
			return nil, err
		}
		key := int(inmateID)
		out[key] = append(out[key], unit)
	}

	return out, rows.Err()
}

func DestinationUnitIDs(units []entity.PrisonUnit) []int {
	ids := make([]int, len(units))
	for i, unit := range units {
		ids[i] = unit.ID
	}
	return ids
}
