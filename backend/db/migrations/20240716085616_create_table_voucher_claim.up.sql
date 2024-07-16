CREATE TABLE IF NOT EXISTS voucher_claim (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    voucher_id uuid NOT NULL,
    tanggal_claim TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY(voucher_id) 
      REFERENCES voucher(id) 
      ON DELETE CASCADE
)