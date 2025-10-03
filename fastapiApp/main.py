from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles
from fastapi.templating import Jinja2Templates

from routers import length, weight, temperature

app = FastAPI()

app.mount("/static", StaticFiles(directory="static"), name="static")
templates = Jinja2Templates(directory="templates")

app.include_router(length.router, prefix="/length", tags=["Length"])
app.include_router(weight.router, prefix="/weight", tags=["Weight"])
app.include_router(temperature.router, prefix="/temperature", tags=["Temperature"])

@app.get("/")
def home():
    return {"message": "Welcome to Unit Converter API. Go to /length, /weight, or /temperature."}


@app.get("/healthcheck")
async def health_check():
    """Checks if server is active"""
    return {"status": "active"}
