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

func (class *Class) CreateClass(input model.NewClass) (*model.Class, error) {
	// todo 曜日と時限がかぶってる授業が既にデータベースにないか確認する(フロント側にまかせてもいいかも)
	timetableId, err := strconv.Atoi(input.TimetableID)
	if err != nil {
		log.Printf("action=create class, status=failed, err=%s", err)
		return nil, err
	}
	newClass := &Class{
		Name:        input.Name,
		Days:        input.Day,
		Periods:     input.Period,
		TimetableID: uint(timetableId),
	}
	if input.Style != nil {
		newClass.Style = *input.Style
	}
	if input.Teacher != nil {
		newClass.Teacher = *input.Teacher
	}
	if input.RoomOrURL != nil {
		newClass.RoomOrUrl = *input.RoomOrURL
	}
	result := database.Db.Create(newClass)
	if result.Error != nil {
		log.Printf("action=create class, status=failed, err=%s", result.Error)
		return nil, result.Error
	}
	graphClass := ConvertClassFromDbToGraph(newClass)

	log.Printf("action=create class, status=success")
	return graphClass, nil
}

func (class *Class) UpdateClass(input model.UpdateClass) (*model.Class, error) {
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
		log.Printf("action=update class, status=failed, err=update data is not set")
		return nil, fmt.Errorf("update data must be set")
	}

	id, err := strconv.Atoi(input.ID)
	if err != nil {
		log.Printf("action=update class, status=failed, err=%s", err)
		return nil, err
	}
	dbClass := new(Class)
	dbClass.ID = uint(id)

	result := database.Db.Model(dbClass).Updates(updateData)
	if result.Error != nil {
		log.Printf("action=update class, status=failed, err=%s", result.Error)
		return nil, result.Error
	}
	graphClass := ConvertClassFromDbToGraph(dbClass)

	log.Printf("action=update class, status=success")
	return graphClass, nil
}

func (class *Class) DeleteClass(input string) (bool, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		log.Printf("action=delete class, status=failed, err=%s", err)
		return false, err
	}
	dbClass := new(Class)
	dbClass.ID = uint(id)

	result := database.Db.Delete(dbClass)
	if result.Error != nil {
		log.Printf("action=delete class, status=failed, err=%s", result.Error)
		return false, result.Error
	}

	log.Printf("action=delete class, status=success")
	return true, nil
}

func FetchClassesByTimetableId(timetableId int) ([]*Class, error) {
	var classes []*Class
	result := database.Db.Where("timetable_id = ?", timetableId).Find(&classes)
	if result.Error != nil {
		return nil, result.Error
	}
	return classes, nil
}

func ConvertClassFromDbToGraph(dbClass *Class) *model.Class {
	graphClass := &model.Class{
		ID:        strconv.Itoa(int(dbClass.ID)),
		Name:      dbClass.Name,
		Day:       dbClass.Days,
		Period:    dbClass.Periods,
		Color:     dbClass.Color,
		Style:     dbClass.Style,
		Teacher:   dbClass.Teacher,
		Credit:    &dbClass.Credit,
		Memo:      dbClass.Memo,
		RoomOrURL: dbClass.RoomOrUrl,
	}
	return graphClass
}

func ConvertClassesFromDbToGraph(dbClasses []*Class) []*model.Class {
	var graphClasses []*model.Class
	for _, dbClass := range dbClasses {
		graphClass := ConvertClassFromDbToGraph(dbClass)
		graphClasses = append(graphClasses, graphClass)
	}
	return graphClasses
}

func GetGraphClasses(timetableId int) []*model.Class {
	dbClasses, err := FetchClassesByTimetableId(timetableId)
	if err != nil {
		log.Printf("action=fetch classes data, status=failed, err=%s", err)
		return nil
	}
	return ConvertClassesFromDbToGraph(dbClasses)
}
