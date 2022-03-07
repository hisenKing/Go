package utils

import (
	context2 "context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"strings"
	"xiaoshijie.com/web/config"
	"xiaoshijie.com/web/internal/logger"
)

func NewKuaidi100() *Kuaidi100 {
	return &Kuaidi100{}
}

type Kuaidi100 struct {
}

func (m *Kuaidi100) Query(company string, number string) string {
	url := "https://poll.kuaidi100.com/poll/query.do"
	params := map[string]string{
		"com": company,
		"num": number,
	}

	var dataParam = make(map[string]string)
	dataParam["customer"] = config.Project.Kuaidi100.CustomId
	param, _ := json.Marshal(params)
	dataParam["param"] = string(param)

	h := md5.New()
	h.Write([]byte(dataParam["param"] + config.Project.Kuaidi100.AppKey + dataParam["customer"]))
	dataParam["sign"] = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))

	resp, err := HttpPost(url, dataParam)
	if err != nil {
		logger.GetZap(context2.TODO()).Error("快递100 订单查询失败" + err.Error())
	}

	return resp
}

func (m *Kuaidi100) Autonumber(number string) string {
	url := "http://www.kuaidi100.com/autonumber/auto"

	resp, err := HttpPost(url, map[string]string{"num": number, "key": config.Project.Kuaidi100.AppKey})
	if err != nil {
		logger.GetZap(context2.TODO()).Error("快递100 订单查询失败" + err.Error())
	}

	return resp
}

func (m *Kuaidi100) Poll(company string, number string) string {
	url := "https://poll.kuaidi100.com/poll"
	parameters := map[string]string{
		"callbackurl": config.Project.Kuaidi100.CallbackUrl,
	}
	params := map[string]interface{} {
		"company": company,
		"number": number,
		"key": config.Project.Kuaidi100.AppKey,
		"parameters": parameters,
	}
	var dataParam = make(map[string]string)
	param, _ := json.Marshal(params)
	dataParam["param"] = string(param)
	dataParam["schema"] = "JSON"

	resp, err := HttpPost(url, dataParam)
	if err != nil {
		logger.GetZap(context2.TODO()).Error("快递100 订单订阅失败" + err.Error())
	}

	return resp
}
