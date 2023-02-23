
def safe_get_or_default(iterable, key, default=None) -> str:
    return iterable[key] if key in iterable else default


def get_first_not_none(first_value, second_value) -> str:
    return first_value if first_value is not None else second_value