WEIGHT_FACTORS = {
    "milligram": 0.001,
    "gram": 1,
    "kilogram": 1000,
    "ounce": 28.3495,
    "pound": 453.592,
}

def convert_weight(value: float, from_unit: str, to_unit: str) -> float:
    """Convert weigth between units"""
    if from_unit not in WEIGHT_FACTORS or to_unit not in WEIGHT_FACTORS:
        raise ValueError("Invalid unit")
    grams = value * weight_factors[from_unit]
    return grams / weight_factors[to_unit]
