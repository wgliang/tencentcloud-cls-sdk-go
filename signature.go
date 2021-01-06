package cls

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

func calSha1sum(msg string) string {
	h := sha1.New()
	h.Write([]byte(msg))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func calSha1HMACDigest(key, msg string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(msg))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Signature 计算步骤：
// 1,按照指定格式拼接 HTTP 请求的相关信息为字符串 HttpRequestInfo。
// 2,对 HttpRequestInfo 使用 sha1 算法计算哈希值，与其他指定参数按照指定格式组成签名原串 StringToSign。
// 3,使用 SecretKey 对 q-key-time 进行加密，得到 SignKey。
// 4,使用 SignKey 对 StringToSign 进行加密，生成 Signature。
func Signature(secretID, secretKey, method, path string, params, headers url.Values, expire int64) string {
	var signedHeaderList []string
	var signedParameterList []string
	hs := url.Values{}
	for key, values := range headers {
		for _, value := range values {
			var lowerKey = strings.ToLower(key)
			if lowerKey == "content-type" || lowerKey == "content-md5" || lowerKey == "host" || lowerKey[0] == 'x' {
				hs.Add(lowerKey, value)
				signedHeaderList = append(signedHeaderList, lowerKey)
			}
		}
	}

	var formatHeaders = hs.Encode()
	sort.Strings(signedHeaderList)
	ps := url.Values{}
	for key, values := range params {
		for _, value := range values {
			var lowerKey = strings.ToLower(key)
			ps.Add(lowerKey, value)
			signedParameterList = append(signedParameterList, lowerKey)
		}
	}

	var formatParameters = ps.Encode()
	sort.Strings(signedParameterList)
	var formatString = fmt.Sprintf("%s\n%s\n%s\n%s\n", strings.ToLower(method), path, formatParameters, formatHeaders)
	var signTime = fmt.Sprintf("%d;%d", time.Now().Unix()-60, time.Now().Unix()+expire)
	var stringToSign = fmt.Sprintf("sha1\n%s\n%s\n", signTime, calSha1sum(formatString))
	var signKey = calSha1HMACDigest(secretKey, signTime)
	var signature = calSha1HMACDigest(signKey, stringToSign)

	return strings.Join([]string{
		"q-sign-algorithm=sha1",
		"q-ak=" + secretID,
		"q-sign-time=" + signTime,
		"q-key-time=" + signTime,
		"q-header-list=" + strings.Join(signedHeaderList, ";"),
		"q-url-param-list=" + strings.Join(signedParameterList, ";"),
		"q-signature=" + signature,
	}, "&")
}
