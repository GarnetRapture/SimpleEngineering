@rule domain-structure
### Good — the concept lives under its area and function, with a declaration and implementation pair

```
order/pricing/OrderPricing.hpp
order/pricing/OrderPricing.cpp
```

Why: area `order`, function `pricing`, one declaration and implementation pair per type (C04, C10, C05 C++ row).

### Bad — a common bucket

```
common/Utils.h
common/Utils.cpp
```

Why: `common` and `Utils` give every concept a second owner and name no responsibility (C04, C05).

@rule domain-naming
### Good — action and target in the name

```cpp
class OrderPricing {
public:
    Money calculateOrderTotal(std::span<const OrderLine> lines) const;
};
```

Why: PascalCase type, lowerCamelCase method stating action and target, a typed view of the input (C05 C++ row).

### Bad — a name that hides the responsibility

```cpp
void process(void* data);
```

Why: `process` and `data` state nothing, and `void*` erases the input contract and its lifetime (C05, C03).
