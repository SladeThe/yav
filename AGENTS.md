# YAV

Go validation library at `github.com/SladeThe/yav`.
Priorities: fast validation with as few heap allocations as possible, then a statically checked API.
Use explicit typed checks; field tags and reflective struct traversal are not the API.
Use the latest go-playground/validator checks as a behavior reference. Fix confirmed bugs even when upstream shares them.

## Layout

- `yav.go` defines validation functions and composition; `errors.go` defines structured errors and aggregation.
- `v*` packages implement checks for each value kind. `vslice`, `vmap`, and `vpointer` also compose checks over contained values.
- `accumulators` builds cross-field presence conditions; `common` holds character predicates.
- `internal` contains shared factory-cache registration. `examples` contains executable API usage.
- Keep this module self-contained. External consumers must not become its dependencies or required test fixtures.

## API contracts

`yav.ValidateFunc[T]` is `func(name string, value T) (stop bool, err error)`.
The results are independent: `false, err` reports an error and continues; `true, nil` skips later checks.
Combinators must collect an error before honoring `stop`.

Use `yav.Chain` for one value and `yav.Join` for independent fields.
For example: `yav.Chain("login", login, vstring.Required, vstring.Between(4, 20))`.
Names only label errors; they do not look up struct fields. A successful chain returns nil; a failing chain returns `yav.Errors`.

Use `yav.Nested` or `yav.NestedValidate` for namespaced errors.
Reusable `Validate` methods use relative names: empty for the value itself, local field/index names for its contents. Containing validators add their field names with `Nested`.
`yav.Or` selects the first nil-error result, otherwise the last failure and its stop flag.
Keep `CheckName`, `Parameter`, `ValueName`, and `Value` consistent with neighboring checks.
Nested and NamedCheck copy validation entries before rewriting them. Collection conversion appends transformed entries directly into the destination.
Append clips adopted slice capacity to prevent append overwrites. AppendConverted applies convertFunc only to validation errors; nil uses Append.
Merging an empty error slice preserves the destination's nilness and capacity.
Public slices and error payloads can still alias; direct mutation requires caller-managed ownership.

`Error.Is` compares `CheckName`, `Parameter`, and `ValueName`, ignoring `Value` on both sides.
For `errors.Is(err, target)`, use an `Error` target with `Value` left nil; `err` may retain any payload.

`Required` rejects the value kind's empty value; `OmitEmpty` skips remaining checks for it.
Strings, slices, maps, and bytes use length; pointers use nil; numbers and durations use zero; bool uses false; time uses `IsZero`.
Inactive conditional required checks use `OmitEmpty`. Inactive excluded checks use `yav.Next`; exclusion alone does not skip later checks.

Accumulators are value builders: keep the result of each fluent call.
`Names` supplies diagnostic labels and returns a validator.
Conditions and compared values are captured during construction. Recompute input-dependent rules for each object; reuse rules with fixed configuration.

`vslice.Items` and `vmap.Keys`/`Values`/`Entries` validate contained values.
`vpointer.Dereference` requires a non-nil pointer; precede it with `vpointer.OmitEmpty` for optional pointers.
Preserve paths such as `field[0].child` and `field[0][1]`; do not assume map iteration order.
Collection names match complete field/index prefixes; an empty collection name also accepts relative field names.

String length checks count bytes. Alpha/digit helpers are ASCII; case and whitespace checks use Unicode.
`Lowercase` and `Uppercase` accept valid UTF-8 unchanged by `strings.ToLower`/`strings.ToUpper`, including empty strings.
Keep these semantics explicit in tests and documentation.

## Numeric generation

`vnumber/_*.go` files are templates, including test templates.
`vnumber/gen.go` pins generation and import-formatting tools and expands twelve primitive numeric types.
Edit templates, run `go generate ./vnumber`, and review the generated diff together. Do not edit generated outputs alone.
`vnumber/required.go` and its tests are handwritten generics.
Primitive-specific factories intentionally require exact types. Convert defined numeric values at the validation boundary, e.g. `int(age)` for `type Age int`.

## Performance and testing

Preserve valid-input performance first; discuss error-construction cost separately. Build error payloads and format diagnostics only where needed.
Implement built-in fixed formats with native checks; regexp use is reserved for the explicit public Regexp adapter.
When replacing a regexp, document an expression matching the native helper's behavior, including any separate length limits.
Measure changes instead of assuming a closure, interface conversion, or byte conversion allocates.
When performance varies by input, prefer substantially simpler or shorter code; saving one or two lines alone is rarely decisive.
Do not manually inline helper bodies at call sites to tune compiler output; keep the helper calls.
Separate fixed-rule reuse from construction, warm caches from cold misses, and valid from invalid inputs.
Global caches need bounded-retention reasoning and synchronized reads and publication; copying a map alone is insufficient.

Follow existing table-driven test style and check both return values plus complete error metadata.
Cover composition, boundary values, nil/empty inputs, named types, Unicode, and NaN where relevant.
For hot-path changes, record `ns/op`, `B/op`, and `allocs/op`; use focused allocation assertions for stable guarantees.
Give each benchmark a clear unit, exclude fixture setup, retain results, and compare representative sizes and failure positions.
Compare implementations within one process; this machine's two CCDs make separate-process timings unsuitable for comparison.

Run commands from the repository root using the `toolchain` version in `go.mod`. Make exports it as `GOTOOLCHAIN`; CI reads it through setup-go.
For direct Go commands, invoke the executable matching `toolchain` in `go.mod`; the directive does not downgrade a newer installed Go.
Keep each invocation standalone; do not prepend environment assignments that prevent reusable command approvals.

```text
go test -mod=readonly ./...
go vet -mod=readonly ./...
go test -mod=readonly -run '^$' -bench . -benchmem ./...
go test -mod=readonly -race ./...
```

The race detector requires an appropriate CGO toolchain. Include concurrent cold-construction tests for cache changes.
`make test` also produces coverage but uses POSIX cleanup commands; direct Go commands work in PowerShell.
Keep CI, the linter version, and its configuration aligned with `go.mod`.
Run lint exactly as defined in `Makefile`, including its `GOTOOLCHAIN` setting. Update the recipe before changing lint flags.

## Repository hygiene

For every multiline delimiter pair (`()`, `[]`, `{}`), give the opening and closing lines the same indentation.
Separate ordinary statements from large adjacent control-flow blocks with a blank line. Always leave a blank line after a block before a following ordinary statement.
Separate adjacent code blocks with a blank line. A short `if` (body at most three lines) may stay next to a related preceding assignment, such as an error check.
Prefer self-documenting code. Outside public API docs, add comments only to prevent likely mistakes by future readers or editors.
Enforce only style rules linters can check without false positives; keep unsupported cases in review.
WSL checks blank lines from control-flow blocks to expression statements; its other rules are excluded.
Use bracketed Go doc links only when they resolve correctly and preserve IDE navigation. Keep declaration opening names and local parameter/field names plain.
Keep references to unimported packages and unsupported members qualified and plain to preserve IDE navigation; do not use full import paths or explicit URLs.
Preserve unrelated working-tree edits. Keep generated code, templates, tests, examples, and docs synchronized with relevant changes.
Keep local agent state and review artifacts in ignored directories. The root `AGENTS.md` belongs in version control.
Keep agent-document lines at most 280 characters, preferably at most 240.
