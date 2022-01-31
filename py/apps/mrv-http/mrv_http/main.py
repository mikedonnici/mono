from fastapi import FastAPI

from mono_http.foo import Foo

from hello_two.hello import say

app = FastAPI()


@app.get("/")
async def root():
    f = Foo()

    return {
        "message": f.hey(),
        "lib": say(),
    }
