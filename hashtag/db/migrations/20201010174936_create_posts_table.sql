-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS posts (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	index SERIAL,
	post_name TEXT NOT NULL,
  post_data TEXT NOT NULL,
  tags TEXT[] NOT NULL DEFAULT '{}',
	created_at timestamptz DEFAULT current_timestamp,
	updated_at timestamptz DEFAULT current_timestamp
) WITH (
  OIDS=FALSE
);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE posts;