package main

import "C"
import (
	"encoding/json"
	"github.com/staparx/go_showstart/client"
	"github.com/staparx/go_showstart/monitor"
	"github.com/staparx/go_showstart/order"
	"github.com/staparx/go_showstart/util"
)

//export GenerateSignForiOS
func GenerateSignForiOS(
	path *C.char,
	data *C.char,
	cusat *C.char,
	sign *C.char,
	cusit *C.char,
	cusid *C.char,
	traceId *C.char,
	token *C.char,
	cterminal *C.char,
) *C.char {
	req := &util.GenerateSignReq{
		Path:      C.GoString(path),
		Data:      C.GoString(data),
		Cusat:     C.GoString(cusat),
		Sign:      C.GoString(sign),
		Cusit:     C.GoString(cusit),
		Cusid:     C.GoString(cusid),
		TraceId:   C.GoString(traceId),
		Token:     C.GoString(token),
		Cterminal: C.GoString(cterminal),
	}
	signResult := util.GenerateSign(req)
	return C.CString(signResult)
}

//export GenerateTraceIdForiOS
func GenerateTraceIdForiOS(length C.int) *C.char {
	return C.CString(util.GenerateTraceId(int(length)))
}

//export GenerateAESKeyForiOS
func GenerateAESKeyForiOS(traceId *C.char, token *C.char) *C.char {
	return C.CString(util.GenerateKey(C.GoString(traceId), C.GoString(token)))
}

//export AESEncryptForiOS
func AESEncryptForiOS(plainText *C.char, key *C.char) *C.char {
	enc, err := util.AESEncrypt(C.GoString(plainText), C.GoString(key))
	if err != nil {
		return C.CString("")
	}
	return C.CString(enc)
}

//export GetTokenForiOS
func GetTokenForiOS(configJSON *C.char) *C.char {
	var cfg client.ShowStartConfig
	if err := json.Unmarshal([]byte(C.GoString(configJSON)), &cfg); err != nil {
		return C.CString(`{"success":false}`)
	}
	cli := client.NewShowStartClient(&cfg)
	_ = cli.GetToken()
	out, _ := json.Marshal(map[string]any{
		"success": true,
		"cusat":   cli.Cusat,
	})
	return C.CString(string(out))
}

//export ActivitySearchListForiOS
func ActivitySearchListForiOS(configJSON, cityCode, keyword *C.char) *C.char {
	var cfg client.ShowStartConfig
	if err := json.Unmarshal([]byte(C.GoString(configJSON)), &cfg); err != nil {
		return C.CString(`{"success":false}`)
	}
	cli := client.NewShowStartClient(&cfg)
	resp, _ := cli.ActivitySearchList(C.GoString(cityCode), C.GoString(keyword))
	out, _ := json.Marshal(resp)
	return C.CString(string(out))
}

//export ActivityDetailForiOS
func ActivityDetailForiOS(configJSON, activityId *C.char) *C.char {
	var cfg client.ShowStartConfig
	if err := json.Unmarshal([]byte(C.GoString(configJSON)), &cfg); err != nil {
		return C.CString(`{"success":false}`)
	}
	cli := client.NewShowStartClient(&cfg)
	resp, _ := cli.ActivityDetail(C.GoString(activityId))
	out, _ := json.Marshal(resp)
	return C.CString(string(out))
}

//export ActivityTicketListForiOS
func ActivityTicketListForiOS(configJSON, activityId *C.char) *C.char {
	var cfg client.ShowStartConfig
	if err := json.Unmarshal([]byte(C.GoString(configJSON)), &cfg); err != nil {
		return C.CString(`{"success":false}`)
	}
	cli := client.NewShowStartClient(&cfg)
	resp, _ := cli.ActivityTicketList(C.GoString(activityId))
	out, _ := json.Marshal(resp)
	return C.CString(string(out))
}

//export ConfirmOrderForiOS
func ConfirmOrderForiOS(configJSON, activityId, ticketId, ticketNum *C.char) *C.char {
	var cfg client.ShowStartConfig
	if err := json.Unmarshal([]byte(C.GoString(configJSON)), &cfg); err != nil {
		return C.CString(`{"success":false}`)
	}
	cli := client.NewShowStartClient(&cfg)
	resp, _ := cli.Confirm(C.GoString(activityId), C.GoString(ticketId), C.GoString(ticketNum))
	out, _ := json.Marshal(resp)
	return C.CString(string(out))
}

//export GetCpListForiOS
func GetCpListForiOS(configJSON, ticketId *C.char) *C.char {
	var cfg client.ShowStartConfig
	if err := json.Unmarshal([]byte(C.GoString(configJSON)), &cfg); err != nil {
		return C.CString(`{"success":false}`)
	}
	cli := client.NewShowStartClient(&cfg)
	resp, _ := cli.CpList(C.GoString(ticketId))
	out, _ := json.Marshal(resp)
	return C.CString(string(out))
}

//export SubmitOrderForiOS
func SubmitOrderForiOS(configJSON, orderReqJSON *C.char) *C.char {
	var cfg client.ShowStartConfig
	if err := json.Unmarshal([]byte(C.GoString(configJSON)), &cfg); err != nil {
		return C.CString(`{"success":false}`)
	}
	var req order.OrderRequest
	_ = json.Unmarshal([]byte(C.GoString(orderReqJSON)), &req)
	cli := client.NewShowStartClient(&cfg)
	resp, _ := cli.Order(&req)
	out, _ := json.Marshal(resp)
	return C.CString(string(out))
}

//export GetOrderResultForiOS
func GetOrderResultForiOS(configJSON, orderJobKey *C.char) *C.char {
	var cfg client.ShowStartConfig
	if err := json.Unmarshal([]byte(C.GoString(configJSON)), &cfg); err != nil {
		return C.CString(`{"success":false}`)
	}
	cli := client.NewShowStartClient(&cfg)
	resp, _ := cli.GetOrderResult(C.GoString(orderJobKey))
	out, _ := json.Marshal(resp)
	return C.CString(string(out))
}

var gMonitor *monitor.Monitor

//export StartMonitorForiOS
func StartMonitorForiOS(configJSON *C.char) *C.char {
	var cfg monitor.MonitorConfig
	if err := json.Unmarshal([]byte(C.GoString(configJSON)), &cfg); err != nil {
		return C.CString(`{"success":false}`)
	}
	gMonitor = monitor.NewMonitor(&cfg)
	_ = gMonitor.Start()
	return C.CString(`{"success":true}`)
}

//export StopMonitorForiOS
func StopMonitorForiOS() *C.char {
	if gMonitor != nil {
		gMonitor.Stop()
	}
	return C.CString(`{"success":true}`)
}

func main() {
}