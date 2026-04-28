-- lorem ipsum
-- lorem ipsum
-- lorem ipsum

-- for docker
SELECT 'CREATE DATABASE frascati'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'frascati')
\gexec
;