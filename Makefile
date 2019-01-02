APP_NAME := goa2_sample
GOP := /Users/masato.tonochi/go/my-project
REPO := github.com/tonouchi510/goa2-sample
BACKUP := bak
ARG = db

# 環境構築
install:
	@go get -u goa.design/goa/...
	@go get -u github.com/golang/dep/cmd/dep
	@dep init

# クイックスタート
all: docker-build docker-up run


# goa関連
goagen:
	@goa gen $(REPO)/design
	@rm -rf $(BACKUP)/*
	@mv -f cmd bak
	@mv -f controllers/* ./
	@goa example $(REPO)/design
	@mv -f *.go controllers/
	@find ./controllers/*.go | xargs sed -i '' 's|package goa2sample|package controllers|g'
	@sed -i '' 's|goa2sample "$(REPO)"|goa2sample "$(REPO)/controllers"|g' cmd/$(APP_NAME)/main.go
	@sed -i '' 's|swaggersvr.Mount(mux)|swaggersvr.Mount(mux, swaggerServer)|g' cmd/$(APP_NAME)/main.go

run:
	@cd cmd/$(APP_NAME) && go build
	@cd cmd/$(APP_NAME) && ./$(APP_NAME)

run-cli:
	@cd cmd/$(APP_NAME)-cli && go build

clean:
	@rm -rf cmd/
	@rm -rf gen/
	@rm ./*.go


# docker用 Makefile
docker-all: rm build up ps

docker-build:
	@docker-compose build

docker-up:
	@docker-compose up -d

docker-ps:
	@docker ps -a && echo "\n"
	@docker-compose ps

docker-rm:
	@docker-compose stop
	@docker-compose rm -f

exec:
	docker-compose exec $(ARG) bash