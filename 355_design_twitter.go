// TODO Try Another Better Approach
package main

import (
	"fmt"
	"sort"
)

type Post struct {
	id        int
	timestamp int
}

type Posts []Post

func (p Posts) Len() int {
	return len(p)
}

func (p Posts) Less(i, j int) bool {
	return p[i].timestamp > p[j].timestamp
}

func (p Posts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Twitter struct {
	postMap   map[int][]Post
	followMap map[int]map[int]bool
	timestamp int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		postMap:   make(map[int][]Post),
		followMap: make(map[int]map[int]bool),
		timestamp: 0,
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	_, is_exist := this.postMap[userId]
	if !is_exist {
		this.postMap[userId] = []Post{}
	}
	this.timestamp++
	this.postMap[userId] = append(this.postMap[userId], Post{tweetId, this.timestamp})
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	feeds := []Post{}

	if list, is_exist := this.postMap[userId]; is_exist {
		for i := 0; i < len(list); i++ {
			feeds = append(feeds, list[i])
		}
	}

	if fm, is_exist := this.followMap[userId]; is_exist {
		for f := range fm {
			if list, is_exist := this.postMap[f]; is_exist {
				for i := 0; i < len(list); i++ {
					feeds = append(feeds, list[i])
				}
			}
		}
	}

	sort.Sort(Posts(feeds))

	results := []int{}
	for i := 0; i < 10 && i < len(feeds); i++ {
		results = append(results, feeds[i].id)
	}
	return results
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	_, is_exist := this.followMap[followerId]
	if !is_exist {
		this.followMap[followerId] = make(map[int]bool)
	}
	this.followMap[followerId][followeeId] = true
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	if fm, is_exist := this.followMap[followerId]; is_exist {
		delete(fm, followeeId)
	}
}

func main() {
	twitter := Constructor()

	// User 1 posts a new tweet (id = 5).
	twitter.PostTweet(1, 5)

	// User 1's news feed should return a list with 1 tweet id -> [5].
	fmt.Println(twitter.GetNewsFeed(1))

	// User 1 follows user 2.
	twitter.Follow(1, 2)

	// User 2 posts a new tweet (id = 6).
	twitter.PostTweet(2, 6)

	// User 1's news feed should return a list with 2 tweet ids -> [6, 5].
	// Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
	fmt.Println(twitter.GetNewsFeed(1))

	// User 1 unfollows user 2.
	twitter.Unfollow(1, 2)

	// User 1's news feed should return a list with 1 tweet id -> [5],
	// since user 1 is no longer following user 2.
	fmt.Println(twitter.GetNewsFeed(1))
}
