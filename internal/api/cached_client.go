package api

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/boltdb/bolt"
	"encoding/json"
	"time"
	"github.com/herval/twtr/internal/util"
)

type CachedClient struct {
	client Client
	cache  *bolt.DB
}

func NewCachedClient(underlying Client) (*CachedClient, error) {
	db, err := bolt.Open("cache.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	c := CachedClient{
		client: underlying,
		cache:  db,
	}

	return &c, nil
}

func (c *CachedClient) Finish() {
	c.cache.Close()
}

func saveTweets(db *bolt.DB, collection string, tweets *TweetSet) error {
	return db.Update(func(tx *bolt.Tx) error {
		util.Log.Printf("Saving %d Tweets to %s", len(tweets.Tweets), collection)
		// store sets of tweets....
		b, err := tx.CreateBucketIfNotExists([]byte("tweets"))
		if err != nil {
			return err
		}

		buf, err := json.Marshal(tweets)
		if err != nil {
			return err
		}

		// ... identified by the collection var (eg timelines, a search resultset, etc)
		return b.Put([]byte(collection), buf)
	})
}

func getTweets(db *bolt.DB, collection string) (*TweetSet, error) {
	tweets := TweetSet{}

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tweets"))
		if b == nil {
			return nil
		}
		v := b.Get([]byte(collection))
		return json.Unmarshal(v, &tweets)
	})

	return &tweets, err
}

func (c *CachedClient) GetUser() (*twitter.User, error) {
	return c.client.GetUser()
}

func (c *CachedClient) GetTimeline() (*TweetSet, error) {
	t, err := getTweets(c.cache, "timeline")
	if err != nil {
		return nil, err
	}

	if t == nil || isExpired(t.UpdatedAt) {
		util.Log.Println("Cached timeline expired, refetching")
		fetched, err := c.client.GetTimeline()
		if err != nil {
			return nil, err
		}
		return fetched, saveTweets(c.cache, "timeline", fetched)
	}

	return t, err
}

func isExpired(t time.Time) bool {
	return t.Add(time.Minute * 1).Before(time.Now())
}
