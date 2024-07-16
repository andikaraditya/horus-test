package claim

import (
	"context"
	"errors"

	"github.com/andikaraditya/horus-test/backend/internal/api"
	"github.com/andikaraditya/horus-test/backend/internal/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ClaimVoucherService interface {
	createClaimVoucher(req *ClaimVoucher) error
	getClaimVouchers() ([]ClaimVoucher, error)
	getClaimVoucher(req *ClaimVoucher) error
	deleteClaimVoucher(req *ClaimVoucher) error
	getClaimSummary() ([]struct {
		Kategori string
		Total    int
	}, error)
}

type srv struct {
	db db.DBService
}

var (
	Service ClaimVoucherService
)

func init() {
	Service = New(db.Service)
}

func New(db db.DBService) ClaimVoucherService {
	return &srv{db}
}

func (s *srv) createClaimVoucher(req *ClaimVoucher) error {
	if err := s.db.Commit(nil, func(tx pgx.Tx) error {
		req.ID = uuid.NewString()
		_, err := tx.Exec(
			context.Background(),
			`INSERT INTO voucher_claim (
				id,
				voucher_id
			) VALUES ($1, $2)`,
			req.ID,
			req.VoucherId,
		)
		if err != nil {
			return err
		}

		_, err = tx.Exec(
			context.Background(),
			`UPDATE voucher
				SET status = 'claimed',
						updated_at = NOW()
				WHERE id = $1`,
			req.VoucherId,
		)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil
	}
	return s.getClaimVoucher(req)
}

func (s *srv) getClaimVouchers() ([]ClaimVoucher, error) {
	c := []ClaimVoucher{}

	rows, err := s.db.Query(
		`SELECT 
			id,
			tanggal_claim,
			(SELECT json_build_object(
				'id', id,
				'nama', nama,
				'foto', foto,
				'kategori', kategori,
				'status', status,
				'created_at', created_at,
				'updated_at', updated_at
				)
				FROM voucher v 
				WHERE v.id = vc.voucher_id 
			) AS voucher
		FROM voucher_claim vc`,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var o ClaimVoucher
		if err := rows.Scan(
			&o.ID,
			&o.TanggalClaim,
			&o.Voucher,
		); err != nil {
			return nil, err
		}

		c = append(c, o)
	}
	return c, nil
}

func (s *srv) getClaimVoucher(req *ClaimVoucher) error {
	if err := s.db.QueryRow(
		`SELECT 
			id,
			tanggal_claim,
			(SELECT json_build_object(
				'id', id,
				'nama', nama,
				'foto', foto,
				'kategori', kategori,
				'status', status,
				'created_at', created_at,
				'updated_at', updated_at
				)
				FROM voucher v 
				WHERE v.id = vc.voucher_id 
			) AS voucher
		FROM voucher_claim vc 
		WHERE vc.id = $1;`,
		req.ID,
	).Scan(
		&req.ID,
		&req.TanggalClaim,
		&req.Voucher,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return api.ErrNotFound
		}
		return err
	}
	return nil
}

func (s *srv) deleteClaimVoucher(req *ClaimVoucher) error {
	if err := s.db.Commit(nil, func(tx pgx.Tx) error {
		if err := s.db.QueryRow(
			`SELECT 
				voucher_id
			FROM voucher_claim
			WHERE id = $1`,
			req.ID,
		).Scan(&req.VoucherId); err != nil {
			return err
		}
		_, err := tx.Exec(
			context.Background(),
			`DELETE FROM voucher_claim
				WHERE id = $1`,
			req.ID,
		)
		if err != nil {
			return err
		}

		_, err = tx.Exec(
			context.Background(),
			`UPDATE voucher
				SET status = 'active',
						updated_at = NOW()
				WHERE id = $1`,
			req.VoucherId,
		)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (s *srv) getClaimSummary() ([]struct {
	Kategori string
	Total    int
}, error) {
	req := []struct {
		Kategori string
		Total    int
	}{}
	rows, err := s.db.Query(
		`SELECT v.kategori AS kategori , count(v.id) AS total
			FROM voucher v 
			WHERE v.status = 'claimed'
			GROUP BY v.kategori;`,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var o struct {
			Kategori string
			Total    int
		}

		if err := rows.Scan(
			&o.Kategori,
			&o.Total,
		); err != nil {
			return nil, err
		}

		req = append(req, o)
	}
	return req, nil
}
