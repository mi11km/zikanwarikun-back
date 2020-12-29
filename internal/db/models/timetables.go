package models

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mi11km/zikanwarikun-back/graph/model"
	database "github.com/mi11km/zikanwarikun-back/internal/db"
	"gorm.io/gorm"
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

/* CreateTimetable 時間割を新しくDBに作成してそれをデフォルトにする */
func (t *Timetable) CreateTimetable(input model.NewTimetable, user User) (*model.Timetable, error) {
	// todo userのtimetableの中でtrueになっているものをfalseにする
	//database.Db.Model(&Timetable{}).Where("is_default = ?", true).Update("is_default", false)

	newTimetable := &Timetable{
		Name:         input.Name,
		ClassDays:    input.Days,
		ClassPeriods: input.Periods,
		IsDefault:    true,
		Users:        []User{user},
	}
	result := database.Db.Create(newTimetable)
	if result.Error != nil {
		log.Printf("action=create timetable, status=failed, err=%s", result.Error)
		return nil, result.Error
	}

	dbTimetables, err := FetchTimetablesByUserId(user.ID)
	if err != nil {
		log.Printf("action=create timetable, status=failed, err=%s", err)
		return nil, err
	}
	graphUser := &model.User{
		ID:     user.ID,
		Email:  user.Email,
		School: user.School,
		Name:   user.Name,
	}
	graphTimetables := ConvertTimetablesFromDbToGraph(dbTimetables, graphUser)
	graphUser.Timetables = graphTimetables

	graphTimetable := ConvertTimetableFromDbToGraph(newTimetable, graphUser)

	log.Printf("action=create timetable, status=success")
	return graphTimetable, nil
}

func (t *Timetable) UpdateTimetable(input model.UpdateTimetable, user User) (*model.Timetable, error) {
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
			database.Db.Model(&Timetable{}).
				Where("is_default = ?", true).Update("is_default", false)
		}
	}
	if len(updateData) == 0 {
		log.Printf("action=update dbTimetable, status=failed, err=update data is not set")
		return nil, fmt.Errorf("update data must be set")
	}

	id, err := strconv.Atoi(input.ID)
	if err != nil {
		log.Printf("action=update timetable, status=failed, err=%s", err)
		return nil, err
	}
	dbTimetable := new(Timetable)
	dbTimetable.ID = uint(id)

	result := database.Db.Model(dbTimetable).Updates(updateData)
	if result.Error != nil {
		log.Printf("action=update timetable, status=failed, err=%s", result.Error)
		return nil, result.Error
	}
	graphUser := &model.User{
		ID:     user.ID,
		Email:  user.Email,
		School: user.School,
		Name:   user.Name,
	}
	graphTimetable := ConvertTimetableFromDbToGraph(dbTimetable, graphUser)

	// todo? userにtimetablesも入れとくべきか。入れても使わない気がする

	log.Printf("action=update timetable, status=success")
	return graphTimetable, nil // todo? updateしたデータとidとupdatedAt以外空になってる。
}

// todo 関連レコードも一括削除する
func (t *Timetable) DeleteTimetable(input string) (bool, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		log.Printf("action=delete timetable, status=failed, err=%s", err)
		return false, err
	}
	dbTimetable := new(Timetable)
	dbTimetable.ID = uint(id)

	result := database.Db.Delete(dbTimetable)
	if result.Error != nil {
		log.Printf("action=delete timetable, status=failed, err=%s", result.Error)
		return false, result.Error
	}

	log.Printf("action=delete timetable, status=success")
	return true, nil
}

/* FetchTimetablesByUserId user_idからdbの時間割データを全て取得する */
func FetchTimetablesByUserId(userId string) ([]*Timetable, error) {
	var timetables []*Timetable
	result := database.Db.Where("user_id = ?", userId).Find(&timetables)
	if result.Error != nil {
		return nil, result.Error
	}
	return timetables, nil
}

/* FetchDefaultTimetableByUserId user_idからdbのデフォルトがtrueになってる時間割データを取得する */
func FetchDefaultTimetableByUserId(userId string) (*Timetable, error) {
	var defaultTimetable Timetable
	result := database.Db.Where("user_id = ? AND is_default = ?", userId, true).First(&defaultTimetable)
	if result.Error != nil {
		return nil, result.Error
	}
	return &defaultTimetable, nil
}

/* ConvertTimetableFromDbToGraph １つ、dbの時間割データをgraphql用のモデルに変換する */
func ConvertTimetableFromDbToGraph(dbTimetable *Timetable, user *model.User) *model.Timetable {
	graphTimetable := &model.Timetable{
		ID:        strconv.Itoa(int(dbTimetable.ID)),
		Name:      dbTimetable.Name,
		Days:      dbTimetable.ClassDays,
		Periods:   dbTimetable.ClassPeriods,
		CreatedAt: dbTimetable.CreatedAt.String(),
		UpdatedAt: dbTimetable.UpdatedAt.String(),
		IsDefault: dbTimetable.IsDefault,
		Classes:   GetGraphClasses(int(dbTimetable.ID)),
		//Classtimes: , // todo
		//RowData: ,
		User: user,
	}
	return graphTimetable
}

/* ConvertTimetablesFromDbToGraph 複数のdbの時間割データをgraphql用のモデルに変換する */
func ConvertTimetablesFromDbToGraph(dbTimetables []*Timetable, user *model.User) []*model.Timetable {
	var graphTimetables []*model.Timetable
	for _, dbTimetable := range dbTimetables {
		graphTimetable := ConvertTimetableFromDbToGraph(dbTimetable, user)
		graphTimetables = append(graphTimetables, graphTimetable)
	}
	return graphTimetables
}
