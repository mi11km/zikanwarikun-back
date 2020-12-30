package models

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mi11km/zikanwarikun-back/graph/model"
	database "github.com/mi11km/zikanwarikun-back/internal/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Timetable struct {
	gorm.Model
	Name         string
	ClassDays    int
	ClassPeriods int
	IsDefault    bool
	Users        []User `gorm:"many2many:user_timetables;"`
	ClassTimes   []*ClassTime
	Classes      []*Class
}

const timetablesAssociation = "Timetables"

/* Create 時間割を新しくDBに作成してis_default=trueとなる時間割を新しいやつ１つにする */
func (t *Timetable) Create(input model.NewTimetable, user User) error {
	if err := ChangeIsDefaultToFalse(user); err != nil {
		log.Printf("action=create timetable data, status=failed, err=%s", err)
		return err
	}
	t.Name = input.Name
	t.ClassDays = input.Days
	t.ClassPeriods = input.Periods
	t.IsDefault = true
	if err := database.Db.Model(&user).Association(timetablesAssociation).Append(t); err != nil {
		log.Printf("action=create timetable data, status=failed, err=%s", err)
		return err
	}
	log.Printf("action=create timetable data, status=success")
	return nil
}

func (t *Timetable) Update(input model.UpdateTimetable, user User) error {
	updateData := make(map[string]interface{})
	if input.Name != nil {
		updateData["name"] = *input.Name
	}
	if input.Days != nil {
		updateData["class_days"] = *input.Days
	}
	if input.Periods != nil {
		updateData["class_periods"] = *input.Periods
	}
	if input.IsDefault != nil {
		updateData["is_default"] = *input.IsDefault
		if *input.IsDefault {
			if err := ChangeIsDefaultToFalse(user); err != nil {
				log.Printf("action=update timetable data, status=failed, err=%s", err)
				return err
			}
		}
	}
	if len(updateData) == 0 {
		log.Printf("action=update timetable data, status=failed, err=update data is not set")
		return fmt.Errorf("update data must be set")
	}

	if err := database.Db.Model(t).Updates(updateData).Error; err != nil {
		log.Printf("action=update timetable data, status=failed, err=%s", err)
		return err
	}
	log.Printf("action=update timetable data, status=success")
	return nil
}

func (t *Timetable) Delete(input string) (bool, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		log.Printf("action=delete timetable data, status=failed, err=%s", err)
		return false, err
	}
	t.ID = uint(id)

	if err := database.Db.Select("ClassTime", "Class", clause.Associations).Delete(t).Error; err != nil {
		log.Printf("action=delete timetable data, status=failed, err=%s", err)
		return false, err
	}
	log.Printf("action=delete timetable data, status=success")
	return true, nil
}

func FetchTimetableById(id string) *Timetable {
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("action=fetch timetable by id, status=failed, err=%s", err)
		return nil
	}
	timetable := &Timetable{}
	timetable.ID = uint(i)
	if err := database.Db.First(timetable).Error; err != nil {
		log.Printf("action=fetch timetable by id, status=failed, err=%s", err)
		return nil
	}
	return timetable
}

/* FetchDefaultTimetableByUserId userからis_defaultがtrueになってる時間割データを取得する(最初の一件分) */
func FetchDefaultTimetableByUser(user User) (*Timetable, error) {
	var defaultTimetable Timetable
	if err := database.Db.Model(&user).Where("is_default = ?", true).
		Association(timetablesAssociation).Find(&defaultTimetable); err != nil {
		return nil, err
	}
	return &defaultTimetable, nil
}

/* FetchTimetablesByUser userの時間割データを全て取得する */
func FetchTimetablesByUser(user User) ([]*Timetable, error) {
	var timetables []*Timetable
	if err := database.Db.Model(&user).Association(timetablesAssociation).Find(&timetables); err != nil {
		return nil, err
	}
	return timetables, nil
}

/* userのtimetableの中でis_defaultがtrueになっているものをfalseにする */
func ChangeIsDefaultToFalse(user User) error {
	timetables, err := FetchTimetablesByUser(user)
	if err == nil {
		for _, t := range timetables {
			if t.IsDefault == true {
				database.Db.Model(t).Updates(map[string]interface{}{"is_default": false})
			}
		}
		return nil
	}
	return err
}
