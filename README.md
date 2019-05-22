# jekyll-to-hugo

> Migrate jekyll posts to Hugo

# Getting started

```bash
$ go run main.go --src jekyll-posts/ --dst hugo-posts/

jekyll-posts/2011-03-30-hello-world.md -> hugo-posts/hello-world.md
```

Use the `--dry` flag to only see which files will be copied:

```bash
$ go run main.go --src jekyll-posts/ --dst hugo-posts/ --dry
```

## License

[MIT](LICENSE)
