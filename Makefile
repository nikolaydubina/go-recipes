.PHONY: all
all:
	go run github.com/nikolaydubina/mdpage -page page.yaml -config render_config.json > README.md