-- TRIGGERS
-- Drop triggers first since they depend on the tables and function
DROP TRIGGER IF EXISTS user_updated_at_trigger ON users;
DROP TRIGGER IF EXISTS user_verification_updated_at_trigger ON user_verifications;
-- FUNCTIONS
-- Drop the function used by the triggers
DROP FUNCTION IF EXISTS update_timestamp();
-- TABLES
-- Drop tables in reverse order of creation
DROP TABLE IF EXISTS user_verifications;
DROP TABLE IF EXISTS users;
-- EXTENSIONS
-- (IMPORTANT) Drop extension, but only if you're sure no other parts are using it
DROP EXTENSION IF EXISTS "uuid-ossp";