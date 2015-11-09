.PHONY: minever

minever:
	docker run --rm -v "${PWD}":/usr/src/minectl -w /usr/src/minectl golang:1.5.1 go build -o minever -v ./cmd/minever
