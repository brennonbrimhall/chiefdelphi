package main

import (
	"chiefdelphi/chiefdelphi"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func initalizeDB(file string) *sql.DB {
	database, _ := sql.Open("sqlite3", file)

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER UNIQUE NOT NULL, username TEXT NOT NULL, team INT)")
	statement.Exec()

	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS topics (id INTEGER UNIQUE NOT NULL, title TEXT NOT NULL)")
	statement.Exec()

	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS posts (id INTEGER UNIQUE NOT NULL, userid INTEGER NOT NULL, topicid INTEGER NOT NULL, timestamp TIMESTAMP NOT NULL, body TEXT NOT NULL)")
	statement.Exec()

	return database
}

func populateUsers(database *sql.DB) {
	statement, _ := database.Prepare("INSERT INTO users (id, username, team) VALUES (?, ?, ?)")

	users := chiefdelphi.GetUsers()

	for user := range users {
		fmt.Printf("%d,%s,%d\n", user.ID, user.Username, user.Team)
		statement.Exec(user.ID, user.Username, user.Team)
	}
}

func populateTopics(database *sql.DB) {
	lastTopicID := 0
	database.QueryRow("SELECT MAX(id) FROM topics").Scan(&lastTopicID)

	statement, _ := database.Prepare("INSERT INTO topics (id, title) VALUES (?, ?)")

	topics := chiefdelphi.GetTopicsSince(lastTopicID)

	for topic := range topics {
		fmt.Printf("%d,%s\n", topic.ID, topic.Title)
		statement.Exec(topic.ID, topic.Title)
	}
}

func populatePosts(database *sql.DB) {
	lastPostID := 0
	database.QueryRow("SELECT MAX(id) FROM posts").Scan(&lastPostID)

	statement, _ := database.Prepare("INSERT INTO posts (id, userid, topicid, `timestamp`, body) VALUES (?, ?, ?, ?, ?)")

	posts := chiefdelphi.GetPostsSince(lastPostID)

	for post := range posts {
		fmt.Printf("%d,%d,%s\n", post.ID, post.UserID, post.Timestamp.String())
		statement.Exec(post.ID, post.UserID, post.TopicID, post.Timestamp, post.Body)
	}
}

func main() {
	scrapeUsers := flag.Bool("users", false, "Scrape CD users")
	scrapeTopics := flag.Bool("topics", false, "Scrape CD topics")
	scrapePosts := flag.Bool("posts", false, "Scrape CD posts")

	flag.Parse()

	database := initalizeDB("./chiefdelphi.db")

	defer database.Close()

	if *scrapeUsers {
		populateUsers(database)
	}

	if *scrapeTopics {
		populateTopics(database)
	}

	if *scrapePosts {
		populatePosts(database)
	}
}
