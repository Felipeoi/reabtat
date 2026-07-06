package entity

type MatchResult struct {
	MyInmate      InmateMatchInfo `json:"my_inmate"`
	MatchedInmate InmateMatchInfo `json:"matched_inmate"`
	MatchScore    float64         `json:"match_score"`
	Custody       string          `json:"custody"`
}

type InmateMatchInfo struct {
	ID                   int                `json:"id"`
	OriginUnitID         int                `json:"origin_unit_id"`
	OriginUnit           *PrisonUnit        `json:"origin_unit,omitempty"`
	DestinationUnitIDs   []int              `json:"destination_unit_ids"`
	DestinationUnits     []PrisonUnit       `json:"destination_units"`
	Custody              string             `json:"custody"`
	UserID               int64              `json:"user_id"`
	Responsible          InmatesResponsible `json:"responsible"`
}
