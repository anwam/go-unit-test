# Overview

This is a demonstration of how to extract the 3rd-party libraries into wrapped packages with exported interfaces.

It is useful when you want to control every package at testing by mocking the interfaces via `mockery`.

## Prerequisite

- mockery
```shell
go install github.com/vektra/mockery/v2@latest
```
## Generate mocks

```shell
mockery --all --with-expecter --inpackage
```

## Running

```shell
go test ./...
```