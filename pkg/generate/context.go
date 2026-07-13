// Package generate creates Go request-body models from OpenAPI documents.
package generate

import (
	"fmt"
	"maps"
	"slices"

	"github.com/getkin/kin-openapi/openapi3"
)

// GenerateContext contains an OpenAPI document and generation selections.
//
//nolint:revive // Existing public API keeps its original generator-specific name.
type GenerateContext struct {
	Document                  *openapi3.T
	OpenAPISource             []byte
	JSONRequestBodyOperations []JSONRequestBodyOperation
}

// Schema is a model schema understood by the code generator.
type Schema interface {
	Base() *BaseSchema
	ChildSchemas() []Schema
	Generate() (string, error)
	SchemaTypeName() string
	SetTypeName(name string)
}

// BaseSchema contains fields common to every generated schema.
type BaseSchema struct {
	Name             string
	Nullable         bool
	EmptyBodyAllowed bool
}

// JSONRequestBodyOperation identifies one generated operation test.
type JSONRequestBodyOperation struct {
	OperationID string
	TypeName    string
}

// TestName returns the operation's exported Go test suffix.
func (o JSONRequestBodyOperation) TestName() string {
	return exportedName(o.OperationID)
}

// Base returns b.
func (b *BaseSchema) Base() *BaseSchema {
	return b
}

// SchemaTypeName returns the generated Go type name.
func (b *BaseSchema) SchemaTypeName() string {
	if b == nil {
		return ""
	}

	return b.Name
}

// SetTypeName sets the generated Go type name.
func (b *BaseSchema) SetTypeName(name string) {
	b.Name = name
}

// LocalName returns the unexported form of the generated type name.
func (b *BaseSchema) LocalName() string {
	return unexportedName(b.SchemaTypeName())
}

// ObjectSchema describes a generated JSON object.
type ObjectSchema struct {
	BaseSchema

	AdditionalProperties       bool
	AdditionalPropertiesSchema Schema
	Properties                 []ObjectFieldContext
}

var (
	_ Schema = new(ObjectSchema)
	_ Schema = new(StringSchema)
	_ Schema = new(ArraySchema)
	_ Schema = new(AllOfSchema)
	_ Schema = new(BoolSchema)
	_ Schema = new(NumberSchema)
)

// ObjectFieldContext describes one generated object property.
type ObjectFieldContext struct {
	Schema

	PropertyName string
	Required     bool
}

// StringSchema describes a generated JSON string.
type StringSchema struct {
	BaseSchema

	Format string
}

// BoolSchema describes a generated JSON boolean.
type BoolSchema struct {
	BaseSchema
}

// NumberSchema describes a generated JSON number.
type NumberSchema struct {
	BaseSchema
}

// ArraySchema describes a generated JSON array.
type ArraySchema struct {
	BaseSchema

	Items Schema
}

// AllOfSchema describes a generated OpenAPI allOf schema.
type AllOfSchema struct {
	BaseSchema

	Schemas []Schema
}

// ChildSchemas returns property and additional-property schemas.
func (o *ObjectSchema) ChildSchemas() []Schema {
	children := make([]Schema, 0, len(o.Properties)+1)
	for _, property := range o.Properties {
		children = append(children, property.Schema)
	}

	if o.AdditionalPropertiesSchema != nil {
		children = append(children, o.AdditionalPropertiesSchema)
	}

	return children
}

// ChildSchemas returns no children for a string schema.
func (s *StringSchema) ChildSchemas() []Schema {
	return nil
}

// ChildSchemas returns no children for a boolean schema.
func (b *BoolSchema) ChildSchemas() []Schema {
	return nil
}

// ChildSchemas returns no children for a number schema.
func (n *NumberSchema) ChildSchemas() []Schema {
	return nil
}

// ChildSchemas returns the array item schema.
func (a *ArraySchema) ChildSchemas() []Schema {
	if a.Items == nil {
		return nil
	}

	return []Schema{a.Items}
}

// ChildSchemas returns every allOf member.
func (a *AllOfSchema) ChildSchemas() []Schema {
	return a.Schemas
}

// FieldType returns the generated Go field type.
func (p *ObjectFieldContext) FieldType() string {
	if p.Required {
		return p.SchemaTypeName()
	}

	return "*" + p.SchemaTypeName()
}

// JSONTag returns the generated JSON struct tag.
func (p *ObjectFieldContext) JSONTag() string {
	if p.Required {
		return fmt.Sprintf("`json:%q`", p.PropertyName)
	}

	return fmt.Sprintf("`json:%q`", p.PropertyName+",omitzero")
}

// LocalName returns the property's unexported Go identifier.
func (p *ObjectFieldContext) LocalName() string {
	return unexportedName(p.PropertyName)
}

// FieldName returns the property's exported Go identifier.
func (p *ObjectFieldContext) FieldName() string {
	return exportedName(p.PropertyName)
}

