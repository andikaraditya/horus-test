package claim

import (
	"github.com/andikaraditya/horus-test/backend/internal/voucher"
	"github.com/jackc/pgx/v5/pgtype"
)

type ClaimVoucher struct {
	ID           string             `json:"id"`
	VoucherId    string             `json:"-"`
	TanggalClaim pgtype.Timestamptz `json:"tanggal_claim"`
	Voucher      voucher.Voucher    `json:"voucher"`
}
