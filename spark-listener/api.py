from typing import Union
from fastapi import FastAPI
from spark_handler import *


app = FastAPI()

@app.get("/")
def read_root():
    return {"Hello": "World"}

@app.get("/pupu")
def pupu_end() -> None:
    return pupu()

@app.get('/count')
def count() -> int:
    return countPeople()

@app.get('/count/{level}')
def countInfoMessage(level)-> int:
    return countLevelMessage(level).count()