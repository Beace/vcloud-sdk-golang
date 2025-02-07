package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/TTvcloud/vcloud-sdk-golang/base"

	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func main() {
	inlinePolicy := new(base.Policy)

	// 给一个权限是所有action的allow的statement
	statement := base.NewAllowStatement([]string{"iam:*"}, []string{})
	inlinePolicy.Statement = append(inlinePolicy.Statement, statement)

	vod.NewInstance().SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	ret, _ := vod.Instance.SignSts2(inlinePolicy, time.Hour)
	b, _ := json.Marshal(ret)
	fmt.Println(string(b))
}
