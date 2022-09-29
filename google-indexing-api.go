package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/api/indexing/v3"
	"google.golang.org/api/option"
	"os"
	"strings"
)

func main() {
	var urlsPath, credentials string
	flag.StringVar(&urlsPath, "urls", "./urls", "path to file with links to pages")
	flag.StringVar(&credentials, "credentials", "./service_account.json", "path to file with credentials downloaded from https://console.developers.google.com/iam-admin/serviceaccounts")
	flag.Parse()

	ctx := context.Background()
	client, err := indexing.NewService(ctx, option.WithCredentialsFile(credentials))

	if err != nil {
		fmt.Println(err)
	}

	b, err := os.ReadFile(urlsPath)

	if err != nil {
		fmt.Println(err)
	}

	str := string(b)
	urls := strings.Split(str, "\n")

	for _, url := range urls {
		notification := indexing.UrlNotification{
			Type: "URL_UPDATED",
			Url:  url,
		}

		res, err := client.UrlNotifications.Publish(&notification).Do()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(res)
	}
}
