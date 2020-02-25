package repository

import (
	"database/sql"
	"psql/src/modules/profile/model"
)

type ProfileRepositoryPostgres struct{
	db *sql.DB
}



func NewProfileRepositoryPostgres(db *sql.DB) *ProfileRepositoryPostgres {
	return &ProfileRepositoryPostgres{db}
}
func (r *ProfileRepositoryPostgres) Save(profile *model.Profile) error{
	query := `INSERT INTO "profile"("ID","FirsName","LastName","Email","Password","CreatedAt","UpdatedAt")
	VALUES ($1, $2 ,$3 , $4, $5, $6 , $7)`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()
	_ , err = statement.Exec( profile.ID,profile.FirsName,profile.LastName,profile.Email,profile.Password,profile.CreatedAt,profile.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProfileRepositoryPostgres) Update (id string , profile model.Profile) error {
	query := `UPDATE "profile" SET "id"=$1 , "FirsName"=$2 , "LastName"=$3 , "Email"=$4 , "Password"=$5 , "CreatedAt"=$5 , "UpdatedAt"=$6 `

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()
	_ , err = statement.Exec( profile.FirsName,profile.LastName,profile.Email,profile.Password,profile.CreatedAt,profile.UpdatedAt,id)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProfileRepositoryPostgres) Delete (id string) error{
	query := `DELETE FROM "profile" WHERE "id" = $1 `

	statement ,err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_,  err = statement.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProfileRepositoryPostgres) FindById(id string) (*model.Profile ,error){
	query := `SELECT * FROM "profile" WHERE "id"=$1 `

	var profile model.Profile

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil,err
	}

	defer statement.Close()
	err = statement.QueryRow(id).Scan(&profile.ID, &profile.FirsName, &profile.LastName , &profile.Email, &profile.Password , &profile.CreatedAt, &profile.UpdatedAt)

	if err != nil {
		return nil,err
	}

	return &profile,nil

}

func (r *ProfileRepositoryPostgres) FindAll() (model.Profiles ,error) {
	query := `SELECT * FROM "profile"`

	var profiles model.Profiles

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next(){
		var profile model.Profile

		err = rows.Scan(&profile.ID, &profile.FirsName, &profile.LastName , &profile.Email , &profile.Password , &profile.CreatedAt , &profile.UpdatedAt)

		if err != nil {
			return nil,err
		}
		profiles = append(profiles , profile)


	}

	return profiles, nil


}



















