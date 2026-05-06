package entity

type InmatesResponsible struct {
	Attorney string `json:"attorney"`
	Phone    string `json:"phone"`
}

type InmatesFlat struct {
	Id            *int               `json:"id"`
	OriginID      int                `json:"origin_id"`
	Custody       string             `json:"custody"` // CLOSED, SEMI_OPEN, OPEN
	DestinationID int                `json:"destination_id"`
	Responsible   InmatesResponsible `json:"responsible"`
	UserID        int64              `json:"-"` // Não vem do JSON, injetado do JWT
}

type InmatesResp struct {
	Id            int                `json:"id"`
	OriginID      int                `json:"origin_id"`
	Origin        *City              `json:"origin,omitempty"`
	Custody       string             `json:"custody"` // CLOSED, SEMI_OPEN, OPEN
	DestinationID int                `json:"destination_id"`
	Destination   *City              `json:"destination,omitempty"`
	Responsible   InmatesResponsible `json:"responsible"`
	UserID        int64              `json:"user_id"`
}

type InmatesList struct {
	Id            int    `json:"id"`
	OriginID      int    `json:"origin_id"`
	Origin        *City  `json:"origin,omitempty"`
	DestinationID int    `json:"destination_id"`
	Destination   *City  `json:"destination,omitempty"`
	Custody       string `json:"custody"`
}
