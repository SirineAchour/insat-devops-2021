

# Format:
# Target: dependance
# 	Commande1
# 	Commande2
# 	Commande3n


install:
	go get ./...

test: install
	go test ./...

start_dev: install
	echo "Server started"
	go run main.go

