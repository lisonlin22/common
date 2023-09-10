package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
	TraceId string      `json:"trace_id"`
}

func Response(w http.ResponseWriter, resp interface{}, traceId string, err error) {
	var body Body
	body.TraceId = traceId
	if err != nil {
		body.Code = -1
		body.Msg = err.Error()
		body.Data = ""
	} else {
		body.Msg = ""
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
