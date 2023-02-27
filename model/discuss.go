package model

import "time"

type Tabler interface {
	TableName() string
}

type Discuss struct {
	Id         int       `json:"id"`
	StudentId  int       `json:"studentId" ,gorm:"column:student_id"`
	Content    string    `json:"content" ,gorm:"column:content"`
	CreateTime time.Time `json:"createTime" ,gorm:"column:create_time"`
}

// TableName 会将 User 的表名重写为 `profiles`
func (Discuss) TableName() string {
	return "discuss"
}

func InsertData(studentId int, content string) error {
	discuss := Discuss{
		StudentId:  studentId,
		Content:    content,
		CreateTime: time.Now(),
	}
	result := db.Debug().Create(&discuss)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func QueryData() ([]Discuss, error) {
	var discusses []Discuss
	// SELECT * FROM `discuss` ORDER BY create_time desc
	result := db.Debug().Order("create_time desc").Find(&discusses)
	if result.Error != nil {
		return nil, result.Error
	}
	return discusses, nil
}
