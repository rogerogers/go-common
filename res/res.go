package res

type Res struct {
	Data         interface{} `json:"data"`
	Success      bool        `json:"success"`
	ErrorCode    int         `json:"errorCode,omitempty"`
	ErrorMessage string      `json:"errorMessage,omitempty"`
	ShowType     int         `json:"showType,omitempty"`
	TraceId      string      `json:"traceId,omitempty"`
	Host         string      `json:"host,omitempty"`
	Total        int         `json:"total,omitempty"`
	PageSize     int         `json:"pageSize,omitempty"`
	Current      int         `json:"current,omitempty"`
}

func ErrorRes(code int, message string) (res *Res) {
	Res := Res{
		Success:      false,
		ErrorCode:    code,
		ErrorMessage: message,
	}
	return &Res
}

func SuccessRes(data interface{}) (res *Res) {
	Res := Res{
		Data:    data,
		Success: true,
	}
	return &Res
}

func ListRes(data interface{}, total int, pageSize int, current int) (res *Res) {

	Res := Res{
		Data:     data,
		Total:    total,
		PageSize: pageSize,
		Current:  current,
		Success:  true,
	}
	return &Res
}
