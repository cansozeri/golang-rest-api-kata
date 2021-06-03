.PHONY: all
all: build
FORCE: ;

# ==============================================================================
# Main

run:
	go run ./cmd/api/main.go

build:
	go build ./cmd/api/main.go

test:
	go test -cover ./...

# ==============================================================================
# Build Mocks

build-mocks:
	@go get github.com/golang/mock/gomock
	@~/go/bin/mockgen -source=internal/records/usecase/interface.go -destination=internal/records/mock/record.go -package=mock -build_flags=-mod=mod
	@~/go/bin/mockgen -source=internal/memory/usecase/interface.go -destination=internal/memory/mock/memory.go -package=mock -build_flags=-mod=mod

# ==============================================================================
# Docker

docker-run:
	docker build . -t clean-rest-api:dev
	docker run -it -p 8080:8080 clean-rest-api:dev
