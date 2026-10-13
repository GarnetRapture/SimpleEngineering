@rule domain-structure
### Good — the concept lives under its domain

```
src/order/logic/calculateOrderTotal.ts
src/order/types/OrderLine.ts
src/order/view/OrderSummaryPanel.tsx
```

Why: domain `order`, class `logic`/`types`/`view`, and an export file named after its symbol (C04, C05 TS row).

### Bad — a technical bucket at the top level

```
src/utils/orderStuff.ts
src/components/Helpers.tsx
```

Why: `utils`, `components`, and `Helpers` name a technique, not a domain; the order concept gets owners in several buckets (C04).

@rule domain-naming
### Good — action and target in the name

```ts
export function calculateOrderTotal(lines: readonly OrderLine[]): Money {
  return lines.reduce(addLineAmount, ZERO_MONEY);
}
```

Why: lowerCamelCase action + target for a function, PascalCase for the types (C05 TS row).

### Bad — a name that hides the responsibility

```ts
export function process(data: any) {
  return data.items.map(handle);
}
```

Why: `process`, `data`, and `handle` state no action, target, or result; `any` hides the input contract (C05).
