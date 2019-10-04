package models

type Department struct {
	DeptNo   int32  `gorm:"dept_no;PRIMARY_KEY" json:"dept_no"`
	DeptName string `gorm:"dept_name" json:"dept_name"`
}
