module example.com/user-service

go 1.18

require (
	example.com/models v0.0.0-00010101000000-000000000000
	example.com/router v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)

// uncomment, don't commit
replace example.com/models => ../../models

replace example.com/router => ../../router
