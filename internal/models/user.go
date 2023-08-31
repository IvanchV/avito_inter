package models

type ReqUser struct {
	Add    []string `json:"add"`
	Delete []string `json:"delete"`
}

type ResUser struct {
	Seg []string `json:"segments"`
}
