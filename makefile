
GO_MAKE = cd go && $(MAKE)
PY_MAKE = cd py && $(MAKE)

test:
	$(GO_MAKE) test
	$(PY_MAKE) test

build:
	$(GO_MAKE) build

run:
	$(GO_MAKE) run

all:
	$(GO_MAKE)

clean:
	$(GO_MAKE) clean

pb:
	bash ./proto/protoc-go.sh $(CURDIR)
	bash ./proto/protoc-py.sh $(CURDIR)
