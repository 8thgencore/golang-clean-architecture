.PHONY: run

default: run

define set-default-container
	ifndef c
	c = app
	else ifeq (${c},all)
	override c=
	endif
endef

set-container:
	$(eval $(call set-default-container))

build: set-container
	docker compose build ${c}

dev:
	docker compose up -d --force-recreate ${c}

restart: set-container
	docker compose restart ${c}

down:
	docker compose down

exec: set-container
	docker compose exec ${c} /bin/bash

log: set-container
	docker compose logs -f ${c}

local_run:
	go run services/contact/cmd/app/main.go
