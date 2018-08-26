package mipush

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/mmcdole/gofeed"
	"sync"
	"time"
)

var managerUrlCN = "http://mi.greencms.net/%s/manager.apk"
var xmsfServiceUrlCN = "http://mi.greencms.net/%s/xmsf_service.apk"

var managerUrl = "https://github.com/Trumeet/MiPushFramework/releases/download/%s/manager.apk"
var xmsfServiceUrl = "https://github.com/Trumeet/MiPushFramework/releases/download/%s/xmsf_service.apk"

var releaseCache = struct {
	sync.RWMutex
	items ReleaseItems
}{}

func init() {

}

func UpdateFeed() {
	logs.Info("Updating Feed %s", time.Now().String())
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://github.com/Trumeet/MiPushFramework/releases.atom")

	var tempItems ReleaseItems

	for _, item := range feed.Items {
		releaseItem := ReleaseItem{
			Version: item.Title,
			Title:   item.Title,
			Content: item.Content,
			Link:    item.Link,
		}

		tempItems = append(tempItems, releaseItem)
	}
	releaseCache.Lock()
	releaseCache.items = tempItems
	releaseCache.Unlock()

}

func GetReleases() ReleaseItems {
	releaseCache.RLock()
	defer releaseCache.RUnlock()
	return releaseCache.items[:]
}

type ReleaseItem struct {
	Version string
	Title   string
	Content string
	Link    string
}

func (c *ReleaseItem) GetManagerUrl() string {
	return fmt.Sprintf(managerUrl, c.Version)
}
func (c *ReleaseItem) GetXmsfUrl() string {
	return fmt.Sprintf(xmsfServiceUrl, c.Version)
}

func (c *ReleaseItem) GetManagerUrlCN() string {
	return fmt.Sprintf(managerUrlCN, c.Version)
}
func (c *ReleaseItem) GetXmsfUrlCN() string {
	return fmt.Sprintf(xmsfServiceUrlCN, c.Version)
}

func (c *ReleaseItem) String() string {
	bJson, _ := json.MarshalIndent(*c, "", "    ")
	return string(bJson)
}

type ReleaseItems []ReleaseItem

func (a ReleaseItems) Len() int {
	return len(a)
}

func (a ReleaseItems) Less(i, j int) bool {
	return a[i].Version < a[j].Version
}

func (a ReleaseItems) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
