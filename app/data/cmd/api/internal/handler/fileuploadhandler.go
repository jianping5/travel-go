package handler

import (
	"net/http"
	"travel/common/result"

	"travel/app/data/cmd/api/internal/logic"
	"travel/app/data/cmd/api/internal/svc"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileType := r.FormValue("fileType")
		file := r.MultipartForm.File
		if v, ok := file["file"]; !ok {
			result.ParamErrorResult(r, w, nil)
		} else {
			if len(v) == 0 {
				result.ParamErrorResult(r, w, nil)
			}
		}

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(fileType, file["file"][0])
		result.HttpResult(r, w, resp, err)
	}
}
