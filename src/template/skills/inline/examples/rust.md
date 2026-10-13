@rule no-inline-comment
### Good — the name carries the intent

```rust
let should_retry_upload = attempt_count < MAX_UPLOAD_ATTEMPTS;
```

Why: the names state what a comment would have said (C04 `no-inline-comment`).

### Bad — a comment explains a vague line

```rust
let ok = n < 3; // retry while under the limit
```

Why: the hook denies the added comment; `ok`, `n`, and `3` stay vague (C04, C05).

@rule no-magic-literal
### Good — a named constant in the domain constant owner

```rust
use crate::order::constant::paging::{ORDER_PAGE_SIZE, ORDER_REQUEST_TIMEOUT};

let page = order_store.read_page(ORDER_PAGE_SIZE, ORDER_REQUEST_TIMEOUT)?;
```

`src/order/constant/paging.rs` defines `pub const ORDER_PAGE_SIZE: usize = 50;` and `pub const ORDER_REQUEST_TIMEOUT: Duration = Duration::from_secs(7);`.

Why: each meaningful value has one named owner (C04 `no-magic-literal`).

### Bad — literals in the call

```rust
let page = order_store.read_page(50, Duration::from_secs(7))?;
```

Why: the page size and the timeout are unnamed and repeated wherever a page is read (C04).
