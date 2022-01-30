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

