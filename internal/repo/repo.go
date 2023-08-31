package repo

import "database/sql"

type Repo struct {
	*UserRepo
	*SegmentRepo
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		UserRepo:    NewUserRepo(db),
		SegmentRepo: NewSegmentRepo(db),
	}
}
