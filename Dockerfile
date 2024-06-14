FROM golang:1.21 as Builder

WORKDIR /app
COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
    go build -o apifocusfinance .

FROM golang:1.21-alpine as Runner

COPY --from=Builder /app/apifocusfinance ./apifocusfinance
EXPOSE 5000

CMD ["./apifocusfinance"]
