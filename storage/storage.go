package storage

import (
	"database/sql"
	"time"
)

// User steps
const (
	EnterFullnameStep    string = "enter_fullname"
	EnterPhoneNumberStep string = "enter_phone_number"
	RegisteredStep       string = "registered"
)

type User struct {
	TgID      int64
	Firstname string
	Lastname  string
	City      string
	CreatedAt *time.Time
}

type Request struct {
	City      string
	CreatedAt string
}

type StorageI interface {
	CreateRequest(u *User) (*User, error)
	CreateUser(u *User) (*User, error)
	GetFirstRequest(id int64) (*Request, error)
	GetAllRequests(tgID int64) ([]*Request, error)
	CheckUserExistence(tgID int64) (bool, error)
}

type storagePg struct {
	db *sql.DB
}

func NewStoragePg(db *sql.DB) StorageI {
	return &storagePg{
		db: db,
	}
}

func (s *storagePg) CreateUser(user *User) (*User, error) {
	query := `
		INSERT INTO users(
			tg_id,
			first_name,
			last_name
		) VALUES($1, $2, $3)`

	row := s.db.QueryRow(
		query,
		user.TgID,
		user.Firstname,
		user.Lastname,
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return user, nil
}

func (s *storagePg) CreateRequest(user *User) (*User, error) {
	query := `
		INSERT INTO requests(
			city,
			user_id
		) VALUES($1, $2)`

	row := s.db.QueryRow(
		query,
		user.City,
		user.TgID,
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return user, nil
}

func (s *storagePg) GetFirstRequest(id int64) (*Request, error) {
	var (
		createdAt time.Time
		result    Request
	)
	query := `
		SELECT 
			city,
			created_at
		FROM requests WHERE user_id = $1 
		ORDER BY created_at ASC LIMIT 1`

	row := s.db.QueryRow(query, id)
	err := row.Scan(
		&result.City,
		&createdAt,
	)
	if err != nil {
		return nil, err
	}
	result.CreatedAt = createdAt.Format("2 января 2006 г. 15:04")
	return &result, nil
}

func (s *storagePg) GetAllRequests(tgID int64) ([]*Request, error) {
	query := `
	SELECT 
		city,
		created_at
	FROM requests WHERE user_id = $1 ORDER BY created_at DESC`

	rows, err := s.db.Query(query, tgID)
	if err != nil {
		return nil, err
	}

	requests := []*Request{}
	for rows.Next() {
		var (
			createdAt time.Time
			request   Request
		)
		err = rows.Scan(
			&request.City,
			&createdAt,
		)
		if err != nil {
			return nil, err
		}
		request.CreatedAt = createdAt.Format("2 января 2006 г. 15:04")
		requests = append(requests, &request)

	}

	return requests, nil
}

func (s *storagePg) CheckUserExistence(tgID int64) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE tg_id = $1)`
	var exists bool
	err := s.db.QueryRow(query, tgID).Scan(&exists)
	if err != nil || err == sql.ErrNoRows {
		return false, err
	}

	return exists, nil
}
