module rush-free-server

go 1.23.4

require (
	// Router and Middleware
	github.com/gorilla/mux v1.8.1

	// Logging
	go.uber.org/zap v1.27.0
)

require go.uber.org/multierr v1.11.0 // indirect; Indirect for zap

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/lib/pq v1.10.9
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
)
