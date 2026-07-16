---
title: Architecture
description: How OpenAPI schemas become runtime validations and targeted random JSON tests.
---

## Validation

`validation.Parse` selects JSON request bodies and query parameters by `operationId`, resolves reachable local references, and compiles each schema into a `Validation` tree.

For example:

```yaml
openapi: 3.0.3
paths:
  /things:
    post:
      operationId: createThing
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              required: [name]
              additionalProperties: false
              properties:
                name:
                  type: string
```

Compiles to data shaped like this:

```go
var createThing = &validation.Validation{
	SchemaPointer: "#/paths/~1things/post/requestBody/content/application~1json/schema",
	BodyRequired:  true,
	KindValidation: validation.KindValidation{
		Type: "object",
	},
	ObjectValidation: validation.ObjectValidation{
		Required: []string{"name"},
		Properties: []validation.PropertyValidation{{
			Name: "name",
			Validation: &validation.Validation{
				SchemaPointer: "#/paths/~1things/post/requestBody/content/application~1json/schema/properties/name",
				KindValidation: validation.KindValidation{
					Type: "string",
				},
				ObjectValidation: validation.ObjectValidation{
					AdditionalPropertiesAllowed: true,
				},
			},
		}},
		AdditionalPropertiesAllowed: false,
	},
}
```

Runtime parsing and generated literals produce the same compiled model. Generated source contains data, not generated validation functions. `Validate` walks that model while retaining raw JSON at every nested value.

For runtime validation, `allOf` stays as separate child validations. Each branch checks the same raw value, matching OpenAPI's rule that its schemas are [validated independently but together](https://spec.openapis.org/oas/v3.0.3.html#composition-and-inheritance-polymorphism).

## Test generation

The test generator does more than draw arbitrary JSON:

1. It compiles a schema into a constructive **Domain**: reachable JSON kinds plus exact constraints for numbers, strings, arrays, objects, and enums.
2. It plans focused **cases** from that domain: an aggregate valid case, useful valid partitions such as boundaries, and rejected cases that fail one constraint while satisfying the others when possible.
3. It attaches a Rapid generator to each case. The generator recursively draws random JSON values from that case's domain.
4. Each case runs as a Rapid property. The value is marshalled as exact JSON and passed to the validator callback. The planned accepted or rejected result is checked against that callback.

The cases provide coverage intent; Rapid provides variation and shrinking inside each case. For the object schema above, planned cases cover valid objects, a missing required `name`, a wrong type for `name`, and an unknown property. Each property can produce many concrete JSON bodies rather than one fixed fixture.

Schema parsing is fuzzed too. A separate Rapid generator builds supported OpenAPI documents, then makes independently mutated invalid copies. This tests both successful compilation and precise rejection of malformed schemas.

### Why `allOf` must merge before generating

Consider a schema where neither branch describes the final valid values:

```yaml
allOf:
  - type: integer
    minimum: 4
  - maximum: 10
    multipleOf: 3
```

At the `allOf` level there is no ready-made value to draw. Choosing from either branch alone is unsafe: `4` satisfies the first branch but not `multipleOf: 3`; `3` satisfies the second branch but not `minimum: 4`.

The generator handles it step by step:

1. Compile the first branch as integers greater than or equal to `4`.
2. Compile the second branch as all JSON kinds, with numbers no greater than `10` and restricted to multiples of `3`.
3. Intersect both domains. The numeric result is integers from `4` through `10` that are multiples of `3`.
4. Build the aggregate valid generator from that merged domain. It can draw `6` or `9`.
5. Build isolated rejected cases from the same context. For example, `3` fails only the minimum, `12` fails only the maximum, and values such as `4` or `10` fail only `multipleOf`.

The merge happens before Rapid generation. This is what lets random values remain meaningful: every accepted draw satisfies all branches, while rejected draws target a specific rule without accidentally failing unrelated siblings.
