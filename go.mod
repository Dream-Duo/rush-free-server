module rush-free-server

go 1.23.4

require (
	// Database migrations
	github.com/golang-migrate/migrate/v4 v4.17.0
	// Router and Middleware
	github.com/gorilla/mux v1.8.1

	// Logging
	go.uber.org/zap v1.27.0
)

require go.uber.org/multierr v1.11.0 // indirect; Indirect for zap

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect; Indirect
	github.com/lib/pq v1.10.9
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect; Indirect
	github.com/stretchr/testify v1.8.4 // indirect; Indirect
)

require (
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
)
