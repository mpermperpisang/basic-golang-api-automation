all: build-package run-package

build-package:
	docker build -t basic-golang-api-automation-docker .
	@echo "Build basic-golang-api-automation-docker successfully"

run-package:
	docker run basic-golang-api-automation-docker
	@echo "Run basic-golang-api-automation-docker successfully"