// Generate renders an object property decoder.
func (p *ObjectFieldContext) Generate() (string, error) {
	if p == nil {
		return "", fmt.Errorf("nil object property")
	}

	return executeGoTemplate("object_property.go.tmpl", p)
}

// Generate renders an object schema.
func (o *ObjectSchema) Generate() (string, error) {
	if o == nil {
		return "", fmt.Errorf("nil object schema")
	}

	return executeGoTemplate("object.go.tmpl", o)
}

// AdditionalPropertiesTypeName returns the additional-property Go type.
func (o *ObjectSchema) AdditionalPropertiesTypeName() string {
	return schemaTypeName(o.AdditionalPropertiesSchema)
}

// Generate renders a string schema.
func (s *StringSchema) Generate() (string, error) {
	if s == nil {
		return "", fmt.Errorf("nil string schema")
	}

	return executeGoTemplate("string.go.tmpl", s)
}

// Generate renders a boolean schema.
func (b *BoolSchema) Generate() (string, error) {
	if b == nil {
		return "", fmt.Errorf("nil bool schema")
	}

	return executeGoTemplate("bool.go.tmpl", b)
}

// Generate renders a number schema.
func (n *NumberSchema) Generate() (string, error) {
	if n == nil {
		return "", fmt.Errorf("nil number schema")
	}

	return executeGoTemplate("number.go.tmpl", n)
}

// Generate renders an array schema.
func (a *ArraySchema) Generate() (string, error) {
	if a == nil {
		return "", fmt.Errorf("nil array schema")
	}

	return executeGoTemplate("array.go.tmpl", a)
}

// Generate renders an allOf schema.
func (a *AllOfSchema) Generate() (string, error) {
	if a == nil {
		return "", fmt.Errorf("nil allOf schema")
	}

	return executeGoTemplate("all_of.go.tmpl", a)
}

// ItemsTypeName returns the array item Go type.
func (a *ArraySchema) ItemsTypeName() string {
	return schemaTypeName(a.Items)
}

// schemaTypeName returns schema's generated Go type name.
func schemaTypeName(schema Schema) string {
	if schema == nil {
		return ""
	}

	base := schema.Base()
	if base == nil {
		return ""
	}

	return base.SchemaTypeName()
}

// jsonRequestBody pairs an operation with its resolved JSON schema.
type jsonRequestBody struct {
	operation *openapi3.Operation
	schema    *openapi3.Schema
	required  bool
}

// JSONRequestBodySchemas returns operation request-body JSON schemas.
func (c *GenerateContext) JSONRequestBodySchemas() (map[*openapi3.Operation]*openapi3.Schema, error) {
	requestBodies, err := c.jsonRequestBodies()
	if err != nil {
		return nil, err
	}

	schemas := make(map[*openapi3.Operation]*openapi3.Schema, len(requestBodies))
	for _, requestBody := range requestBodies {
		schemas[requestBody.operation] = requestBody.schema
	}

	return schemas, nil
}

// JSONRequestBodyModelSchemas converts selected request bodies to model schemas.
func (c *GenerateContext) JSONRequestBodyModelSchemas() ([]Schema, error) {
	requestBodies, err := c.jsonRequestBodies()
	if err != nil {
		return nil, err
	}

	c.JSONRequestBodyOperations = nil

	var schemas []Schema

	for _, requestBody := range requestBodies {
		definitions, operation, conversionErr := modelSchemasForRequestBody(requestBody)
		if conversionErr != nil {
			return nil, conversionErr
		}

		schemas = append(schemas, definitions...)
		c.JSONRequestBodyOperations = append(c.JSONRequestBodyOperations, operation)
	}

	return schemas, nil
}

// jsonRequestBodies collects resolved application/json request bodies.
func (c *GenerateContext) jsonRequestBodies() ([]jsonRequestBody, error) {
	if c.Document == nil || c.Document.Paths == nil {
		return nil, fmt.Errorf("openapi document has no paths")
	}

	var requestBodies []jsonRequestBody

	for _, path := range c.Document.Paths.InMatchingOrder() {
		pathItem := c.Document.Paths.Value(path)
		if pathItem == nil {
			continue
		}

		operations := pathItem.Operations()
		for _, method := range slices.Sorted(maps.Keys(operations)) {
			requestBody, ok := requestBodyForOperation(operations[method])
			if ok {
				requestBodies = append(requestBodies, requestBody)
			}
		}
	}

	return requestBodies, nil
}

