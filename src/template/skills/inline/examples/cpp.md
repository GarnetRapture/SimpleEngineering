@rule no-inline-comment
### Good — the name carries the intent

```cpp
const bool shouldRetryUpload = attemptCount < MAX_UPLOAD_ATTEMPTS;
```

Why: the variable and constant names state what a comment would have said (C04 `no-inline-comment`).

### Bad — a comment explains a vague line

```cpp
bool ok = n < 3; // retry while under the limit
```

Why: the hook denies the added comment; `ok`, `n`, and `3` stay vague (C04, C05).

@rule no-magic-literal
### Good — a named constant in the domain constant owner

```cpp
#include "order/constant/OrderPaging.hpp"

const auto page = orderStore.readPage(ORDER_PAGE_SIZE);
```

`order/constant/OrderPaging.hpp` declares `inline constexpr std::size_t ORDER_PAGE_SIZE = 50;`.

Why: the page size has one named owner (C04 `no-magic-literal`, C05 constant row).

### Bad — a literal in the call

```cpp
const auto page = orderStore.readPage(50);
```

Why: `50` is unnamed and repeated wherever a page is read (C04).
