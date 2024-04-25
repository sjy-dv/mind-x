package entity

type InsertObject struct {
	ID       []byte
	Vector   []float32
	Metadata map[string]string
}