// requestBodyForOperation returns an operation's resolved JSON request body.
func requestBodyForOperation(operation *openapi3.Operation) (jsonRequestBody, bool) {
	if operation == nil || operation.RequestBody == nil {
		return jsonRequestBody{}, false
	}

	requestBody := operation.RequestBody.Value
	if requestBody == nil {
		return jsonRequestBody{}, false
	}

	mediaType := requestBody.Content.Get("application/json")
	if mediaType == nil || mediaType.Schema == nil {
		return jsonRequestBody{}, false
	}

	if mediaType.Schema.Value == nil {
		return jsonRequestBody{}, false
	}

	return jsonRequestBody{
		operation: operation,
		schema:    mediaType.Schema.Value,
		required:  requestBody.Required,
	}, true
}

// modelSchemasForRequestBody converts and names one request-body schema.
func modelSchemasForRequestBody(requestBody jsonRequestBody) ([]Schema, JSONRequestBodyOperation, error) {
	operationID := requestBody.operation.OperationID
	if operationID == "" {
		return nil, JSONRequestBodyOperation{}, fmt.Errorf("json request body operation has no operationId")
	}

	schema, err := SchemaFromOpenAPISchema(requestBody.schema)
	if err != nil {
		return nil, JSONRequestBodyOperation{}, fmt.Errorf("operation %q request body schema: %w", operationID, err)
	}

	schema.Base().EmptyBodyAllowed = !requestBody.required
	if schema.SchemaTypeName() == "" {
		name := requestBody.schema.Title
		if name == "" {
			name = operationID
		}

		schema.SetTypeName(exportedName(name))
	}

	definitions, err := namedSchemaDefinitions(schema)
	if err != nil {
		return nil, JSONRequestBodyOperation{}, fmt.Errorf("operation %q request body schema names: %w", operationID, err)
	}

	return definitions, JSONRequestBodyOperation{
		OperationID: operationID,
		TypeName:    schema.SchemaTypeName(),
	}, nil
}

// SchemaFromOpenAPISchema converts an OpenAPI schema to a generator schema.
func SchemaFromOpenAPISchema(schema *openapi3.Schema) (Schema, error) {
	if schema == nil {
		return nil, fmt.Errorf("openapi schema is nil")
	}

	base := BaseSchema{
		Nullable: schema.PermitsNull(),
	}
	if schema.Title != "" {
		base.Name = exportedName(schema.Title)
	}

	if len(schema.AllOf) != 0 {
		allOfSchema, err := allOfSchemaFromOpenAPISchema(schema)
		if err != nil {
			return nil, err
		}

		allOfSchema.BaseSchema = base

		return allOfSchema, nil
	}

	schemaType, err := schemaType(schema)
	if err != nil {
		return nil, err
	}

	return schemaFromType(schema, schemaType, base)
}

// schemaFromType converts a schema whose type has already been selected.
func schemaFromType(schema *openapi3.Schema, schemaType string, base BaseSchema) (Schema, error) {
	switch schemaType {
	case openapi3.TypeObject:
		objectSchema, err := objectSchemaFromOpenAPISchema(schema)
		if err != nil {
			return nil, err
		}

		objectSchema.BaseSchema = base

		return objectSchema, nil
	case openapi3.TypeArray:
		arraySchema, err := arraySchemaFromOpenAPISchema(schema)
		if err != nil {
			return nil, err
		}

		arraySchema.BaseSchema = base

		return arraySchema, nil
	case openapi3.TypeString:
		return &StringSchema{BaseSchema: base, Format: schema.Format}, nil
	case openapi3.TypeBoolean:
		return &BoolSchema{BaseSchema: base}, nil
	case openapi3.TypeNumber:
		return &NumberSchema{BaseSchema: base}, nil
	default:
		return nil, fmt.Errorf("unsupported schema type %q", schemaType)
	}
}

// allOfSchemaFromOpenAPISchema converts allOf members.
func allOfSchemaFromOpenAPISchema(schema *openapi3.Schema) (*AllOfSchema, error) {
	if len(schema.AllOf) == 0 {
		return nil, fmt.Errorf("allOf schema has no items")
	}

	allOfSchema := &AllOfSchema{
		Schemas: make([]Schema, 0, len(schema.AllOf)),
	}

	for i, schemaRef := range schema.AllOf {
		childSchema, err := schemaFromOpenAPISchemaRef(schemaRef)
		if err != nil {
			return nil, fmt.Errorf("allOf schema %d: %w", i+1, err)
		}

		err = setSchemaTypeNameIfEmpty(childSchema, fmt.Sprintf("AllOf%d", i+1))
		if err != nil {
			return nil, fmt.Errorf("allOf schema %d name: %w", i+1, err)
		}

		allOfSchema.Schemas = append(allOfSchema.Schemas, childSchema)
	}

	return allOfSchema, nil
}

