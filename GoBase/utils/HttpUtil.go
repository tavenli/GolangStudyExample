package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"bytes"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func GetIP(c *beego.Controller) string {
	//utils.GetIP(&c.Controller)
	//也可以直接用 c.Ctx.Input.IP() 取真实IP
	ip := c.Ctx.Request.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	ip = c.Ctx.Request.Header.Get("Remote_addr")
	if ip == "" {
		ip = c.Ctx.Request.RemoteAddr
	}
	return ip
}

func HttpPostJson(url string, json string) (string, error) {
	resp, err := http.Post(url, "application/json", strings.NewReader(json))
	if err != nil {
		logs.Error("HttpPostJson error: ", err)
		return "", err
	}

	if resp == nil {
		return "", errors.New("返回对象为空")
	}

	defer resp.Body.Close()
	result := ""
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		result = string(body)
		//logs.Debug("HttpPostJson result: ", result)
	} else {
		logs.Error("HttpPostJson error: ", err)
	}

	return result, nil
}

func HttpPostJsonReturnByte(url string, json string) ([]byte, error) {
	resp, err := http.Post(url, "application/json", strings.NewReader(json))
	if err != nil {
		logs.Error("HttpPostJson error: ", err)
		return nil, err
	}

	if resp == nil {
		return nil, errors.New("返回对象为空")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		return body, err
		//logs.Debug("HttpPostJson result: ", result)
	} else {
		logs.Error("HttpPostJson error: ", err)
		return nil, err
	}

}

func HttpPost(url string, param map[string]string) (string, error) {

	var paramBuf bytes.Buffer
	paramBuf.WriteString("curTime=" + GetCurrentTime())
	for k, v := range param {
		paramBuf.WriteString("&" + k + "=" + v)
	}

	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(paramBuf.String()))
	if err != nil {
		logs.Error("HttpPost error: ", err)
		return "", err
	}

	if resp == nil {
		return "", errors.New("返回对象为空")
	}

	defer resp.Body.Close()
	result := ""
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		result = string(body)
		logs.Debug("HttpPost result: ", result)
	} else {
		logs.Error("HttpPost error: ", err)
	}

	return result, nil
}

func HttpRequestRaw(filePath string, https bool) (string, error) {
	client := &http.Client{}

	rawPayload, err := ioutil.ReadFile(filePath)
	if err != nil {
		//logs.Error("", err)
		return "", err
	}

	rawReq, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(rawPayload)))
	if err != nil && err != io.ErrUnexpectedEOF {
		return "", err
	}

	var reqUrl string
	if https {
		reqUrl = "https://" + rawReq.Host + rawReq.RequestURI
	} else {
		reqUrl = "http://" + rawReq.Host + rawReq.RequestURI
	}

	httpReq, err := http.NewRequest(rawReq.Method, reqUrl, rawReq.Body)
	httpReq.Header = rawReq.Header
	response, err := client.Do(httpReq)
	fmt.Println(err)
	bodyBytes, err := ioutil.ReadAll(response.Body)

	return string(bodyBytes), nil

}

func UrlEncode(input string) string {
	if IsEmpty(input) {
		return ""
	}
	return url.QueryEscape(input)
}

func UrlDecode(input string) string {
	if IsEmpty(input) {
		return ""
	}
	result, err := url.QueryUnescape(input)
	if err != nil {
		return input
	} else {
		return result
	}
}
