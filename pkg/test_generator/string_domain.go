package testgenerator

type StringDomain struct {
	Nullable bool

	Enum []string

	Pattern *string
	Format  *string

	MinLength int
	MaxLength *int
}
