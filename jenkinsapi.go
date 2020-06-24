package jenkinsapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackcipher/quickrequest"
)

type _JenkinsParam struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type _JenkinsParameter struct {
	Parameter []_JenkinsParam `json:"parameter"`
}

type JenkinsAccount struct {
	Host 	string
	Username 	string
	Token 		string
}

func New(host, username, token string) JenkinsAccount {
	return JenkinsAccount{
		Username: username,
		Token: token,
		Host: host,
	}
}

func formatParams(params map[string]string) _JenkinsParameter {
	var jenkinsParams []_JenkinsParam
	for k,v := range params {
		jenkinsParams = append(jenkinsParams, _JenkinsParam{
			Name:  k,
			Value: v,
		})
	}
	return _JenkinsParameter{
		Parameter: jenkinsParams,
	}
}

func (this JenkinsAccount)Build(jobName string, params map[string]string) (bool,error){
	jenkinsParams := formatParams(params)
	jsonByte,_ := json.Marshal(jenkinsParams)
	url := fmt.Sprintf("http://%s:%s@%s/job/%s/build", this.Username, this.Token, this.Host, jobName)
	resp,code := quickrequest.PostForm(url, map[string]string{
		"json": string(jsonByte),
	}, map[string]string{})
	if code != 201 {
		return false, errors.New(string(resp))
	}
	return true,nil
}