package service

import (
	"Project/model"
	"strconv"
)

func DiscussActionService(studentId string, content string) error {
	student_id, err := strconv.Atoi(studentId)
	if err != nil {
		return err
	}
	if err = model.InsertData(student_id, content); err != nil {
		return err
	}

	return nil
}

func DiscussListService() ([]model.Discuss, error) {
	discusses, err := model.QueryData()
	if err != nil {
		return nil, err
	}
	return discusses, nil
}
