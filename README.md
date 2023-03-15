# fmultierr ![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square) ![Travis](https://img.shields.io/travis/gostaticanalysis/dupimport.svg?style=flat-square) [![Go Report Card](https://goreportcard.com/badge/github.com/gostaticanalysis/dupimport)](https://goreportcard.com/report/github.com/gostaticanalysis/dupimport) [![codecov](https://codecov.io/gh/gostaticanalysis/dupimport/branch/master/graph/badge.svg)](https://codecov.io/gh/gostaticanalysis/dupimport)


multierr.Errors関数が呼び出されている部分を特定する


## install

```sh
go install github.com/mahiro72/multierr/cmd/multierr@latest
```

## Useage

```sh
go vet -vettool=`which multierr` pkgname
```
