package models

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/mi11km/zikanwarikun-back/graph/model"
	database "github.com/mi11km/zikanwarikun-back/internal/db"
	parsetime "github.com/mi11km/zikanwarikun-back/pkg/time"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Kind       string
	Deadline   time.Time
	IsDone     bool
	Memo       *string
	IsRepeated bool
	ClassID    uint
}

func (todo *Todo) Create(input model.NewTodo) error {
	id, err := strconv.Atoi(input.ClassID)
	if err != nil {
		log.Printf("action=create todo data, status=failed, err=%s", err)
		return err
	}
	deadline := parsetime.StringToTime(input.Deadline)
	if deadline == nil {
		err = fmt.Errorf("failed to parse deadline value")
		log.Printf("action=create todo data, status=failed, err=%s", err)
		return err
	}
	todo.Kind = input.Kind
	todo.Deadline = *deadline
	todo.IsDone = false
	todo.IsRepeated = input.IsRepeated
	todo.ClassID = uint(id)
	if err := database.Db.Create(todo).Error; err != nil {
		log.Printf("action=create todo data, status=failed, err=%s", err)
		return err
	}
	log.Printf("action=create todo data, status=success")
	return nil
}

func (todo *Todo) Update(input model.UpdateTodo) error {
	updateData := make(map[string]interface{})
	if input.Kind != nil && *input.Kind != todo.Kind {
		updateData["kind"] = *input.Kind
	}
	if input.Deadline != nil {
		deadline := parsetime.StringToTime(*input.Deadline)
		if deadline == nil {
			err := fmt.Errorf("failed to parse deadline value")
			log.Printf("action=create todo data, status=failed, err=%s", err)
			return err
		}
		if *deadline != todo.Deadline {
			updateData["deadline"] = *deadline
		}
	}
	if input.IsDone != nil && *input.IsDone != todo.IsDone {
		updateData["is_done"] = *input.IsDone
	}
	if input.Memo != nil && input.Memo != todo.Memo {
		updateData["memo"] = *input.Memo
	}
	if input.IsRepeated != nil && *input.IsRepeated != todo.IsRepeated {
		updateData["is_repeated"] = *input.IsRepeated
	}
	if len(updateData) == 0 {
		log.Printf("action=update todo data, status=failed, err=update data is not set or the only same data id set")
		return fmt.Errorf("update data must be set or the only same data id set")
	}

	if err := database.Db.Model(todo).Updates(updateData).Error; err != nil {
		log.Printf("action=update todo data, status=failed, err=%s", err)
		return err
	}
	log.Printf("action=update todo data, status=success")
	return nil
}

func (todo *Todo) Delete(input string) (bool, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		log.Printf("action=delete todo data, status=failed, err=%s", err)
		return false, err
	}
	todo.ID = uint(id)
	if err := database.Db.Delete(todo).Error; err != nil {
		log.Printf("action=delete todo data, status=failed, err=%s", err)
		return false, err
	}
	log.Printf("action=delete todo data, status=success")
	return true, nil
}

func FetchTodoById(id string) *Todo {
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("action=fetch todo data by id, status=failed, err=%s", err)
		return nil
	}
	todo := &Todo{}
	todo.ID = uint(i)
	if err := database.Db.First(todo).Error; err != nil {
		log.Printf("action=fetch todo data by id, status=failed, err=%s", err)
		return nil
	}
	return todo
}

/* FetchTodosByClass 指定した科目のTodoデータを全て取得する */
func FetchTodosByClass(class Class) ([]*Todo, error) {
	var todos []*Todo
	if err := database.Db.Where("class_id = ?", class.ID).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

/* GetTodosFromTimetableId 時間割のidからその時間割のtodo一覧を取得する */
func GetTodosFromTimetableId(timetableId string) ([]*Todo, error) {
	var allTodos []*Todo
	id, err := strconv.Atoi(timetableId)
	if err != nil {
		log.Printf("action=get todos from timetable id, status=failed, err=%s", err)
		return nil, err
	}
	timetable := &Timetable{}
	timetable.ID = uint(id)
	classes, err := FetchClassesByTimetable(*timetable)
	for _, class := range classes {
		todos, err := FetchTodosByClass(*class)
		if err != nil {
			log.Printf("action=get todos from timetable id, status=failed, err=%s", err)
			return nil, err
		}
		allTodos = append(allTodos, todos...)
	}
	// todo 締め切り近い順でsort + is_done=falseのみを返す？
	return allTodos, nil
}
