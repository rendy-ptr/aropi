CREATE TABLE products (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(100) NOT NULL,
  price BIGINT NOT NULL,
  stock INT NOT NULL DEFAULT 0,
  category VARCHAR(50),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
