# API Contract

| Method | Path | Auth | Success | Errors |
|---|---|---|---|---|
| GET | `/health` | public | 200 | 405 |
| GET | `/ready` | public | 200 | 503 |
| POST | `/projects/{id}/tasks` | required | 201 | |
| GET | `/projects/{id}/tasks` | required | 200 | |
| GET | `/projects/{id}/tasks/{taskID}` | required | 200 | |
| PUT | `/projects/{id}/tasks/{taskID}` | required | 200 | |
| DELETE | `/projects/{id}/tasks/{taskID}` | required | 204 | |

## Pagination

- Order:
- Cursor:
- Maximum page size:
- Invalid cursor behavior:

## Error Body

```json
{"error":"message","requestId":"req-id"}
```

