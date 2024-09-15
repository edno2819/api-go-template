runa_app:
	go run main.go

build_app:
	go build main.go

test_carga:
	k6 run load_test.js

run_compose:
	docker-compose up -d

refresh_swager:
	swag init

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"