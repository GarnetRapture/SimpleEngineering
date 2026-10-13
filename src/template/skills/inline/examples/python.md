@rule no-inline-comment
### Good — the name carries the intent

```python
unpaid_orders = [order for order in orders if is_awaiting_payment(order)]
```

Why: the names state what a comment would have said (C04 `no-inline-comment`).

### Bad — a comment explains a vague line

```python
lst = [o for o in orders if o.s == 1]  # unpaid orders
```

Why: the hook denies the added `#` comment; the vague names stay (C04, C05).

@rule no-magic-literal
### Good — a named constant in the domain constant owner

```python
from order.constant.paging import ORDER_API_URL, ORDER_PAGE_SIZE, ORDER_REQUEST_TIMEOUT_SECONDS

response = session.get(ORDER_API_URL, params={"limit": ORDER_PAGE_SIZE}, timeout=ORDER_REQUEST_TIMEOUT_SECONDS)
```

Why: each meaningful value has one named owner (C04 `no-magic-literal`).

### Bad — literals in the call

```python
response = session.get("https://api.example.com/orders", params={"limit": 50}, timeout=7)
```

Why: the URL, page size, and timeout are unnamed and repeated wherever orders are fetched (C04).
