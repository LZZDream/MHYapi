package mhyapp

import (
	"github.com/LZZDream/MHYapi/cookies"
	"github.com/LZZDream/MHYapi/define"
	"github.com/LZZDream/MHYapi/tools"
	"math/rand"
	"net/http"
	"time"
)

type AppCore struct {
	Cookies *cookies.CookiesCore
	headers http.Header
}

func (t *AppCore) getHeaders() http.Header {
	t.updateHeader()
	return t.headers
}

func (t *AppCore) updateHeader() {
	if t.headers == nil {
		t.headers = make(map[string][]string)
	}
	rand.Seed(time.Now().UnixNano())
	t.headers["Accept"] = []string{"*/*"}
	t.headers["DS"] = []string{tools.GetDs(false)}
	t.headers["x-rpc-client_type"] = []string{define.APPCLIENT_TYPE_ANDROID}
	t.headers["x-rpc-app_version"] = []string{define.APPCLIENT_VERSIONS}
	t.headers["x-rpc-sys_version"] = []string{"6.0.1"}
	t.headers["x-rpc-channel"] = []string{"miyousheluodi"}
	t.headers["x-rpc-device_id"] = []string{tools.GetDevicesID()}
	t.headers["x-rpc-device_name"] = []string{tools.RandString(rand.Intn(5) + 5)}
	t.headers["x-rpc-device_model"] = []string{"Mi 10"}
	t.headers["Referer"] = []string{"https://app.mihoyo.com"}
	t.headers["User-Agent"] = []string{"okhttp/4.8.0"}
	t.headers["cookie"] = []string{t.Cookies.GetStr()}
	t.headers["Content-Type"] = []string{"application/json"}
}
