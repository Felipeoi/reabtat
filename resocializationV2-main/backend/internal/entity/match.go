package entity

// MatchResult representa um match encontrado entre dois inmates
type MatchResult struct {
	// Dados do inmate do usuário logado
	MyInmate InmateMatchInfo `json:"my_inmate"`

	// Dados do inmate que deu match (de outro usuário)
	MatchedInmate InmateMatchInfo `json:"matched_inmate"`

	// Informações do match
	MatchScore int    `json:"match_score"` // Pode ser usado para ranking no futuro
	Custody    string `json:"custody"`     // Regime em comum
}

// InmateMatchInfo contém informações resumidas de um inmate para exibição no match
type InmateMatchInfo struct {
	ID            int    `json:"id"`
	OriginID      int    `json:"origin_id"`
	Origin        *City  `json:"origin,omitempty"`
	DestinationID int    `json:"destination_id"`
	Destination   *City  `json:"destination,omitempty"`
	Custody       string `json:"custody"`
	UserID        int64  `json:"user_id"`

	// Informações do responsável (útil para contato)
	Responsible InmatesResponsible `json:"responsible"`
}
