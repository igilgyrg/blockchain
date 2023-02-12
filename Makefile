build:
	go build -o ./dist/blockhain

run: build
	./dist/blockhain

test:
	go test -v ./... -cover