package main;


import "sync"


type TwitterAccount struct {
	TwitterId            string `json:"twitter_id"`
	BulliedTweet         []string
	mu                   sync.RWMutex
}


func (ta *TwitterAccount) init(twitterId string) {
	ta.TwitterId = twitterId;
}

func (ta *TwitterAccount) monitorAccount() {
	
}


func (ta *TwitterAccount) getTweets()  {
	
}