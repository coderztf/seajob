package mail

import (
	"spider/entity"
	"bytes"
	"strconv"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"net/url"
	"fmt"
)

func SendMail(list []entity.JobInfo) {
	content, err := json.Marshal(map[string]interface{}{
		"to": []string{"coderztf@hotmail.com"},
		"sub": map[string][]string{
			"%count%": {strconv.Itoa(len(list))},
			"%list%":  {makeContent(list)},
		},
	})
	RequestURI := "http://api.sendcloud.net/apiv2/mail/sendtemplate"
	PostParams := url.Values{
		"apiUser":            {"coderztf_test_4a477n"},
		"apiKey":             {"12wb9356Ivn4Dg9r"},
		"from":               {"seajob@coderztf.fun"},
		"fromName":           {"海投助手"},
		"templateInvokeName": {"seajob"},
		"xsmtpapi":           {string(content)},
	}
	PostBody := bytes.NewBufferString(PostParams.Encode())
	ResponseHandler, err := http.Post(RequestURI, "application/x-www-form-urlencoded", PostBody)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ResponseHandler.Body.Close()
	BodyByte, err := ioutil.ReadAll(ResponseHandler.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(BodyByte))
}

func makeContent(list []entity.JobInfo) string {
	var buf bytes.Buffer
	for index, item := range list {
		buf.WriteString(fmt.Sprintf("<a href=\"%s\">%d.%s</a></br>", item.Url, index+1, item.Title))
	}
	return buf.String()
}
