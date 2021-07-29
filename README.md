# Bibtex online reference generator

This repo contains a small script that given an url will generate a Bibtex or Biblatex reference with the most informations possible included to help you gain time.

So you will be able to generate something like this:

```
@misc{2021-07-python-vulnerabilities-code-execution-in-jinja-templates-podalirius,
  author = "Podalirius",
  title = "Python vulnerabilities : Code execution in jinja templates · Podalirius",
  year = "2021",
  month = "Juillet",
  howpublished = "\url{https://podalirius.net/en/articles/python-vulnerabilities-code-execution-in-jinja-templates/}",
  note = "[En ligne, accédé le 27 Juillet 2021]"
}
```

With only the link: https://podalirius.net/en/articles/python-vulnerabilities-code-execution-in-jinja-templates/ (poke [@Podalirius](https://twitter.com/podalirius_/) by the way, thanks for your article 😇).

# How do I do that?

This is quite simple (provided that you have Go installed on your machine)! Simply follow the next steps:

1. Get the code:
  ```bash
  git clone https://github.com/bastantoine/bibtex-reference-generator
  ```
2. Get the dependencies:
  ```bash
  go get .
  ```
3. Run the script. Two possibilities here:
   1. You compile it to a standalone executable and then use it to generate the references:
        ```bash
        go build -o bibtex-reference-generator .
        ./bibtex-reference-generator -bibtex -url "https://podalirius.net/en/articles/python-vulnerabilities-code-execution-in-jinja-templates/"
        ```
   2. You use the `go run` method to run it:
        ```bash
        go run main.go -bibtex -url "https://podalirius.net/en/articles/python-vulnerabilities-code-execution-in-jinja-templates/"
        ```
        Note that if you are using this method, you must not move the `main.go` file away from the `go.mod` and `go.sum` files.

*Note :* switch `-bibtex` with `-biblatex` to use a Biblatex format.

# Why some informations might be missing sometimes?

This script uses the meta informations of the page you provide to generate the reference. These meta informations are provided using the HTML `<meta>` tags and the [OpenGraph protocol](https://ogp.me). The *problem* here is that these meta tags are not mandatory, so sometimes there are some informations that can't be retreived, especially the author(s) and year and month of publication date.

For those specially cases, you'll have to fill them by hand, sorry.