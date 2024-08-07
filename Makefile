.PHONY:
#.SILENT:

# COLORS.
GREEN:=\033[0;1;32m
NOCOLOR:=\033[0m

# GO ENV.
export GOSUMDB=off
export GO111MODULE=on
export GOPROXY=https://proxy.golang.org,direct
export ENV=local

BIN := $(CURDIR)/bin
MIGRATE_BIN := ./bin/migrate

.build-govulncheck:
	go install golang.org/x/vuln/cmd/govulncheck@latest

lint: .build-golangci-lint
	@echo "${GREEN}# Running configured linters...${NOCOLOR}"
	./bin/golangci-lint run --config=.golangci.yml ./...

.build-golangci-lint:
	@if [ ! -f ./bin/golangci-lint ]; then \
    	echo "${GREEN}# Installing golangci-lint binary...${NOCOLOR}"; \
        curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin; \
    fi

vuln: .build-govulncheck
	govulncheck ./...
