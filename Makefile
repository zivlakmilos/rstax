all: run

run: build
	@./bin/rstax

.PHONY: build
build:
	@go build -o bin/rstax ./main.go

clean:
	rm -Rf bin/
