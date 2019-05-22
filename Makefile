.PHONY: example
example:
	go run cmd/jekyll2hugo/main.go --src example/jekyll-posts/ --dst example/hugo-posts/
