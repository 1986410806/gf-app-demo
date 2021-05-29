package resources

const (
	SuccessCode = "E0000000"
	ErrorCode   = "E0000001"
)

type ResultData struct {
	Code     string
	Msg      string
	ErrorMsg string
	Data     interface{}
}
