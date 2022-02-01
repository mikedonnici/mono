import grpc
from fastapi import FastAPI, HTTPException

import status_grpc.status_pb2
import status_grpc.status_pb2_grpc

app = FastAPI()


@app.get("/")
async def root():
    channel = grpc.insecure_channel('localhost:50051')
    stub = status_grpc.status_pb2_grpc.StatusServiceStub(channel)
    req = status_grpc.status_pb2.StatusRequest()
    try:
        res = stub.FetchStatus(req)
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"{e}")

    return {"msg": res.text}
