package mipush

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

func GetFeed()  {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://github.com/Trumeet/MiPushFramework/releases.atom")
	fmt.Println(feed.String())
}