# Python Projects

- `apps/???` is the root folder of a python application
- `pkg/???` is a python package for internal use across projects / applications

- install `pyenv` to set python version in each `app/project`


- install poetry

```shell
curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py | python
```


## Generating gRPC Code

Ref: https://grpc.io/docs/languages/python/quickstart/

Install gRPC and tools:

```shell
python -m pip install grpcio
python -m pip install grpcio-tools
```

Then, instead of running `protoc` directly, python code can be generated with:

```shell
python -m grpc_tools.protoc ...
```

See `py/makefile` for examples.

## Building and including local packages

Some of these steps can be applied to any package, but is primarily aimed at the generation and subsequent inclusion 
of gRPC client packages by Python apps.

The gRPC packages are located in the [pkg](./pkg) folder and need to be rebuilt when the `.proto` files are updated.

Python packages can be a bit confusing but this seems to work and `poetry` makes it pretty easy to build and include 
local packages that are located outside the application's main project dir.

### Creating a new package

- Using `poetry` makes it easy to create a new package from scratch:

```shell
cd py
poetry new pkg-name
```
- This creates a `pyproject.toml` file in the package root - edit as required, but it should build as is.

- If it's a gRPC package then add the package name to the list of packages in [proto/protoc-py.sh](../proto/protoc-py.sh)

```shell
declare -a PackageNames=("pkg-one" "pkg-two" "pkg-name")
```

From here, the following steps are required each time the package needs to be rebuilt. 

- Update `.proto` file(s)
- Regenerate gRPC code (use `py/makefile` to only generate Python code):
```shell
make pb
```
- Build the package
```shell
cd pkg/pkg-name
poetry build
```

### Including a local package

If the application is using `poetry` can just add the package with a relative file path:

```shell
poetry add ../../pkg/pkg-name
``` 

This adds a line like this to `pyproject.toml`

```toml
[tool.poetry.dependencies]
pkg-name = {path = "../../pkg/pkg-name"}
```

By default, a `poetry` project `pkg-name` would create a package dir `pkg-name/pkg_name`, so the import is thus:

```py
import pkg_name
```

