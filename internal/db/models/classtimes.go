package models

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mi11km/zikanwarikun-back/graph/model"
	database "github.com/mi11km/zikanwarikun-back/internal/db"
	"gorm.io/gorm"
)

type ClassTime struct {
	gorm.Model
	Periods     int
	StartTime   string
	EndTime     string
	TimetableID uint
}

// todo 時刻計算をバックエンドでやるべきかわからない。現在単純なstring型で管理保存。

func (ct *ClassTime) Create(input model.NewClassTime) error {
	// todo? periodsとtimetable_idが完全一致するものがないか確認(フロント側にまかせてもいいかも)
	id, err := strconv.Atoi(input.TimetableID)
	if err != nil {
		log.Printf("action=create class_time data, status=failed, err=%s", err)
		return err
	}
	ct.Periods = input.Period
	ct.StartTime = input.StartTime
	ct.EndTime = input.EndTime
	ct.TimetableID = uint(id)
	if err := database.Db.Create(ct).Error; err != nil {
		log.Printf("action=create class_time data, status=failed, err=%s", err)
		return err
	}
	log.Printf("action=create class_time data, status=success")
	return nil
}
func (ct *ClassTime) Update(input model.UpdateClassTime) error {
	updateData := make(map[string]interface{})
	if input.StartTime != nil {
		updateData["start_time"] = *input.StartTime
	}
	if input.EndTime != nil {
		updateData["end_time"] = *input.EndTime
	}
	if len(updateData) == 0 {
		log.Printf("action=update class_time data, status=failed, err=update data is not set")
		return fmt.Errorf("update data must be set")
	}

	if err := database.Db.Model(ct).Updates(updateData).Error; err != nil {
		log.Printf("action=update class_time data, status=failed, err=%s", err)
		return err
	}
	log.Printf("action=update class_time data, status=success")
	return nil
}

func FetchClassTimeById(id string) *ClassTime {
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("action=fetch class_time by id, status=failed, err=%s", err)
		return nil
	}
	classTime := &ClassTime{}
	classTime.ID = uint(i)
	if err := database.Db.First(classTime).Error; err != nil {
		log.Printf("action=fetch class_time by id, status=failed, err=%s", err)
		return nil
	}
	return classTime
}

/* FetchClassTimesByTimetable 指定した時間割の時限ごとの授業時間データを全て取得する */
func FetchClassTimesByTimetable(timetable Timetable) ([]*ClassTime, error) {
	var classTimes []*ClassTime
	if err := database.Db.Where("timetable_id = ?", timetable.ID).Find(&classTimes).Error; err != nil {
		return nil, err
	}
	return classTimes, nil
}

func SetClassTimesToEachTimetable(timetables []*Timetable) {
	if len(timetables) == 0 {
		return
	}
	for _, t := range timetables {
		classTimes, err := FetchClassTimesByTimetable(*t)
		if err == nil {
			t.ClassTimes = classTimes
		}
	}
}
