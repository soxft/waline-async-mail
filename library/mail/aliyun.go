package mail

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/soxft/waline-async-mail/config"
	"time"
)

func sendByAliyun(mail Mail) error {
	client, err := sdk.NewClientWithAccessKey(config.Aliyun.Region, config.Aliyun.AccessKey, config.Aliyun.AccessSecret)
	if err != nil {
		return err
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = config.Aliyun.Domain
	request.Version = config.Aliyun.Version
	request.ApiName = "SingleSendMail"
	request.QueryParams["ToAddress"] = mail.ToAddress
	request.QueryParams["Subject"] = mail.Subject
	request.QueryParams["HtmlBody"] = mail.Content
	request.QueryParams["FromAlias"] = config.BlogInfo.Title
	request.QueryParams["AccountName"] = config.Aliyun.Email
	request.QueryParams["AddressType"] = "1"
	request.QueryParams["ReplyToAddress"] = "true"

	client.SetConnectTimeout(15 * time.Second)
	client.SetReadTimeout(15 * time.Second)
	_, err = client.ProcessCommonRequest(request)
	return err
}
