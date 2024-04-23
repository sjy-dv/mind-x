package entity

type InsertObject struct {
	ID       string
	Vector   []float32
	Metadata map[string]string
}
