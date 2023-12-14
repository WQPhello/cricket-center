# <center>微震平台中心端-前后端API设计</center>


### 1.获取验证码
- GET /v1/captcha
- 备注：若登录时验证码验证失败，前端需要立即刷新验证码
- 响应 body：json
```go
{
    code int       //0-成功 非0-失败
    message string //提示信息
    data {
        captcha_id string    //验证码身份标识
        captcha_image string //验证码图片
    }
}
```


### 2.登录
- POST /v1/login
- 备注：除了获取验证码和登录请求外，其他请求需携带 token
- 请求 body：json
```go
{
  captcha_id:string
  captcha_value:string
  username:string
  password:string
}
```
- 响应 body：json
```go
{
    code int       //0-成功 非0-失败
    message string //提示信息
    data {
        expire int64 //token有效时间戳(秒级)
        token string //token字符串
    }
}
```


### 3.查询项目列表
- GET /v1/projects
- 请求参数：
  - project_number?:string //项目编号，模糊查询
  - project_name?:string   //项目名称，模糊查询
  - enterprise?:string     //单位名称，模糊查询
  - project_status?:int    //项目状态，1-检测中 2-已停用
  - sensor_type?:int       //传感器类型，1-速度传感器 2-加速度传感器
  - page?:int              //页码，默认1
  - page_size?:int         //每页条数，默认10
- 响应 body：json
```go
{
    code int       //0-成功 非0-失败
    message string //提示信息
    data {
        total int64 //项目总数
        list [{
            project_number string    //项目编号
            project_name string      //项目名称
            enterprise string        //单位名称
            project_status int       //项目状态，1-检测中 2-已停用
            sensor_type int          //传感器类型，1-速度传感器 2-加速度传感器
            longitude float64        //项目经度
            latitude float64         //项目纬度
            sensor_number int        //传感器数量
            noise_sensor float64     //噪声传感器
            sampling int             //传感器采样频率，Hz
            response_frequency [int] //传感器响应频率
            sensitivity float64      //传感器灵敏度
            p_wave_velocity float64  //P波速度
            s_wave_velocity float64  //S波速度
        }]
    }
}
```


### 4.查询传感器列表
- GET /v1/sensors
- 请求参数：
  - project_number?:string //所属项目编号，不传递时查询全部项目的数据
  - sensor_code?:string    //传感器编号，模糊查询
  - sensor_type?:int       //传感器类型，1-速度传感器 2-加速度传感器
  - page?:int              //页码，默认1
  - page_size?:int         //每页条数，默认10
- 响应 body：json
```go
{
    code int       //0-成功 非0-失败
    message string //提示信息
    data {
        total int64 //传感器总数
        list [{
            sensor_code string    //传感器编号
            used_status bool      //使用状态 true-使用中 false-停用中
            traffic_status bool   //通信状态 true-正常 false-异常
            created_at string     //创建时间
            positions [{          //位置记录
                x float64           //x坐标(m)
                y float64           //y坐标(m)
                z float64           //z坐标(m)
                updated_at string   //更新时间，格式如：2006-01-02T15:04:05+08:00
            }]
        }]
    }
}
```


### 5.查询指定事件列表
- GET /v1/events
- 请求参数：
  - project_number?:string //所属项目编号，不传递时查询全部项目的数据
  - quake_type?:int        //事件类型，1-微震 2-爆破 3-未知
  - happend_at_from?:string //发震时间范围，格式如"2006-01-02T15:04:05.000000Z08:00"
  - happend_at_to?:string //发震时间范围，格式如"2006-01-02T15:04:05.000000Z08:00"
  - seismic_moment_from?:float64 //地震矩范围，N·m
  - seismic_moment_to?:float64 //地震矩范围，N·m
  - moment_magnitude_from?:float64 //局震级范围，级
  - moment_magnitude_to?:float64 //局震级范围，级
  - local_magnitude_from?:float64 //本地震级范围，级
  - local_magnitude_to?:float64 //本地震级范围，级
  - seismic_energy_from?:float64 //地震能量范围，焦耳
  - seismic_energy_to?:float64 //地震能量范围，焦耳
  - seismic_radius_from?:float64 //地震半径范围，米
  - seismic_radius_to?:float64 //地震半径范围，米
  - stress_drop_from?:float64 //应力降范围，帕
  - stress_drop_to?:float64 //应力降范围，帕
  - x_min?:float64 //事件范围
  - x_max?:float64 //事件范围
  - y_min?:float64 //事件范围
  - y_max?:float64 //事件范围
  - z_min?:float64 //事件范围
  - z_max?:float64 //事件范围
  - page?:int       //页码，默认1
  - page_size?:int  //每页条数，默认10
- 响应 body：json
```go
{
    code int       //0-成功 非0-失败
    message string //提示信息
    data {
        total int64 //事件总数
        list [{
            event_code string     //事件编号
            project_number string //所属项目编号
            quake_type int        //事件类型，1-微震 2-爆破 3-未知
            happend_at string     //发震时间
            coordinate {          //定位坐标
                x float64 //x轴，m
                y float64 //y轴，m
                z float64 //z轴，m
            }
            head_at string           //事件波形开始时间
            tail_at string           //事件波形结束时间
            p_wave_velocity float64  //p波速度，m/s
            s_wave_velocity float64  //s波速度，m/s
            sampling int             //本事件的数据采用率，一秒钟多少次采样
            trigger_count int        //传感器触发数量
            seismic_moment float64   //地震矩，N·m
            moment_magnitude float64 //矩震级，级
            local_magnitude float64  //本地震级，级
            seismic_energy float64   //地震能量，焦耳
            seismic_radius float64   //地震半径，米
            stress_drop float64      //应力降，帕
            manual_at string         //人工干预时间，为空时表示非人工干预，不为空时表示人工干预
            updated_at string        //计算时间，格式如"2006-01-02T15:04:05.000000+08:00"
            sensors [{ //传感器信息
                sn string  //传感器编号
                pt string  //P波理论到时
                pa string  //P波拾取到时，没有p波到时的传感器无此字段
                pp float64 //P波到时概率，没有p波到时的传感器无此字段
                pe float64 //P波到时误差(us)，没有p波到时的传感器无此字段
                cd { //传感器坐标
                    x float64 //x轴，m
                    y float64 //y轴，m
                    z float64 //z轴，m
                }
            }]
        }]
    }
}
```


### 6.查询指定事件波形
- GET /v1/events/waves?event_code=string&project_number=string 
- 响应body：json
```go
{
    code    int
    message string
    data [{
        sn string          //传感器编号
        head_at int64      //(理论)数据开始时间戳(us)
        tail_at int64      //(理论)数据结束时间戳(us)
        wave {
            t []int64   //时间戳序列(us)
            x []float64 //x轴数据，没有时不传
            y []float64 //y轴数据，没有时不传
            z []float64 //z轴数据，没有时不传
        }
    }]
}
```


### 7.下载指定事件的波形文件
- GET /v1/events/waves/download?event_code=string&project_number=string 
- 响应：
  成功时状态码为 200，并返回压缩文件，Content-Type: application/octet-stream
  失败时状态码为 400，并返回提示信息

