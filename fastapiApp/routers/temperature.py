from fastapi import APIRouter, Request, Form
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from services.temperature import convert_temperature

templates = Jinja2Templates(directory="templates")

router = APIRouter()
@router.get("/", response_class=HTMLResponse)
async def temp_form(request: Request):
    return templates.TemplateResponse(
        "index.html",
        {"request": request, "result": None, "units": ["Celsius", "Fahrenheit", "Kelvin"], "active_tab": "temperature"}
    )

@router.post("/", response_class=HTMLResponse)
async def temp_convert(
    request: Request,
    value: float = Form(...),
    from_unit: str = Form(...),
    to_unit: str = Form(...),
):
    result = f"{value} {from_unit} = {round(convert_temperature(value, from_unit, to_unit), 2)} {to_unit}"
    return templates.TemplateResponse(
        "index.html",
        {"request": request, "result": result, "units": ["Celsius", "Fahrenheit", "Kelvin"], "active_tab": "temperature"}
    )