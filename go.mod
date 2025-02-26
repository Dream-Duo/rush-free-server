module rush-free-server

go 1.24

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
	github.com/go-redis/redis/v8 v8.11.5
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
)
