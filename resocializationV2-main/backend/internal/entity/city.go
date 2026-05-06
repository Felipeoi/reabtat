package entity

type City struct {
	ID       int    `json:"id"`
	IBGECode string `json:"ibge_code"`
	Name     string `json:"name"`
	UFCode   string `json:"uf_code"`
}
