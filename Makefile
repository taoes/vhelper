build:
	@go build ./main.go
	@echo "编译完成......"
install: build
	@mv ./main ./v
	@sudo mv ./v /usr/bin/


time: build
	./main time -h