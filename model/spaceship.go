package model

import (
	"starwars_api/database"

	"gorm.io/gorm"
)

type Spaceship struct {
	gorm.Model
	ID       uint       `gorm:"primaryKey"`
	Name     string     `json:"name"`
	Class    Class      `json:"class"`
	Crew     int32      `json:"crew"`
	Image    string     `json:"image"`
	Value    int32      `json:"value"`
	Status   string     `json:"status"`
	Armanent []Armanent `json:"armanent"`
}

func (spaceship *Spaceship) Save() error {

	err := database.Database.Create(&spaceship).Error
	if err != nil {
		return err
	}
	return nil
}

func (spaceship *Spaceship) Update() error {

	err := database.Database.Save(&spaceship).Error

	if err != nil {
		return err
	}

	return nil
}

func FindByShipName(name string) (Spaceship, error) {
	var list Spaceship

	err := database.Database.Where("name=?", name).Find(&list).Error
	if err != nil {
		return Spaceship{}, err
	}
	return list, nil
}

func Delete(id string) error {

	err := database.Database.Delete(&Spaceship{}, id).Error

	if err != nil {
		return err
	}
	return nil
}
