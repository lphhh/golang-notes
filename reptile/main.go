package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/tealeg/xlsx"
)

var count int

func main() {
	var client = &http.Client{
		Timeout: time.Second * 5,
	}
	// size == 200, pageMax == 74
	// followerMax == 149448694
	var start int = 1
	var followerCount string
	for i := 79000; i < 149448695; i += 10000 {
		followerCount = strconv.Itoa(start) + "-" + strconv.Itoa(i)
		getData(client, followerCount, "desc")
		getData(client, followerCount, "asc")
		start = i
		time.Sleep(1 * time.Second)
	}

}

func getData(client *http.Client, followerCount, orderBy string) {
	for p := 1; ; p++ {
		rqst, err := http.NewRequest("GET",
			"https://api-service.chanmama.com/v2/home/author/search?page="+
				strconv.Itoa(p)+
				"&follower_count="+followerCount+
				"&sort=follower_count"+
				"&order_by="+orderBy+
				"&size=200",
			nil)
		if err != nil {
			fmt.Println("New request failed:", err)
			return
		}

		rqst.Header.Add("Cookie", "CMM_A_C_ID=7bad0832-a2f3-11ec-88e3-ee93e641e198; Hm_lvt_1f19c27e7e3e3255a5c79248a7f4bdf1=1647192718; LOGIN-TOKEN-FORSNS=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBJZCI6MTAwMDAsImV4cGlyZV90aW1lIjoxNjQ3ODAyODAwLCJpYXQiOjE2NDcyNzMxODksImlkIjozNTMyMDY4fQ.edrJOkEkagmd6jEsrsm9YTohLEwag96TPsAEojczNxM; Hm_lpvt_1f19c27e7e3e3255a5c79248a7f4bdf1=1647273191; Qs_lvt_411649=1647192718%2C1647273191; Qs_pv_411649=2387410460744642000%2C1410934070754850300%2C2618486341181278700%2C4372307893197005300%2C3507494727943831600")
		rqst.Header.Add("origin", "https://www.chanmama.com")
		rqst.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
		rqst.Header.Add("referer", "https://www.chanmama.com/bloggerRank?keyword=")

		rsps, err := client.Do(rqst)
		if err != nil {
			fmt.Println("Request failed:", err)
			return
		}
		defer rsps.Body.Close()

		body, _ := ioutil.ReadAll(rsps.Body)

		var author Author
		json.Unmarshal(body, &author)

		// fmt.Println(author.Data.List)

		Export(author.Data.List)
		if len(author.Data.List) == 0 {
			break
		}
	}
}

func Export(data []List) {
	count += len(data)
	// file := xlsx.NewFile()
	// sheet, _ := file.AddSheet("Sheet1")
	file, err := xlsx.OpenFile("data.xlsx")
	if err != nil {
		panic(err)
	}
	sheet := file.Sheets[0]
	for _, v := range data {
		row := sheet.AddRow()
		// row.SetHeightCM(1) //设置每行的高度
		cell := row.AddCell()
		cell.Value = v.AuthorID
		cell = row.AddCell()
		cell.Value = v.Nickname
		fmt.Println("authorID:" + v.AuthorID + "   nickname:" + v.Nickname + "   followerCount:" + strconv.Itoa(v.FollowerCount) + "   limit:" + strconv.Itoa(len(data)) + "  count:" + strconv.Itoa(count))
	}

	err = file.Save("data.xlsx")
	if err != nil {
		panic(err)
	}
}

type Author struct {
	Data    Data `json:"data"`
	ErrCode int  `json:"errCode"`
}
type SingleTags struct {
	First   string        `json:"first"`
	Second  []interface{} `json:"second"`
	Primary bool          `json:"primary"`
}
type List struct {
	AuthorID                string        `json:"author_id"`
	Nickname                string        `json:"nickname"`
	Avatar                  string        `json:"avatar"`
	ProductCount            int           `json:"product_count"`
	ID                      string        `json:"id"`
	Label                   string        `json:"label"`
	IsFav                   int           `json:"is_fav"`
	SingleTags              SingleTags    `json:"single_tags"`
	VerificationType        int           `json:"verification_type"`
	Gender                  int           `json:"gender"`
	UniqueID                string        `json:"unique_id"`
	ShortID                 string        `json:"short_id"`
	ProductVideoCount       int           `json:"product_video_count"`
	FollowerCount           int           `json:"follower_count"`
	TotalFavorited          int64         `json:"total_favorited"`
	MarkDelete              int           `json:"mark_delete"`
	AwemeCount              int           `json:"aweme_count"`
	FollowerIncr            int           `json:"follower_incr"`
	DiggIncr                int           `json:"digg_incr"`
	Commerce                int           `json:"commerce"`
	LiveRoomID              string        `json:"live_room_id"`
	LiveRoomStatus          int           `json:"live_room_status"`
	ReputationPercentage    int           `json:"reputation_percentage"`
	ReputationScore         int           `json:"reputation_score"`
	ReputationRecentSales   string        `json:"reputation_recent_sales"`
	ReputationLevel         int           `json:"reputation_level"`
	LiveAverageVolume30     int           `json:"live_average_volume_30"`
	LiveAverageInteract30   float64       `json:"live_average_interact_30"`
	LiveAverageAmount30     int           `json:"live_average_amount_30"`
	TotalAmount30           int           `json:"total_amount_30"`
	AwemeDiggFollowerRation float64       `json:"aweme_digg_follower_ration"`
	AwemeDiggMedium         int           `json:"aweme_digg_medium"`
	LivePvMedium            int           `json:"live_pv_medium"`
	BringLiveProducts       bool          `json:"bring_live_products"`
	BringProducts30         bool          `json:"bring_products_30"`
	WithGoodsCategory       []interface{} `json:"with_goods_category"`
	CollectProcessStatus    int           `json:"collect_process_status"`
}
type PageInfo struct {
	Page       int `json:"page"`
	TotalCount int `json:"totalCount"`
	TotalPage  int `json:"totalPage"`
	Size       int `json:"size"`
}
type Data struct {
	List             []List        `json:"list"`
	PageInfo         PageInfo      `json:"page_info"`
	RequestID        string        `json:"request_id"`
	RecommendKeyword string        `json:"recommend_keyword"`
	DyAuthors        []interface{} `json:"dy_authors"`
}
