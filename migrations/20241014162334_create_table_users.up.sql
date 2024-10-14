CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY,
  first_name VARCHAR(100) NOT NULL,
  last_name VARCHAR(100),
  phone_number VARCHAR(15) UNIQUE NOT NULL,
  pin VARCHAR(255) NOT NULL,
  balance FLOAT NOT NULL,
  address TEXT,
  created_date TIMESTAMP,
  updated_date TIMESTAMP
);