module github.com/WilliamCWhite/sharescout_backend

go 1.23.0

toolchain go1.23.12

require (
	github.com/golang-jwt/jwt/v5 v5.3.0
	github.com/gorilla/mux v1.8.1
	github.com/joho/godotenv v1.5.1
	github.com/piquette/finance-go v1.1.0
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
)

require golang.org/x/net v0.42.0 // indirect

replace github.com/piquette/finance-go => github.com/psanford/finance-go v0.0.0-20250222221941-906a725c60a0
