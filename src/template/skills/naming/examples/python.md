@rule domain-structure
### Good — the concept lives under its domain

```
src/order/logic/calculate_order_total.py
src/order/types/order_line.py
```

Why: domain `order`, class `logic`/`types`, snake_case module files (C04, C05 Python row).

### Bad — a technical bucket

```
src/utils.py
src/helpers/misc.py
```

Why: `utils`, `helpers`, and `misc` hide which domain owns each function, so every concept can land anywhere (C04, C05).

@rule domain-naming
### Good — action and target in the name

```python
def calculate_order_total(lines: list[OrderLine]) -> Money:
    return sum((line.amount for line in lines), ZERO_MONEY)
```

Why: snake_case action + target for a function, PascalCase for classes, typed input and result (C05 Python row).

### Bad — a name that hides the responsibility

```python
def do_it(data):
    return [handle(x) for x in data]
```

Why: `do_it`, `data`, `handle`, and `x` state no action, target, or result (C05).
