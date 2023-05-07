// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	file "CloudMind/app/fileservice/cmd/api/internal/handler/file"
	"CloudMind/app/fileservice/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/file/FileUpload",
				Handler: file.FileuploadHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/file/FileDownload",
				Handler: file.FiledownloadHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/file/FileList",
				Handler: file.FilelistHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/file/FileNameUpdate",
				Handler: file.FilenameupdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/file/FileCreate",
				Handler: file.FilecreateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/file/FileDeletion",
				Handler: file.FiledeletionHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/file/FileMove",
				Handler: file.FilemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/file/FileShare",
				Handler: file.FileshareHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/file/FileShareSave",
				Handler: file.FilesharesaveHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/file/ViewFileDetails",
				Handler: file.ViewfiledetailsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/file-api/v1"),
	)
}
