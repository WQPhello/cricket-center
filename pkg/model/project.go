package model

// 项目列表
type Project struct {
	ID                int     `query:"id" xorm:"pk autoincr" json:"id"`
	ProjectNumber     string  `json:"project_number" xorm:"unique" query:"project_number"` //项目的编号
	ProjectName       string  `json:"project_name"`                                        //项目名称
	Enterprise        string  `json:"enterprise" query:"enterprise"`                       //单位名称
	ProjectStatus     int     `json:"project_status"`                                      //项目状态 1-监测中 2-已停用
	Longitude         float64 `json:"longitude"`                                           //项目标定经度，初始为 0.0
	Latitude          float64 `json:"latitude"`                                            //项目标定纬度，初始为 0.0
	SensorType        int     `json:"sensor_type"`                                         //传感器类型: 1-速度传感器 2-加速度传感器
	SensorNumber      int     `json:"sensor_number"`                                       //传感器数量
	Sampling          int     `json:"sampling"`                                            //(传感器)采样率，一般是 2000 或 5000
	ResponseFrequency []int   `json:"response_frequency"`                                  //(传感器)响应频率范围
	Sensitivity       float64 `json:"sensitivity"`                                         //(传感器)灵敏度
	PWaveVelocity     float64 `json:"p_wave_velocity"`                                     //P波速度，单位：m/s，初始值为 4000.0
	SWaveVelocity     float64 `json:"s_wave_velocity"`                                     //S波速度，单位：m/s，初始值为 3000.0
	NoiseSensor       string  `json:"noise_sensor"`                                        //噪声传感器(环境传感器)
}
