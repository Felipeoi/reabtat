package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	inmatesrepo "github.com/jeje-gab/resocializationV2/backend/internal/collections/inmates/repository"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type FindMatchesRepo struct {
	pool *pgxpool.Pool
}

func NewFindMatchesRepo(pool *pgxpool.Pool) *FindMatchesRepo {
	return &FindMatchesRepo{pool: pool}
}

func (r *FindMatchesRepo) Execute(ctx context.Context, userID int64) ([]entity.MatchResult, error) {
	rows, err := r.pool.Query(ctx, findMatchesSQL, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]entity.MatchResult, 0)
	inmateIDs := make([]int, 0)

	for rows.Next() {
		var (
			myID, myOriginUnitID int32
			myCustody            string
			myAttorney, myPhone  pgtype.Text
			myUserID             int64
			myOriginUnitIDJoin   int32
			myOriginUnitName     string
			myOriginUnitUF       string

			matchID, matchOriginUnitID int32
			matchCustody               string
			matchAttorney, matchPhone  pgtype.Text
			matchUserID                int64
			matchOriginUnitIDJoin      int32
			matchOriginUnitName        string
			matchOriginUnitUF          string
		)

		if err := rows.Scan(
			&myID, &myOriginUnitID, &myCustody, &myAttorney, &myPhone, &myUserID,
			&myOriginUnitIDJoin, &myOriginUnitName, &myOriginUnitUF,
			&matchID, &matchOriginUnitID, &matchCustody, &matchAttorney, &matchPhone, &matchUserID,
			&matchOriginUnitIDJoin, &matchOriginUnitName, &matchOriginUnitUF,
		); err != nil {
			return nil, err
		}

		inmateIDs = append(inmateIDs, int(myID), int(matchID))

		results = append(results, entity.MatchResult{
			MyInmate: entity.InmateMatchInfo{
				ID:           int(myID),
				OriginUnitID: int(myOriginUnitID),
				Custody:      myCustody,
				UserID:       myUserID,
				OriginUnit: &entity.PrisonUnit{
					ID: int(myOriginUnitIDJoin), Name: myOriginUnitName, UFCode: myOriginUnitUF,
				},
				Responsible: entity.InmatesResponsible{
					Attorney: myAttorney.String,
					Phone:    myPhone.String,
				},
			},
			MatchedInmate: entity.InmateMatchInfo{
				ID:           int(matchID),
				OriginUnitID: int(matchOriginUnitID),
				Custody:      matchCustody,
				UserID:       matchUserID,
				OriginUnit: &entity.PrisonUnit{
					ID: int(matchOriginUnitIDJoin), Name: matchOriginUnitName, UFCode: matchOriginUnitUF,
				},
				Responsible: entity.InmatesResponsible{
					Attorney: matchAttorney.String,
					Phone:    matchPhone.String,
				},
			},
			MatchScore: 100,
			Custody:    myCustody,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	destMap, err := inmatesrepo.LoadDestinationUnitsMap(ctx, r.pool, inmateIDs)
	if err != nil {
		return nil, err
	}

	for i := range results {
		attachMatchDestinations(&results[i].MyInmate, destMap)
		attachMatchDestinations(&results[i].MatchedInmate, destMap)
	}

	return results, nil
}

func attachMatchDestinations(info *entity.InmateMatchInfo, destMap map[int][]entity.PrisonUnit) {
	units := destMap[info.ID]
	if units == nil {
		units = []entity.PrisonUnit{}
	}
	info.DestinationUnits = units
	info.DestinationUnitIDs = inmatesrepo.DestinationUnitIDs(units)
}

var findMatchesSQL = `
	SELECT
		my.id, my.origin_unit_id, my.custody,
		my.responsible_attorney, my.responsible_phone, my.user_id,
		my_ou.id, my_ou.name, my_ou.uf_code,
		other.id, other.origin_unit_id, other.custody,
		other.responsible_attorney, other.responsible_phone, other.user_id,
		other_ou.id, other_ou.name, other_ou.uf_code
	FROM resocialization.inmates AS my
	LEFT JOIN public.prison_units AS my_ou ON my.origin_unit_id = my_ou.id
	INNER JOIN resocialization.inmates AS other
		ON my.custody = other.custody
		AND my.user_id != other.user_id
		AND EXISTS (
			SELECT 1
			FROM resocialization.inmate_destination_units AS my_dest
			WHERE my_dest.inmate_id = my.id
			  AND my_dest.prison_unit_id = other.origin_unit_id
		)
		AND EXISTS (
			SELECT 1
			FROM resocialization.inmate_destination_units AS other_dest
			WHERE other_dest.inmate_id = other.id
			  AND other_dest.prison_unit_id = my.origin_unit_id
		)
	LEFT JOIN public.prison_units AS other_ou ON other.origin_unit_id = other_ou.id
	WHERE my.user_id = $1
	ORDER BY my.id, other.id;
`
