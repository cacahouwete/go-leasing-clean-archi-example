package dto

type Customer struct {
	Name string `binding:"required,max=255" maximum:"255"`
}
