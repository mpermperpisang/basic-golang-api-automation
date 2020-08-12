all: install-package

install-package:
	@GO111MODULE=on go get github.com/cucumber/godog/cmd/godog@v0.10.0
	@echo "Install package successfully"
