package model

import "time"

// 传感器表
type Sensor struct {
	ID            int         `query:"id" xorm:"pk autoincr" json:"id"`
	ProjectNumber string      `json:"project_number" xorm:"unique(prj-sn)"`  //所属项目编号
	SensorCode    string      `json:"sensor_code" xorm:"unique(prj-sn)"`     //传感器编号
	CardName      string      `json:"card_name" xorm:"unique(prj-sn)"`       //数采编号
	UsedStatus    *bool       `json:"used_status"`                           //使用状态 true-使用中 false-停用中
	TrafficStatus *bool       `json:"traffic_status"`                        //通讯状态 true-正常 false-异常
	Positions     []*Position `json:"positions" xorm:"json"`                 //传感器坐标信息
	CreatedAt     *time.Time  `json:"created_at" xorm:"created DATETIME(6)"` //创建时间
}
