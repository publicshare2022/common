package xerr

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func Response(ctx context.Context, w http.ResponseWriter, resp interface{}, err error) {
	body := NewErrorCode(OK)
	if err != nil {
		if x, ok := err.(*Error); !ok {
			body = NewError(UNKNOW_ERROR, err.Error())
		} else {
			body = x
		}
	} else {
		body.Data = resp
	}

	httpx.OkJsonCtx(ctx, w, body)
}
