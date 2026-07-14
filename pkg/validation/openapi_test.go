//nolint:dupl,lll,maintidx // Complete inline expectations intentionally contain large, repeated graphs.
package validation_test

import (
	"os"
	"testing"

	"decode_and_validate_generator/pkg/validation"

	"github.com/stretchr/testify/require"
)

// TestParseOpenAPI verifies every compiled validation in the example OpenAPI document.
func TestParseOpenAPI(t *testing.T) {
	t.Parallel()

	spec, err := os.ReadFile("../../resources/openapi.yaml")
	require.NoError(t, err)

	tests := map[string]struct {
		expectedValidation *validation.Validation
	}{
		"allOfObject": {
			expectedValidation: &validation.Validation{
				SchemaPointer:    "#/paths/~1all-of-object/post/requestBody/content/application~1json/schema",
				BodyRequired:     true,
				KindValidation:   validation.KindValidation{Type: "", Nullable: false},
				EnumValidation:   validation.EnumValidation{},
				NumberValidation: validation.NumberValidation{},
				StringValidation: validation.StringValidation{Format: ""},
				ArrayValidation: validation.ArrayValidation{
					UniqueItems: false,
				},
				ObjectValidation: validation.ObjectValidation{
					AdditionalPropertiesAllowed: true,
				},
				AllOfValidations: []*validation.Validation{
					{
						SchemaPointer:    "#/paths/~1all-of-object/post/requestBody/content/application~1json/schema/allOf/0",
						BodyRequired:     false,
						KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
						EnumValidation:   validation.EnumValidation{},
						NumberValidation: validation.NumberValidation{},
						StringValidation: validation.StringValidation{Format: ""},
						ArrayValidation: validation.ArrayValidation{
							UniqueItems: false,
						},
						ObjectValidation: validation.ObjectValidation{
							Required: []string{"first"},
							Properties: []validation.PropertyValidation{
								{Name: "first", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1all-of-object/post/requestBody/content/application~1json/schema/allOf/0/properties/first",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
							},
							AdditionalPropertiesAllowed: true,
						},
					},
					{
						SchemaPointer:    "#/paths/~1all-of-object/post/requestBody/content/application~1json/schema/allOf/1",
						BodyRequired:     false,
						KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
						EnumValidation:   validation.EnumValidation{},
						NumberValidation: validation.NumberValidation{},
						StringValidation: validation.StringValidation{Format: ""},
						ArrayValidation: validation.ArrayValidation{
							UniqueItems: false,
						},
						ObjectValidation: validation.ObjectValidation{
							Required: []string{"second"},
							Properties: []validation.PropertyValidation{
								{Name: "second", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1all-of-object/post/requestBody/content/application~1json/schema/allOf/1/properties/second",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "boolean", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
							},
							AdditionalPropertiesAllowed: true,
						},
					},
					{
						SchemaPointer:    "#/paths/~1all-of-object/post/requestBody/content/application~1json/schema/allOf/2",
						BodyRequired:     false,
						KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
						EnumValidation:   validation.EnumValidation{},
						NumberValidation: validation.NumberValidation{},
						StringValidation: validation.StringValidation{Format: ""},
						ArrayValidation: validation.ArrayValidation{
							UniqueItems: false,
						},
						ObjectValidation: validation.ObjectValidation{
							Required: []string{"last"},
							Properties: []validation.PropertyValidation{
								{Name: "last", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1all-of-object/post/requestBody/content/application~1json/schema/allOf/2/properties/last",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "number", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
							},
							AdditionalPropertiesAllowed: true,
						},
					},
				},
			},
		},
		"arrayNotNullable": {
			expectedValidation: &validation.Validation{
				SchemaPointer:    "#/paths/~1array-not-nullable/post/requestBody/content/application~1json/schema",
				BodyRequired:     true,
				KindValidation:   validation.KindValidation{Type: "array", Nullable: false},
				EnumValidation:   validation.EnumValidation{},
				NumberValidation: validation.NumberValidation{},
				StringValidation: validation.StringValidation{Format: ""},
				ArrayValidation: validation.ArrayValidation{
					Items: &validation.Validation{
						SchemaPointer:    "#/paths/~1array-not-nullable/post/requestBody/content/application~1json/schema/items",
						BodyRequired:     false,
						KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
						EnumValidation:   validation.EnumValidation{},
						NumberValidation: validation.NumberValidation{},
						StringValidation: validation.StringValidation{Format: ""},
						ArrayValidation: validation.ArrayValidation{
							UniqueItems: false,
						},
						ObjectValidation: validation.ObjectValidation{
							AdditionalPropertiesAllowed: true,
						},
					},
					UniqueItems: false,
				},
				ObjectValidation: validation.ObjectValidation{
					AdditionalPropertiesAllowed: true,
				},
			},
		},
		"arrayNullable": {
			expectedValidation: &validation.Validation{
				SchemaPointer:    "#/paths/~1array-nullable/post/requestBody/content/application~1json/schema",
				BodyRequired:     true,
				KindValidation:   validation.KindValidation{Type: "array", Nullable: true},
				EnumValidation:   validation.EnumValidation{},
				NumberValidation: validation.NumberValidation{},
				StringValidation: validation.StringValidation{Format: ""},
				ArrayValidation: validation.ArrayValidation{
					Items: &validation.Validation{
						SchemaPointer:    "#/paths/~1array-nullable/post/requestBody/content/application~1json/schema/items",
						BodyRequired:     false,
						KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
						EnumValidation:   validation.EnumValidation{},
						NumberValidation: validation.NumberValidation{},
						StringValidation: validation.StringValidation{Format: ""},
						ArrayValidation: validation.ArrayValidation{
							UniqueItems: false,
						},
						ObjectValidation: validation.ObjectValidation{
							AdditionalPropertiesAllowed: true,
						},
					},
					UniqueItems: false,
				},
				ObjectValidation: validation.ObjectValidation{
					AdditionalPropertiesAllowed: true,
				},
			},
		},
		"compositeObject": {
			expectedValidation: &validation.Validation{
				SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema",
				BodyRequired:     true,
				KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
				EnumValidation:   validation.EnumValidation{},
				NumberValidation: validation.NumberValidation{},
				StringValidation: validation.StringValidation{Format: ""},
				ArrayValidation: validation.ArrayValidation{
					UniqueItems: false,
				},
				ObjectValidation: validation.ObjectValidation{
					Required: []string{"arrayNotNullableItemsNotNullable", "arrayNotNullableItemsNullable", "arrayNullableItemsNotNullable", "arrayNullableItemsNullable", "boolNotNullable", "boolNullable", "numberNotNullable", "numberNullable", "objectAdditionalPropertiesImplicit", "objectAdditionalPropertiesSchema", "objectAdditionalPropertiesTrue", "stringFormatNotNullable", "stringFormatNullable"},
					Properties: []validation.PropertyValidation{
						{Name: "arrayNotNullableItemsNotNullable", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/arrayNotNullableItemsNotNullable",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "array", Nullable: false},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								Items: &validation.Validation{
									SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/arrayNotNullableItemsNotNullable/items",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								},
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
						{Name: "arrayNotNullableItemsNullable", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/arrayNotNullableItemsNullable",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "array", Nullable: false},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								Items: &validation.Validation{
									SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/arrayNotNullableItemsNullable/items",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								},
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
						{Name: "arrayNullableItemsNotNullable", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/arrayNullableItemsNotNullable",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "array", Nullable: true},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								Items: &validation.Validation{
									SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/arrayNullableItemsNotNullable/items",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								},
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
						{Name: "arrayNullableItemsNullable", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/arrayNullableItemsNullable",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "array", Nullable: true},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								Items: &validation.Validation{
									SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/arrayNullableItemsNullable/items",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								},
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
						{Name: "boolNotNullable", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/boolNotNullable",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "boolean", Nullable: false},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
						{Name: "boolNullable", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/boolNullable",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "boolean", Nullable: true},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
						{Name: "numberNotNullable", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/numberNotNullable",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "number", Nullable: false},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
						{Name: "numberNullable", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/numberNullable",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "number", Nullable: true},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
						{Name: "objectAdditionalPropertiesImplicit", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/objectAdditionalPropertiesImplicit",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								Properties: []validation.PropertyValidation{
									{Name: "known", Validation: &validation.Validation{
										SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/objectAdditionalPropertiesImplicit/properties/known",
										BodyRequired:     false,
										KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
										EnumValidation:   validation.EnumValidation{},
										NumberValidation: validation.NumberValidation{},
										StringValidation: validation.StringValidation{Format: ""},
										ArrayValidation: validation.ArrayValidation{
											UniqueItems: false,
										},
										ObjectValidation: validation.ObjectValidation{
											AdditionalPropertiesAllowed: true,
										},
									}},
								},
								AdditionalPropertiesAllowed: true,
							},
						}},
						{Name: "objectAdditionalPropertiesSchema", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/objectAdditionalPropertiesSchema",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								Properties: []validation.PropertyValidation{
									{Name: "known", Validation: &validation.Validation{
										SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/objectAdditionalPropertiesSchema/properties/known",
										BodyRequired:     false,
										KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
										EnumValidation:   validation.EnumValidation{},
										NumberValidation: validation.NumberValidation{},
										StringValidation: validation.StringValidation{Format: ""},
										ArrayValidation: validation.ArrayValidation{
											UniqueItems: false,
										},
										ObjectValidation: validation.ObjectValidation{
											AdditionalPropertiesAllowed: true,
										},
									}},
								},
								AdditionalPropertiesAllowed: true,
								AdditionalPropertiesValidation: &validation.Validation{
									SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/objectAdditionalPropertiesSchema/additionalProperties",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								},
							},
						}},
						{Name: "objectAdditionalPropertiesTrue", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/objectAdditionalPropertiesTrue",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								Properties: []validation.PropertyValidation{
									{Name: "known", Validation: &validation.Validation{
										SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/objectAdditionalPropertiesTrue/properties/known",
										BodyRequired:     false,
										KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
										EnumValidation:   validation.EnumValidation{},
										NumberValidation: validation.NumberValidation{},
										StringValidation: validation.StringValidation{Format: ""},
										ArrayValidation: validation.ArrayValidation{
											UniqueItems: false,
										},
										ObjectValidation: validation.ObjectValidation{
											AdditionalPropertiesAllowed: true,
										},
									}},
								},
								AdditionalPropertiesAllowed: true,
							},
						}},
						{Name: "stringFormatNotNullable", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/stringFormatNotNullable",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: "date-time"},
							ArrayValidation: validation.ArrayValidation{
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
						{Name: "stringFormatNullable", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1composite-object/post/requestBody/content/application~1json/schema/properties/stringFormatNullable",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: "date-time"},
							ArrayValidation: validation.ArrayValidation{
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
					},
					AdditionalPropertiesAllowed: false,
				},
			},
		},
		"nullableObjectKeysAdditionalPropertiesFalse": {
			expectedValidation: &validation.Validation{
				SchemaPointer:    "#/paths/~1nullable-object-keys-additional-properties-false/post/requestBody/content/application~1json/schema",
				BodyRequired:     true,
				KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
				EnumValidation:   validation.EnumValidation{},
				NumberValidation: validation.NumberValidation{},
				StringValidation: validation.StringValidation{Format: ""},
				ArrayValidation: validation.ArrayValidation{
					UniqueItems: false,
				},
				ObjectValidation: validation.ObjectValidation{
					Required: []string{"requiredNotNullableString", "requiredNullableString"},
					Properties: []validation.PropertyValidation{
						{Name: "optionalNotNullableString", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1nullable-object-keys-additional-properties-false/post/requestBody/content/application~1json/schema/properties/optionalNotNullableString",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
						{Name: "optionalNullableString", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1nullable-object-keys-additional-properties-false/post/requestBody/content/application~1json/schema/properties/optionalNullableString",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
						{Name: "requiredNotNullableString", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1nullable-object-keys-additional-properties-false/post/requestBody/content/application~1json/schema/properties/requiredNotNullableString",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
						{Name: "requiredNullableString", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1nullable-object-keys-additional-properties-false/post/requestBody/content/application~1json/schema/properties/requiredNullableString",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
					},
					AdditionalPropertiesAllowed: false,
				},
			},
		},
		"objectKeysAdditionalPropertiesFalse": {
			expectedValidation: &validation.Validation{
				SchemaPointer:    "#/paths/~1object-keys-additional-properties-false/post/requestBody/content/application~1json/schema",
				BodyRequired:     true,
				KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
				EnumValidation:   validation.EnumValidation{},
				NumberValidation: validation.NumberValidation{},
				StringValidation: validation.StringValidation{Format: ""},
				ArrayValidation: validation.ArrayValidation{
					UniqueItems: false,
				},
				ObjectValidation: validation.ObjectValidation{
					Required: []string{"requiredNotNullableString", "requiredNullableString"},
					Properties: []validation.PropertyValidation{
						{Name: "optionalNotNullableString", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1object-keys-additional-properties-false/post/requestBody/content/application~1json/schema/properties/optionalNotNullableString",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
						{Name: "optionalNullableString", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1object-keys-additional-properties-false/post/requestBody/content/application~1json/schema/properties/optionalNullableString",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
						{Name: "requiredNotNullableString", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1object-keys-additional-properties-false/post/requestBody/content/application~1json/schema/properties/requiredNotNullableString",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
						{Name: "requiredNullableString", Validation: &validation.Validation{
							SchemaPointer:    "#/paths/~1object-keys-additional-properties-false/post/requestBody/content/application~1json/schema/properties/requiredNullableString",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
					},
					AdditionalPropertiesAllowed: false,
				},
			},
		},
		"optionalArrayNullable": {
			expectedValidation: &validation.Validation{
				SchemaPointer:    "#/paths/~1optional-array-nullable/post/requestBody/content/application~1json/schema",
				BodyRequired:     false,
				KindValidation:   validation.KindValidation{Type: "array", Nullable: true},
				EnumValidation:   validation.EnumValidation{},
				NumberValidation: validation.NumberValidation{},
				StringValidation: validation.StringValidation{Format: ""},
				ArrayValidation: validation.ArrayValidation{
					Items: &validation.Validation{
						SchemaPointer:    "#/paths/~1optional-array-nullable/post/requestBody/content/application~1json/schema/items",
						BodyRequired:     false,
						KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
						EnumValidation:   validation.EnumValidation{},
						NumberValidation: validation.NumberValidation{},
						StringValidation: validation.StringValidation{Format: ""},
						ArrayValidation: validation.ArrayValidation{
							UniqueItems: false,
						},
						ObjectValidation: validation.ObjectValidation{
							AdditionalPropertiesAllowed: true,
						},
					},
					UniqueItems: false,
				},
				ObjectValidation: validation.ObjectValidation{
					AdditionalPropertiesAllowed: true,
				},
			},
		},
		"refObject": {
			expectedValidation: &validation.Validation{
				SchemaPointer:    "#/components/schemas/RefObjectRequest",
				BodyRequired:     true,
				KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
				EnumValidation:   validation.EnumValidation{},
				NumberValidation: validation.NumberValidation{},
				StringValidation: validation.StringValidation{Format: ""},
				ArrayValidation: validation.ArrayValidation{
					UniqueItems: false,
				},
				ObjectValidation: validation.ObjectValidation{
					Required: []string{"refRequiredString"},
					Properties: []validation.PropertyValidation{
						{Name: "refOptionalBool", Validation: &validation.Validation{
							SchemaPointer:    "#/components/schemas/RefObjectRequest/properties/refOptionalBool",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "boolean", Nullable: true},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
						{Name: "refRequiredString", Validation: &validation.Validation{
							SchemaPointer:    "#/components/schemas/RefObjectRequest/properties/refRequiredString",
							BodyRequired:     false,
							KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
							EnumValidation:   validation.EnumValidation{},
							NumberValidation: validation.NumberValidation{},
							StringValidation: validation.StringValidation{Format: ""},
							ArrayValidation: validation.ArrayValidation{
								UniqueItems: false,
							},
							ObjectValidation: validation.ObjectValidation{
								AdditionalPropertiesAllowed: true,
							},
						}},
					},
					AdditionalPropertiesAllowed: false,
				},
			},
		},
		"refStressObject": {
			expectedValidation: &validation.Validation{
				SchemaPointer:    "#/paths/~1ref-stress-object/post/requestBody/content/application~1json/schema",
				BodyRequired:     true,
				KindValidation:   validation.KindValidation{Type: "", Nullable: false},
				EnumValidation:   validation.EnumValidation{},
				NumberValidation: validation.NumberValidation{},
				StringValidation: validation.StringValidation{Format: ""},
				ArrayValidation: validation.ArrayValidation{
					UniqueItems: false,
				},
				ObjectValidation: validation.ObjectValidation{
					AdditionalPropertiesAllowed: true,
				},
				AllOfValidations: []*validation.Validation{
					{
						SchemaPointer:    "#/components/schemas/RefStressFirstAllOf",
						BodyRequired:     false,
						KindValidation:   validation.KindValidation{Type: "", Nullable: false},
						EnumValidation:   validation.EnumValidation{},
						NumberValidation: validation.NumberValidation{},
						StringValidation: validation.StringValidation{Format: ""},
						ArrayValidation: validation.ArrayValidation{
							UniqueItems: false,
						},
						ObjectValidation: validation.ObjectValidation{
							AdditionalPropertiesAllowed: true,
						},
						AllOfValidations: []*validation.Validation{
							{
								SchemaPointer:    "#/components/schemas/RefStressFinal",
								BodyRequired:     false,
								KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
								EnumValidation:   validation.EnumValidation{},
								NumberValidation: validation.NumberValidation{},
								StringValidation: validation.StringValidation{Format: ""},
								ArrayValidation: validation.ArrayValidation{
									UniqueItems: false,
								},
								ObjectValidation: validation.ObjectValidation{
									Required: []string{"finalCode", "sharedName"},
									Properties: []validation.PropertyValidation{
										{Name: "finalCode", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressFinal/properties/finalCode",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
										{Name: "nested", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressNestedBase",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												Required: []string{"sameName"},
												Properties: []validation.PropertyValidation{
													{Name: "leaf", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "sameName", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
												},
												AdditionalPropertiesAllowed: true,
											},
										}},
										{Name: "optionalShared", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressFinal/properties/optionalShared",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
										{Name: "sharedName", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressFinal/properties/sharedName",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
									},
									AdditionalPropertiesAllowed: true,
								},
							},
							{
								SchemaPointer:    "#/components/schemas/RefStressViaMiddle",
								BodyRequired:     false,
								KindValidation:   validation.KindValidation{Type: "", Nullable: false},
								EnumValidation:   validation.EnumValidation{},
								NumberValidation: validation.NumberValidation{},
								StringValidation: validation.StringValidation{Format: ""},
								ArrayValidation: validation.ArrayValidation{
									UniqueItems: false,
								},
								ObjectValidation: validation.ObjectValidation{
									AdditionalPropertiesAllowed: true,
								},
								AllOfValidations: []*validation.Validation{
									{
										SchemaPointer:    "#/components/schemas/RefStressMiddleAllOf",
										BodyRequired:     false,
										KindValidation:   validation.KindValidation{Type: "", Nullable: false},
										EnumValidation:   validation.EnumValidation{},
										NumberValidation: validation.NumberValidation{},
										StringValidation: validation.StringValidation{Format: ""},
										ArrayValidation: validation.ArrayValidation{
											UniqueItems: false,
										},
										ObjectValidation: validation.ObjectValidation{
											AdditionalPropertiesAllowed: true,
										},
										AllOfValidations: []*validation.Validation{
											{
												SchemaPointer:    "#/components/schemas/RefStressFinal",
												BodyRequired:     false,
												KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
												EnumValidation:   validation.EnumValidation{},
												NumberValidation: validation.NumberValidation{},
												StringValidation: validation.StringValidation{Format: ""},
												ArrayValidation: validation.ArrayValidation{
													UniqueItems: false,
												},
												ObjectValidation: validation.ObjectValidation{
													Required: []string{"finalCode", "sharedName"},
													Properties: []validation.PropertyValidation{
														{Name: "finalCode", Validation: &validation.Validation{
															SchemaPointer:    "#/components/schemas/RefStressFinal/properties/finalCode",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																AdditionalPropertiesAllowed: true,
															},
														}},
														{Name: "nested", Validation: &validation.Validation{
															SchemaPointer:    "#/components/schemas/RefStressNestedBase",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																Required: []string{"sameName"},
																Properties: []validation.PropertyValidation{
																	{Name: "leaf", Validation: &validation.Validation{
																		SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																		BodyRequired:     false,
																		KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																		EnumValidation:   validation.EnumValidation{},
																		NumberValidation: validation.NumberValidation{},
																		StringValidation: validation.StringValidation{Format: ""},
																		ArrayValidation: validation.ArrayValidation{
																			UniqueItems: false,
																		},
																		ObjectValidation: validation.ObjectValidation{
																			AdditionalPropertiesAllowed: true,
																		},
																	}},
																	{Name: "sameName", Validation: &validation.Validation{
																		SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
																		BodyRequired:     false,
																		KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																		EnumValidation:   validation.EnumValidation{},
																		NumberValidation: validation.NumberValidation{},
																		StringValidation: validation.StringValidation{Format: ""},
																		ArrayValidation: validation.ArrayValidation{
																			UniqueItems: false,
																		},
																		ObjectValidation: validation.ObjectValidation{
																			AdditionalPropertiesAllowed: true,
																		},
																	}},
																},
																AdditionalPropertiesAllowed: true,
															},
														}},
														{Name: "optionalShared", Validation: &validation.Validation{
															SchemaPointer:    "#/components/schemas/RefStressFinal/properties/optionalShared",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																AdditionalPropertiesAllowed: true,
															},
														}},
														{Name: "sharedName", Validation: &validation.Validation{
															SchemaPointer:    "#/components/schemas/RefStressFinal/properties/sharedName",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																AdditionalPropertiesAllowed: true,
															},
														}},
													},
													AdditionalPropertiesAllowed: true,
												},
											},
											{
												SchemaPointer:    "#/components/schemas/RefStressMiddleAllOf/allOf/1",
												BodyRequired:     false,
												KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
												EnumValidation:   validation.EnumValidation{},
												NumberValidation: validation.NumberValidation{},
												StringValidation: validation.StringValidation{Format: ""},
												ArrayValidation: validation.ArrayValidation{
													UniqueItems: false,
												},
												ObjectValidation: validation.ObjectValidation{
													Required: []string{"sharedName"},
													Properties: []validation.PropertyValidation{
														{Name: "optionalCode", Validation: &validation.Validation{
															SchemaPointer:    "#/components/schemas/RefStressMiddleAllOf/allOf/1/properties/optionalCode",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																AdditionalPropertiesAllowed: true,
															},
														}},
														{Name: "sharedName", Validation: &validation.Validation{
															SchemaPointer:    "#/components/schemas/RefStressMiddleAllOf/allOf/1/properties/sharedName",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																AdditionalPropertiesAllowed: true,
															},
														}},
													},
													AdditionalPropertiesAllowed: true,
												},
											},
										},
									},
									{
										SchemaPointer:    "#/components/schemas/RefStressViaMiddle/allOf/1",
										BodyRequired:     false,
										KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
										EnumValidation:   validation.EnumValidation{},
										NumberValidation: validation.NumberValidation{},
										StringValidation: validation.StringValidation{Format: ""},
										ArrayValidation: validation.ArrayValidation{
											UniqueItems: false,
										},
										ObjectValidation: validation.ObjectValidation{
											Required: []string{"middleFlag", "sharedName"},
											Properties: []validation.PropertyValidation{
												{Name: "middleFlag", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressViaMiddle/allOf/1/properties/middleFlag",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "boolean", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														AdditionalPropertiesAllowed: true,
													},
												}},
												{Name: "nested", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressNestedCombined",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														AdditionalPropertiesAllowed: true,
													},
													AllOfValidations: []*validation.Validation{
														{
															SchemaPointer:    "#/components/schemas/RefStressNestedBase",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																Required: []string{"sameName"},
																Properties: []validation.PropertyValidation{
																	{Name: "leaf", Validation: &validation.Validation{
																		SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																		BodyRequired:     false,
																		KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																		EnumValidation:   validation.EnumValidation{},
																		NumberValidation: validation.NumberValidation{},
																		StringValidation: validation.StringValidation{Format: ""},
																		ArrayValidation: validation.ArrayValidation{
																			UniqueItems: false,
																		},
																		ObjectValidation: validation.ObjectValidation{
																			AdditionalPropertiesAllowed: true,
																		},
																	}},
																	{Name: "sameName", Validation: &validation.Validation{
																		SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
																		BodyRequired:     false,
																		KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																		EnumValidation:   validation.EnumValidation{},
																		NumberValidation: validation.NumberValidation{},
																		StringValidation: validation.StringValidation{Format: ""},
																		ArrayValidation: validation.ArrayValidation{
																			UniqueItems: false,
																		},
																		ObjectValidation: validation.ObjectValidation{
																			AdditionalPropertiesAllowed: true,
																		},
																	}},
																},
																AdditionalPropertiesAllowed: true,
															},
														},
														{
															SchemaPointer:    "#/components/schemas/RefStressNestedOverlay",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																Required: []string{"sameName"},
																Properties: []validation.PropertyValidation{
																	{Name: "leaf", Validation: &validation.Validation{
																		SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																		BodyRequired:     false,
																		KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																		EnumValidation:   validation.EnumValidation{},
																		NumberValidation: validation.NumberValidation{},
																		StringValidation: validation.StringValidation{Format: ""},
																		ArrayValidation: validation.ArrayValidation{
																			UniqueItems: false,
																		},
																		ObjectValidation: validation.ObjectValidation{
																			AdditionalPropertiesAllowed: true,
																		},
																	}},
																	{Name: "sameName", Validation: &validation.Validation{
																		SchemaPointer:    "#/components/schemas/RefStressNestedOverlay/properties/sameName",
																		BodyRequired:     false,
																		KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																		EnumValidation:   validation.EnumValidation{},
																		NumberValidation: validation.NumberValidation{},
																		StringValidation: validation.StringValidation{Format: ""},
																		ArrayValidation: validation.ArrayValidation{
																			UniqueItems: false,
																		},
																		ObjectValidation: validation.ObjectValidation{
																			AdditionalPropertiesAllowed: true,
																		},
																	}},
																},
																AdditionalPropertiesAllowed: true,
															},
														},
														{
															SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																Required: []string{"sameName", "sealed"},
																Properties: []validation.PropertyValidation{
																	{Name: "sameName", Validation: &validation.Validation{
																		SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2/properties/sameName",
																		BodyRequired:     false,
																		KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																		EnumValidation:   validation.EnumValidation{},
																		NumberValidation: validation.NumberValidation{},
																		StringValidation: validation.StringValidation{Format: ""},
																		ArrayValidation: validation.ArrayValidation{
																			UniqueItems: false,
																		},
																		ObjectValidation: validation.ObjectValidation{
																			AdditionalPropertiesAllowed: true,
																		},
																	}},
																	{Name: "sealed", Validation: &validation.Validation{
																		SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2/properties/sealed",
																		BodyRequired:     false,
																		KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
																		EnumValidation:   validation.EnumValidation{},
																		NumberValidation: validation.NumberValidation{},
																		StringValidation: validation.StringValidation{Format: ""},
																		ArrayValidation: validation.ArrayValidation{
																			UniqueItems: false,
																		},
																		ObjectValidation: validation.ObjectValidation{
																			Required: []string{"locked"},
																			Properties: []validation.PropertyValidation{
																				{Name: "locked", Validation: &validation.Validation{
																					SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2/properties/sealed/properties/locked",
																					BodyRequired:     false,
																					KindValidation:   validation.KindValidation{Type: "boolean", Nullable: false},
																					EnumValidation:   validation.EnumValidation{},
																					NumberValidation: validation.NumberValidation{},
																					StringValidation: validation.StringValidation{Format: ""},
																					ArrayValidation: validation.ArrayValidation{
																						UniqueItems: false,
																					},
																					ObjectValidation: validation.ObjectValidation{
																						AdditionalPropertiesAllowed: true,
																					},
																				}},
																			},
																			AdditionalPropertiesAllowed: false,
																		},
																	}},
																},
																AdditionalPropertiesAllowed: true,
															},
														},
													},
												}},
												{Name: "sharedName", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressViaMiddle/allOf/1/properties/sharedName",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														AdditionalPropertiesAllowed: true,
													},
												}},
											},
											AdditionalPropertiesAllowed: true,
										},
									},
								},
							},
							{
								SchemaPointer:    "#/components/schemas/RefStressFirstAllOf/allOf/2",
								BodyRequired:     false,
								KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
								EnumValidation:   validation.EnumValidation{},
								NumberValidation: validation.NumberValidation{},
								StringValidation: validation.StringValidation{Format: ""},
								ArrayValidation: validation.ArrayValidation{
									UniqueItems: false,
								},
								ObjectValidation: validation.ObjectValidation{
									Required: []string{"final", "nested", "nullableRequired"},
									Properties: []validation.PropertyValidation{
										{Name: "final", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressFinal",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												Required: []string{"finalCode", "sharedName"},
												Properties: []validation.PropertyValidation{
													{Name: "finalCode", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressFinal/properties/finalCode",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "nested", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressNestedBase",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															Required: []string{"sameName"},
															Properties: []validation.PropertyValidation{
																{Name: "leaf", Validation: &validation.Validation{
																	SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																	BodyRequired:     false,
																	KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																	EnumValidation:   validation.EnumValidation{},
																	NumberValidation: validation.NumberValidation{},
																	StringValidation: validation.StringValidation{Format: ""},
																	ArrayValidation: validation.ArrayValidation{
																		UniqueItems: false,
																	},
																	ObjectValidation: validation.ObjectValidation{
																		AdditionalPropertiesAllowed: true,
																	},
																}},
																{Name: "sameName", Validation: &validation.Validation{
																	SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
																	BodyRequired:     false,
																	KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																	EnumValidation:   validation.EnumValidation{},
																	NumberValidation: validation.NumberValidation{},
																	StringValidation: validation.StringValidation{Format: ""},
																	ArrayValidation: validation.ArrayValidation{
																		UniqueItems: false,
																	},
																	ObjectValidation: validation.ObjectValidation{
																		AdditionalPropertiesAllowed: true,
																	},
																}},
															},
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "optionalShared", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressFinal/properties/optionalShared",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "sharedName", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressFinal/properties/sharedName",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
												},
												AdditionalPropertiesAllowed: true,
											},
										}},
										{Name: "nested", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressNestedCombined",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
											AllOfValidations: []*validation.Validation{
												{
													SchemaPointer:    "#/components/schemas/RefStressNestedBase",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														Required: []string{"sameName"},
														Properties: []validation.PropertyValidation{
															{Name: "leaf", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "sameName", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
														},
														AdditionalPropertiesAllowed: true,
													},
												},
												{
													SchemaPointer:    "#/components/schemas/RefStressNestedOverlay",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														Required: []string{"sameName"},
														Properties: []validation.PropertyValidation{
															{Name: "leaf", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "sameName", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressNestedOverlay/properties/sameName",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
														},
														AdditionalPropertiesAllowed: true,
													},
												},
												{
													SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														Required: []string{"sameName", "sealed"},
														Properties: []validation.PropertyValidation{
															{Name: "sameName", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2/properties/sameName",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "sealed", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2/properties/sealed",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	Required: []string{"locked"},
																	Properties: []validation.PropertyValidation{
																		{Name: "locked", Validation: &validation.Validation{
																			SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2/properties/sealed/properties/locked",
																			BodyRequired:     false,
																			KindValidation:   validation.KindValidation{Type: "boolean", Nullable: false},
																			EnumValidation:   validation.EnumValidation{},
																			NumberValidation: validation.NumberValidation{},
																			StringValidation: validation.StringValidation{Format: ""},
																			ArrayValidation: validation.ArrayValidation{
																				UniqueItems: false,
																			},
																			ObjectValidation: validation.ObjectValidation{
																				AdditionalPropertiesAllowed: true,
																			},
																		}},
																	},
																	AdditionalPropertiesAllowed: false,
																},
															}},
														},
														AdditionalPropertiesAllowed: true,
													},
												},
											},
										}},
										{Name: "nullableRequired", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressFirstAllOf/allOf/2/properties/nullableRequired",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
										{Name: "optionalShared", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressFirstAllOf/allOf/2/properties/optionalShared",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
										{Name: "sharedName", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressFirstAllOf/allOf/2/properties/sharedName",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
									},
									AdditionalPropertiesAllowed: true,
								},
							},
						},
					},
					{
						SchemaPointer:    "#/components/schemas/RefStressSecondAllOf",
						BodyRequired:     false,
						KindValidation:   validation.KindValidation{Type: "", Nullable: false},
						EnumValidation:   validation.EnumValidation{},
						NumberValidation: validation.NumberValidation{},
						StringValidation: validation.StringValidation{Format: ""},
						ArrayValidation: validation.ArrayValidation{
							UniqueItems: false,
						},
						ObjectValidation: validation.ObjectValidation{
							AdditionalPropertiesAllowed: true,
						},
						AllOfValidations: []*validation.Validation{
							{
								SchemaPointer:    "#/components/schemas/RefStressOtherMiddle",
								BodyRequired:     false,
								KindValidation:   validation.KindValidation{Type: "", Nullable: false},
								EnumValidation:   validation.EnumValidation{},
								NumberValidation: validation.NumberValidation{},
								StringValidation: validation.StringValidation{Format: ""},
								ArrayValidation: validation.ArrayValidation{
									UniqueItems: false,
								},
								ObjectValidation: validation.ObjectValidation{
									AdditionalPropertiesAllowed: true,
								},
								AllOfValidations: []*validation.Validation{
									{
										SchemaPointer:    "#/components/schemas/RefStressFinal",
										BodyRequired:     false,
										KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
										EnumValidation:   validation.EnumValidation{},
										NumberValidation: validation.NumberValidation{},
										StringValidation: validation.StringValidation{Format: ""},
										ArrayValidation: validation.ArrayValidation{
											UniqueItems: false,
										},
										ObjectValidation: validation.ObjectValidation{
											Required: []string{"finalCode", "sharedName"},
											Properties: []validation.PropertyValidation{
												{Name: "finalCode", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressFinal/properties/finalCode",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														AdditionalPropertiesAllowed: true,
													},
												}},
												{Name: "nested", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressNestedBase",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														Required: []string{"sameName"},
														Properties: []validation.PropertyValidation{
															{Name: "leaf", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "sameName", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
														},
														AdditionalPropertiesAllowed: true,
													},
												}},
												{Name: "optionalShared", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressFinal/properties/optionalShared",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														AdditionalPropertiesAllowed: true,
													},
												}},
												{Name: "sharedName", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressFinal/properties/sharedName",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														AdditionalPropertiesAllowed: true,
													},
												}},
											},
											AdditionalPropertiesAllowed: true,
										},
									},
									{
										SchemaPointer:    "#/components/schemas/RefStressOtherMiddle/allOf/1",
										BodyRequired:     false,
										KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
										EnumValidation:   validation.EnumValidation{},
										NumberValidation: validation.NumberValidation{},
										StringValidation: validation.StringValidation{Format: ""},
										ArrayValidation: validation.ArrayValidation{
											UniqueItems: false,
										},
										ObjectValidation: validation.ObjectValidation{
											Required: []string{"metadata", "rootFlag"},
											Properties: []validation.PropertyValidation{
												{Name: "final", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressFinal",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														Required: []string{"finalCode", "sharedName"},
														Properties: []validation.PropertyValidation{
															{Name: "finalCode", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressFinal/properties/finalCode",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "nested", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressNestedBase",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	Required: []string{"sameName"},
																	Properties: []validation.PropertyValidation{
																		{Name: "leaf", Validation: &validation.Validation{
																			SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																			BodyRequired:     false,
																			KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																			EnumValidation:   validation.EnumValidation{},
																			NumberValidation: validation.NumberValidation{},
																			StringValidation: validation.StringValidation{Format: ""},
																			ArrayValidation: validation.ArrayValidation{
																				UniqueItems: false,
																			},
																			ObjectValidation: validation.ObjectValidation{
																				AdditionalPropertiesAllowed: true,
																			},
																		}},
																		{Name: "sameName", Validation: &validation.Validation{
																			SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
																			BodyRequired:     false,
																			KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																			EnumValidation:   validation.EnumValidation{},
																			NumberValidation: validation.NumberValidation{},
																			StringValidation: validation.StringValidation{Format: ""},
																			ArrayValidation: validation.ArrayValidation{
																				UniqueItems: false,
																			},
																			ObjectValidation: validation.ObjectValidation{
																				AdditionalPropertiesAllowed: true,
																			},
																		}},
																	},
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "optionalShared", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressFinal/properties/optionalShared",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "sharedName", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressFinal/properties/sharedName",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
														},
														AdditionalPropertiesAllowed: true,
													},
												}},
												{Name: "metadata", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressOtherMiddle/allOf/1/properties/metadata",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														AdditionalPropertiesAllowed: true,
														AdditionalPropertiesValidation: &validation.Validation{
															SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																AdditionalPropertiesAllowed: true,
															},
														},
													},
												}},
												{Name: "rootFlag", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressOtherMiddle/allOf/1/properties/rootFlag",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "boolean", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														AdditionalPropertiesAllowed: true,
													},
												}},
											},
											AdditionalPropertiesAllowed: true,
										},
									},
								},
							},
							{
								SchemaPointer:    "#/components/schemas/RefStressSecondAllOf/allOf/1",
								BodyRequired:     false,
								KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
								EnumValidation:   validation.EnumValidation{},
								NumberValidation: validation.NumberValidation{},
								StringValidation: validation.StringValidation{Format: ""},
								ArrayValidation: validation.ArrayValidation{
									UniqueItems: false,
								},
								ObjectValidation: validation.ObjectValidation{
									Required: []string{"count", "finals", "metadata", "rootFlag"},
									Properties: []validation.PropertyValidation{
										{Name: "count", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressSecondAllOf/allOf/1/properties/count",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "number", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
										{Name: "finals", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressSecondAllOf/allOf/1/properties/finals",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "array", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												Items: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressFinal",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														Required: []string{"finalCode", "sharedName"},
														Properties: []validation.PropertyValidation{
															{Name: "finalCode", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressFinal/properties/finalCode",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "nested", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressNestedBase",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	Required: []string{"sameName"},
																	Properties: []validation.PropertyValidation{
																		{Name: "leaf", Validation: &validation.Validation{
																			SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																			BodyRequired:     false,
																			KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																			EnumValidation:   validation.EnumValidation{},
																			NumberValidation: validation.NumberValidation{},
																			StringValidation: validation.StringValidation{Format: ""},
																			ArrayValidation: validation.ArrayValidation{
																				UniqueItems: false,
																			},
																			ObjectValidation: validation.ObjectValidation{
																				AdditionalPropertiesAllowed: true,
																			},
																		}},
																		{Name: "sameName", Validation: &validation.Validation{
																			SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
																			BodyRequired:     false,
																			KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																			EnumValidation:   validation.EnumValidation{},
																			NumberValidation: validation.NumberValidation{},
																			StringValidation: validation.StringValidation{Format: ""},
																			ArrayValidation: validation.ArrayValidation{
																				UniqueItems: false,
																			},
																			ObjectValidation: validation.ObjectValidation{
																				AdditionalPropertiesAllowed: true,
																			},
																		}},
																	},
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "optionalShared", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressFinal/properties/optionalShared",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "sharedName", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressFinal/properties/sharedName",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
														},
														AdditionalPropertiesAllowed: true,
													},
												},
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
										{Name: "metadata", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressSecondAllOf/allOf/1/properties/metadata",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
												AdditionalPropertiesValidation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														AdditionalPropertiesAllowed: true,
													},
												},
											},
										}},
										{Name: "rootFlag", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressSecondAllOf/allOf/1/properties/rootFlag",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "boolean", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
										{Name: "sharedName", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressSecondAllOf/allOf/1/properties/sharedName",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
									},
									AdditionalPropertiesAllowed: true,
								},
							},
						},
					},
					{
						SchemaPointer:    "#/paths/~1ref-stress-object/post/requestBody/content/application~1json/schema/allOf/2",
						BodyRequired:     false,
						KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
						EnumValidation:   validation.EnumValidation{},
						NumberValidation: validation.NumberValidation{},
						StringValidation: validation.StringValidation{Format: ""},
						ArrayValidation: validation.ArrayValidation{
							UniqueItems: false,
						},
						ObjectValidation: validation.ObjectValidation{
							Required: []string{"count", "final", "finalCode", "finals", "metadata", "middleFlag", "nested", "nullableRequired", "rootFlag", "sharedName"},
							Properties: []validation.PropertyValidation{
								{Name: "count", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object/post/requestBody/content/application~1json/schema/allOf/2/properties/count",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "number", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
								{Name: "final", Validation: &validation.Validation{
									SchemaPointer:    "#/components/schemas/RefStressFinal",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										Required: []string{"finalCode", "sharedName"},
										Properties: []validation.PropertyValidation{
											{Name: "finalCode", Validation: &validation.Validation{
												SchemaPointer:    "#/components/schemas/RefStressFinal/properties/finalCode",
												BodyRequired:     false,
												KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
												EnumValidation:   validation.EnumValidation{},
												NumberValidation: validation.NumberValidation{},
												StringValidation: validation.StringValidation{Format: ""},
												ArrayValidation: validation.ArrayValidation{
													UniqueItems: false,
												},
												ObjectValidation: validation.ObjectValidation{
													AdditionalPropertiesAllowed: true,
												},
											}},
											{Name: "nested", Validation: &validation.Validation{
												SchemaPointer:    "#/components/schemas/RefStressNestedBase",
												BodyRequired:     false,
												KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
												EnumValidation:   validation.EnumValidation{},
												NumberValidation: validation.NumberValidation{},
												StringValidation: validation.StringValidation{Format: ""},
												ArrayValidation: validation.ArrayValidation{
													UniqueItems: false,
												},
												ObjectValidation: validation.ObjectValidation{
													Required: []string{"sameName"},
													Properties: []validation.PropertyValidation{
														{Name: "leaf", Validation: &validation.Validation{
															SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																AdditionalPropertiesAllowed: true,
															},
														}},
														{Name: "sameName", Validation: &validation.Validation{
															SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																AdditionalPropertiesAllowed: true,
															},
														}},
													},
													AdditionalPropertiesAllowed: true,
												},
											}},
											{Name: "optionalShared", Validation: &validation.Validation{
												SchemaPointer:    "#/components/schemas/RefStressFinal/properties/optionalShared",
												BodyRequired:     false,
												KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
												EnumValidation:   validation.EnumValidation{},
												NumberValidation: validation.NumberValidation{},
												StringValidation: validation.StringValidation{Format: ""},
												ArrayValidation: validation.ArrayValidation{
													UniqueItems: false,
												},
												ObjectValidation: validation.ObjectValidation{
													AdditionalPropertiesAllowed: true,
												},
											}},
											{Name: "sharedName", Validation: &validation.Validation{
												SchemaPointer:    "#/components/schemas/RefStressFinal/properties/sharedName",
												BodyRequired:     false,
												KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
												EnumValidation:   validation.EnumValidation{},
												NumberValidation: validation.NumberValidation{},
												StringValidation: validation.StringValidation{Format: ""},
												ArrayValidation: validation.ArrayValidation{
													UniqueItems: false,
												},
												ObjectValidation: validation.ObjectValidation{
													AdditionalPropertiesAllowed: true,
												},
											}},
										},
										AdditionalPropertiesAllowed: true,
									},
								}},
								{Name: "finalCode", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object/post/requestBody/content/application~1json/schema/allOf/2/properties/finalCode",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
								{Name: "finals", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object/post/requestBody/content/application~1json/schema/allOf/2/properties/finals",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "array", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										Items: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressFinal",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												Required: []string{"finalCode", "sharedName"},
												Properties: []validation.PropertyValidation{
													{Name: "finalCode", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressFinal/properties/finalCode",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "nested", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressNestedBase",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															Required: []string{"sameName"},
															Properties: []validation.PropertyValidation{
																{Name: "leaf", Validation: &validation.Validation{
																	SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																	BodyRequired:     false,
																	KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																	EnumValidation:   validation.EnumValidation{},
																	NumberValidation: validation.NumberValidation{},
																	StringValidation: validation.StringValidation{Format: ""},
																	ArrayValidation: validation.ArrayValidation{
																		UniqueItems: false,
																	},
																	ObjectValidation: validation.ObjectValidation{
																		AdditionalPropertiesAllowed: true,
																	},
																}},
																{Name: "sameName", Validation: &validation.Validation{
																	SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
																	BodyRequired:     false,
																	KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																	EnumValidation:   validation.EnumValidation{},
																	NumberValidation: validation.NumberValidation{},
																	StringValidation: validation.StringValidation{Format: ""},
																	ArrayValidation: validation.ArrayValidation{
																		UniqueItems: false,
																	},
																	ObjectValidation: validation.ObjectValidation{
																		AdditionalPropertiesAllowed: true,
																	},
																}},
															},
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "optionalShared", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressFinal/properties/optionalShared",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "sharedName", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressFinal/properties/sharedName",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
												},
												AdditionalPropertiesAllowed: true,
											},
										},
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
								{Name: "metadata", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object/post/requestBody/content/application~1json/schema/allOf/2/properties/metadata",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
										AdditionalPropertiesValidation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										},
									},
								}},
								{Name: "middleFlag", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object/post/requestBody/content/application~1json/schema/allOf/2/properties/middleFlag",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "boolean", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
								{Name: "nested", Validation: &validation.Validation{
									SchemaPointer:    "#/components/schemas/RefStressNestedCombined",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
									AllOfValidations: []*validation.Validation{
										{
											SchemaPointer:    "#/components/schemas/RefStressNestedBase",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												Required: []string{"sameName"},
												Properties: []validation.PropertyValidation{
													{Name: "leaf", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "sameName", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
												},
												AdditionalPropertiesAllowed: true,
											},
										},
										{
											SchemaPointer:    "#/components/schemas/RefStressNestedOverlay",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												Required: []string{"sameName"},
												Properties: []validation.PropertyValidation{
													{Name: "leaf", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "sameName", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressNestedOverlay/properties/sameName",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
												},
												AdditionalPropertiesAllowed: true,
											},
										},
										{
											SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												Required: []string{"sameName", "sealed"},
												Properties: []validation.PropertyValidation{
													{Name: "sameName", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2/properties/sameName",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "sealed", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2/properties/sealed",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															Required: []string{"locked"},
															Properties: []validation.PropertyValidation{
																{Name: "locked", Validation: &validation.Validation{
																	SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2/properties/sealed/properties/locked",
																	BodyRequired:     false,
																	KindValidation:   validation.KindValidation{Type: "boolean", Nullable: false},
																	EnumValidation:   validation.EnumValidation{},
																	NumberValidation: validation.NumberValidation{},
																	StringValidation: validation.StringValidation{Format: ""},
																	ArrayValidation: validation.ArrayValidation{
																		UniqueItems: false,
																	},
																	ObjectValidation: validation.ObjectValidation{
																		AdditionalPropertiesAllowed: true,
																	},
																}},
															},
															AdditionalPropertiesAllowed: false,
														},
													}},
												},
												AdditionalPropertiesAllowed: true,
											},
										},
									},
								}},
								{Name: "nullableRequired", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object/post/requestBody/content/application~1json/schema/allOf/2/properties/nullableRequired",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
								{Name: "optionalCode", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object/post/requestBody/content/application~1json/schema/allOf/2/properties/optionalCode",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
								{Name: "optionalShared", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object/post/requestBody/content/application~1json/schema/allOf/2/properties/optionalShared",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
								{Name: "rootFlag", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object/post/requestBody/content/application~1json/schema/allOf/2/properties/rootFlag",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "boolean", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
								{Name: "sharedName", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object/post/requestBody/content/application~1json/schema/allOf/2/properties/sharedName",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
							},
							AdditionalPropertiesAllowed: false,
						},
					},
				},
			},
		},
		"refStressObjectPut": {
			expectedValidation: &validation.Validation{
				SchemaPointer:    "#/paths/~1ref-stress-object-put/put/requestBody/content/application~1json/schema",
				BodyRequired:     true,
				KindValidation:   validation.KindValidation{Type: "", Nullable: false},
				EnumValidation:   validation.EnumValidation{},
				NumberValidation: validation.NumberValidation{},
				StringValidation: validation.StringValidation{Format: ""},
				ArrayValidation: validation.ArrayValidation{
					UniqueItems: false,
				},
				ObjectValidation: validation.ObjectValidation{
					AdditionalPropertiesAllowed: true,
				},
				AllOfValidations: []*validation.Validation{
					{
						SchemaPointer:    "#/components/schemas/RefStressFirstAllOf",
						BodyRequired:     false,
						KindValidation:   validation.KindValidation{Type: "", Nullable: false},
						EnumValidation:   validation.EnumValidation{},
						NumberValidation: validation.NumberValidation{},
						StringValidation: validation.StringValidation{Format: ""},
						ArrayValidation: validation.ArrayValidation{
							UniqueItems: false,
						},
						ObjectValidation: validation.ObjectValidation{
							AdditionalPropertiesAllowed: true,
						},
						AllOfValidations: []*validation.Validation{
							{
								SchemaPointer:    "#/components/schemas/RefStressFinal",
								BodyRequired:     false,
								KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
								EnumValidation:   validation.EnumValidation{},
								NumberValidation: validation.NumberValidation{},
								StringValidation: validation.StringValidation{Format: ""},
								ArrayValidation: validation.ArrayValidation{
									UniqueItems: false,
								},
								ObjectValidation: validation.ObjectValidation{
									Required: []string{"finalCode", "sharedName"},
									Properties: []validation.PropertyValidation{
										{Name: "finalCode", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressFinal/properties/finalCode",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
										{Name: "nested", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressNestedBase",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												Required: []string{"sameName"},
												Properties: []validation.PropertyValidation{
													{Name: "leaf", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "sameName", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
												},
												AdditionalPropertiesAllowed: true,
											},
										}},
										{Name: "optionalShared", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressFinal/properties/optionalShared",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
										{Name: "sharedName", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressFinal/properties/sharedName",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
									},
									AdditionalPropertiesAllowed: true,
								},
							},
							{
								SchemaPointer:    "#/components/schemas/RefStressViaMiddle",
								BodyRequired:     false,
								KindValidation:   validation.KindValidation{Type: "", Nullable: false},
								EnumValidation:   validation.EnumValidation{},
								NumberValidation: validation.NumberValidation{},
								StringValidation: validation.StringValidation{Format: ""},
								ArrayValidation: validation.ArrayValidation{
									UniqueItems: false,
								},
								ObjectValidation: validation.ObjectValidation{
									AdditionalPropertiesAllowed: true,
								},
								AllOfValidations: []*validation.Validation{
									{
										SchemaPointer:    "#/components/schemas/RefStressMiddleAllOf",
										BodyRequired:     false,
										KindValidation:   validation.KindValidation{Type: "", Nullable: false},
										EnumValidation:   validation.EnumValidation{},
										NumberValidation: validation.NumberValidation{},
										StringValidation: validation.StringValidation{Format: ""},
										ArrayValidation: validation.ArrayValidation{
											UniqueItems: false,
										},
										ObjectValidation: validation.ObjectValidation{
											AdditionalPropertiesAllowed: true,
										},
										AllOfValidations: []*validation.Validation{
											{
												SchemaPointer:    "#/components/schemas/RefStressFinal",
												BodyRequired:     false,
												KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
												EnumValidation:   validation.EnumValidation{},
												NumberValidation: validation.NumberValidation{},
												StringValidation: validation.StringValidation{Format: ""},
												ArrayValidation: validation.ArrayValidation{
													UniqueItems: false,
												},
												ObjectValidation: validation.ObjectValidation{
													Required: []string{"finalCode", "sharedName"},
													Properties: []validation.PropertyValidation{
														{Name: "finalCode", Validation: &validation.Validation{
															SchemaPointer:    "#/components/schemas/RefStressFinal/properties/finalCode",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																AdditionalPropertiesAllowed: true,
															},
														}},
														{Name: "nested", Validation: &validation.Validation{
															SchemaPointer:    "#/components/schemas/RefStressNestedBase",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																Required: []string{"sameName"},
																Properties: []validation.PropertyValidation{
																	{Name: "leaf", Validation: &validation.Validation{
																		SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																		BodyRequired:     false,
																		KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																		EnumValidation:   validation.EnumValidation{},
																		NumberValidation: validation.NumberValidation{},
																		StringValidation: validation.StringValidation{Format: ""},
																		ArrayValidation: validation.ArrayValidation{
																			UniqueItems: false,
																		},
																		ObjectValidation: validation.ObjectValidation{
																			AdditionalPropertiesAllowed: true,
																		},
																	}},
																	{Name: "sameName", Validation: &validation.Validation{
																		SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
																		BodyRequired:     false,
																		KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																		EnumValidation:   validation.EnumValidation{},
																		NumberValidation: validation.NumberValidation{},
																		StringValidation: validation.StringValidation{Format: ""},
																		ArrayValidation: validation.ArrayValidation{
																			UniqueItems: false,
																		},
																		ObjectValidation: validation.ObjectValidation{
																			AdditionalPropertiesAllowed: true,
																		},
																	}},
																},
																AdditionalPropertiesAllowed: true,
															},
														}},
														{Name: "optionalShared", Validation: &validation.Validation{
															SchemaPointer:    "#/components/schemas/RefStressFinal/properties/optionalShared",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																AdditionalPropertiesAllowed: true,
															},
														}},
														{Name: "sharedName", Validation: &validation.Validation{
															SchemaPointer:    "#/components/schemas/RefStressFinal/properties/sharedName",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																AdditionalPropertiesAllowed: true,
															},
														}},
													},
													AdditionalPropertiesAllowed: true,
												},
											},
											{
												SchemaPointer:    "#/components/schemas/RefStressMiddleAllOf/allOf/1",
												BodyRequired:     false,
												KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
												EnumValidation:   validation.EnumValidation{},
												NumberValidation: validation.NumberValidation{},
												StringValidation: validation.StringValidation{Format: ""},
												ArrayValidation: validation.ArrayValidation{
													UniqueItems: false,
												},
												ObjectValidation: validation.ObjectValidation{
													Required: []string{"sharedName"},
													Properties: []validation.PropertyValidation{
														{Name: "optionalCode", Validation: &validation.Validation{
															SchemaPointer:    "#/components/schemas/RefStressMiddleAllOf/allOf/1/properties/optionalCode",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																AdditionalPropertiesAllowed: true,
															},
														}},
														{Name: "sharedName", Validation: &validation.Validation{
															SchemaPointer:    "#/components/schemas/RefStressMiddleAllOf/allOf/1/properties/sharedName",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																AdditionalPropertiesAllowed: true,
															},
														}},
													},
													AdditionalPropertiesAllowed: true,
												},
											},
										},
									},
									{
										SchemaPointer:    "#/components/schemas/RefStressViaMiddle/allOf/1",
										BodyRequired:     false,
										KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
										EnumValidation:   validation.EnumValidation{},
										NumberValidation: validation.NumberValidation{},
										StringValidation: validation.StringValidation{Format: ""},
										ArrayValidation: validation.ArrayValidation{
											UniqueItems: false,
										},
										ObjectValidation: validation.ObjectValidation{
											Required: []string{"middleFlag", "sharedName"},
											Properties: []validation.PropertyValidation{
												{Name: "middleFlag", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressViaMiddle/allOf/1/properties/middleFlag",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "boolean", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														AdditionalPropertiesAllowed: true,
													},
												}},
												{Name: "nested", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressNestedCombined",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														AdditionalPropertiesAllowed: true,
													},
													AllOfValidations: []*validation.Validation{
														{
															SchemaPointer:    "#/components/schemas/RefStressNestedBase",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																Required: []string{"sameName"},
																Properties: []validation.PropertyValidation{
																	{Name: "leaf", Validation: &validation.Validation{
																		SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																		BodyRequired:     false,
																		KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																		EnumValidation:   validation.EnumValidation{},
																		NumberValidation: validation.NumberValidation{},
																		StringValidation: validation.StringValidation{Format: ""},
																		ArrayValidation: validation.ArrayValidation{
																			UniqueItems: false,
																		},
																		ObjectValidation: validation.ObjectValidation{
																			AdditionalPropertiesAllowed: true,
																		},
																	}},
																	{Name: "sameName", Validation: &validation.Validation{
																		SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
																		BodyRequired:     false,
																		KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																		EnumValidation:   validation.EnumValidation{},
																		NumberValidation: validation.NumberValidation{},
																		StringValidation: validation.StringValidation{Format: ""},
																		ArrayValidation: validation.ArrayValidation{
																			UniqueItems: false,
																		},
																		ObjectValidation: validation.ObjectValidation{
																			AdditionalPropertiesAllowed: true,
																		},
																	}},
																},
																AdditionalPropertiesAllowed: true,
															},
														},
														{
															SchemaPointer:    "#/components/schemas/RefStressNestedOverlay",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																Required: []string{"sameName"},
																Properties: []validation.PropertyValidation{
																	{Name: "leaf", Validation: &validation.Validation{
																		SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																		BodyRequired:     false,
																		KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																		EnumValidation:   validation.EnumValidation{},
																		NumberValidation: validation.NumberValidation{},
																		StringValidation: validation.StringValidation{Format: ""},
																		ArrayValidation: validation.ArrayValidation{
																			UniqueItems: false,
																		},
																		ObjectValidation: validation.ObjectValidation{
																			AdditionalPropertiesAllowed: true,
																		},
																	}},
																	{Name: "sameName", Validation: &validation.Validation{
																		SchemaPointer:    "#/components/schemas/RefStressNestedOverlay/properties/sameName",
																		BodyRequired:     false,
																		KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																		EnumValidation:   validation.EnumValidation{},
																		NumberValidation: validation.NumberValidation{},
																		StringValidation: validation.StringValidation{Format: ""},
																		ArrayValidation: validation.ArrayValidation{
																			UniqueItems: false,
																		},
																		ObjectValidation: validation.ObjectValidation{
																			AdditionalPropertiesAllowed: true,
																		},
																	}},
																},
																AdditionalPropertiesAllowed: true,
															},
														},
														{
															SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																Required: []string{"sameName", "sealed"},
																Properties: []validation.PropertyValidation{
																	{Name: "sameName", Validation: &validation.Validation{
																		SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2/properties/sameName",
																		BodyRequired:     false,
																		KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																		EnumValidation:   validation.EnumValidation{},
																		NumberValidation: validation.NumberValidation{},
																		StringValidation: validation.StringValidation{Format: ""},
																		ArrayValidation: validation.ArrayValidation{
																			UniqueItems: false,
																		},
																		ObjectValidation: validation.ObjectValidation{
																			AdditionalPropertiesAllowed: true,
																		},
																	}},
																	{Name: "sealed", Validation: &validation.Validation{
																		SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2/properties/sealed",
																		BodyRequired:     false,
																		KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
																		EnumValidation:   validation.EnumValidation{},
																		NumberValidation: validation.NumberValidation{},
																		StringValidation: validation.StringValidation{Format: ""},
																		ArrayValidation: validation.ArrayValidation{
																			UniqueItems: false,
																		},
																		ObjectValidation: validation.ObjectValidation{
																			Required: []string{"locked"},
																			Properties: []validation.PropertyValidation{
																				{Name: "locked", Validation: &validation.Validation{
																					SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2/properties/sealed/properties/locked",
																					BodyRequired:     false,
																					KindValidation:   validation.KindValidation{Type: "boolean", Nullable: false},
																					EnumValidation:   validation.EnumValidation{},
																					NumberValidation: validation.NumberValidation{},
																					StringValidation: validation.StringValidation{Format: ""},
																					ArrayValidation: validation.ArrayValidation{
																						UniqueItems: false,
																					},
																					ObjectValidation: validation.ObjectValidation{
																						AdditionalPropertiesAllowed: true,
																					},
																				}},
																			},
																			AdditionalPropertiesAllowed: false,
																		},
																	}},
																},
																AdditionalPropertiesAllowed: true,
															},
														},
													},
												}},
												{Name: "sharedName", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressViaMiddle/allOf/1/properties/sharedName",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														AdditionalPropertiesAllowed: true,
													},
												}},
											},
											AdditionalPropertiesAllowed: true,
										},
									},
								},
							},
							{
								SchemaPointer:    "#/components/schemas/RefStressFirstAllOf/allOf/2",
								BodyRequired:     false,
								KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
								EnumValidation:   validation.EnumValidation{},
								NumberValidation: validation.NumberValidation{},
								StringValidation: validation.StringValidation{Format: ""},
								ArrayValidation: validation.ArrayValidation{
									UniqueItems: false,
								},
								ObjectValidation: validation.ObjectValidation{
									Required: []string{"final", "nested", "nullableRequired"},
									Properties: []validation.PropertyValidation{
										{Name: "final", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressFinal",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												Required: []string{"finalCode", "sharedName"},
												Properties: []validation.PropertyValidation{
													{Name: "finalCode", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressFinal/properties/finalCode",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "nested", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressNestedBase",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															Required: []string{"sameName"},
															Properties: []validation.PropertyValidation{
																{Name: "leaf", Validation: &validation.Validation{
																	SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																	BodyRequired:     false,
																	KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																	EnumValidation:   validation.EnumValidation{},
																	NumberValidation: validation.NumberValidation{},
																	StringValidation: validation.StringValidation{Format: ""},
																	ArrayValidation: validation.ArrayValidation{
																		UniqueItems: false,
																	},
																	ObjectValidation: validation.ObjectValidation{
																		AdditionalPropertiesAllowed: true,
																	},
																}},
																{Name: "sameName", Validation: &validation.Validation{
																	SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
																	BodyRequired:     false,
																	KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																	EnumValidation:   validation.EnumValidation{},
																	NumberValidation: validation.NumberValidation{},
																	StringValidation: validation.StringValidation{Format: ""},
																	ArrayValidation: validation.ArrayValidation{
																		UniqueItems: false,
																	},
																	ObjectValidation: validation.ObjectValidation{
																		AdditionalPropertiesAllowed: true,
																	},
																}},
															},
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "optionalShared", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressFinal/properties/optionalShared",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "sharedName", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressFinal/properties/sharedName",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
												},
												AdditionalPropertiesAllowed: true,
											},
										}},
										{Name: "nested", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressNestedCombined",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
											AllOfValidations: []*validation.Validation{
												{
													SchemaPointer:    "#/components/schemas/RefStressNestedBase",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														Required: []string{"sameName"},
														Properties: []validation.PropertyValidation{
															{Name: "leaf", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "sameName", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
														},
														AdditionalPropertiesAllowed: true,
													},
												},
												{
													SchemaPointer:    "#/components/schemas/RefStressNestedOverlay",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														Required: []string{"sameName"},
														Properties: []validation.PropertyValidation{
															{Name: "leaf", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "sameName", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressNestedOverlay/properties/sameName",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
														},
														AdditionalPropertiesAllowed: true,
													},
												},
												{
													SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														Required: []string{"sameName", "sealed"},
														Properties: []validation.PropertyValidation{
															{Name: "sameName", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2/properties/sameName",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "sealed", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2/properties/sealed",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	Required: []string{"locked"},
																	Properties: []validation.PropertyValidation{
																		{Name: "locked", Validation: &validation.Validation{
																			SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2/properties/sealed/properties/locked",
																			BodyRequired:     false,
																			KindValidation:   validation.KindValidation{Type: "boolean", Nullable: false},
																			EnumValidation:   validation.EnumValidation{},
																			NumberValidation: validation.NumberValidation{},
																			StringValidation: validation.StringValidation{Format: ""},
																			ArrayValidation: validation.ArrayValidation{
																				UniqueItems: false,
																			},
																			ObjectValidation: validation.ObjectValidation{
																				AdditionalPropertiesAllowed: true,
																			},
																		}},
																	},
																	AdditionalPropertiesAllowed: false,
																},
															}},
														},
														AdditionalPropertiesAllowed: true,
													},
												},
											},
										}},
										{Name: "nullableRequired", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressFirstAllOf/allOf/2/properties/nullableRequired",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
										{Name: "optionalShared", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressFirstAllOf/allOf/2/properties/optionalShared",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
										{Name: "sharedName", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressFirstAllOf/allOf/2/properties/sharedName",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
									},
									AdditionalPropertiesAllowed: true,
								},
							},
						},
					},
					{
						SchemaPointer:    "#/components/schemas/RefStressSecondAllOf",
						BodyRequired:     false,
						KindValidation:   validation.KindValidation{Type: "", Nullable: false},
						EnumValidation:   validation.EnumValidation{},
						NumberValidation: validation.NumberValidation{},
						StringValidation: validation.StringValidation{Format: ""},
						ArrayValidation: validation.ArrayValidation{
							UniqueItems: false,
						},
						ObjectValidation: validation.ObjectValidation{
							AdditionalPropertiesAllowed: true,
						},
						AllOfValidations: []*validation.Validation{
							{
								SchemaPointer:    "#/components/schemas/RefStressOtherMiddle",
								BodyRequired:     false,
								KindValidation:   validation.KindValidation{Type: "", Nullable: false},
								EnumValidation:   validation.EnumValidation{},
								NumberValidation: validation.NumberValidation{},
								StringValidation: validation.StringValidation{Format: ""},
								ArrayValidation: validation.ArrayValidation{
									UniqueItems: false,
								},
								ObjectValidation: validation.ObjectValidation{
									AdditionalPropertiesAllowed: true,
								},
								AllOfValidations: []*validation.Validation{
									{
										SchemaPointer:    "#/components/schemas/RefStressFinal",
										BodyRequired:     false,
										KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
										EnumValidation:   validation.EnumValidation{},
										NumberValidation: validation.NumberValidation{},
										StringValidation: validation.StringValidation{Format: ""},
										ArrayValidation: validation.ArrayValidation{
											UniqueItems: false,
										},
										ObjectValidation: validation.ObjectValidation{
											Required: []string{"finalCode", "sharedName"},
											Properties: []validation.PropertyValidation{
												{Name: "finalCode", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressFinal/properties/finalCode",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														AdditionalPropertiesAllowed: true,
													},
												}},
												{Name: "nested", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressNestedBase",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														Required: []string{"sameName"},
														Properties: []validation.PropertyValidation{
															{Name: "leaf", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "sameName", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
														},
														AdditionalPropertiesAllowed: true,
													},
												}},
												{Name: "optionalShared", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressFinal/properties/optionalShared",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														AdditionalPropertiesAllowed: true,
													},
												}},
												{Name: "sharedName", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressFinal/properties/sharedName",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														AdditionalPropertiesAllowed: true,
													},
												}},
											},
											AdditionalPropertiesAllowed: true,
										},
									},
									{
										SchemaPointer:    "#/components/schemas/RefStressOtherMiddle/allOf/1",
										BodyRequired:     false,
										KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
										EnumValidation:   validation.EnumValidation{},
										NumberValidation: validation.NumberValidation{},
										StringValidation: validation.StringValidation{Format: ""},
										ArrayValidation: validation.ArrayValidation{
											UniqueItems: false,
										},
										ObjectValidation: validation.ObjectValidation{
											Required: []string{"metadata", "rootFlag"},
											Properties: []validation.PropertyValidation{
												{Name: "final", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressFinal",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														Required: []string{"finalCode", "sharedName"},
														Properties: []validation.PropertyValidation{
															{Name: "finalCode", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressFinal/properties/finalCode",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "nested", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressNestedBase",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	Required: []string{"sameName"},
																	Properties: []validation.PropertyValidation{
																		{Name: "leaf", Validation: &validation.Validation{
																			SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																			BodyRequired:     false,
																			KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																			EnumValidation:   validation.EnumValidation{},
																			NumberValidation: validation.NumberValidation{},
																			StringValidation: validation.StringValidation{Format: ""},
																			ArrayValidation: validation.ArrayValidation{
																				UniqueItems: false,
																			},
																			ObjectValidation: validation.ObjectValidation{
																				AdditionalPropertiesAllowed: true,
																			},
																		}},
																		{Name: "sameName", Validation: &validation.Validation{
																			SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
																			BodyRequired:     false,
																			KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																			EnumValidation:   validation.EnumValidation{},
																			NumberValidation: validation.NumberValidation{},
																			StringValidation: validation.StringValidation{Format: ""},
																			ArrayValidation: validation.ArrayValidation{
																				UniqueItems: false,
																			},
																			ObjectValidation: validation.ObjectValidation{
																				AdditionalPropertiesAllowed: true,
																			},
																		}},
																	},
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "optionalShared", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressFinal/properties/optionalShared",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "sharedName", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressFinal/properties/sharedName",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
														},
														AdditionalPropertiesAllowed: true,
													},
												}},
												{Name: "metadata", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressOtherMiddle/allOf/1/properties/metadata",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														AdditionalPropertiesAllowed: true,
														AdditionalPropertiesValidation: &validation.Validation{
															SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																AdditionalPropertiesAllowed: true,
															},
														},
													},
												}},
												{Name: "rootFlag", Validation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressOtherMiddle/allOf/1/properties/rootFlag",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "boolean", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														AdditionalPropertiesAllowed: true,
													},
												}},
											},
											AdditionalPropertiesAllowed: true,
										},
									},
								},
							},
							{
								SchemaPointer:    "#/components/schemas/RefStressSecondAllOf/allOf/1",
								BodyRequired:     false,
								KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
								EnumValidation:   validation.EnumValidation{},
								NumberValidation: validation.NumberValidation{},
								StringValidation: validation.StringValidation{Format: ""},
								ArrayValidation: validation.ArrayValidation{
									UniqueItems: false,
								},
								ObjectValidation: validation.ObjectValidation{
									Required: []string{"count", "finals", "metadata", "rootFlag"},
									Properties: []validation.PropertyValidation{
										{Name: "count", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressSecondAllOf/allOf/1/properties/count",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "number", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
										{Name: "finals", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressSecondAllOf/allOf/1/properties/finals",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "array", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												Items: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressFinal",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														Required: []string{"finalCode", "sharedName"},
														Properties: []validation.PropertyValidation{
															{Name: "finalCode", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressFinal/properties/finalCode",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "nested", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressNestedBase",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	Required: []string{"sameName"},
																	Properties: []validation.PropertyValidation{
																		{Name: "leaf", Validation: &validation.Validation{
																			SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																			BodyRequired:     false,
																			KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																			EnumValidation:   validation.EnumValidation{},
																			NumberValidation: validation.NumberValidation{},
																			StringValidation: validation.StringValidation{Format: ""},
																			ArrayValidation: validation.ArrayValidation{
																				UniqueItems: false,
																			},
																			ObjectValidation: validation.ObjectValidation{
																				AdditionalPropertiesAllowed: true,
																			},
																		}},
																		{Name: "sameName", Validation: &validation.Validation{
																			SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
																			BodyRequired:     false,
																			KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																			EnumValidation:   validation.EnumValidation{},
																			NumberValidation: validation.NumberValidation{},
																			StringValidation: validation.StringValidation{Format: ""},
																			ArrayValidation: validation.ArrayValidation{
																				UniqueItems: false,
																			},
																			ObjectValidation: validation.ObjectValidation{
																				AdditionalPropertiesAllowed: true,
																			},
																		}},
																	},
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "optionalShared", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressFinal/properties/optionalShared",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
															{Name: "sharedName", Validation: &validation.Validation{
																SchemaPointer:    "#/components/schemas/RefStressFinal/properties/sharedName",
																BodyRequired:     false,
																KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																EnumValidation:   validation.EnumValidation{},
																NumberValidation: validation.NumberValidation{},
																StringValidation: validation.StringValidation{Format: ""},
																ArrayValidation: validation.ArrayValidation{
																	UniqueItems: false,
																},
																ObjectValidation: validation.ObjectValidation{
																	AdditionalPropertiesAllowed: true,
																},
															}},
														},
														AdditionalPropertiesAllowed: true,
													},
												},
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
										{Name: "metadata", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressSecondAllOf/allOf/1/properties/metadata",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
												AdditionalPropertiesValidation: &validation.Validation{
													SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
													BodyRequired:     false,
													KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
													EnumValidation:   validation.EnumValidation{},
													NumberValidation: validation.NumberValidation{},
													StringValidation: validation.StringValidation{Format: ""},
													ArrayValidation: validation.ArrayValidation{
														UniqueItems: false,
													},
													ObjectValidation: validation.ObjectValidation{
														AdditionalPropertiesAllowed: true,
													},
												},
											},
										}},
										{Name: "rootFlag", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressSecondAllOf/allOf/1/properties/rootFlag",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "boolean", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
										{Name: "sharedName", Validation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressSecondAllOf/allOf/1/properties/sharedName",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										}},
									},
									AdditionalPropertiesAllowed: true,
								},
							},
						},
					},
					{
						SchemaPointer:    "#/paths/~1ref-stress-object-put/put/requestBody/content/application~1json/schema/allOf/2",
						BodyRequired:     false,
						KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
						EnumValidation:   validation.EnumValidation{},
						NumberValidation: validation.NumberValidation{},
						StringValidation: validation.StringValidation{Format: ""},
						ArrayValidation: validation.ArrayValidation{
							UniqueItems: false,
						},
						ObjectValidation: validation.ObjectValidation{
							Required: []string{"count", "final", "finalCode", "finals", "metadata", "middleFlag", "nested", "nullableRequired", "rootFlag", "sharedName"},
							Properties: []validation.PropertyValidation{
								{Name: "count", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object-put/put/requestBody/content/application~1json/schema/allOf/2/properties/count",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "number", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
								{Name: "final", Validation: &validation.Validation{
									SchemaPointer:    "#/components/schemas/RefStressFinal",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										Required: []string{"finalCode", "sharedName"},
										Properties: []validation.PropertyValidation{
											{Name: "finalCode", Validation: &validation.Validation{
												SchemaPointer:    "#/components/schemas/RefStressFinal/properties/finalCode",
												BodyRequired:     false,
												KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
												EnumValidation:   validation.EnumValidation{},
												NumberValidation: validation.NumberValidation{},
												StringValidation: validation.StringValidation{Format: ""},
												ArrayValidation: validation.ArrayValidation{
													UniqueItems: false,
												},
												ObjectValidation: validation.ObjectValidation{
													AdditionalPropertiesAllowed: true,
												},
											}},
											{Name: "nested", Validation: &validation.Validation{
												SchemaPointer:    "#/components/schemas/RefStressNestedBase",
												BodyRequired:     false,
												KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
												EnumValidation:   validation.EnumValidation{},
												NumberValidation: validation.NumberValidation{},
												StringValidation: validation.StringValidation{Format: ""},
												ArrayValidation: validation.ArrayValidation{
													UniqueItems: false,
												},
												ObjectValidation: validation.ObjectValidation{
													Required: []string{"sameName"},
													Properties: []validation.PropertyValidation{
														{Name: "leaf", Validation: &validation.Validation{
															SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																AdditionalPropertiesAllowed: true,
															},
														}},
														{Name: "sameName", Validation: &validation.Validation{
															SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
															BodyRequired:     false,
															KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
															EnumValidation:   validation.EnumValidation{},
															NumberValidation: validation.NumberValidation{},
															StringValidation: validation.StringValidation{Format: ""},
															ArrayValidation: validation.ArrayValidation{
																UniqueItems: false,
															},
															ObjectValidation: validation.ObjectValidation{
																AdditionalPropertiesAllowed: true,
															},
														}},
													},
													AdditionalPropertiesAllowed: true,
												},
											}},
											{Name: "optionalShared", Validation: &validation.Validation{
												SchemaPointer:    "#/components/schemas/RefStressFinal/properties/optionalShared",
												BodyRequired:     false,
												KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
												EnumValidation:   validation.EnumValidation{},
												NumberValidation: validation.NumberValidation{},
												StringValidation: validation.StringValidation{Format: ""},
												ArrayValidation: validation.ArrayValidation{
													UniqueItems: false,
												},
												ObjectValidation: validation.ObjectValidation{
													AdditionalPropertiesAllowed: true,
												},
											}},
											{Name: "sharedName", Validation: &validation.Validation{
												SchemaPointer:    "#/components/schemas/RefStressFinal/properties/sharedName",
												BodyRequired:     false,
												KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
												EnumValidation:   validation.EnumValidation{},
												NumberValidation: validation.NumberValidation{},
												StringValidation: validation.StringValidation{Format: ""},
												ArrayValidation: validation.ArrayValidation{
													UniqueItems: false,
												},
												ObjectValidation: validation.ObjectValidation{
													AdditionalPropertiesAllowed: true,
												},
											}},
										},
										AdditionalPropertiesAllowed: true,
									},
								}},
								{Name: "finalCode", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object-put/put/requestBody/content/application~1json/schema/allOf/2/properties/finalCode",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
								{Name: "finals", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object-put/put/requestBody/content/application~1json/schema/allOf/2/properties/finals",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "array", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										Items: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressFinal",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												Required: []string{"finalCode", "sharedName"},
												Properties: []validation.PropertyValidation{
													{Name: "finalCode", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressFinal/properties/finalCode",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "nested", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressNestedBase",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															Required: []string{"sameName"},
															Properties: []validation.PropertyValidation{
																{Name: "leaf", Validation: &validation.Validation{
																	SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
																	BodyRequired:     false,
																	KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
																	EnumValidation:   validation.EnumValidation{},
																	NumberValidation: validation.NumberValidation{},
																	StringValidation: validation.StringValidation{Format: ""},
																	ArrayValidation: validation.ArrayValidation{
																		UniqueItems: false,
																	},
																	ObjectValidation: validation.ObjectValidation{
																		AdditionalPropertiesAllowed: true,
																	},
																}},
																{Name: "sameName", Validation: &validation.Validation{
																	SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
																	BodyRequired:     false,
																	KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
																	EnumValidation:   validation.EnumValidation{},
																	NumberValidation: validation.NumberValidation{},
																	StringValidation: validation.StringValidation{Format: ""},
																	ArrayValidation: validation.ArrayValidation{
																		UniqueItems: false,
																	},
																	ObjectValidation: validation.ObjectValidation{
																		AdditionalPropertiesAllowed: true,
																	},
																}},
															},
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "optionalShared", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressFinal/properties/optionalShared",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "sharedName", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressFinal/properties/sharedName",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
												},
												AdditionalPropertiesAllowed: true,
											},
										},
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
								{Name: "metadata", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object-put/put/requestBody/content/application~1json/schema/allOf/2/properties/metadata",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
										AdditionalPropertiesValidation: &validation.Validation{
											SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												AdditionalPropertiesAllowed: true,
											},
										},
									},
								}},
								{Name: "middleFlag", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object-put/put/requestBody/content/application~1json/schema/allOf/2/properties/middleFlag",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "boolean", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
								{Name: "nested", Validation: &validation.Validation{
									SchemaPointer:    "#/components/schemas/RefStressNestedCombined",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
									AllOfValidations: []*validation.Validation{
										{
											SchemaPointer:    "#/components/schemas/RefStressNestedBase",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "object", Nullable: true},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												Required: []string{"sameName"},
												Properties: []validation.PropertyValidation{
													{Name: "leaf", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "sameName", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressNestedBase/properties/sameName",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
												},
												AdditionalPropertiesAllowed: true,
											},
										},
										{
											SchemaPointer:    "#/components/schemas/RefStressNestedOverlay",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												Required: []string{"sameName"},
												Properties: []validation.PropertyValidation{
													{Name: "leaf", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressMetadataValue",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "sameName", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressNestedOverlay/properties/sameName",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
												},
												AdditionalPropertiesAllowed: true,
											},
										},
										{
											SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2",
											BodyRequired:     false,
											KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
											EnumValidation:   validation.EnumValidation{},
											NumberValidation: validation.NumberValidation{},
											StringValidation: validation.StringValidation{Format: ""},
											ArrayValidation: validation.ArrayValidation{
												UniqueItems: false,
											},
											ObjectValidation: validation.ObjectValidation{
												Required: []string{"sameName", "sealed"},
												Properties: []validation.PropertyValidation{
													{Name: "sameName", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2/properties/sameName",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															AdditionalPropertiesAllowed: true,
														},
													}},
													{Name: "sealed", Validation: &validation.Validation{
														SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2/properties/sealed",
														BodyRequired:     false,
														KindValidation:   validation.KindValidation{Type: "object", Nullable: false},
														EnumValidation:   validation.EnumValidation{},
														NumberValidation: validation.NumberValidation{},
														StringValidation: validation.StringValidation{Format: ""},
														ArrayValidation: validation.ArrayValidation{
															UniqueItems: false,
														},
														ObjectValidation: validation.ObjectValidation{
															Required: []string{"locked"},
															Properties: []validation.PropertyValidation{
																{Name: "locked", Validation: &validation.Validation{
																	SchemaPointer:    "#/components/schemas/RefStressNestedCombined/allOf/2/properties/sealed/properties/locked",
																	BodyRequired:     false,
																	KindValidation:   validation.KindValidation{Type: "boolean", Nullable: false},
																	EnumValidation:   validation.EnumValidation{},
																	NumberValidation: validation.NumberValidation{},
																	StringValidation: validation.StringValidation{Format: ""},
																	ArrayValidation: validation.ArrayValidation{
																		UniqueItems: false,
																	},
																	ObjectValidation: validation.ObjectValidation{
																		AdditionalPropertiesAllowed: true,
																	},
																}},
															},
															AdditionalPropertiesAllowed: false,
														},
													}},
												},
												AdditionalPropertiesAllowed: true,
											},
										},
									},
								}},
								{Name: "nullableRequired", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object-put/put/requestBody/content/application~1json/schema/allOf/2/properties/nullableRequired",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
								{Name: "optionalCode", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object-put/put/requestBody/content/application~1json/schema/allOf/2/properties/optionalCode",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
								{Name: "optionalShared", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object-put/put/requestBody/content/application~1json/schema/allOf/2/properties/optionalShared",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
								{Name: "rootFlag", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object-put/put/requestBody/content/application~1json/schema/allOf/2/properties/rootFlag",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "boolean", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
								{Name: "sharedName", Validation: &validation.Validation{
									SchemaPointer:    "#/paths/~1ref-stress-object-put/put/requestBody/content/application~1json/schema/allOf/2/properties/sharedName",
									BodyRequired:     false,
									KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
									EnumValidation:   validation.EnumValidation{},
									NumberValidation: validation.NumberValidation{},
									StringValidation: validation.StringValidation{Format: ""},
									ArrayValidation: validation.ArrayValidation{
										UniqueItems: false,
									},
									ObjectValidation: validation.ObjectValidation{
										AdditionalPropertiesAllowed: true,
									},
								}},
							},
							AdditionalPropertiesAllowed: false,
						},
					},
				},
			},
		},
		"stringNoFormatNotNullable": {
			expectedValidation: &validation.Validation{
				SchemaPointer:    "#/paths/~1string-no-format-not-nullable/post/requestBody/content/application~1json/schema",
				BodyRequired:     true,
				KindValidation:   validation.KindValidation{Type: "string", Nullable: false},
				EnumValidation:   validation.EnumValidation{},
				NumberValidation: validation.NumberValidation{},
				StringValidation: validation.StringValidation{Format: ""},
				ArrayValidation: validation.ArrayValidation{
					UniqueItems: false,
				},
				ObjectValidation: validation.ObjectValidation{
					AdditionalPropertiesAllowed: true,
				},
			},
		},
		"stringNoFormatNullable": {
			expectedValidation: &validation.Validation{
				SchemaPointer:    "#/paths/~1string-no-format-nullable/post/requestBody/content/application~1json/schema",
				BodyRequired:     true,
				KindValidation:   validation.KindValidation{Type: "string", Nullable: true},
				EnumValidation:   validation.EnumValidation{},
				NumberValidation: validation.NumberValidation{},
				StringValidation: validation.StringValidation{Format: ""},
				ArrayValidation: validation.ArrayValidation{
					UniqueItems: false,
				},
				ObjectValidation: validation.ObjectValidation{
					AdditionalPropertiesAllowed: true,
				},
			},
		},
	}

	for operationID, test := range tests {
		t.Run(operationID, func(t *testing.T) {
			t.Parallel()

			actualValidation, parseErr := validation.Parse(spec, operationID)
			require.NoError(t, parseErr)
			require.Equal(t, test.expectedValidation, actualValidation)
		})
	}
}
