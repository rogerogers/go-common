package res

type res struct {
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

func NewRes() *res {
	return &res{}
}

func (r *res) ErrorRes(code int, message string) (res *res) {
	r.Success = false
	r.ErrorCode = code
	r.ErrorMessage = message
	return r
}

func (r *res) SuccessRes(data interface{}) (res *res) {
	r.Data = data
	r.Success = true
	return r
}

func (r *res) ListRes(data interface{}, total int, pageSize int, current int) (res *res) {
	r.Data = data
	r.Total = total
	r.PageSize = pageSize
	r.Current = current
	r.Success = true
	return res
}
