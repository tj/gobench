
# gobench

Tiny utility around [benchcmp](https://godoc.org/golang.org/x/tools/cmd/benchcmp).

## Usage

Simple old/new scenario:

```
# benchmark
$ gobench old
$ gobench new

# compare
$ gobench old new
```

Many branches:

```
# benchmark
$ gobench foo
$ gobench bar
$ gobench baz

# compare
$ gobench foo bar
$ gobench foo baz
$ gobench baz bar
```


## Badges

[![GoDoc](https://godoc.org/github.com/tj/gobench?status.svg)](https://godoc.org/github.com/tj/gobench)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-stable-green.svg)
[![](http://apex.sh/images/badge.svg)](https://apex.sh/)

---

> [tjholowaychuk.com](http://tjholowaychuk.com) &nbsp;&middot;&nbsp;
> GitHub [@tj](https://github.com/tj) &nbsp;&middot;&nbsp;
> Twitter [@tjholowaychuk](https://twitter.com/tjholowaychuk)
