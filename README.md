# Red Badger: Robot

## Running

If you have go installed

First install the support applications

```bash
go get github.com/golang/dep/cmd/dep \
           github.com/onsi/ginkgo/ginkgo \
           github.com/onsi/gomega
```

Then download the dependencies

```
dep ensure
```

Then you can build and run the tests

This will run the unit tests
```
go test
```

This will run the end to end tests

```
(cd red-badger-robot/ && go test)
```

You can also build the binary

```
(cd red-badger-robot/ && go install)
``` 

and the binary will be called `red-badger-robot`.

Alternatively you could use the docker container

```
docker build -t billie/red-badger-robot:latest .
docker run --rm -it billie/red-badger-robot:latest
```