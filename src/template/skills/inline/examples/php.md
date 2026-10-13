@rule no-inline-comment
### Good — the name carries the intent

```php
$unpaidOrders = array_filter($orders, $this->isAwaitingPayment(...));
```

Why: names state what a comment would have said (C04 `no-inline-comment`).

### Bad — a comment explains a vague line

```php
$list = array_filter($orders, fn ($o) => $o->s === 1); // unpaid
```

Why: the hook denies the added comment; the vague names stay (C04, C05).

@rule no-inline-style
### Good — a named class from the domain stylesheet

```php
<section class="order-summary-panel">
```

Why: the style lives in the order domain stylesheet under a name (C04 `no-inline-style`).

### Bad — an inline style attribute

```php
<section style="padding: 8px; color: #c00">
```

Why: the hook denies the `style="` attribute (C04).

@rule no-inline-handler
### Good — a handler bound from the domain script owner

```php
<button class="order-select" data-order-id="<?= $order->id ?>">
```

The domain script `assets/order/selectOrder.js` binds the click handler to `.order-select` by name.

Why: the behavior has one named owner (C05 `no-inline-handler`).

### Bad — an inline event attribute

```php
<button onclick="selectOrder(<?= $order->id ?>)">
```

Why: the hook denies the `onclick="` attribute (C05).

@rule no-magic-literal
### Good — a named class constant

```php
final class OrderRepository
{
    private const PAGE_SIZE = 50;
}
```

Why: the page size has one named owner (C04 `no-magic-literal`).

### Bad — literals inside the query

```php
$rows = $db->query('SELECT * FROM orders LIMIT 50');
```

Why: the table name and the page size are unnamed and repeated wherever orders are read (C04).
