-- EXTENSIONS:
-- "uuid-ossp" to set UUID data values.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- TABLES:
-- users
-- user_verifications
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100),
    email_addr VARCHAR(255) NOT NULL,
    email_addr_verified_at TIMESTAMPTZ,
    phone_num VARCHAR(15),
    phone_num_verified_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);
CREATE TABLE IF NOT EXISTS user_verifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    user_id UUID NOT NULL REFERENCES users (id),
    verification_entity VARCHAR(20) CHECK (
        verification_entity IN ('email_addr', 'phone_num')
    ),
    token_hash VARCHAR(255) NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    revoked_at TIMESTAMPTZ
);
-- FUNCTIONS
-- `updated_at` field with current timestamp
-- used with triggers
CREATE OR REPLACE FUNCTION update_timestamp() RETURNS TRIGGER AS $$ BEGIN NEW.updated_at = CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- TRIGGERS
-- update `updated_at` datetime whenever there is an update to
-- any rows of users and user_verifications tables
CREATE TRIGGER user_updated_at_trigger BEFORE
UPDATE ON users FOR EACH ROW EXECUTE FUNCTION update_timestamp();
CREATE TRIGGER user_verification_updated_at_trigger BEFORE
UPDATE ON user_verifications FOR EACH ROW EXECUTE FUNCTION update_timestamp();