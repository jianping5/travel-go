package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io"
	"net/http"
	"strings"
	"travel/app/intelligence/cmd/api/internal/svc"
	"travel/app/intelligence/cmd/api/internal/types"
	"travel/common/result"
)

const (
	ak = "EFeXyFMYcg56CP4hxmXrLopV"
	sk = "HblCZXVxyUFQFyu1M1mN49WiS1TS5vtt"
)

type ChatCompletion struct {
	ID               string `json:"id"`
	Object           string `json:"object"`
	Created          int64  `json:"created"`
	SentenceID       int    `json:"sentence_id"`
	IsEnd            bool   `json:"is_end"`
	IsTruncated      bool   `json:"is_truncated"`
	Result           string `json:"result"`
	NeedClearHistory bool   `json:"need_clear_history"`
	Usage            struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func ConversationGenerateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var clientReq types.ConversationGenerateReq
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
		reqBody := fmt.Sprintf(`{"messages":[{"role":"user","content":"%s"}],"stream":true}`, clientReq.Content)
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

type AccessTokenResp struct {
	AccessToken string `json:"access_token"`
}

func getAccessToken() (string, error) {
	url := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s", ak, sk)
	resp, err := http.Post(url, "application/json", strings.NewReader(""))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var tokenResp AccessTokenResp
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", err
	}
	return tokenResp.AccessToken, nil
}
