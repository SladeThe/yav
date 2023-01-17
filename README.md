# Yet Another Validator

Go struct and field validation.

The project is inspired by and is a subset of [go-playground/validator](https://github.com/go-playground/validator). \
YAV aims to provide playground-like validator configuration and produce compatible errors whenever possible. \
At the same time, the introduced chained validator improves validation speed dramatically. \
The playground validator's performance is quite poor due to heavy reflection usage, which YAV strives not to use.

The main drawback of other builder-like validators (e.g. [ozzo](https://github.com/go-ozzo/ozzo-validation)) is that
they are `interface{}`-based and allocate a lot of memory upon each run.
Sometimes, it even makes them slower than the playground validator.

The repo also includes playground validator wrapper, which converts playground errors to YAV errors,
so that you can use chained validation along with the playground validator in the same project,
i.e. using the first one in hot paths and the second one in the places, where YAV still lacks features.
Also, this approach allows to migrate a project from playground to chained validation part-by-part.

The future plans are to introduce more playground and probably to add YAV's own features.

## Examples

#### Single field struct validation

The field name passed to `yav.Chain` doesn't affect the validation process and
is only necessary for building a validation error. 

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

#### Combine validation errors 

Use [multierr](https://github.com/uber-go/multierr) to combine multiple validation errors.

```go
type Password struct {
	Salt, Hash []byte
}

func (p Password) Validate() error {
	return multierr.Combine(
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
	return multierr.Combine(
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

#### Compare YAV and playground validator

Despite the YAV's playground wrapper registers JSON tags as field names using `RegisterTagNameFunc`,
the validator ignores it in `required_with/without` constructs.
Therefore, we pass uppercase `"Login"` to `vstring.RequiredWithAny` in order to produce playground-compatible errors.
If full compatibility is not required, pass the field names in whatever style you prefer.

```go
type Account struct {
    ID string `json:"id" validate:"required,uuid"`
    
    Login    string `json:"login" validate:"required,min=4,max=20,alphanum,lowercase"`
    Password string `json:"password" validate:"required_with=Login,omitempty,min=8,max=32,text"`
    
    Email string `json:"email" validate:"required,min=6,max=100,email"`
    Phone string `json:"phone" validate:"required,min=8,max=16,e164"`
}

func (a Account) Validate() error {
	return multierr.Combine(
		yav.Chain(
			"id", a.ID,
			vstring.Required,
			vstring.UUID,
		),
		yav.Chain(
			"login", a.Login,
			vstring.Required,
			vstring.Min(4),
			vstring.Max(20),
			vstring.Alphanumeric,
			vstring.Lowercase,
		),
		yav.Chain(
			"password", a.Password,
			vstring.RequiredWithAny().String(a.Login).Names("Login"),
			vstring.OmitEmpty,
			vstring.Between(8, 32),
			vstring.Text,
		),
		yav.Chain(
			"email", a.Email,
			vstring.Required,
			vstring.Between(6, 100),
			vstring.Email,
		),
		yav.Chain(
			"phone", a.Phone,
			vstring.Required,
			vstring.Between(8, 16),
			vstring.E164,
		),
	)
}
```

## Available validations

#### Bytes

```
OmitEmpty
Required
RequiredWithAny
RequiredWithoutAny
RequiredWithAll
RequiredWithoutAll

Min
Max
Between
```

#### Map

```
OmitEmpty
Required
RequiredWithAny
RequiredWithoutAny
RequiredWithAll
RequiredWithoutAll

Unique

Min
Max
Between

Keys
Values
```

#### Number

```
OmitEmpty
Required

Min
Max
Between
LessThan
LessThanOrEqual
GreaterThan
GreaterThanOrEqual
```

#### Slice

```
OmitEmpty
Required
RequiredWithAny
RequiredWithoutAny
RequiredWithAll
RequiredWithoutAll

Unique

Min
Max

Items
```

#### String

```
OmitEmpty
Required
RequiredWithAny
RequiredWithoutAny
RequiredWithAll
RequiredWithoutAll

Min
Max
Between

Equal
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
URI
URL
UUID
```

#### Time

```
OmitEmpty
Required
```

## Benchmarks

```
goos: windows
goarch: amd64
cpu: Intel(R) Core(TM) i9-10850K CPU @ 3.60GHz
```

#### Tiny struct validation

```
BenchmarkYAV              12907930       92.15 ns/op          0 B/op        0 allocs/op
BenchmarkOzzo              1334562       890.1 ns/op       1248 B/op       20 allocs/op
BenchmarkPlayground        1324868       911.8 ns/op         40 B/op        2 allocs/op
```

#### Account struct validation

```
BenchmarkYAV                562638        2127 ns/op        123 B/op        4 allocs/op
BenchmarkOzzo*               64515       17206 ns/op      15487 B/op      253 allocs/op
BenchmarkPlayground         189013        6204 ns/op        587 B/op       21 allocs/op
```

#### Notes

* The Account in the Examples section is a reduced version of the [benchmarked structure](tests/account_test.go).
* Ozzo validator lacks some features available in both YAV and playground validator.
  Therefore, those validation steps were not enabled for ozzo.
* The YAV is still slower, than a manually written validation boilerplate, but the amount of code differs dramatically.
