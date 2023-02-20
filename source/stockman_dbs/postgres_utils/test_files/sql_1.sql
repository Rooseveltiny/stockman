CREATE TABLE test_table (
	link uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
    last_login TIMESTAMP DEFAULT now()
);