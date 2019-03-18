package chiefdelphi

import (
	"fmt"
	"log"
	"time"
)

const FIRST_TOPIC_ID = 29276

type getTopicResponse struct {
	PostStream struct {
		Posts []struct {
			ID                       int         `json:"id"`
			Name                     string      `json:"name"`
			Username                 string      `json:"username"`
			AvatarTemplate           string      `json:"avatar_template"`
			CreatedAt                time.Time   `json:"created_at"`
			Cooked                   string      `json:"cooked"`
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
			Read                     bool        `json:"read"`
			UserTitle                string      `json:"user_title"`
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
			UserCustomFields   struct {
				UserField4 string `json:"user_field_4"`
			} `json:"user_custom_fields,omitempty"`
			UserCreatedAt   time.Time   `json:"user_created_at"`
			UserDateOfBirth interface{} `json:"user_date_of_birth"`
			ReplyToUser     struct {
				Username       string `json:"username"`
				AvatarTemplate string `json:"avatar_template"`
			} `json:"reply_to_user,omitempty"`
			LinkCounts []struct {
				URL        string `json:"url"`
				Internal   bool   `json:"internal"`
				Reflection bool   `json:"reflection"`
				Title      string `json:"title"`
				Clicks     int    `json:"clicks"`
			} `json:"link_counts,omitempty"`
		} `json:"posts"`
		Stream []int `json:"stream"`
	} `json:"post_stream"`
	TimelineLookup  [][]int `json:"timeline_lookup"`
	SuggestedTopics []struct {
		ID                 int           `json:"id"`
		Title              string        `json:"title"`
		FancyTitle         string        `json:"fancy_title"`
		Slug               string        `json:"slug"`
		PostsCount         int           `json:"posts_count"`
		ReplyCount         int           `json:"reply_count"`
		HighestPostNumber  int           `json:"highest_post_number"`
		ImageURL           interface{}   `json:"image_url"`
		CreatedAt          time.Time     `json:"created_at"`
		LastPostedAt       time.Time     `json:"last_posted_at"`
		Bumped             bool          `json:"bumped"`
		BumpedAt           time.Time     `json:"bumped_at"`
		Unseen             bool          `json:"unseen"`
		LastReadPostNumber int           `json:"last_read_post_number"`
		Unread             int           `json:"unread"`
		NewPosts           int           `json:"new_posts"`
		Pinned             bool          `json:"pinned"`
		Unpinned           interface{}   `json:"unpinned"`
		Visible            bool          `json:"visible"`
		Closed             bool          `json:"closed"`
		Archived           bool          `json:"archived"`
		NotificationLevel  int           `json:"notification_level"`
		Bookmarked         bool          `json:"bookmarked"`
		Liked              bool          `json:"liked"`
		Tags               []interface{} `json:"tags"`
		Archetype          string        `json:"archetype"`
		LikeCount          int           `json:"like_count"`
		Views              int           `json:"views"`
		CategoryID         int           `json:"category_id"`
		FeaturedLink       interface{}   `json:"featured_link"`
		Posters            []struct {
			Extras      interface{} `json:"extras"`
			Description string      `json:"description"`
			User        struct {
				ID             int    `json:"id"`
				Username       string `json:"username"`
				Name           string `json:"name"`
				AvatarTemplate string `json:"avatar_template"`
			} `json:"user"`
		} `json:"posters"`
	} `json:"suggested_topics"`
	Tags           []interface{} `json:"tags"`
	ID             int           `json:"id"`
	Title          string        `json:"title"`
	FancyTitle     string        `json:"fancy_title"`
	PostsCount     int           `json:"posts_count"`
	CreatedAt      time.Time     `json:"created_at"`
	Views          int           `json:"views"`
	ReplyCount     int           `json:"reply_count"`
	LikeCount      int           `json:"like_count"`
	LastPostedAt   time.Time     `json:"last_posted_at"`
	Visible        bool          `json:"visible"`
	Closed         bool          `json:"closed"`
	Archived       bool          `json:"archived"`
	HasSummary     bool          `json:"has_summary"`
	Archetype      string        `json:"archetype"`
	Slug           string        `json:"slug"`
	CategoryID     int           `json:"category_id"`
	WordCount      int           `json:"word_count"`
	DeletedAt      interface{}   `json:"deleted_at"`
	UserID         int           `json:"user_id"`
	FeaturedLink   interface{}   `json:"featured_link"`
	PinnedGlobally bool          `json:"pinned_globally"`
	PinnedAt       interface{}   `json:"pinned_at"`
	PinnedUntil    interface{}   `json:"pinned_until"`
	Draft          interface{}   `json:"draft"`
	DraftKey       string        `json:"draft_key"`
	DraftSequence  int           `json:"draft_sequence"`
	Unpinned       interface{}   `json:"unpinned"`
	Pinned         bool          `json:"pinned"`
	Details        struct {
		CreatedBy struct {
			ID             int    `json:"id"`
			Username       string `json:"username"`
			Name           string `json:"name"`
			AvatarTemplate string `json:"avatar_template"`
		} `json:"created_by"`
		LastPoster struct {
			ID             int    `json:"id"`
			Username       string `json:"username"`
			Name           string `json:"name"`
			AvatarTemplate string `json:"avatar_template"`
		} `json:"last_poster"`
		Participants []struct {
			ID                       int         `json:"id"`
			Username                 string      `json:"username"`
			Name                     string      `json:"name"`
			AvatarTemplate           string      `json:"avatar_template"`
			PostCount                int         `json:"post_count"`
			PrimaryGroupName         string      `json:"primary_group_name"`
			PrimaryGroupFlairURL     interface{} `json:"primary_group_flair_url"`
			PrimaryGroupFlairColor   interface{} `json:"primary_group_flair_color"`
			PrimaryGroupFlairBgColor interface{} `json:"primary_group_flair_bg_color"`
		} `json:"participants"`
		NotificationLevel  int  `json:"notification_level"`
		CanCreatePost      bool `json:"can_create_post"`
		CanReplyAsNewTopic bool `json:"can_reply_as_new_topic"`
		CanFlagTopic       bool `json:"can_flag_topic"`
	} `json:"details"`
	CurrentPostNumber int         `json:"current_post_number"`
	HighestPostNumber int         `json:"highest_post_number"`
	DeletedBy         interface{} `json:"deleted_by"`
	ActionsSummary    []struct {
		ID     int  `json:"id"`
		Count  int  `json:"count"`
		Hidden bool `json:"hidden"`
		CanAct bool `json:"can_act"`
	} `json:"actions_summary"`
	ChunkSize         int         `json:"chunk_size"`
	Bookmarked        interface{} `json:"bookmarked"`
	TopicTimer        interface{} `json:"topic_timer"`
	PrivateTopicTimer interface{} `json:"private_topic_timer"`
	MessageBusLastID  int         `json:"message_bus_last_id"`
	ParticipantCount  int         `json:"participant_count"`
}

