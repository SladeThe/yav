# Yet Another Validator

Go struct and field validation.

The project is inspired by [go-playground/validator](https://github.com/go-playground/validator) and uses some of its codebase. \
YAV aims to provide Playground-like validator configuration and produce compatible errors when possible. \
At the same time, the introduced chained validator improves validation speed dramatically. \
The Playground validator's performance is quite poor due to heavy reflection usage, which YAV strives not to use.

YAV's key principles:
- mimic Playground validator when possible;
- make fewer to zero allocations;
- work fast.

The main drawback of other builder-like validators (e.g. [ozzo](https://github.com/go-ozzo/ozzo-validation)) is that
they are `interface{}`-based and allocate a lot of memory upon each run.
Sometimes, it even makes them slower than the Playground validator.

Unlike in earlier versions, the repo no longer includes Playground validator wrapper in order to reduce the number
of 3rd-party dependencies. The removed code still can be found in [yav-tests](https://github.com/SladeThe/yav-tests).

## Examples

#### Single field struct validation

The field name passed to `yav.Chain` doesn't affect the validation process and is only necessary
for building a validation error, so that you may use whatever naming style you like, i.e. `id`, `ID`, `Id`.

```go
type AccountID struct {
	ID string
}

func (id AccountID) Validate() error {
	return yav.Chain(
		"id", id.ID,
		vstring.Required,
		vstring.UUID,
	)
}
```

#### Defined numeric types

Type-specific numeric factories, such as `vnumber.MinInt`, accept their exact primitive type.
Convert defined numeric values at the validation boundary.
A value validator leaves its name empty; the containing validator adds the field name with `yav.Nested`:

```go
type Age int

func (age Age) Validate() error {
	return yav.Chain("", int(age), vnumber.Required[int], vnumber.MinInt(18))
}

type User struct {
	Age Age
}

func (u User) Validate() error {
	return yav.Nested("age", u.Age.Validate())
}
```

Checks that report the value store the converted type in `Error.Value`.

#### Combine validation errors

Use `yav.Join` to combine multiple validation errors.

```go
type Password struct {
	Salt, Hash []byte
}

func (p Password) Validate() error {
	return yav.Join(
		yav.Chain(
			"salt", p.Salt,
			vbytes.Required,
			vbytes.Max(200),
		),
		yav.Chain(
			"hash", p.Hash,
			vbytes.Required,
			vbytes.Max(200),
		),
	)
}
```

#### Validate nested structs

Use `yav.Nested` to add value namespace, i.e. to get `password.salt` error instead of just `salt`. \
Contrary, any possible `id` error is returned as if the field were in the `Account` struct directly.

```go
type Account struct {
	AccountID

	Login    string
	Password Password
}

func (a Account) Validate() error {
	return yav.Join(
		a.AccountID.Validate(),
		yav.Chain(
			"login", a.Login,
			vstring.Required,
			vstring.Between(4, 20),
			vstring.Alphanumeric,
			vstring.Lowercase,
		),
		yav.Nested("password", a.Password.Validate()),
	)
}
```

#### Match validation errors

`yav.Error.Is` matches `CheckName`, `Parameter`, and `ValueName` exactly, ignoring `Value` on both sides.

For [`errors.Is`](https://pkg.go.dev/errors#Is), the right-hand `target` must be comparable.
Leave `Value` nil in YAV targets; the actual error may retain any `Value`.

```go
err := yav.ErrUnique("items", []int{1, 1})
target := yav.ErrUnique("items", nil)

errors.Is(err, target) // true
```

`errors.Is` may compare values before calling `Error.Is`.
Using an error with a slice, map, or other non-comparable `Value` as the target can panic,
including `errors.Is(err, err)`.

#### Compare YAV and Playground validator

Here we pass to YAV Go-like field names in order to produce Playground-compatible errors. \
YAV doesn't anyhow use it, except while building validation errors. \
If compatibility is not required, pass the field names in whatever style you prefer.

```go
type Account struct {
    ID string `validate:"required,uuid"`
    
    Login    string `validate:"required,min=4,max=20,alphanum,lowercase"`
    Password string `validate:"required_with=Login,omitempty,min=8,max=32,text"`
    
    Email string `validate:"required,min=6,max=100,email"`
    Phone string `validate:"required,min=8,max=16,e164"`
}

func (a Account) Validate() error {
	return yav.Join(
		yav.Chain(
			"ID", a.ID,
			vstring.Required,
			vstring.UUID,
		),
		yav.Chain(
			"Login", a.Login,
			vstring.Required,
			vstring.Min(4),
			vstring.Max(20),
			vstring.Alphanumeric,
			vstring.Lowercase,
		),
		yav.Chain(
			"Password", a.Password,
			vstring.RequiredWithAny().String(a.Login).Names("Login"),
			vstring.Between(8, 32),
			vstring.Text,
		),
		yav.Chain(
			"Email", a.Email,
			vstring.Required,
			vstring.Between(6, 100),
			vstring.Email,
		),
		yav.Chain(
			"Phone", a.Phone,
			vstring.Required,
			vstring.Between(8, 16),
			vstring.E164,
		),
	)
}
```

## Available validations

#### Common

```
OmitEmpty
Required
RequiredIf
RequiredUnless
RequiredWithAny
RequiredWithoutAny
RequiredWithAll
RequiredWithoutAll
ExcludedIf
ExcludedUnless
ExcludedWithAny
ExcludedWithoutAny
ExcludedWithAll
ExcludedWithoutAll
```

#### Bool

```
Equal
NotEqual
```

#### Bytes

```
Min
Max
Between
```

#### Duration

```
Min
Max
Between
LessThan
LessThanOrEqual
GreaterThan
GreaterThanOrEqual

LessThanNamed
LessThanOrEqualNamed
GreaterThanNamed
GreaterThanOrEqualNamed
```

#### Map

```
Min
Max
Between

Unique

Keys
Values
```

#### Number

```
Min
Max
Between
LessThan
LessThanOrEqual
GreaterThan
GreaterThanOrEqual

Equal
NotEqual
OneOf
```

#### Slice

```
Min
Max
Between

Unique

Items
```

#### String

```
Min
Max
Between

Equal
NotEqual
OneOf

Alpha
Alphanumeric
Lowercase
Uppercase
ContainsAlpha
ContainsLowerAlpha
ContainsUpperAlpha
ContainsDigit
ContainsSpecialCharacter
ExcludesWhitespace
StartsWithAlpha
StartsWithLowerAlpha
StartsWithUpperAlpha
StartsWithDigit
StartsWithSpecialCharacter
EndsWithAlpha
EndsWithLowerAlpha
EndsWithUpperAlpha
EndsWithDigit
EndsWithSpecialCharacter

Text
Title

E164
Email
Hostname
HostnameRFC1123
HostnamePort
FQDN
URI
URL
UUID

Regexp
```

#### Time

```
Min
Max
Between
LessThan
LessThanOrEqual
GreaterThan
GreaterThanOrEqual

LessThanNamed
LessThanOrEqualNamed
GreaterThanNamed
GreaterThanOrEqualNamed
```

## Benchmarks

Valid Account validation, measured on 2026-09-18 with Go 1.27.1 on Windows/amd64,
AMD Ryzen 9 9900X, `GOAMD64=v1`, and `GOMAXPROCS=24`.
Versions: YAV v0.17.0, go-playground/validator v10.15.1, and ozzo-validation v4.4.1.

Each operation validates one account. Results are medians of five 300 ms samples collected in one process.

### Sequential

| Validator | ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| YAV | 438.7 | 0 | 0 |
| Preallocated YAV | 375.7 | 0 | 0 |
| Playground v10.15.1 | 3078 | 559 | 15 |
| Ozzo v4.4.1 | 6324 | 6255 | 107 |

### Parallel

| Validator | ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| YAV | 32.74 | 0 | 0 |
| Preallocated YAV | 29.44 | 0 | 0 |
| Playground v10.15.1 | 509.1 | 552 | 15 |
| Ozzo v4.4.1 | 2042 | 6211 | 107 |

Parallel ns/op reflects aggregate throughput across 24 workers.
Preallocated YAV reuses rules prepared before validation.

All 32 compatibility tests pass against Playground v10.15.1.
The Ozzo fixture omits some password/name rules and avatar key/value checks.

The Account in the Examples section is a reduced version of the [benchmarked structure](https://github.com/SladeThe/yav-tests/blob/main/account_test.go).
