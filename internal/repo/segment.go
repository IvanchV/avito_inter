package repo

import (
	"avito/internal/models"
	"database/sql"
	"fmt"
)

type SegmentRepo struct {
	db *sql.DB
}

func NewSegmentRepo(db *sql.DB) *SegmentRepo {
	return &SegmentRepo{db: db}
}

func (s *SegmentRepo) CreateSegment(segment *models.Segment) error {
	query := `INSERT INTO segments(name) VALUES ($1);`
	_, err := s.db.Query(query, segment.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *SegmentRepo) DeleteSegment(name string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()
	check, err := tx.Query(`select name from segments`)
	chek := make(map[string]struct{})
	for check.Next() {
		var tmp string
		err = check.Scan(&tmp)
		if err != nil {
			return err
		}
		chek[tmp] = struct{}{}
	}
	if _, ok := chek[name]; !ok {
		return fmt.Errorf("error: name of seggment is not exists")
	}
	query := `DELETE FROM segments WHERE name = $1`
	_, err = s.db.Exec(query, name)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
