# API Engineering Quick Reference

Use this page for the recurring HTTP and Go shapes. The numbered projects still
define the required behavior and evidence.

## Request To Response

```text
client -> method + URL + headers + body
server -> parse -> validate -> authorize -> business rule -> storage
server -> status + headers + body
```

Keep those responsibilities visible. A handler should translate HTTP; it should
not become the place where every business rule and database query lives.

## Basic Go Handler

```go
func getStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
```

Write headers before the body. After `WriteHeader`, the status cannot be
changed. Return after an error response.

## JSON Input And Output

```go
type CreateLabelRequest struct {
	Name string `json:"name"`
}

var input CreateLabelRequest
if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
	http.Error(w, "invalid JSON", http.StatusBadRequest)
	return
}
```

Decode into a request type, validate required fields, and reject malformed or
unknown input according to the contract. Do not trust JSON fields merely
because they decoded successfully.

## Common Status Choices

| Situation | Status |
|---|---:|
| successful read | `200 OK` |
| successful create | `201 Created` |
| accepted asynchronous work | `202 Accepted` |
| malformed input | `400 Bad Request` |
| unauthenticated | `401 Unauthorized` |
| authenticated but forbidden | `403 Forbidden` |
| resource missing | `404 Not Found` |
| method unsupported for resource | `405 Method Not Allowed` |
| conflicting version or duplicate | `409 Conflict` |
| unexpected server failure | `500 Internal Server Error` |

Choose the status from the observable contract, not from the error string.

## Query Parameters And Pagination

Parse query parameters once at the boundary. Validate limits and ordering. For
stable pagination, use a deterministic sort and a cursor or a documented
offset rule; do not rely on accidental database order.

## Context And Timeouts

```go
ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
defer cancel()

result, err := service.Load(ctx, id)
```

Pass the request context downward. Every external call needs a deadline. A
timeout is a controlled failure; an unbounded wait is a capacity problem.

## Best Practices

- Define request, response, status, and error examples before coding.
- Validate at the HTTP boundary, authorize before data access, and check
  ownership using server-side identity.
- Make retries safe by defining idempotency behavior.
- Do not expose database errors or secrets in responses.
- Test through the HTTP boundary with valid, invalid, missing, and repeated
  requests.

## Official Resources

- [Go `net/http`](https://pkg.go.dev/net/http)
- [Go `encoding/json`](https://pkg.go.dev/encoding/json)
- [Go `context`](https://pkg.go.dev/context)
- [OpenAPI specification](https://spec.openapis.org/oas/latest.html)
- [HTTP semantics, RFC 9110](https://www.rfc-editor.org/rfc/rfc9110)
- [HTTP method reference](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Methods)
