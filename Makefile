build:
	@go build ./main.go
install: build
	@mv ./main ./v
	@sudo mv ./v /usr/bin/


time: build
	./main time -h