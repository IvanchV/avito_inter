package models

type Segment struct {
	Name string `json:"name"`
}

func NewSegment(name string) *Segment {
	return &Segment{Name: name}
}
