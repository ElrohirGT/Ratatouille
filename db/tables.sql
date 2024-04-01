CREATE ROLE postgres WITH LOGIN SUPERUSER;

CREATE DATABASE ratatouille;
\c ratatouille;

CREATE USER root WITH PASSWORD 'root';
GRANT ALL PRIVILEGES ON DATABASE ratatouille TO root;

CREATE USER backend WITH PASSWORD 'backend';
REVOKE ALL ON DATABASE ratatouille FROM backend;
GRANT CONNECT ON DATABASE ratatouille TO backend;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO backend;
GRANT USAGE ON ALL SEQUENCES IN SCHEMA public TO backend;

ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE ON SEQUENCES TO backend;
