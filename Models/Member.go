package Models

import (
	"fmt"

	"../Config"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllMembers(member *[]Member) (err error) {
	if err = Config.DB.Find(member).Error; err != nil {
		return err
	}
	return nil
}

func CreateAMember(member *Member) (err error) {
	if err = Config.DB.Create(member).Error; err != nil {
		return err
	}
	return nil
}

func GetAMember(member *Member, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(member).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAMember(member *Member, id string) (err error) {
	fmt.Println(member)
	Config.DB.Save(member)
	return nil
}

func DeleteAMember(member *Member, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(member)
	return nil
}
