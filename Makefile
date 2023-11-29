build:
	go build -o bin/soccerfriend

run: build
	./bin/soccerfriend

test:
	go test -v ./... -count=1
