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
			vstring.IsUUID,
		),
		yav.Chain(
			"login", a.Login,
			vstring.Required,
			vstring.Min(4),
			vstring.Max(20),
			vstring.IsAlphanumeric,
			vstring.IsLowercase,
		),
		yav.Chain(
			"password", a.Password,
			vstring.RequiredWithAny("Login", yav.RequiredWithAny().String(a.Login)),
			vstring.OmitEmpty,
			vstring.Min(8),
			vstring.Max(32),
			vstring.IsText,
		),
		yav.Chain(
			"email", a.Email,
			vstring.Required,
			vstring.Min(6),
			vstring.Max(100),
			vstring.IsEmail,
		),
		yav.Chain(
			"phone", a.Phone,
			vstring.Required,
			vstring.Min(8),
			vstring.Max(16),
			vstring.IsE164,
		),
	)
}
```

#### Notes

* YAV doesn't need the validate field attribute. It is added, so that you can compare YAV and playground code.
* Despite the YAV's playground wrapper registers JSON tags as field names using `RegisterTagNameFunc`,
  the validator ignores it in `required_with/without` constructs.
  Therefore, we pass uppercase `"Login"` to `vstring.RequiredWithAny` in order to produce playground-compatible errors.
  If full compatibility is not required, pass the field names in whatever style you prefer.

## Benchmarks

```
goos: windows
goarch: amd64
cpu: Intel(R) Core(TM) i9-10850K CPU @ 3.60GHz
```

#### Tiny struct validation

```
BenchmarkChain            12907930       92.15 ns/op          0 B/op        0 allocs/op
BenchmarkOzzo*             1334562       890.1 ns/op       1248 B/op       20 allocs/op
BenchmarkPlayground        1324868       911.8 ns/op         40 B/op        2 allocs/op
```

#### Account struct validation

```
BenchmarkChain              805762        1526 ns/op          0 B/op        0 allocs/op
BenchmarkOzzo*               92058       13223 ns/op      11908 B/op      194 allocs/op
BenchmarkPlayground         365997        3216 ns/op        189 B/op        4 allocs/op
```

#### Notes

* The Account in the Examples section is a reduced version of the structure.
* Ozzo validator lacks some playground features. Therefore, those validation steps were not enabled for ozzo.
* The YAV is still slower, than a manually written validation boilerplate, but the amount of code differs dramatically.
