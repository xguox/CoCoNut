package util

const (
	MarkOk   = 0
	MarkWarn = 1
)

type Resp struct {
	Mark int         `json:"mark"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResp() *Resp {
	return &Resp{}
}
