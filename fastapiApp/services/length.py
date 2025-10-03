LENGTH_FACTORS = {
    "millimeter": 0.001,
    "centimeter": 0.01,
    "meter": 1,
    "kilometer": 1000,
    "inch": 0.0254,
    "foot": 0.3048,
    "yard": 0.9144,
    "mile": 1609.34,
}
def convert_length(value: float, from_unit: str, to_unit: str) -> float:
    """Convert length between units."""
    if from_unit not in LENGTH_FACTORS or to_unit not in LENGTH_FACTORS:
        raise ValueError("Invalid unit")
    meters = value * LENGTH_FACTORS[from_unit]  
    return meters / LENGTH_FACTORS[to_unit] 
