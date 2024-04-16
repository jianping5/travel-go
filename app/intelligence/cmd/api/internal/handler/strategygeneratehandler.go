package handler

import (
	"bytes"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"travel/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"travel/app/intelligence/cmd/api/internal/svc"
	"travel/app/intelligence/cmd/api/internal/types"
)

func StrategyGenerateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var clientReq types.StrategyGenerateReq
		if err := httpx.Parse(r, &clientReq); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		accessToken, err := getAccessToken()
		if err != nil {
			logx.Error(err)
			http.Error(w, "Failed to get access token", http.StatusInternalServerError)
			return
		}

		url := fmt.Sprintf("https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/yi_34b_chat?access_token=%s", accessToken)
		fmt.Println(url)

		// 构建请求主体
		reqBody := fmt.Sprintf(`{"messages":[{"role":"user","content":"请你帮我制定一份详细的旅游攻略，目的地为%s，持续时长为%s，预算为%s，旅行团队为%s，旅行类型为%s。现在假如你是一名阅历丰富的旅行家，请你根据这些帮我指定一份详细且细致的旅游攻略，每个景点的安排都要有至少500字的详细说明，并给出预算，具体时间安排以及相关注意事项（格式：总标题为当前旅游攻略，分标题为时间，再一个分标题为上午、中午、下午、晚上、预算、注意事项等）"}],"stream":true}`,
			clientReq.Destination, clientReq.Duration, clientReq.Budget, clientReq.TripGroup, clientReq.TripMood)
		//reqBody := `{"messages":[{"role":"user","content":"帮我制定一份去云南旅游的3日攻略"}],"stream":true}`
		reqBuffer := bytes.NewBufferString(reqBody)

		// 创建带请求主体长度的 POST 请求
		req, err := http.NewRequest(http.MethodPost, url, reqBuffer)
		if err != nil {
			logx.Error(err)
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.ContentLength = int64(reqBuffer.Len())

		// 发送请求
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			logx.Error(err)
			http.Error(w, "Failed to post request", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// 创建一个通道，用于接收从 goroutine 中读取的数据块
		dataCh := make(chan []byte)

		// 启动一个 goroutine，用于从响应体中读取数据，并将其发送到通道中
		go func() {
			defer close(dataCh)

			buf := make([]byte, 8072)
			for {
				n, err := resp.Body.Read(buf)
				if err != nil {
					if err != io.EOF {
						logx.Error(err)
					}
					return
				}
				if n > 0 {
					dataCh <- buf[:n]
				}
			}
		}()

		// 从通道中循环读取数据，并将其写入响应流中
		for data := range dataCh {
			_, err := w.Write(data)
			if err != nil {
				logx.Error(err)
				return
			}
			// 刷新响应流，确保数据及时发送给客户端
			if flusher, ok := w.(http.Flusher); ok {
				flusher.Flush()
			}
		}
	}
}
