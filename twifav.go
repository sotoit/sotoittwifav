package main

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"net/http"
	)

func main() {
	r := gin.Default()
	http.ListenAndServe(":8080",r)

	//Twitterの認証情報を取得
 api := GetTwitterApi()

 //タイムライン取得用の引数を設定
 v := url.Values{}
 v.Set("count", "20")

 //タイムラインを20件取得する
 tweets, err := api.GetHomeTimeline(v)
 if err != nil {
	  panic(err) //error時の処理
 }

 //取得したツイート分繰り返し
 for _, tweet := range tweets {
	  //ツイートがいいねされているかで条件分岐
	  if tweet.Favorited {
		   //いいねされていれば何もしない
		   fmt.Println("いいね済みでした")

	  //いいねされていない場合
	  } else {
	  _, err := api.Favorite(tweet.Id) //いいね処理
	  if err != nil {
		   panic(err) //error時の処理
	  }
 }        
}

//Twitterの認証を行うメソッド
func GetTwitterApi() *anaconda.TwitterApi {
 anaconda.SetConsumerKey("XXXXXX")
 anaconda.SetConsumerSecret("XXXXXXXX")
 api := anaconda.NewTwitterApi("XXXXXX", "XXXXXXXX")
 return api
}