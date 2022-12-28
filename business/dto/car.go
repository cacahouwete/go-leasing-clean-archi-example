package dto

type Car struct {
	Name string `binding:"required,max=255" maximum:"255"`
}
