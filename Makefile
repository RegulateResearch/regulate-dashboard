include .env

init:
	cp .env.empty .env
	cp go/.env.empty go/.env
	cp sveltekit/.env.empty sveltekit/.env
	chmod 777 .env

build-run:
	docker compose build
	docker compose up

run:
	docker compose up

# after modifying db schema (for example, adding new table or adding column)
# always execute this
initdb-sync:
	cd sqlmerger && \
	go run main.go

# execute this to apply db changes in your container
# since this execute directly into container
# only execute this when the db container is still run (not stopped)
apply-migrate:
	docker exec -it ${DB_CONTAINER_NAME} \
	psql -U ${BACKEND_DB_USERNAME} \
	-d ${BACKEND_DB_NAME} \
	-f docker-entrypoint-initdb.d/initdb.sql

# execute this if you want to modify/sync initdb.sql 
# and mmediately apply it in one command
# caution: comments on each command is relevant here
# check those comments first before execute this command
db-sync-apply: initdb-sync apply-migrate