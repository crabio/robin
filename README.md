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

### Auth providers secrets

All auth providers secrets are stored in the `secrets` folder.

#### Google

For creating OAuth2 secrets, follow steps:

1. Go to [Google API console](https://console.cloud.google.com)
2. Create project
3. Add `OAuth 2.0 Client IDs` creadentials
4. Add scopes for getting user profile info:
    * `https://www.googleapis.com/auth/userinfo.email`
	* `https://www.googleapis.com/auth/userinfo.profile`
5. Add redirecting URL, which is link to your server API. (example: `http://localhost:9000`)

Google OAuth2 secrets should be downloaded as JSON file from [Google API console](https://console.cloud.google.com/apis/).

## Lint

For linting we are using [golangci-lint](https://github.com/golangci/golangci-lint)

To run this we recomend use Docker:

```sh
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.40-alpine golangci-lint run
```

## Test

Before pull request, please run `test.sh` script from the `scripts` folder for checking errors.

### Code coverage

After executing `test.sh` file `coverage.out` with code coverage statistic will be generated.
For view this file in HTML format use:
```sh
go tool cover -html=coverage.out
```

## Run

For running use:

```sh
go run .
```

## Deployment

### Docker

For building docker container use:

```sh
sh scripts/build_docker.sh
```

For deploy app with docker run script in `docker` folder:
```sh
docker-compose up -d
```