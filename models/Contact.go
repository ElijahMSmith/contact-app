package models

type Contact struct {
	Name Name
	Age  int
}

type Name struct {
	FirstName  string
	MiddleName string
	LastName   string
}
