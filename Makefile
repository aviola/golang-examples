PKGS = $(shell go list ./... | grep -v /test)

lint:
	golint $(PKGS) 
.PHONY: lint

test-unit:
	go test --race --cover -v $(PKGS)
.PHONY: test-unit

test: lint test-unit
.PHONY: test