-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	index SERIAL,
	name TEXT,
	college_name TEXT,
  address TEXT,
  mobile_no int,
  image_path TEXT NOT NULL,
	image_uid TEXT UNIQUE NOT NULL,
	is_deleted BOOLEAN DEFAULT false,
	created_at timestamptz DEFAULT current_timestamp,
	updated_at timestamptz DEFAULT current_timestamp
) WITH (
  OIDS=FALSE
);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;