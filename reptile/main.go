package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reptile/database"
	"strconv"
	"sync"
	"time"

	"github.com/tealeg/xlsx"
)

var count int

func main() {
	var wg sync.WaitGroup
	var client = &http.Client{
		Timeout: time.Second * 5,
	}
	// size == 200, pageMax == 74
	// followerMax == 149448694
	var start int = 1
	var followerCount string
	wg.Add(1)
	go func() {
		for i := 70000; i < 15000000; i += 10000 {
			followerCount = strconv.Itoa(start) + "-" + strconv.Itoa(i)
			getData(client, followerCount, "desc")
			getData(client, followerCount, "asc")
			start = i
			//time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := 15000000; i < 30000000; i += 10000 {
			followerCount = strconv.Itoa(start) + "-" + strconv.Itoa(i)
			getData(client, followerCount, "desc")
			getData(client, followerCount, "asc")
			start = i
			//time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := 30000000; i < 45000000; i += 10000 {
			followerCount = strconv.Itoa(start) + "-" + strconv.Itoa(i)
			getData(client, followerCount, "desc")
			getData(client, followerCount, "asc")
			start = i
			//time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := 45000000; i < 60000000; i += 10000 {
			followerCount = strconv.Itoa(start) + "-" + strconv.Itoa(i)
			getData(client, followerCount, "desc")
			getData(client, followerCount, "asc")
			start = i
			//time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := 60000000; i < 75000000; i += 10000 {
			followerCount = strconv.Itoa(start) + "-" + strconv.Itoa(i)
			getData(client, followerCount, "desc")
			getData(client, followerCount, "asc")
			start = i
			//time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := 75000000; i < 90000000; i += 10000 {
			followerCount = strconv.Itoa(start) + "-" + strconv.Itoa(i)
			getData(client, followerCount, "desc")
			getData(client, followerCount, "asc")
			start = i
			//time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := 90000000; i < 105000000; i += 10000 {
			followerCount = strconv.Itoa(start) + "-" + strconv.Itoa(i)
			getData(client, followerCount, "desc")
			getData(client, followerCount, "asc")
			start = i
			//time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()
	//for i := 70000; i < 149450000; i += 10000 {
	//	followerCount = strconv.Itoa(start) + "-" + strconv.Itoa(i)
	//	getData(client, followerCount, "desc", quit, i)
	//	getData(client, followerCount, "asc", quit, i)
	//	start = i
	//	//time.Sleep(1 * time.Second)
	//}
	wg.Wait()
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

		rqst.Header.Add("Cookie", "CMM_A_C_ID=93153ab3-a355-11ec-88e3-ee93e641e198; Hm_lvt_1f19c27e7e3e3255a5c79248a7f4bdf1=1647234848; UM_distinctid=17f86dcd430213-0489d25e830a6b-113f645d-1fa400-17f86dcd4313aa; Qs_lvt_411649=1647234849%2C1647235091%2C1647235096%2C1647241604%2C1647312516; LOGIN-TOKEN-FORSNS=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBJZCI6MTAwMDAsImV4cGlyZV90aW1lIjoxNjQ3ODg5MjAwLCJpYXQiOjE2NDczMTI1NTQsImlkIjozNTMyMDY4fQ.OtyPYW8ClnYvCBucjfnS1OOJlsI3-cKFjhtgBTKsmIk; Hm_lpvt_1f19c27e7e3e3255a5c79248a7f4bdf1=1647312582; Qs_pv_411649=3063137796106164000%2C3019850414182804500%2C1381674642571697200%2C1105717118416939100%2C3780877900242174000")
		rqst.Header.Add("origin", "https://www.chanmama.com")
		rqst.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
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

		//Export(author.Data.List)
		SaveDatabase(author.Data.List)
		if len(author.Data.List) == 0 {
			break
		}
	}
}

type ReptileDatas struct {
	ID            int    `json:"id"`
	AuthorID      string `json:"author_id"`
	Nickname      string `json:"nickname"`
	FollowerCount int    `json:"follower_count"`
}

func SaveDatabase(data []List) {
	count += len(data)
	for _, i2 := range data {
		fmt.Println("authorID:" + i2.AuthorID + "   nickname:" + i2.Nickname + "   followerCount:" + strconv.Itoa(i2.FollowerCount) + "   limit:" + strconv.Itoa(len(data)) + "  count:" + strconv.Itoa(count))
		database.DB.Save(&ReptileDatas{AuthorID: i2.AuthorID, Nickname: i2.Nickname, FollowerCount: i2.FollowerCount})
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
