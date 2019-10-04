package models

type Department struct {
	DeptNo int32 `gorm:"dept_no;PRIMARY_KEY" json:"dept_no"`
}
