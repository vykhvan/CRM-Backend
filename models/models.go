package models

type Customer struct {
	ID        string
	Name      string
	Role      string
	Email     string
	Phone     uint32
	Contacted bool
}
