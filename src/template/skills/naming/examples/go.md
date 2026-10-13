@rule domain-structure
### Good — the concept lives in its domain package

```
order/pricing/order_total.go
order/types/order_line.go
```

Why: domain `order`, class `pricing`/`types`, lowercase single-word package names (C04, C05 Go row).

### Bad — a meaningless package

```
util/util.go
common/helpers.go
```

Why: the Go blog "Package names" calls `util`, `common`, and `misc` meaningless package names because they tell clients nothing about the contents (C04, C05).

@rule domain-naming
### Good — action and target in the name

```go
func CalculateOrderTotal(lines []OrderLine) Money {
	total := ZeroMoney
	for _, line := range lines {
		total = total.Add(line.Amount)
	}
	return total
}
```

Why: MixedCaps action + target; the exported name starts with an upper-case letter, which is the Go export rule (C05 Go row).

### Bad — a name that hides the responsibility

```go
func Process(data interface{}) error {
	return handle(data)
}
```

Why: `Process`, `data`, and `handle` state nothing, and `interface{}` erases the input contract (C05).
