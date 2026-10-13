@rule domain-structure
### Good — the concept lives under its domain namespace

```
src/Order/Pricing/OrderTotalCalculator.php
src/Order/Type/OrderLine.php
```

Why: domain `Order`, class `Pricing`/`Type`, the class file named after its class (C04, C05 PHP row).

### Bad — a global functions file

```
src/functions.php
src/Helpers/Util.php
```

Why: `functions`, `Helpers`, and `Util` name no domain, so every concept can land there (C04, C05).

@rule domain-naming
### Good — action and target in the name

```php
final class OrderTotalCalculator
{
    public function calculateOrderTotal(OrderLine ...$lines): Money
    {
        return array_reduce($lines, Money::add(...), Money::zero());
    }
}
```

Why: PascalCase class, lowerCamelCase method stating action and target, typed input and result (C05 PHP row).

### Bad — a name that hides the responsibility

```php
function doStuff($data)
{
    return handle($data);
}
```

Why: `doStuff`, `$data`, and `handle` state nothing, and the untyped input hides the contract (C05).
