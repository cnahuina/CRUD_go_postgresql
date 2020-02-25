package main

import (
	"fmt"
	"psql/config"
	"psql/src/modules/profile/model"
	"psql/src/modules/profile/repository"
)

func main(){
	fmt.Println("go postgressql")

	db , err := config.GetPostgresDB()

	if err != nil {
		fmt.Println(err)
	}
	insertando := model.NewProfile()
	insertando.ID = "P2"
	insertando.FirsName = "Rob"
	insertando.LastName = "Pike"
	insertando.Email = "rob@yahoo.com"
	insertando.Password = "123456"
	fmt.Println(insertando)

	profileRepositoryPostgres := repository.NewProfileRepositoryPostgres(db)

	//profiles, err := getProfiles(profileRepositoryPostgres)

	err = saveProfile(insertando, profileRepositoryPostgres)


	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("=======================")

	//for _, v := range profiles {
	//	fmt.Println(v)
	//}

}

func saveProfile(p *model.Profile, repo *repository.ProfileRepositoryPostgres) error {
	err := repo.Save(p)
	if err != nil {
		return err
	}

	return nil

}

func updateProfile(p *model.Profile, repo repository.ProfileRepository) error{
	err := repo.Update(p.ID, p)

	if err != nil {
		return err
	}

	return nil
}

func deleteProfile(id string, repo repository.ProfileRepository) error{
	err := repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func getProfile(id string, repo repository.ProfileRepository) (*model.Profile, error) {
	profile, err := repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

func getProfiles(repo *repository.ProfileRepositoryPostgres) (model.Profiles, error) {
	profiles, err := repo.FindAll()

	if err != nil {
		return nil, err
	}

	return profiles, nil
}