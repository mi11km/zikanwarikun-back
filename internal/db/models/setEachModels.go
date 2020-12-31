package models

import "sort"

func SetTodosToClass(class *Class) {
	todos, err := FetchTodosByClass(*class)
	if err == nil {
		sort.Slice(todos, func(i, j int) bool {
			return todos[i].Deadline.Before(todos[j].Deadline)
		})
		class.Todos = todos
	}
}

func SetUrlsToClass(class *Class) {
	urls, err := FetchUrlsByClass(*class)
	if err == nil {
		class.Urls = urls
	}
}

func SetTodosAndUrlsToEachClasses(classes []*Class) {
	if len(classes) == 0 {
		return
	}
	for _, c := range classes {
		SetTodosToClass(c)
		SetUrlsToClass(c)
	}
}

func SetClassesToTimetable(timetable *Timetable) {
	classes, err := FetchClassesByTimetable(*timetable)
	if err == nil {
		SetTodosAndUrlsToEachClasses(classes)
		timetable.Classes = classes
	}
}

func SetClassTimesToTimetable(timetable *Timetable) {
	classTimes, err := FetchClassTimesByTimetable(*timetable)
	if err == nil {
		timetable.ClassTimes = classTimes
	}
}

func SetClassesAndClassTimesToEachTimetables(timetables []*Timetable) {
	if len(timetables) == 0 {
		return
	}
	for _, t := range timetables {
		SetClassesToTimetable(t)
		SetClassTimesToTimetable(t)
	}
}
