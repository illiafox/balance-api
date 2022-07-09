include .env
export

BUILD=./app/cmd/api

.PHONY: all
all: clean build run


.PHONY: run
run:
	(cd $(BUILD) && ./app)


.PHONY: build
build: clean
	go build -o $(BUILD)/app $(BUILD)


.PHONY: clean
clean:
	if [ -f $(BUILD)/app ]; then rm $(BUILD)/app; fi


.PHONY: compose
compose: compose-down
	docker-compose up -d --wait

.PHONY: compose-debug
compose-debug: compose-down swagger
	docker-compose up -d --build


.PHONY: compose-down
compose-down:
	docker-compose down

.PHONY: test
test:
	go test -v ./...

# swagger

.PHONY: swagger
swagger:
	 swag init --parseDependency --parseInternal \
	 --parseDepth 1 --generatedTime=true -o=docs \
     -g=./app/cmd/api/main.go --outputTypes=yaml,go


# gci
.PHONY: gci
gci:
	find . -name "*.go" -exec gci write {} \;