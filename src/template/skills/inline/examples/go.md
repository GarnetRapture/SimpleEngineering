@rule no-inline-comment
### Good — the name carries the intent

```go
shouldRetryUpload := attemptCount < MaxUploadAttempts
```

Why: the names state what a comment would have said (C04 `no-inline-comment`).

### Bad — a comment explains a vague line

```go
ok := n < 3 // retry while under the limit
```

Why: the hook denies the added comment; `ok`, `n`, and `3` stay vague (C04, C05).

@rule no-magic-literal
### Good — a named constant in the domain constant owner

```go
page, err := orderStore.ReadPage(ctx, paging.OrderPageSize)
```

`order/paging/paging.go` declares `const OrderPageSize = 50`.

Why: the page size has one named owner, written in MixedCaps as Effective Go names constants (C04 `no-magic-literal`, C05 Go row).

### Bad — a literal in the call

```go
page, err := orderStore.ReadPage(ctx, 50)
```

Why: `50` is unnamed and repeated wherever a page is read (C04).
