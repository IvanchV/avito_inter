package repo

import (
	"database/sql"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (s *UserRepo) ChangeUserSegment(id int, addSegments []string, deleteSegments []string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()
	_, err = tx.Exec(`insert into users(user_id) select $1 where not exists(select user_id from users where user_id = $2)`, id, id)
	if err != nil {
		return err
	}
	buff, err := tx.Query(`select name from segments`)
	res := make(map[string]struct{})
	for buff.Next() {
		var tmp string
		err = buff.Scan(&tmp)
		if err != nil {
			return err
		}
		res[tmp] = struct{}{}
	}
	rows, err := tx.Query(`select name from usersegment where user_id=$1`, id)
	if err != nil {
		return err
	}
	res1 := make(map[string]struct{})
	for rows.Next() {
		var tmp string
		err = rows.Scan(&tmp)
		if err != nil {
			return err
		}
		res1[tmp] = struct{}{}
	}
	for _, name := range addSegments {
		if _, ok := res[name]; !ok {
			continue
		}
		if _, ok := res1[name]; ok {
			break
		}
		_, err = tx.Exec(`insert into usersegment (user_id, name) values ($1, $2)`, id, name)
		if err != nil {
			return err
		}
	}

	for _, name := range deleteSegments {
		if _, ok := res[name]; !ok {
			continue
		}
		_, err = tx.Exec(`delete from usersegment where user_id=$1 and name=$2`, id, name)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s *UserRepo) GetUserSegment(id int) ([]string, error) {
	rows, err := s.db.Query(`select name from usersegment where user_id=$1`, id)
	if err != nil {
		return nil, err
	}
	res := make([]string, 0)
	for rows.Next() {
		var tmp string
		err = rows.Scan(&tmp)
		if err != nil {
			return nil, err
		}
		res = append(res, tmp)
	}
	return res, nil
}
