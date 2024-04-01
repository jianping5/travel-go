package logic

import (
	"context"
	"github.com/google/uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
	"travel/app/data/cmd/api/internal/svc"
	"travel/app/data/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	client *cos.Client
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶 region 可以在 COS 控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse("https://travel-7-1310703557.cos.ap-shanghai.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: svcCtx.Config.COS.SecretID, // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: svcCtx.Config.COS.SecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
		},
	})

	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		client: client,
	}
}

func (l *FileUploadLogic) FileUpload(fileType string, file *multipart.FileHeader) (resp *types.UploadResp, err error) {
	// todo: add your logic here and delete this line
	srcFile, err := file.Open()

	fileName := uuid.New().String() + "-" + file.Filename

	if fileType == "image" {
		_, err = l.client.Object.Put(context.Background(), fileName, srcFile, &cos.ObjectPutOptions{
			ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
				ContentType:        "image/jpeg", // 设置文件类型为 JPEG 图片
				ContentDisposition: "inline",     // 访问时直接显示，不会自动下载（首次还是会）
			},
		})
	} else if fileType == "video" {
		_, err = l.client.Object.Put(context.Background(), fileName, srcFile, &cos.ObjectPutOptions{
			ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
				ContentType:        "video/mp4", // 设置文件类型为 video 视频
				ContentDisposition: "inline",    // 访问时直接显示，不会自动下载（首次还是会）
			},
		})
	}

	if err != nil {
		// ERROR
	}
	url := l.client.Object.GetObjectURL(fileName).String()

	return &types.UploadResp{
		FileUrl: url,
	}, nil
}
