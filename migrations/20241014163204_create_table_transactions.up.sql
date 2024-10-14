CREATE TABLE IF NOT EXISTS transactions (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL,
  target_user UUID,
  amount FLOAT NOT NULL,
  remarks VARCHAR(255),
  category VARCHAR(10) NOT NULL,
  status VARCHAR(10) NOT NULL,
  created_date TIMESTAMP,
  updated_date TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id)
)