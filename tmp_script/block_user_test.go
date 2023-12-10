package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type bodyResp struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

var uids = []int{
	677387873,
	710942306,
	666902125,
	698359397,
	580918884,
	578821735,
	662707805,
	580918897,
	572530275,
	767565397,
	580918850,
	633347613,
	685776346,
	658513083,
	543169734,
	622861473,
	782245022,
	769662061,
	771758871,
	738203452,
}

func TestBlockUser(t *testing.T) {

	client := http.Client{}
	num := 0
	for _, uid := range uids {
		body := map[string]interface{}{
			"user_id":         uid,
			"forbid_duration": 6,
			"forbid_type":     1,
			"token":           "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3ODM5NDY3ODksImlzcyI6ImFkbWluX2lzc3VlciIsInVpZCI6MzgsImFkbWluIjpmYWxzZX0.jk8D3YXlp7NcpGm2tdktXqczRAWfbT2JOlgfEVv1ZDA",
		}
		params, _ := json.Marshal(body)
		buffer := bytes.NewBuffer(params)

		request, err := http.NewRequest("POST", "https://api.wuhubaoshi.com/admin/block/addBlock", buffer)
		if err != nil {
			t.Errorf("err:%v", err)
			return
		}
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
		response, err := client.Do(request)
		if err != nil {
			t.Errorf("err:%v", err)
			return
		}
		//t.Logf("info:%s", response.Status)
		bodyCon, _ := io.ReadAll(response.Body)

		bodyRet := &bodyResp{}
		json.Unmarshal(bodyCon, bodyRet)
		if bodyRet.Code != 20022 {
			num += 1
		}
		t.Logf("uid:%d body:%s code:%d", uid, bodyCon, bodyRet.Code)
	}
	t.Logf("处理完成！待封禁用户数量:%d 已经处理:%d,本次处理数量:%d", len(uids), len(uids)-num, num)
}
