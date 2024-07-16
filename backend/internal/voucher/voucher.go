package voucher

import "github.com/jackc/pgx/v5/pgtype"

type Voucher struct {
	ID        string             `json:"id"`
	Name      string             `json:"nama" validate:"required"`
	Foto      string             `json:"foto"`
	Category  string             `json:"kategori" validate:"required"`
	Status    string             `json:"status"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
