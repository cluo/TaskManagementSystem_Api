package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	TaskList map[string]*Task
)

func init() {
	TaskList = make(map[string]*Task)
	u := Task{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	TaskList["user_11111"] = &u
}

type Task struct {
	Id       string
	Taskname string
	Password string
	Profile  Profile
}

func AddTask(u Task) string {
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	TaskList[u.Id] = &u
	return u.Id
}

func GetTask(uid string) (u *Task, err error) {
	if u, ok := TaskList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("Task not exists")
}

func GetAllTasks() map[string]*Task {
	return TaskList
}

func UpdateTask(uid string, uu *Task) (a *Task, err error) {
	if u, ok := TaskList[uid]; ok {
		if uu.Taskname != "" {
			u.Taskname = uu.Taskname
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("Task Not Exist")
}

func DeleteTask(uid string) {
	delete(TaskList, uid)
}
