package main

import (
	"fmt"
	"os"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {
	//sess := session.Must(session.NewSessionWithOptions(session.Options{
	//	SharedConfigState: session.SharedConfigEnable,
	//}))
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := sns.New(sess)
	result, err := svc.ListTopics(nil)
	fmt.Println("err: ", err)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _,t := range result.Topics{
		fmt.Println((*t.TopicArn))
	}

	fmt.Println("len:", len(result.Topics))
}
