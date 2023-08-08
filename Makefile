clean:
	rm -rf bin/*

dependencies:
	go mod download

build-mocks:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@~/go/bin/mockgen -source=usecase/user/interface.go -destination=usecase/user/mock/user.go -package=mock

test:
	go test -tags testing ./...