// objectSchemaFromOpenAPISchema converts object properties and policies.
func objectSchemaFromOpenAPISchema(schema *openapi3.Schema) (*ObjectSchema, error) {
	objectSchema := &ObjectSchema{
		AdditionalProperties: true,
		Properties:           make([]ObjectFieldContext, 0, len(schema.Properties)),
	}

	if hasAdditionalProperties := schema.AdditionalProperties.Has; hasAdditionalProperties != nil {
		objectSchema.AdditionalProperties = *hasAdditionalProperties
	}

	if additionalPropertiesSchema := schema.AdditionalProperties.Schema; additionalPropertiesSchema != nil {
		additionalPropertiesObject, err := schemaFromOpenAPISchemaRef(additionalPropertiesSchema)
		if err != nil {
			return nil, fmt.Errorf("additionalProperties schema: %w", err)
		}

		err = setSchemaTypeNameIfEmpty(additionalPropertiesObject, "AdditionalProperty")
		if err != nil {
			return nil, fmt.Errorf("additionalProperties schema name: %w", err)
		}

		objectSchema.AdditionalProperties = true
		objectSchema.AdditionalPropertiesSchema = additionalPropertiesObject
	}

	required := make(map[string]struct{}, len(schema.Required))
	for _, propertyName := range schema.Required {
		required[propertyName] = struct{}{}
	}

	for _, propertyName := range slices.Sorted(maps.Keys(schema.Properties)) {
		propertySchema := schema.Properties[propertyName]

		propertyObject, err := schemaFromOpenAPISchemaRef(propertySchema)
		if err != nil {
			return nil, fmt.Errorf("property %q schema: %w", propertyName, err)
		}

		err = setSchemaTypeNameIfEmpty(propertyObject, propertyName)
		if err != nil {
			return nil, fmt.Errorf("property %q schema name: %w", propertyName, err)
		}

		_, isRequired := required[propertyName]
		objectSchema.Properties = append(objectSchema.Properties, ObjectFieldContext{
			PropertyName: propertyName,
			Schema:       propertyObject,
			Required:     isRequired,
		})
	}

	return objectSchema, nil
}

// arraySchemaFromOpenAPISchema converts an array item schema.
func arraySchemaFromOpenAPISchema(schema *openapi3.Schema) (*ArraySchema, error) {
	if schema.Items == nil {
		return nil, fmt.Errorf("array schema has no items")
	}

	items, err := schemaFromOpenAPISchemaRef(schema.Items)
	if err != nil {
		return nil, fmt.Errorf("array items schema: %w", err)
	}

	err = setSchemaTypeNameIfEmpty(items, "Item")
	if err != nil {
		return nil, fmt.Errorf("array items schema name: %w", err)
	}

	return &ArraySchema{
		Items: items,
	}, nil
}

// setSchemaTypeNameIfEmpty assigns name when schema is unnamed.
func setSchemaTypeNameIfEmpty(schema Schema, name string) error {
	if schema == nil {
		return fmt.Errorf("nil schema")
	}

	base := schema.Base()
	if base == nil {
		return fmt.Errorf("schema %T has nil base", schema)
	}

	if base.Name == "" {
		base.Name = exportedName(name)
	}

	return nil
}

// schemaFromOpenAPISchemaRef converts a resolved OpenAPI schema reference.
func schemaFromOpenAPISchemaRef(schemaRef *openapi3.SchemaRef) (Schema, error) {
	if schemaRef == nil {
		return nil, fmt.Errorf("openapi schema ref is nil")
	}

	if schemaRef.Value == nil {
		if schemaRef.Ref != "" {
			return nil, fmt.Errorf("openapi schema ref %q has no value", schemaRef.Ref)
		}

		return nil, fmt.Errorf("openapi schema ref has no value")
	}

	return SchemaFromOpenAPISchema(schemaRef.Value)
}

// schemaType returns the single supported non-null schema type.
func schemaType(schema *openapi3.Schema) (string, error) {
	if schema.Type == nil || schema.Type.IsEmpty() {
		return inferredSchemaType(schema)
	}

	schemaTypes := schema.Type.Slice()

	nonNullTypes := make([]string, 0, len(schemaTypes))
	for _, schemaType := range schemaTypes {
		if schemaType != openapi3.TypeNull {
			nonNullTypes = append(nonNullTypes, schemaType)
		}
	}

	if len(nonNullTypes) != 1 {
		return "", fmt.Errorf("unsupported schema types %v", schemaTypes)
	}

	return nonNullTypes[0], nil
}

// inferredSchemaType infers object and array schemas from structural keywords.
func inferredSchemaType(schema *openapi3.Schema) (string, error) {
	hasObjectKeywords := len(schema.Properties) != 0 ||
		len(schema.Required) != 0 ||
		schema.AdditionalProperties.Has != nil ||
		schema.AdditionalProperties.Schema != nil
	if hasObjectKeywords {
		return openapi3.TypeObject, nil
	}

	if schema.Items != nil {
		return openapi3.TypeArray, nil
	}

	return "", fmt.Errorf("openapi schema has no type")
}
