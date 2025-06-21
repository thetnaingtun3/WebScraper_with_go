go mod tidy
go mod vendor

# RUN Command

<!-- go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest -->
<!-- goose install -->
<!-- go install github.com/pressly/goose/v3/cmd/goose@latest -->

go build -o app && ./app
