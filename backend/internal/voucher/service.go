package voucher

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/andikaraditya/horus-test/backend/internal/api"
	"github.com/andikaraditya/horus-test/backend/internal/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type VoucherService interface {
	createVoucher(req *Voucher) error
	getVouchers(param string) ([]Voucher, error)
	getVoucher(req *Voucher) error
	updateVoucher(req *Voucher, updatedFields []string) error
	deleteVoucher(req *Voucher) error
}

type srv struct {
	db db.DBService
}

var (
	Service VoucherService
)

func init() {
	Service = New(db.Service)
}

func New(db db.DBService) VoucherService {
	return &srv{db}
}

func (s *srv) createVoucher(req *Voucher) error {
	if err := s.db.Commit(nil, func(tx pgx.Tx) error {
		req.ID = uuid.NewString()
		_, err := tx.Exec(
			context.Background(),
			`INSERT INTO voucher (
				id,
				nama,
				foto,
				kategori
			) VALUES ($1, $2, $3, $4)`,
			req.ID,
			req.Name,
			req.Foto,
			req.Category,
		)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		if strings.Contains(err.Error(), "23505") {
			return api.ErrPayload
		}
		return err
	}
	return s.getVoucher(req)
}

func (s *srv) getVoucher(req *Voucher) error {
	if err := s.db.QueryRow(
		`SELECT 
			id,
			nama,
			foto,
			kategori,
			status,
			created_at,
			updated_at
		FROM voucher 
		WHERE id = $1;`,
		req.ID,
	).Scan(
		&req.ID,
		&req.Name,
		&req.Foto,
		&req.Category,
		&req.Status,
		&req.CreatedAt,
		&req.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return api.ErrNotFound
		}
		return err
	}
	return nil
}

func (s *srv) getVouchers(param string) ([]Voucher, error) {
	v := []Voucher{}
	var sb strings.Builder

	args := []any{"active"}

	if len(param) > 0 {
		sb.WriteString("AND kategori = $2")
		args = append(args, param)
	}
	rows, err := s.db.Query(
		`SELECT 
			id,
			nama,
			foto,
			kategori,
			status,
			created_at,
			updated_at
		FROM voucher 
		WHERE status = $1 `+sb.String(),
		args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var o Voucher
		if err := rows.Scan(
			&o.ID,
			&o.Name,
			&o.Foto,
			&o.Category,
			&o.Status,
			&o.CreatedAt,
			&o.UpdatedAt,
		); err != nil {
			return nil, err
		}
		v = append(v, o)
	}
	return v, nil
}

func (s *srv) updateVoucher(req *Voucher, updatedFields []string) error {
	if err := s.db.Commit(nil, func(tx pgx.Tx) error {
		req.UpdatedAt = pgtype.Timestamptz{Time: time.Now(), Valid: true}
		args := []any{req.ID, req.UpdatedAt}
		var sb strings.Builder

		for _, field := range updatedFields {
			switch field {
			case "nama":
				args = append(args, req.Name)
				sb.WriteString(fmt.Sprintf("nama = $%d,", len(args)))
			case "kategori":
				args = append(args, req.Category)
				sb.WriteString(fmt.Sprintf("kategori = $%d,", len(args)))
			case "status":
				args = append(args, req.Status)
				sb.WriteString(fmt.Sprintf("status = $%d,", len(args)))
			case "foto":
				args = append(args, req.Foto)
				sb.WriteString(fmt.Sprintf("foto = $%d,", len(args)))
			}
		}

		if len(args) > 2 {
			if _, err := tx.Exec(
				context.Background(),
				fmt.Sprintf(
					`UPDATE voucher
					SET %s
						updated_at = $2
					WHERE id = $1;`,
					sb.String(),
				),
				args...,
			); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return s.getVoucher(req)
}

func (s *srv) deleteVoucher(req *Voucher) error {
	if err := s.db.Commit(nil, func(tx pgx.Tx) error {
		tx.Exec(
			context.Background(),
			`DELETE FROM voucher
				WHERE id = $1`,
			req.ID,
		)
		return nil
	}); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return api.ErrNotFound
		}
		return err
	}
	return nil
}
