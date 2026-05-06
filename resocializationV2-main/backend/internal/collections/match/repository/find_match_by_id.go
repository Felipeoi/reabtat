package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type FindMatchByIDRepo struct {
	pool *pgxpool.Pool
}

func NewFindMatchByIDRepo(pool *pgxpool.Pool) *FindMatchByIDRepo {
	return &FindMatchByIDRepo{
		pool: pool,
	}
}

func (r *FindMatchByIDRepo) Execute(ctx context.Context, userID int64, myInmateID, matchedInmateID int) (*entity.MatchResult, error) {
	var (
		// My Inmate
		myID            int32
		myOriginID      int32
		myDestinationID int32
		myCustody       string
		myAttorney      pgtype.Text
		myPhone         pgtype.Text
		myUserID        int64
		// My Origin City
		myOriginCityID int32
		myOriginIBGE   string
		myOriginName   string
		myOriginUF     string
		// My Destination City
		myDestCityID int32
		myDestIBGE   string
		myDestName   string
		myDestUF     string
		// Matched Inmate
		matchID            int32
		matchOriginID      int32
		matchDestinationID int32
		matchCustody       string
		matchAttorney      pgtype.Text
		matchPhone         pgtype.Text
		matchUserID        int64
		// Matched Origin City
		matchOriginCityID int32
		matchOriginIBGE   string
		matchOriginName   string
		matchOriginUF     string
		// Matched Destination City
		matchDestCityID int32
		matchDestIBGE   string
		matchDestName   string
		matchDestUF     string
	)

	err := r.pool.QueryRow(ctx, findMatchByIDSQL, userID, myInmateID, matchedInmateID).Scan(
		// My Inmate
		&myID, &myOriginID, &myDestinationID, &myCustody, &myAttorney, &myPhone, &myUserID,
		// My Origin City
		&myOriginCityID, &myOriginIBGE, &myOriginName, &myOriginUF,
		// My Destination City
		&myDestCityID, &myDestIBGE, &myDestName, &myDestUF,
		// Matched Inmate
		&matchID, &matchOriginID, &matchDestinationID, &matchCustody, &matchAttorney, &matchPhone, &matchUserID,
		// Matched Origin City
		&matchOriginCityID, &matchOriginIBGE, &matchOriginName, &matchOriginUF,
		// Matched Destination City
		&matchDestCityID, &matchDestIBGE, &matchDestName, &matchDestUF,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("match não encontrado")
		}
		return nil, err
	}

	result := &entity.MatchResult{
		MyInmate: entity.InmateMatchInfo{
			ID:            int(myID),
			OriginID:      int(myOriginID),
			DestinationID: int(myDestinationID),
			Custody:       myCustody,
			UserID:        myUserID,
			Origin: &entity.City{
				ID:       int(myOriginCityID),
				IBGECode: myOriginIBGE,
				Name:     myOriginName,
				UFCode:   myOriginUF,
			},
			Destination: &entity.City{
				ID:       int(myDestCityID),
				IBGECode: myDestIBGE,
				Name:     myDestName,
				UFCode:   myDestUF,
			},
			Responsible: entity.InmatesResponsible{
				Attorney: myAttorney.String,
				Phone:    myPhone.String,
			},
		},
		MatchedInmate: entity.InmateMatchInfo{
			ID:            int(matchID),
			OriginID:      int(matchOriginID),
			DestinationID: int(matchDestinationID),
			Custody:       matchCustody,
			UserID:        matchUserID,
			Origin: &entity.City{
				ID:       int(matchOriginCityID),
				IBGECode: matchOriginIBGE,
				Name:     matchOriginName,
				UFCode:   matchOriginUF,
			},
			Destination: &entity.City{
				ID:       int(matchDestCityID),
				IBGECode: matchDestIBGE,
				Name:     matchDestName,
				UFCode:   matchDestUF,
			},
			Responsible: entity.InmatesResponsible{
				Attorney: matchAttorney.String,
				Phone:    matchPhone.String,
			},
		},
		MatchScore: 100, // Score fixo por enquanto
		Custody:    myCustody,
	}

	return result, nil
}

var findMatchByIDSQL = `
	SELECT
		-- My Inmate
		my.id, my.origin_id, my.destination_id, my.custody,
		my.responsible_attorney, my.responsible_phone, my.user_id,
		-- My Origin City
		my_oc.id, my_oc.ibge_code, my_oc.name, my_oc.uf_code,
		-- My Destination City
		my_dc.id, my_dc.ibge_code, my_dc.name, my_dc.uf_code,
		-- Matched Inmate
		other.id, other.origin_id, other.destination_id, other.custody,
		other.responsible_attorney, other.responsible_phone, other.user_id,
		-- Matched Origin City
		other_oc.id, other_oc.ibge_code, other_oc.name, other_oc.uf_code,
		-- Matched Destination City
		other_dc.id, other_dc.ibge_code, other_dc.name, other_dc.uf_code
	FROM resocialization.inmates AS my
	-- My cities
	LEFT JOIN public.cities AS my_oc ON my.origin_id = my_oc.id
	LEFT JOIN public.cities AS my_dc ON my.destination_id = my_dc.id
	-- Find matching inmates
	INNER JOIN resocialization.inmates AS other
		ON my.origin_id = other.destination_id
		AND my.destination_id = other.origin_id
		AND my.custody = other.custody
		AND my.user_id != other.user_id
	-- Other cities
	LEFT JOIN public.cities AS other_oc ON other.origin_id = other_oc.id
	LEFT JOIN public.cities AS other_dc ON other.destination_id = other_dc.id
	WHERE my.user_id = $1
		AND my.id = $2
		AND other.id = $3
	LIMIT 1;
`
