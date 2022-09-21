.PHONY: all
all:
	go install github.com/nikolaydubina/mdpage@latest
	go run github.com/nikolaydubina/mdpage -page page.yaml > README.md
