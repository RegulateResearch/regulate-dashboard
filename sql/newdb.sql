-- create new database if not exists
SELECT 'CREATE DATABASE frascati'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'frascati')
\gexec
;

-- move to new database
\c frascati;