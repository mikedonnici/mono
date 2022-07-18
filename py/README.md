# Python Projects

## Overview

Contains Python projects.

```
py
├─ apps
├─ gen
├─ pkg  
└─ gen-to-pk.sh  <-- copy generated to code to pkg dir
```

## Setup

- install `pyenv` to set python version in each `app/project`
- install poetry

```shell
curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py | python
```

## Generating gRPC Code

- Run `make pb` from repo root
- This will generate code into `py/gen/<proto-name>/<proto-version>/`


## Creating shared packages for generated gRPC code

- To share generated gRPC code across multiple projects need to create packages
- Name packages same as `proto-name`, _without_ the version number
- For example, setting up to share `py/gen/status/v1`:

```shell
cd py/pkg
poetry new status
```

- Poetry creates a package structure like this:

```
py  
└─ pkg
   └─ status             <-- poetry package root
      ├─ status          <-- this is where the generated pb files will go
      ├─ tests           <-- can delete
      ├─ pyproject.toml  <-- poetry pkg config
      └─ README.rst      <-- can delete
```

- Run `py/gen-to-pkg.sh` to copy generated code to a corresponding python package (version) dir:

```
py  
└─ pkg
   └─ status              
      └─ status
         └─ v1   <-- copied from py/gen/status/v1
            ├─ status_pb2.py
            └─ status_pb2_grpc.py
```

- The `pyproject.toml` file in the package root enables poetry to include the package with a relative path (see below) 

## Including local shared packages

- From the poetry project root add the package with a relative file path:

```shell
poetry add ../../pkg/status
```

- Above cmd adds this to `pyproject.toml`, or can add manually

```toml
[tool.poetry.dependencies]
status = {path = "../../pkg/status"}
```

- Package files can be imported thus:

```python
from status.v1 import status_pb2
from status.v1 import status_pb2_grpc
```
