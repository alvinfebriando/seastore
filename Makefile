BINARY=bin/seastore
test: 
        go test -short -cover -covermode=atomic -coverprofile=cover.html ./...

build:
        go build -o ${BINARY} *.go

clean:
        if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

start:
        docker-compose up --build --detach dev

prod:
        docker-compose up --build prod

stop:
        docker-compose down

lint-prepare:
        @echo "Installing golangci-lint" 
        curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
        ./bin/golangci-lint run ./...

.PHONY: clean install build start prod stop vendor lint-prepare lint