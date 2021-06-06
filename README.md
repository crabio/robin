# Robin
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/mindinventory/Golang-Project-Structure/blob/master/LICENSE)

Authentication service, which provide multiple authentication providers

## Structure

Code structure is based on the [Golang-Project-Structure](https://github.com/Mindinventory/Golang-Project-Structure)

## Presequinces

* Install `protobuf compiler`:

```sh
sudo apt-get -y install protobuf-compiler
```

* Install Google protobuf for Golang:

```sh
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

## Configuration

App should be located near configuration file `config.yml`.

Example file can be found in repo's `config.yml`.