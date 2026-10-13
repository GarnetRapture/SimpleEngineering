@rule no-inline-comment
### Good — the name carries the intent

```ts
const unpaidOrders = orders.filter(isAwaitingPayment);
```

Why: the variable and predicate names state what a comment would have said (C04 `no-inline-comment`).

### Bad — a comment explains a vague line

```ts
const list = orders.filter(o => o.s === 1); // unpaid orders
```

Why: the hook denies the added comment; the vague names `list`, `o`, `s`, and `1` remain (C04, C05).

@rule no-inline-type
### Good — the type lives in the domain types folder

```
src/order/types/OrderSummaryPanelProps.ts   export interface OrderSummaryPanelProps { order: Order }
src/order/view/OrderSummaryPanel.tsx         import type { OrderSummaryPanelProps } from '../types/OrderSummaryPanelProps';
```

Why: the `types/` path is the allowed owner path of the `no-inline-type` rule, so the type has one owner (C04).

### Bad — an ad hoc type inside the component file

```tsx
type Props = { order: Order };
export function OrderSummaryPanel({ order }: Props) { … }
```

Why: the hook counts `type Props` as an added inline type in a non-`types/` file and denies the edit (C04 `no-inline-type`).

@rule no-inline-style
### Good — a named class from the domain style owner

```tsx
import orderSummaryStyles from '../style/OrderSummaryPanel.module.css';
<section className={orderSummaryStyles.panel}>
```

Why: the style lives in its domain style owner under a name (C04 `no-inline-style`).

### Bad — an inline style object

```tsx
<section style={{ padding: 8, color: '#c00' }}>
```

Why: the hook denies `style={{`; the values become unnamed literals scattered across components (C04).

@rule no-inline-handler
### Good — a named handler passed by name

```tsx
function selectOrder() { onOrderSelect(order.id); }
<button onClick={selectOrder}>
```

Why: the handler name states action and target and is referenced by name (C05 `no-inline-handler`).

### Bad — an ad hoc arrow function

```tsx
<button onClick={() => onOrderSelect(order.id)}>
```

Why: the hook denies `onClick={() =>`; the behavior has no name to search, test, or reuse (C05).

@rule no-magic-literal
### Good — a named constant in the domain constant owner

```ts
import { ORDER_PAGE_SIZE, ORDER_API_PATH } from '../constant/orderPaging';
const response = await fetch(`${ORDER_API_PATH}?limit=${ORDER_PAGE_SIZE}`);
```

Why: each meaningful value has one named owner (C04 `no-magic-literal`).

### Bad — literals in the call

```ts
const response = await fetch('https://api.example.com/orders?limit=50');
```

Why: the URL and the page size are unnamed and duplicated wherever orders are fetched (C04).
