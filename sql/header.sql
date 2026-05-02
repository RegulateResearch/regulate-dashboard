-- lorem ipsum
-- lorem ipsum
-- lorem ipsum

-- for docker
SELECT 'CREATE DATABASE regulatedb'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'regulatedb')
\gexec
;