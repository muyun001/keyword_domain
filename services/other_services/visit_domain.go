package other_services

import (
	"github.com/panwenbin/ghttpclient"
	"strings"
	"time"
)

// 访问url，获取源码
func VisitDomainGetHtml(searchDomain string) (string, error) {
	var html string
	var err error
	if strings.HasPrefix(searchDomain, "http://") || strings.HasPrefix(searchDomain, "https://") {
		html, err = getHtml(searchDomain)
	} else {
		getStrs := []string{"http://", "https://"}
		for _, str := range getStrs {
			html, err = getHtml(str + searchDomain)
			if err == nil {
				break
			}
		}
	}

	if err != nil {
		return "", err
	}

	return html, nil
}

func getHtml(url string) (html string, err error) {
	client := ghttpclient.NewClient().Timeout(time.Second * 30).Url(url).Headers(nil).Get()
	body, err := client.ReadBodyClose()
	if err != nil {
		return "", err
	}

	return string(body), nil
}