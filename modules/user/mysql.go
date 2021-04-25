package user

import (
	"api-go-hexa/business/user/model"
	"database/sql"
)

type Repository interface {
	GetByID(id int) (*model.UserModel, error)
	UserRegister(u *model.UserModel) (*model.UserModel, error)
	UserLogin(u *model.UserLoginModel) (*model.UserLoginModel, error)
	Update(id int, u *model.UserModel) error
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (m *repository) GetByID(id int) (*model.UserModel, error) {
	userObj := new(model.UserModel)
	row := m.db.QueryRow("SELECT id, username, fullname, email from user WHERE id=?", id)
	err := row.Scan(&userObj.ID, &userObj.Username, &userObj.Fullname, &userObj.Email)
	if err != nil {
		return userObj, err
	}
	return userObj, nil
}

func (m *repository) UserRegister(u *model.UserModel) (*model.UserModel, error) {
	query := "INSERT user SET username=?, fullname=?, password=?, email=?"
	stmt, err := m.db.Prepare(query)
	if err != nil {
		return u, err
	}
	res, err := stmt.Exec(u.Username, u.Fullname, u.Password, u.Email)
	if err != nil {
		return u, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return u, err
	}
	u.ID = int(id)
	return u, nil

}

func (m *repository) UserLogin(u *model.UserLoginModel) (*model.UserLoginModel, error) {
	userObj := new(model.UserLoginModel)
	row := m.db.QueryRow("SELECT password FROM user WHERE username=?", u.Username)
	err := row.Scan(&userObj.Password)
	if err != nil {
		return userObj, err
	}
	return userObj, nil

}

func (m *repository) Update(id int, u *model.UserModel) error {
	query := "UPDATE user SET username=?, fullname=?, email=? WHERE id=?"
	stmt, err := m.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(u.Username, u.Fullname, u.Email, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return err
	}
	return nil

}

func (m *repository) Delete(id int) error {
	query := "DELETE FROM user WHERE id=?"
	stmt, err := m.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return err
	}
	return nil
}
