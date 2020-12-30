package models

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mi11km/zikanwarikun-back/graph/model"
	database "github.com/mi11km/zikanwarikun-back/internal/db"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name        string
	Days        int
	Periods     int
	Style       string
	RoomOrUrl   string
	Teacher     string
	Credit      int
	Memo        *string
	Color       string
	TimetableID uint
	Todos       []*Todo
	Attendances []*Attendance
	Urls        []*Url
}

func (class *Class) Create(input model.NewClass) error {
	// todo? 曜日と時限がかぶってる授業が既にデータベースにないか確認する(フロント側にまかせてもいいかも)
	timetableId, err := strconv.Atoi(input.TimetableID)
	if err != nil {
		log.Printf("action=create class data, status=failed, err=%s", err)
		return err
	}
	class.Name = input.Name
	class.Days = input.Day
	class.Periods = input.Period
	class.TimetableID = uint(timetableId)
	if input.Style != nil {
		class.Style = *input.Style
	}
	if input.Teacher != nil {
		class.Teacher = *input.Teacher
	}
	if input.RoomOrURL != nil {
		class.RoomOrUrl = *input.RoomOrURL
	}
	if err := database.Db.Create(class).Error; err != nil {
		log.Printf("action=create class data, status=failed, err=%s", err)
		return err
	}
	log.Printf("action=create class data, status=success")
	return nil
}

func (class *Class) Update(input model.UpdateClass) error {
	updateData := make(map[string]interface{})
	if input.Name != nil {
		updateData["name"] = *input.Name
	}
	if input.Color != nil {
		updateData["color"] = *input.Color
	}
	if input.Style != nil {
		updateData["style"] = *input.Style
	}
	if input.Teacher != nil {
		updateData["teacher"] = *input.Teacher
	}
	if input.Credit != nil {
		updateData["credit"] = *input.Credit
	}
	if input.Memo != nil {
		updateData["memo"] = *input.Memo
	}
	if input.RoomOrURL != nil {
		updateData["room_or_url"] = *input.RoomOrURL
	}
	if len(updateData) == 0 {
		log.Printf("action=update class data, status=failed, err=update data is not set")
		return fmt.Errorf("update data must be set")
	}

	if err := database.Db.Model(class).Updates(updateData).Error; err != nil {
		log.Printf("action=update class data, status=failed, err=%s", err)
		return err
	}
	log.Printf("action=update class data, status=success")
	return nil
}

func (class *Class) Delete(input string) (bool, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		log.Printf("action=delete class data, status=failed, err=%s", err)
		return false, err
	}
	class.ID = uint(id)

	if err := database.Db.Select("Todo", "Attendance", "Url").Delete(class).Error; err != nil {  // todo 関連レコードも削除できてるか確認
		log.Printf("action=delete class data, status=failed, err=%s", err)
		return false, err
	}
	log.Printf("action=delete class data, status=success")
	return true, nil
}

func FetchClassById(id string) *Class {
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("action=fetch class data by id, status=failed, err=%s", err)
		return nil
	}
	class := &Class{}
	class.ID = uint(i)
	if err := database.Db.First(class).Error; err != nil {
		log.Printf("action=fetch class data by id, status=failed, err=%s", err)
		return nil
	}
	return class
}

/* FetchClassesByTimetable 指定した時間割の授業・科目データを全て取得する */
func FetchClassesByTimetable(timetable Timetable) ([]*Class, error) {
	var classes []*Class
	if err := database.Db.Where("timetable_id = ?", timetable.ID).Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}

func SetClassesToEachTimetable(timetables []*Timetable) {
	if len(timetables) == 0 {
		return
	}
	for _, t := range timetables {
		classes, err := FetchClassesByTimetable(*t)
		if err == nil {
			t.Classes = classes
		}
	}
}
