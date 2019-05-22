# jekyll-to-hugo

> Migrate [jekyll](https://jekyllrb.com/) posts to [Hugo](https://gohugo.io/)

# Installl

```bash
go get -u github.com/miguelmota/go-jekyll-to-hugo/jekyll2hugo
```

# Getting started

```bash
$ jekyll2hugo --src jekyll-posts/ --dst hugo-posts/

jekyll-posts/2011-03-30-hello-world.md -> hugo-posts/hello-world.md
```

Use the `--dry` flag to only see which files will be copied:

```bash
$ jekyll2hugo --src jekyll-posts/ --dst hugo-posts/ --dry
```

## License

[MIT](LICENSE)
