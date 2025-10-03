from fastapi import APIRouter, Form, Request
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from services.weight import convert_weight, WEIGHT_FACTORS

router = APIRouter()
templates = Jinja2Templates(directory="templates")

@router.get("/", response_class=HTMLResponse)
async def weigth_form(request: Request):
    return templates.TemplateResponse(
        "index.html",
        {"request": request, "result": None, "units": WEIGHT_FACTORS.keys(), "active_tab": "weight"}
    )

@router.post("/", response_class=HTMLResponse)
async def weight_convert(request: Request, value: float = Form(...), from_unit: str = Form(...), to_unit: str = Form(...)):
    result = convert_weight(value, from_unit, to_unit)
    return templates.TemplateResponse(
        "index.html", 
        {"request": request, "result": result, "units": length_factors.keys(), "active_tab": "length"}
    )
