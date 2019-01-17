package chiefdelphi

import (
	"fmt"
	"log"
	"time"
)

const FIRST_POST_ID = 36001

type getPostResponse struct {
	ID                       int         `json:"id"`
	Name                     string      `json:"name"`
	Username                 string      `json:"username"`
	AvatarTemplate           string      `json:"avatar_template"`
	CreatedAt                time.Time   `json:"created_at"`
	Cooked                   string      `json:"cooked"`
	Ignored                  bool        `json:"ignored"`
	PostNumber               int         `json:"post_number"`
	PostType                 int         `json:"post_type"`
	UpdatedAt                time.Time   `json:"updated_at"`
	ReplyCount               int         `json:"reply_count"`
	ReplyToPostNumber        interface{} `json:"reply_to_post_number"`
	QuoteCount               int         `json:"quote_count"`
	AvgTime                  int         `json:"avg_time"`
	IncomingLinkCount        int         `json:"incoming_link_count"`
	Reads                    int         `json:"reads"`
	Score                    float64     `json:"score"`
	Yours                    bool        `json:"yours"`
	TopicID                  int         `json:"topic_id"`
	TopicSlug                string      `json:"topic_slug"`
	DisplayUsername          string      `json:"display_username"`
	PrimaryGroupName         string      `json:"primary_group_name"`
	PrimaryGroupFlairURL     interface{} `json:"primary_group_flair_url"`
	PrimaryGroupFlairBgColor interface{} `json:"primary_group_flair_bg_color"`
	PrimaryGroupFlairColor   interface{} `json:"primary_group_flair_color"`
	Version                  int         `json:"version"`
	CanEdit                  bool        `json:"can_edit"`
	CanDelete                bool        `json:"can_delete"`
	CanRecover               bool        `json:"can_recover"`
	CanWiki                  bool        `json:"can_wiki"`
	UserTitle                string      `json:"user_title"`
	Raw                      string      `json:"raw"`
	ActionsSummary           []struct {
		ID     int  `json:"id"`
		CanAct bool `json:"can_act"`
	} `json:"actions_summary"`
	Moderator          bool        `json:"moderator"`
	Admin              bool        `json:"admin"`
	Staff              bool        `json:"staff"`
	UserID             int         `json:"user_id"`
	Hidden             bool        `json:"hidden"`
	TrustLevel         int         `json:"trust_level"`
	DeletedAt          interface{} `json:"deleted_at"`
	UserDeleted        bool        `json:"user_deleted"`
	EditReason         interface{} `json:"edit_reason"`
	CanViewEditHistory bool        `json:"can_view_edit_history"`
	Wiki               bool        `json:"wiki"`
	UserCreatedAt      time.Time   `json:"user_created_at"`
	UserDateOfBirth    interface{} `json:"user_date_of_birth"`
}

type getLatestPostsResponse struct {
	LatestPosts []struct {
		ID                       int         `json:"id"`
		Name                     string      `json:"name"`
		Username                 string      `json:"username"`
		AvatarTemplate           string      `json:"avatar_template"`
		CreatedAt                time.Time   `json:"created_at"`
		Cooked                   string      `json:"cooked"`
		Ignored                  bool        `json:"ignored"`
		PostNumber               int         `json:"post_number"`
		PostType                 int         `json:"post_type"`
		UpdatedAt                time.Time   `json:"updated_at"`
		ReplyCount               int         `json:"reply_count"`
		ReplyToPostNumber        interface{} `json:"reply_to_post_number"`
		QuoteCount               int         `json:"quote_count"`
		AvgTime                  interface{} `json:"avg_time"`
		IncomingLinkCount        int         `json:"incoming_link_count"`
		Reads                    int         `json:"reads"`
		Score                    int         `json:"score"`
		Yours                    bool        `json:"yours"`
		TopicID                  int         `json:"topic_id"`
		TopicSlug                string      `json:"topic_slug"`
		TopicTitle               string      `json:"topic_title"`
		TopicHTMLTitle           string      `json:"topic_html_title"`
		CategoryID               int         `json:"category_id"`
		DisplayUsername          string      `json:"display_username"`
		PrimaryGroupName         string      `json:"primary_group_name"`
		PrimaryGroupFlairURL     interface{} `json:"primary_group_flair_url"`
		PrimaryGroupFlairBgColor interface{} `json:"primary_group_flair_bg_color"`
		PrimaryGroupFlairColor   interface{} `json:"primary_group_flair_color"`
		Version                  int         `json:"version"`
		CanEdit                  bool        `json:"can_edit"`
		CanDelete                bool        `json:"can_delete"`
		CanRecover               bool        `json:"can_recover"`
		CanWiki                  bool        `json:"can_wiki"`
		UserTitle                string      `json:"user_title"`
		Raw                      string      `json:"raw"`
		ActionsSummary           []struct {
			ID     int  `json:"id"`
			CanAct bool `json:"can_act"`
		} `json:"actions_summary"`
		Moderator          bool        `json:"moderator"`
		Admin              bool        `json:"admin"`
		Staff              bool        `json:"staff"`
		UserID             int         `json:"user_id"`
		Hidden             bool        `json:"hidden"`
		TrustLevel         int         `json:"trust_level"`
		DeletedAt          interface{} `json:"deleted_at"`
		UserDeleted        bool        `json:"user_deleted"`
		EditReason         interface{} `json:"edit_reason"`
		CanViewEditHistory bool        `json:"can_view_edit_history"`
		Wiki               bool        `json:"wiki"`
		UserCreatedAt      time.Time   `json:"user_created_at"`
		UserDateOfBirth    interface{} `json:"user_date_of_birth"`
		ReplyToUser        struct {
			Username       string `json:"username"`
			AvatarTemplate string `json:"avatar_template"`
		} `json:"reply_to_user,omitempty"`
	} `json:"latest_posts"`
}

func getLatestPostID() (int, error) {
	result := getLatestPostsResponse{}
	endpoint := fmt.Sprintf("/posts.json")
	err := makeRequest(endpoint, &result)

	if err != nil {
		return 0, err
	}

	biggestID := 0

	for _, post := range result.LatestPosts {
		if biggestID < post.ID {
			biggestID = post.ID
		}
	}

	return biggestID, nil
}

func GetPost(ID int) Post {
	result := getPostResponse{}
	endpoint := fmt.Sprintf("/posts/%d.json", ID)
	err := makeRequest(endpoint, &result)

	if err != nil {
		return Post{}
	}

	return Post{ID, result.UserID, result.TopicID, result.CreatedAt, result.Raw}
}

func GetPosts() chan Post {
	return GetPostsSince(FIRST_POST_ID)
}

func GetPostsSince(i int) chan Post {
	yield := make(chan Post)

	latestPostID, err := getLatestPostID()

	if err != nil {
		log.Print("Unable to get the latest topic ID from Chief Delphi.")
		close(yield)
		return yield
	}

	go func() {
		for ; i <= latestPostID; i++ {
			post := GetPost(i)
			yield <- post

			sleep()
		}

		close(yield)
	}()

	return yield
}