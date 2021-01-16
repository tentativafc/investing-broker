FROM golang:alpine
COPY . /usr/local/go/pkg/github.com/tentativafc/investing-broker/app/backend/sts-service/
WORKDIR /usr/local/go/pkg/github.com/tentativafc/investing-broker/app/backend/sts-service/
RUN go get -d -v
ENV DB_CONNECTION="host=localhost user=postgres password=123456 dbname=postgres port=5433"
CMD ["go", "run", "main.go"]