# goutils

> **Author:** Onur Yaşar ([@onurid](https://github.com/onurid))
> **Part of:** [Rymory](https://rymory.org) — Open Identity Infrastructure
> © 2017–2026 Onur Yaşar. All rights reserved.

---

## What is goutils?

`goutils` is the shared Go utility library used across all Rymory identity services. It provides the common building blocks that every service in the Rymory ecosystem depends on — JWT authentication, role hierarchy, database connection, HTTP response helpers, cookie management and rate limiting.

Every package in `rymory-core` imports from this library.

---

## Packages

### `api` — Core API utilities

| File | Description |
|---|---|
| `model.go` | Shared types: `Response`, `Context`, `Token`, `CustomHttp`, `CustomHeader` |
| `JwtAuthentication.go` | JWT HS512 parsing, signature verification, context population |
| `convertLevel.go` | Role level extraction from JWT claims, project ID parsing |
| `result.go` | HTTP response helpers: `Message`, `Respond`, `ResMessage`, `CheckOk` |
| `httpCookie.go` | Cookie generation, validation, origin allowlist checking |
| `rateToken.go` | HMAC-based stateless rate limiting token system |
| `pasword.go` | Secure random password generation |
| `ipAddr.go` | Local IP address utilities |

### `db` — Database connection

| File | Description |
|---|---|
| `db.go` | PostgreSQL connection via GORM, singleton `GetDB()` |

---

## Role Hierarchy

Defined in `result.go` as constants:

```go
const (
    None = iota   // 0 — unauthenticated
    Root          // 1 — full system access
    MerchantAdmin // 2 — tenant owner
    Admin         // 3 — tenant admin
    Superuser     // 4 — elevated user
    User          // 5 — standard user
    Member        // 6 — application member
)
```

Role level is extracted from the first digit of the `roleId` field in JWT claims via `GetRoleLevel()`. This allows hierarchical permission checks at every service layer.

---

## JWT Token Structure

```go
type Token struct {
    UserId        uuid.UUID
    RoleId        int
    AppId         int
    MerchantId    uuid.UUID
    HasId         bool       // true = acting on behalf of another user
    ProjectId     int
    CustomData    string
    InitCompleted bool
    jwt.StandardClaims
}
```

Tokens are signed with HS512. Two signing keys are supported: `TOKEN_SECRET_KEY` for standard tokens, `TOKEN_ROOT_SECRET_KEY` for root-level tokens.

---

## Installation

```bash
go get github.com/rymory/goutils
```

Or with Go workspace (used in rymory-core):

```go
// go.work
use ./goutils
```

---

## Environment Variables

| Variable | Description |
|---|---|
| `DATABASE_URL` | PostgreSQL connection string |
| `TOKEN_SECRET_KEY` | JWT signing key (standard tokens) |
| `TOKEN_ROOT_SECRET_KEY` | JWT signing key (root tokens) |
| `X_API_KEY` | Internal service-to-service API key |
| `MAIN_DOMAIN` | Main domain for cookie scoping |
| `cookie_allowed_domains` | Comma-separated allowed origins |

---

## Usage Example

```go
import (
    u "github.com/rymory/goutils/api"
    d "github.com/rymory/goutils/db"
)

// JWT validation
context := &u.Context{}
if isOk, res := u.JwtAuthentication(authHeader, context); !isOk {
    return &res, nil
}

// Role check
roleLevel := u.GetRoleLevel(context.RoleId)
if roleLevel != u.Root {
    return u.Respond(u.Message(false, "Access denied"))
}

// DB query
db := d.GetDB()
db.Table("security.accounts").Where("user_id = ?", context.UserId).First(&account)

// Response
return u.Respond(u.Message(true, "Success"))
```

---

## Security Notes

- `rateToken.go` — the `secretKey` constant should be moved to an environment variable before production use. Current value is a placeholder.
- `pasword.go` — uses `math/rand` seeded with time; for cryptographic password generation consider `crypto/rand`.
- JWT parsing falls back to `TOKEN_ROOT_SECRET_KEY` if standard key fails — this is intentional for root token support.

---

## License

Licensed under **GNU AGPL v3** with Commercial Exception.
See [LICENSE](./LICENSE) for full terms.

Commercial licensing: onxorg@proton.me

---

## Part of Rymory

```
goutils           ← you are here (shared library)
rymory-core       ← identity backend (imports goutils)
rymory-gateway    ← edge proxy (imports goutils)
```

→ [rymory.org](https://rymory.org)
