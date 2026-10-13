@rule domain-structure
### Good — the concept lives under its domain module

```
src/order/pricing/order_total.rs
src/order/types/order_line.rs
```

Why: domain module `order`, class `pricing`/`types`, snake_case module files (C04, C05 Rust row).

### Bad — a utility module

```
src/util.rs
src/helpers/mod.rs
```

Why: `util` and `helpers` name no domain, so unrelated functions gather in one module (C04, C05).

@rule domain-naming
### Good — action and target in the name

```rust
pub fn calculate_order_total(lines: &[OrderLine]) -> Money {
    lines.iter().map(OrderLine::amount).sum()
}
```

Why: snake_case action + target for a function, PascalCase for types (C05 Rust row).

### Bad — a name that hides the responsibility

```rust
pub fn do_stuff(data: Vec<String>) -> bool {
    !data.is_empty()
}
```

Why: `do_stuff` and `data` state no action, target, or meaning of the returned boolean (C05).
