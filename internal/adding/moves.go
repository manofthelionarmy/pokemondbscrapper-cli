package adding

// Move refers to the data we're adding to the relational database in this context
type Move struct {
	Name        string
	Type        string
	Category    string
	Power       int
	Accuracy    int
	PowerPoints int
	Effect      string
}
