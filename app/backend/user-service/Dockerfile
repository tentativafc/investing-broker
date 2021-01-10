FROM golang:alpine
COPY . /usr/local/go/src/github.com/tentativafc/investing-broker/app/backend/user-service/
WORKDIR /usr/local/go/src/github.com/tentativafc/investing-broker/app/backend/user-service/
RUN go get -d -v
CMD ["go", "run", "main.go"]