package models

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mi11km/zikanwarikun-back/graph/model"
	database "github.com/mi11km/zikanwarikun-back/internal/db"
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Name    string
	Url     string
	ClassID uint
}

func (url *Url) Create(input model.NewURL) error {
	id, err := strconv.Atoi(input.ClassID)
	if err != nil {
		log.Printf("action=create url data, status=failed, err=%s", err)
		return err
	}
	url.Name = input.Name
	url.Url = input.URL
	url.ClassID = uint(id)
	if err := database.Db.Create(url).Error; err != nil {
		log.Printf("action=create url data, status=failed, err=%s", err)
		return err
	}
	log.Printf("action=create url data, status=success")
	return nil
}

func (url *Url) Update(input model.UpdateURL) error {
	updateData := make(map[string]interface{})
	if input.Name != nil && *input.Name != url.Name {
		updateData["name"] = *input.Name
	}
	if input.URL != nil && *input.URL != url.Url {
		updateData["url"] = *input.URL
	}
	if len(updateData) == 0 {
		log.Printf("action=update url data, status=failed, err=update data is not set or the only same data id set")
		return fmt.Errorf("update data must be set or the only same data id set")
	}

	if err := database.Db.Model(url).Updates(updateData).Error; err != nil {
		log.Printf("action=update url data, status=failed, err=%s", err)
		return err
	}
	log.Printf("action=update url data, status=success")
	return nil
}

func (url *Url) Delete(input string) (bool, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		log.Printf("action=delete url data, status=failed, err=%s", err)
		return false, err
	}
	url.ID = uint(id)
	if err := database.Db.Delete(url).Error; err != nil {
		log.Printf("action=delete url data, status=failed, err=%s", err)
		return false, err
	}
	log.Printf("action=delete url data, status=success")
	return true, nil
}

func FetchUrlById(id string) *Url {
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("action=fetch url data by id, status=failed, err=%s", err)
		return nil
	}
	url := &Url{}
	url.ID = uint(i)
	if err := database.Db.First(url).Error; err != nil {
		log.Printf("action=fetch url data by id, status=failed, err=%s", err)
		return nil
	}
	return url
}
