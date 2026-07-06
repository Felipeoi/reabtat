package entity

type InmatesResponsible struct {
	Attorney string `json:"attorney"`
	Phone    string `json:"phone"`
}

type InmatesFlat struct {
	Id                   *int               `json:"id"`
	OriginUnitID         int                `json:"origin_unit_id"`
	Custody              string             `json:"custody"` // CLOSED, SEMI_OPEN, OPEN
	DestinationUnitIDs   []int              `json:"destination_unit_ids"`
	Responsible          InmatesResponsible `json:"responsible"`
	UserID               int64              `json:"-"` // Não vem do JSON, injetado do JWT
}

type InmatesResp struct {
	Id                   int                `json:"id"`
	OriginUnitID         int                `json:"origin_unit_id"`
	OriginUnit           *PrisonUnit        `json:"origin_unit,omitempty"`
	Custody              string             `json:"custody"`
	DestinationUnitIDs   []int              `json:"destination_unit_ids"`
	DestinationUnits     []PrisonUnit       `json:"destination_units"`
	Responsible          InmatesResponsible `json:"responsible"`
	UserID               int64              `json:"user_id"`
}

type InmatesList struct {
	Id                   int          `json:"id"`
	OriginUnitID         int          `json:"origin_unit_id"`
	OriginUnit           *PrisonUnit  `json:"origin_unit,omitempty"`
	DestinationUnitIDs   []int        `json:"destination_unit_ids"`
	DestinationUnits     []PrisonUnit `json:"destination_units"`
	Custody              string       `json:"custody"`
}
