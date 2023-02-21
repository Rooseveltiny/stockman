CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS video_camera (
	link uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	address VARCHAR ( 50 ) NOT NULL,
    port  VARCHAR ( 50 ) NOT NULL,
    login  VARCHAR ( 50 ) NOT NULL,
    password  VARCHAR ( 50 ) NOT NULL
);
