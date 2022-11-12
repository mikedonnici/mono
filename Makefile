
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

pb-gen:
	buf lint
	buf generate

pb-go: pb-gen
	cd go && bash gen-to-pkg.sh && cd ..

pb-py: pb-gen
	cd py && bash gen-to-pkg.sh && cd ..

pb-all: pb-gen pb-go pb-py
