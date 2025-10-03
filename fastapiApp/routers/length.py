from fastapi import APIRouter, Form, Request
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates

from services.length import convert_length, LENGTH_FACTORS

router = APIRouter()
templates = Jinja2Templates(directory="templates")

@router.get("/", response_class=HTMLResponse)
async def length_form(request: Request):
    return templates.TemplateResponse(
        "index.html",
        {"request": request, "result": None, "units": LENGTH_FACTORS.keys(), "active_tab": "length"}
    )

@router.post("/", response_class=HTMLResponse)
async def length_convert(
    request: Request,
    value: float = Form(...),
    from_unit: str = Form(...),
    to_unit: str = Form(...),
):
    result = f"{value} {from_unit} = {round(convert_length(value, from_unit, to_unit), 4)} {to_unit}"
    return templates.TemplateResponse(
        "index.html",
        {"request": request, "result": result, "units": LENGTH_FACTORS.keys(), "active_tab": "length"}
    )
