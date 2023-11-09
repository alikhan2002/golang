CREATE TABLE IF NOT EXISTS users (
     id bigserial PRIMARY KEY,
     created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    email citext UNIQUE NOT NULL,
    password_hash bytea NOT NULL,
    activated bool NOT NULL,
    version integer NOT NULL DEFAULT 1
    );


-- $ curl -d "{\"name\": \"Alice Jones\", \"email\": \"alice@example.com\", \"password\": \"pa55word\"}" localhost:4000/v1/users
-- BODY='{"name": "Alice Jones", "email": "alice@example.com", "password": "pa55word"}'
-- curl -X PUT -d "{\"token\": \"SJ42SE5V7KA6OHTVR5MWOCEBAY\"}" localhost:4000/v1/users/activated
-- curl -H "Authorization: Bearer CIPZ3OTH2G6RN7PEZSHSOCGGE4"" localhost:4000/v1/strollers/1
-- curl -d "{\"email\": \""faith@example.com\", \"password\": \"pa55word\"}" localhost:4000/v1/tokens/authentication