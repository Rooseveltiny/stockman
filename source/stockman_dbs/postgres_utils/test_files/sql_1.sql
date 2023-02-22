CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS test_table (
	link uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
    last_login TIMESTAMP DEFAULT now()
);

ALTER TABLE test_table DROP CONSTRAINT IF EXISTS "test_table_username_key";