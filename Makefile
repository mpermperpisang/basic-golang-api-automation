all: build-package run-package

build-package:
	docker build -t basic-golang-api-automation-docker .
	@echo "Build golang-automation-docker successfully"

run-package:
	docker run basic-golang-api-automation-docker
	@echo "Run golang-automation-docker successfully"
