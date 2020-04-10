# go-rest-template

This project is a rest api template written in Go.

## Get Started

Development environment needs:

- [Docker](https://www.docker.com/get-started)
- [Golang (v1.14)](https://golang.org/dl/) (used for building exec file, and testing)

Create file `.env` with content copied from `.env.sample`. And then change its values accordingly.

### Build

#### Docker Image

Run:

```
./script.sh build image
```

> Note: this will create docker image with name `go-rest-template`.

#### Executable File

Must install Golang (v1.14) first, and then run:

```
./script.sh build main
```

> Note: the built `main` file must be located next to `.env` file for it to be able to run.

### Development

#### Mongo

Start local mongodb:

```
./script.sh mongo start
```

Stop local mongodb:

```
./script.sh mongo stop
```

Get mongodb connection uri:

```
./script.sh mongo uri
```

#### Run project in dev mode:

```
./script.sh dev
```

#### Clean

Once in a while (weekly), run clean script to clean stuffs and remove redundant docker images:

```
./script.sh clean
```

## Project Structure

This project follow [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

## Others

### Commit Message

Commit message should and must follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/#specification).
