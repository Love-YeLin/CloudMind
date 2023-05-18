package logic

import (
	"CloudMind/app/es/cmd/rpc/internal/svc"
	"CloudMind/app/es/cmd/rpc/pb"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertFileLogic {
	return &InsertFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertFileLogic) InsertFile(in *pb.InsertFileReq) (*pb.InsertFileResp, error) {
	// 序列化
	data, err := json.Marshal(struct {
		Title     string `json:"title"`
		Id        string `json:"id"`
		UserId    int64  `json:"userId"`
		TypeMount string `json:"typeMount"`
	}{Title: in.File.Title,
		Id:        in.File.Id,
		UserId:    in.UserId,
		TypeMount: in.TypeMount,
	})
	if err != nil {
		return &pb.InsertFileResp{
			Error: fmt.Sprintf("Error marshaling the document: %s", err),
		}, nil
	}
	// 构建请求

	req := esapi.IndexRequest{
		Index:   "files",
		Body:    bytes.NewReader(data),
		Refresh: "true",
	}

	// 发请求
	res, err := req.Do(context.Background(), l.svcCtx.Es)
	if err != nil {
		return &pb.InsertFileResp{
			Error: fmt.Sprintf("Error indexing the document: %s", err),
		}, nil
	}
	defer res.Body.Close()

	if res.IsError() {
		return &pb.InsertFileResp{
			Error: fmt.Sprintf("[%s] Error indexing document ID", res.Status()),
		}, nil
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			return &pb.InsertFileResp{
				Error: fmt.Sprintf("Error parsing the response body: %s", err),
			}, nil
		}
	}

	return &pb.InsertFileResp{}, nil
}