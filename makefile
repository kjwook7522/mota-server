build/dev:
	ENV=dev go build -o .out/Mota main.go

build/prod:
	ENV=prod go build -o .out/Mota main.go

run:
	ENV=dev go run main.go