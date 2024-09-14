runa_app:
	go run main.go

build_app:
	go build main.go

test_carga:
	k6 run load_test.js

