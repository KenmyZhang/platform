//accurme--
package oauthwechat

import (
	"testing"
	"github.com/mattermost/platform/model"
	"bytes"
)
func TestGetUserFromJson(t *testing.T) { 
	weChatProvider := &WeChatProvider{}
	userData := []byte(`{"openid":"oLuGXwIQIWuNs6C1BYI-I63j-8Bg","nickname":"鍖垮悕娑堟伅","sex":1,"language":"zh_CN","city":"Shenzhen","province":"Guangdong","country":"CN","headimgurl":"http:\/\/wx.qlogo.cn\/mmopen\/pBfEh6bfnRFfyPRhguoiblsbiciaq4or44s2lP7sRfLbvIepysWSlL7q2I5icvCj8ibCSYicfg9lPicd1cfd5FPqDdiaiaiahZLzr9x2zy\/0","privilege":[],"unionid":"oRZMquMIattaRxD71QahiaBUQVrk"}`)
	result := weChatProvider.GetUserFromJson(bytes.NewReader(userData))
	if *(result.AuthData) != "oLuGXwIQIWuNs6C1BYI-I63j-8Bg" || result.Nickname != "鍖垮悕娑堟伅" || result.AuthService != model.USER_AUTH_SERVICE_WECHAT {

	   		t.Fatal("GetUserFromJson failed!")
	   }

}

func TestGetAuthDataFromJson(t *testing.T) {
	weChatProvider := &WeChatProvider{}
	userData := []byte(`{"openid":"oLuGXwIQIWuNs6C1BYI-I63j-8Bg","nickname":"鍖垮悕娑堟伅","sex":1,"language":"zh_CN","city":"Shenzhen","province":"Guangdong","country":"CN","headimgurl":"http:\/\/wx.qlogo.cn\/mmopen\/pBfEh6bfnRFfyPRhguoiblsbiciaq4or44s2lP7sRfLbvIepysWSlL7q2I5icvCj8ibCSYicfg9lPicd1cfd5FPqDdiaiaiahZLzr9x2zy\/0","privilege":[],"unionid":"oRZMquMIattaRxD71QahiaBUQVrk"}`)
	rauthData := weChatProvider.GetAuthDataFromJson(bytes.NewReader(userData))
	if rauthData != "oLuGXwIQIWuNs6C1BYI-I63j-8Bg" {
		t.Fatal("GetAuthDataFromJson failed!")
	}
}
//--accurme