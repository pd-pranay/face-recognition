-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS admin (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	index SERIAL,
	name TEXT NOT NULL,
	email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
	created_at timestamptz DEFAULT current_timestamp,
	updated_at timestamptz DEFAULT current_timestamp
) WITH (
  OIDS=FALSE
);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE admin;