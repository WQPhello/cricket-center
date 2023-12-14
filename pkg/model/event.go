package model

import "time"

type Event struct {
	ID              uint             `json:"id" xorm:"id autoincr"`
	ProjectNumber   string           `json:"project_number" xorm:"unique(prj-evt)"`           //所属项目编号
	EventCode       string           `json:"event_code" xorm:"unique(prj-evt)"`               //事件编号
	QuakeType       int64            `json:"quake_type"`                                      //事件类型，1-微震 2-爆破 3-未知
	HappendAt       *time.Time       `json:"happend_at" xorm:"happend_at DATETIME(6)"`        // Time of the quake event, eg."2006-01-02T15:04:05.000000+08:00"
	Coordinate      *Coordinate      `json:"coordinate" xorm:"extends"`                       // Coordinate of the quake event
	HeadAt          *time.Time       `json:"head_at" xorm:"DATETIME(6)"`                      // Start time of the event, eg."2006-01-02T15:04:05.000000+08:00"
	TailAt          *time.Time       `json:"tail_at" xorm:"DATETIME(6)"`                      // End time of the event, eg."2006-01-02T15:04:05.000000+08:00"
	ManualAt        *time.Time       `json:"manual_at" xorm:"DATETIME(6)"`                    // Time of manual intervention, eg."2006-01-02T15:04:05.000000+08:00"
	PWaveVelocity   float64          `json:"p_wave_velocity"`                                 // P-wave velocity
	SWaveVelocity   float64          `json:"s_wave_velocity"`                                 // S-wave velocity
	SeismicMoment   float64          `json:"seismic_moment"`                                  // 地震矩，N·m
	MomentMagnitude float64          `json:"moment_magnitude"`                                // 矩震级，级
	LocalMagnitude  float64          `json:"local_magnitude"`                                 // 局部震级，级
	SeismicRadius   float64          `json:"seismic_radius"`                                  // 震源半径，米
	SeismicEnergy   float64          `json:"seismic_energy"`                                  // 地震能量，焦耳
	StressDrop      float64          `json:"stress_drop"`                                     // 应力降，Pa
	TriggerCount    int              `json:"trigger_count"`                                   // Number of triggered sensors
	Sampling        float64          `json:"sampling"`                                        // 采用率，一秒钟多少次采样
	Sensors         []*EventSensor   `json:"sensors,omitempty" xorm:"json"`                   // Details of triggered sensors
	CreatedAt       *time.Time       `json:"created_at,omitempty" xorm:"created DATETIME(6)"` // Creation time
	UpdatedAt       *time.Time       `json:"updated_at" xorm:"updated DATETIME(6)"`           // Last updated time
	WaveMap         map[string]*Wave `json:"-" xorm:"json"`                                   // 波形数据 key-sn value-*Wave
}

type EventSensor struct {
	SensorCode       string      `json:"sn"`           //传感器编号
	PwaveArriveTime  *time.Time  `json:"pa,omitempty"` //P波拾取到时
	PwaveTheoryTime  *time.Time  `json:"pt,omitempty"` //P波理论到时
	PwaveProbability float64     `json:"pp,omitempty"` //P波到时概率
	PwaveTimeError   int64       `json:"pe,omitempty"` //P波到时误差
	Coordinate       *Coordinate `json:"cd,omitempty"` //传感器坐标
}

type Wave struct {
	T []int64   `json:"t,omitempty"` //波形时间序列(us)
	X []float64 `json:"x,omitempty"` //x轴数据序列
	Y []float64 `json:"y,omitempty"` //y轴数据序列
	Z []float64 `json:"z,omitempty"` //z轴数据序列
}

func GetQuakeTypeStr(tp int64) string {
	if tp == 1 {
		return "微震"
	} else if tp == 2 {
		return "爆破"
	} else {
		return "未知"
	}
}
