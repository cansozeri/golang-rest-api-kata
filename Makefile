.PHONY: all
all: build
FORCE: ;

build-mocks:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@~/go/bin/mockgen -source=internal/records/usecase/interface.go -destination=internal/records/mock/record.go -package=mock