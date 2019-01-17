package main

import (
	"chiefdelphi/chiefdelphi"
	"fmt"
)

func populateUsers() {
	users := chiefdelphi.GetUsers()

	for user := range users {
		fmt.Printf("%d,%s,%d\n", user.ID, user.Username, user.Team)
	}
}

func populateTopics() {
	topics := chiefdelphi.GetTopics()

	for topic := range topics {
		fmt.Printf("%d,%s\n", topic.ID, topic.Title)
	}
}

func populatePosts() {
	posts := chiefdelphi.GetPosts()

	for post := range posts {
		fmt.Printf("%d,%d,%s,%s\n", post.ID, post.UserID, post.Timestamp.String(), post.Body)
	}
}

func main() {
	populateUsers()
	populateTopics()
	populatePosts()
}
