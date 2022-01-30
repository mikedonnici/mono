from fastapi import FastAPI

from mrv_http.foo import Foo

from status import status

app = FastAPI()


@app.get("/")
async def root():
    f = Foo()

    return {
        "message": f.hey(),
        "lib": status.hello(),
    }