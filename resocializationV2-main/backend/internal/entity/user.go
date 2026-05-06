package entity

import (
	"strings"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	Role      string    `json:"role"`
	Telefone  string    `json:"telefone,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserAuth struct {
	ID           int64
	Name         string
	Email        string
	Status       string
	Role         string
	PasswordHash string
}

// GetWhatsAppLink gera um link de contato do WhatsApp a partir do número de telefone
// Exemplo: +5511999999999 -> https://wa.me/5511999999999
func (u *User) GetWhatsAppLink() string {
	if u.Telefone == "" {
		return ""
	}
	// Remove caracteres não numéricos exceto o +
	phone := strings.ReplaceAll(u.Telefone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")

	// Remove o + inicial se existir
	phone = strings.TrimPrefix(phone, "+")

	return "https://wa.me/" + phone
}
