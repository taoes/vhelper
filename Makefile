build:
	@go build ./main.go
	@echo "编译完成......"
install: build
	@mv ./main ./v
	@sudo mv ./v /usr/bin/


push:
	@git push github main
	@git push gitee main
time: build
	./main time -h