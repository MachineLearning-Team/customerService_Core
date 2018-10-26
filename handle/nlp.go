package handle

import (
	"git.jsjit.cn/customerService/customerService_Core/wechat/util"
	"log"
	"net/url"
)

// AI语义处理
type AiSemantic struct {
	hostUrl string
}

func NewAiSemantic(aiHost string) *AiSemantic {
	return &AiSemantic{hostUrl: aiHost}
}

func (ai *AiSemantic) Dialogue(msg, token string) string {
	msg = url.QueryEscape(msg)
	bytes, err := util.PostJSON(ai.hostUrl, struct {
		Msg   string `json:"msg"`
		Token string `json:"token"`
	}{msg, token})
	if err != nil {
		log.Printf("AiSemantic.Dialogue is err :%s", err.Error())
		return ""
	} else {
		return string(bytes)
	}
}
