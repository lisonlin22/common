package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	httpx.WriteJson(w, http.StatusUnauthorized, &Body{401, "鉴权失败", ""})
}
