# Cargo

Sample project mixing a few concepts.

## Runing the project

The binaries are compiled on Docker image builds, so running the following should be enough:

```
$ docker-compose up -d
```

The Port list query endpoint should be available on `localhost:9001/ports`.
Query parameters `page` and `perPage` are available as pagination has been implemented.

If rebuilds are necessary, just run

```
$ docker-compose build
```

whenever required.

## Testing the project

To run all tests, run the following shell command from the project's root directory:

```
$ go test ./...
```
