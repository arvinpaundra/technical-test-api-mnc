CREATE TABLE IF NOT EXISTS transaction_histories (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL,
  transaction_type VARCHAR(10) NOT NULL,
  amount FLOAT NOT NULL,
  balance_before FLOAT NOT NULL,
  balance_after FLOAT NOT NULL,
  remarks VARCHAR(255),
  reference_id VARCHAR(255) NOT NULL,
  created_date TIMESTAMP,
  updated_date TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id)
)