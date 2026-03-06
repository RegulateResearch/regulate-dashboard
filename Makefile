init:
	touch .pvkey
	cp .env.empty .env
	cp go/.env.empty go/.env
	cp sveltekit/.env.empty sveltekit/.env

build-run:
	docker compose build
	docker compose up
