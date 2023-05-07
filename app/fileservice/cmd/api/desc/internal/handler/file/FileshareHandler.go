package file

import (
	"net/http"

	"CloudMind/app/fileservice/cmd/api/desc/internal/logic/file"
	"CloudMind/app/fileservice/cmd/api/desc/internal/svc"
	"CloudMind/app/fileservice/cmd/api/desc/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileshareHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileShareReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewFileshareLogic(r.Context(), svcCtx)
		resp, err := l.Fileshare(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
