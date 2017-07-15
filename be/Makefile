build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -o main
	docker build -t wury/say .
	rm -f main
