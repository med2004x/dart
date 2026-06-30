# Project 12 - Data Modeling And Normalization

## Goal

Model an order system so each fact has one owner and historical facts remain
reproducible.

## Domain

Design:

```text
customers
products
orders
order_items
payment_attempts
```

Requirements:

- product price can change
- order item preserves charged unit price
- order contains multiple products
- product can appear in multiple orders
- payment can be attempted more than once
- order total must be reproducible later

## Deliverables

1. `model.md` defining every fact.
2. `schema.sql`.
3. primary/foreign/unique/check constraints.
4. sample data.
5. query reconstructing one invoice.
6. deletion and retention policy.

## Design Questions

- Why is `order_items.unit_price` not accidental duplication?
- Which status values are valid?
- Can products be deleted after sale?
- Which timestamp records each business event?
- Which totals are calculated versus stored?

## Failure Drills

1. store product IDs comma-separated in orders.
2. store only current product price.
3. use customer email as every foreign key.
4. duplicate order total without a consistency policy.

## Done Means

Changing current product/customer data cannot corrupt a historical invoice.

