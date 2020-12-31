package parsetime

import (
	"log"
	"time"
)

/* StringToTime ex) strTime: "2020-06-01 01:32" */
func StringToTime(strTime string) *time.Time {
	layout := "2006-01-02 15:04"
	time, err := time.Parse(layout, strTime)
	if err != nil {
		log.Printf("action=parse string to time, status=failed, err=%s", err)
		return nil
	}
	return &time
}

/* TimeToString ex) return string: "2020-06-01 01:32" */
func TimeToString(time time.Time) string {
	layout := "2006-01-02 15:04"
	return time.Format(layout)
}