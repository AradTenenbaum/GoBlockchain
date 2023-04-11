run:
	go run .

run-cli:
	go run . $(command)

build:
	set CGO_ENABLED=0&& set GOOS=windows&& go build -o blockchain_go.exe