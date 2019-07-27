package tools

type XResult struct {
	Code    int
	Data    interface{}
	Message string
}

var defaultXResult = XResult{
	Code:    0,
	Data:    nil,
	Message: "操作成功！",
}

type XResultOption func(*XResult)

func WithCode(code int) XResultOption {
	return func(result *XResult) {
		result.Code = code
	}
}

func WithData(data interface{}) XResultOption {
	return func(result *XResult) {
		result.Data = data
	}
}

func WithMessage(message string) XResultOption {
	return func(result *XResult) {
		result.Message = message
	}
}

// 函数默认参数的Go实现（真麻烦！(╯﹏╰)）
func GetResult(opts ...XResultOption) XResult {
	result := defaultXResult

	for _, o := range opts {
		o(&result)
	}
	return result
}
