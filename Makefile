APP_NAME := goa2_sample
REPO := github.com/tonouchi510/goa2-sample
SWAGGER_DIR := ./server/swagger-ui/swagger/

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
	@cp -f ./gen/http/openapi.json ${SWAGGER_DIR}
	@cp -f ./gen/http/openapi.yaml ${SWAGGER_DIR}

regen:
	@rm -rf ./cmd
	@goa example $(REPO)/design
	@mv -n *.go controller/
	@./script/fix-goagen-source.sh controller ${REPO} ${APP_NAME}

run:
	@cd cmd/$(APP_NAME) && go build
	@./cmd/$(APP_NAME)/$(APP_NAME)

clean:
	@rm -rf cmd/
	@rm -rf gen/


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
