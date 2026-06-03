package contractor

type Contractor struct {
	ID   int
	Name string
	INN  string
}

type CreateInput struct {
	Name string
	INN  string
}
