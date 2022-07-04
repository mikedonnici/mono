# Python Projects

## Overview

Contains Python projects.

```
py
├── apps
├── gen
├── pkg  
└── gen-to-pk.sh  <-- copy generated to code to pkg dir
```

## Setup

- install `pyenv` to set python version in each `app/project`
- install poetry

```shell
curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py | python
```

## Creating shared packages for generated gRPC code

- To share generated gRPC code across multiple projects need to create packages
- Name packages with `proto-name` + `proto-version`
- For example, generated code from `gen/status/v1` requires a package named `pkg/statusv1`:

```shell
cd py/pkg
poetry new statusv1
```
- The `pyproject.toml` file in the package root enables poetry to include the package with a relative path (see below) 


## Generating gRPC Code and updating packages

- Run `make pb` from repo root
- This will generate code into `py/gen/<proto-name>/<proto-version>/`
- Run `gen-to-pkg.sh` to copy generated code into the corresponding package dir


## Including local shared packages

- From the poetry project root add the package with a relative file path:

```shell
poetry add ../../pkg/statusv1
```

- This adds a line like this to `pyproject.toml`

```toml
[tool.poetry.dependencies]
statusv1 = {path = "../../pkg/statusv1"}
```

- Then, for example, the generated files would be included thus:

```python
import statusv1.status_pb2
import statusv1.status_pb2_grpc
```

