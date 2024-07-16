CREATE TABLE IF NOT EXISTS "voucher" (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    nama VARCHAR(60) NOT NULL,
    foto TEXT,
    kategori VARCHAR(60) NOT NULL,
    status VARCHAR(60) DEFAULT 'active',
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
)