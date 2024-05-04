package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"
	"travel/app/social/cmd/model"
	"travel/common/ctxdata"

	"github.com/ipfs/kubo/client/rpc"
	"github.com/zeromicro/go-zero/core/logx"
)

type CopyrightCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type CopyrightMetadata struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	VideoUrl    string `json:"video_url"`
	Description string `json:"description"`
}

func NewCopyrightCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyrightCreateLogic {
	return &CopyrightCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CopyrightCreateLogic) CopyrightCreate(req *types.CopyrightCreateReq) (*types.CopyrightCreateResp, error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)

	// step1：获取作品信息
	var content model.Content
	l.svcCtx.DB.Model(&model.Content{}).Select("title", "description", "cover_url", "content").Where("id = ?", req.ItemId).Scan(&content)

	// step2：上传到 IPFS
	// 构建版权元数据
	copyrightMetadata := &CopyrightMetadata{
		Name:        content.Title,
		Image:       content.CoverUrl,
		VideoUrl:    content.Content,
		Description: content.Description,
	}
	metadataJson, err := json.Marshal(copyrightMetadata)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil, err
	}
	// 上传到 IPFS，并获取 path
	metadataStr := string(metadataJson)
	ipfsHash, err := upload2IPFS(metadataStr)
	if err != nil {
		fmt.Println("Error uploading:", err)
		return nil, err
	}

	// step3：将版权记录插入数据库
	copyright := &model.Copyright{
		UserId:   loginUserId,
		ItemType: req.ItemType,
		ItemId:   req.ItemId,
		Metadata: metadataStr,
		IpfsHash: ipfsHash,
	}
	l.svcCtx.DB.Model(&model.Copyright{}).Create(copyright)

	return &types.CopyrightCreateResp{
		IpfsHash: ipfsHash,
	}, nil
}

func upload2IPFS(nftStr string) (string, error) {
	// 创建 IPFS 客户端
	client, _ := rpc.NewLocalApi()

	// 将字符串转换为 io.Reader 接口
	reader := strings.NewReader(nftStr)

	// 创建 BlockAPI 实例
	blockAPI := client.Block()

	// 调用 Put 方法上传文件内容到 IPFS
	stat, err := blockAPI.Put(context.Background(), reader)
	if err != nil {
		fmt.Println("Error uploading:", err)
		return "", err
	}

	// 打印上传文件的信息，包括 CID
	fmt.Println("Uploaded file to IPFS with CID:", stat.Path().RootCid())
	ipfsHash := stat.Path().RootCid().String()
	fmt.Println(ipfsHash)
	return ipfsHash, nil
}