type getLatestTopicsResponse struct {
	Users []struct {
		ID             int    `json:"id"`
		Username       string `json:"username"`
		Name           string `json:"name"`
		AvatarTemplate string `json:"avatar_template"`
	} `json:"users"`
	PrimaryGroups []struct {
		ID           int         `json:"id"`
		Name         string      `json:"name"`
		FlairURL     interface{} `json:"flair_url"`
		FlairBgColor interface{} `json:"flair_bg_color"`
		FlairColor   interface{} `json:"flair_color"`
	} `json:"primary_groups"`
	TopicList struct {
		CanCreateTopic bool        `json:"can_create_topic"`
		MoreTopicsURL  string      `json:"more_topics_url"`
		Draft          interface{} `json:"draft"`
		DraftKey       string      `json:"draft_key"`
		DraftSequence  int         `json:"draft_sequence"`
		PerPage        int         `json:"per_page"`
		Topics         []struct {
			ID                 int           `json:"id"`
			Title              string        `json:"title"`
			FancyTitle         string        `json:"fancy_title"`
			Slug               string        `json:"slug"`
			PostsCount         int           `json:"posts_count"`
			ReplyCount         int           `json:"reply_count"`
			HighestPostNumber  int           `json:"highest_post_number"`
			ImageURL           string        `json:"image_url"`
			CreatedAt          time.Time     `json:"created_at"`
			LastPostedAt       time.Time     `json:"last_posted_at"`
			Bumped             bool          `json:"bumped"`
			BumpedAt           time.Time     `json:"bumped_at"`
			Unseen             bool          `json:"unseen"`
			Pinned             bool          `json:"pinned"`
			Unpinned           interface{}   `json:"unpinned"`
			Excerpt            string        `json:"excerpt,omitempty"`
			Visible            bool          `json:"visible"`
			Closed             bool          `json:"closed"`
			Archived           bool          `json:"archived"`
			Bookmarked         interface{}   `json:"bookmarked"`
			Liked              interface{}   `json:"liked"`
			Tags               []interface{} `json:"tags"`
			Views              int           `json:"views"`
			LikeCount          int           `json:"like_count"`
			HasSummary         bool          `json:"has_summary"`
			Archetype          string        `json:"archetype"`
			LastPosterUsername string        `json:"last_poster_username"`
			CategoryID         int           `json:"category_id"`
			PinnedGlobally     bool          `json:"pinned_globally"`
			FeaturedLink       interface{}   `json:"featured_link"`
			Posters            []struct {
				Extras         interface{} `json:"extras"`
				Description    string      `json:"description"`
				UserID         int         `json:"user_id"`
				PrimaryGroupID interface{} `json:"primary_group_id"`
			} `json:"posters"`
			LastReadPostNumber interface{} `json:"last_read_post_number,omitempty"`
			Unread             int         `json:"unread,omitempty"`
			NewPosts           int         `json:"new_posts,omitempty"`
			NotificationLevel  int         `json:"notification_level,omitempty"`
		} `json:"topics"`
	} `json:"topic_list"`
}

func getLatestTopicID() (int, error) {
	result := getLatestTopicsResponse{}
	endpoint := fmt.Sprintf("/latest.json")
	err := makeRequest(endpoint, &result)

	if err != nil {
		return 0, err
	}

	biggestID := 0

	for _, topic := range result.TopicList.Topics {
		if biggestID < topic.ID {
			biggestID = topic.ID
		}
	}

	return biggestID, nil
}

func GetTopic(ID int) Topic {
	result := getTopicResponse{}
	endpoint := fmt.Sprintf("/t/%d.json", ID)
	_ = makeRequest(endpoint, &result)

	return Topic{ID, result.Title}
}

func GetTopics() chan Topic {
	return GetTopicsSince(FIRST_TOPIC_ID)
}

func GetTopicsSince(i int) chan Topic {
	yield := make(chan Topic)

	latestTopicID, err := getLatestTopicID()

	if err != nil {
		log.Print("Unable to get the latest topic ID from Chief Delphi.")
		close(yield)
		return yield
	}

	go func() {
		for ; i <= latestTopicID; i++ {
			topic := GetTopic(i)
			yield <- topic
		}

		close(yield)
	}()

	return yield
}