CREATE TABLE IF NOT EXISTS sessions (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL,
  access_token TEXT NOT NULL,
  refresh_token TEXT,
  created_date TIMESTAMP,
  updated_date TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id)
);