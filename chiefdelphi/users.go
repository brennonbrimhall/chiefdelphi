package chiefdelphi

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type getUsersPeriod int

const (
	Daily getUsersPeriod = iota
	Weekly
	Monthly
	Quarterly
	Yearly
	All
)

func (period getUsersPeriod) String() string {
	switch period {
	case Daily:
		return "daily"
	case Weekly:
		return "weekly"
	case Monthly:
		return "monthly"
	case Quarterly:
		return "quarterly"
	case Yearly:
		return "yearly"
	case All:
		return "all"
	}

	return ""
}

type getUsersOrder int

const (
	LIKES_RECIEVED getUsersOrder = iota
	LIKES_GIVEN
	TOPIC_COUNT
	POST_COUNT
	TOPICS_ENTERED
	POSTS_READ
	DAYS_VISITED
)

func (order getUsersOrder) String() string {
	switch order {
	case LIKES_RECIEVED:
		return "likes_recieved"
	case LIKES_GIVEN:
		return "likes_given"
	case TOPIC_COUNT:
		return "topic_count"
	case POST_COUNT:
		return "post_count"
	case TOPICS_ENTERED:
		return "topics_entered"
	case POSTS_READ:
		return "posts_read"
	case DAYS_VISITED:
		return "days_visited"
	}

	return ""
}

// Representation of the JSON structure returned by getting a page from the  user directory.
// It is specific to the CD API.
type getUsersResponse struct {
	DirectoryItems []struct {
		ID            int `json:"id"`
		TimeRead      int `json:"time_read"`
		LikesReceived int `json:"likes_received"`
		LikesGiven    int `json:"likes_given"`
		TopicsEntered int `json:"topics_entered"`
		TopicCount    int `json:"topic_count"`
		PostCount     int `json:"post_count"`
		PostsRead     int `json:"posts_read"`
		DaysVisited   int `json:"days_visited"`
		User          struct {
			ID               int    `json:"id"`
			Username         string `json:"username"`
			Name             string `json:"name"`
			AvatarTemplate   string `json:"avatar_template"`
			Title            string `json:"title"`
			PrimaryGroupName string `json:"primary_group_name"`
		} `json:"user"`
	} `json:"directory_items"`
	TotalRowsDirectoryItems int    `json:"total_rows_directory_items"`
	LoadMoreDirectoryItems  string `json:"load_more_directory_items"`
}

const numUsersPerPage = 50

// Make a request to the CD directory.
func getDirectoryPage(period getUsersPeriod, order getUsersOrder, page int) getUsersResponse {
	result := getUsersResponse{}
	endpoint := fmt.Sprintf("/directory_items.json?period=%s&order=%s&page=%d", period, order, page)
	_ = makeRequest(endpoint, &result)

	return result
}

// Concurrently get all users in no particular order.  This returns a channel and spawns a
// goroutine to make HTTP requests as needed.
func GetUsers() chan User {

	// Buffer up to a single page in the channel.
	yield := make(chan User, numUsersPerPage)

	go func() {
		for i := 0; ; i++ {
			page := getDirectoryPage(All, DAYS_VISITED, i)

			for _, user := range page.DirectoryItems {
				user := GetUser(user.User.Username)
				yield <- user
			}

			if len(page.DirectoryItems) == 0 {
				close(yield)
				return
			}
		}
	}()

	return yield
}

// Representation of the JSON structure returned by getting a specific user.  It is specific
// to the CD API.
type getUserResponse struct {
	UserBadges []struct {
		ID          int       `json:"id"`
		GrantedAt   time.Time `json:"granted_at"`
		Count       int       `json:"count"`
		BadgeID     int       `json:"badge_id"`
		UserID      int       `json:"user_id"`
		GrantedByID int       `json:"granted_by_id"`
	} `json:"user_badges"`
	Badges []struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		GrantCount        int         `json:"grant_count"`
		AllowTitle        bool        `json:"allow_title"`
		MultipleGrant     bool        `json:"multiple_grant"`
		Icon              string      `json:"icon"`
		Image             interface{} `json:"image"`
		Listable          bool        `json:"listable"`
		Enabled           bool        `json:"enabled"`
		BadgeGroupingID   int         `json:"badge_grouping_id"`
		System            bool        `json:"system"`
		Slug              string      `json:"slug"`
		ManuallyGrantable bool        `json:"manually_grantable"`
		BadgeTypeID       int         `json:"badge_type_id"`
	} `json:"badges"`
	BadgeTypes []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		SortOrder int    `json:"sort_order"`
	} `json:"badge_types"`
	Users []struct {
		ID               int    `json:"id"`
		Username         string `json:"username"`
		Name             string `json:"name"`
		AvatarTemplate   string `json:"avatar_template"`
		PrimaryGroupName string `json:"primary_group_name,omitempty"`
		Moderator        bool   `json:"moderator"`
		Admin            bool   `json:"admin"`
	} `json:"users"`
	User struct {
		ID                          int       `json:"id"`
		Username                    string    `json:"username"`
		Name                        string    `json:"name"`
		AvatarTemplate              string    `json:"avatar_template"`
		LastPostedAt                time.Time `json:"last_posted_at"`
		LastSeenAt                  time.Time `json:"last_seen_at"`
		BioRaw                      string    `json:"bio_raw"`
		BioCooked                   string    `json:"bio_cooked"`
		CreatedAt                   time.Time `json:"created_at"`
		CanEdit                     bool      `json:"can_edit"`
		CanEditUsername             bool      `json:"can_edit_username"`
		CanEditEmail                bool      `json:"can_edit_email"`
		CanEditName                 bool      `json:"can_edit_name"`
		CanSendPrivateMessages      bool      `json:"can_send_private_messages"`
		CanSendPrivateMessageToUser bool      `json:"can_send_private_message_to_user"`
		BioExcerpt                  string    `json:"bio_excerpt"`
		TrustLevel                  int       `json:"trust_level"`
		Moderator                   bool      `json:"moderator"`
		Admin                       bool      `json:"admin"`
		Title                       string    `json:"title"`
		UploadedAvatarID            int       `json:"uploaded_avatar_id"`
		BadgeCount                  int       `json:"badge_count"`
		CustomFields                struct {
		} `json:"custom_fields"`
		UserFields struct {
			Num1 string `json:"1"`
			Num2 string `json:"2"`
			Num3 string `json:"3"`
		} `json:"user_fields"`
		PendingCount             int         `json:"pending_count"`
		ProfileViewCount         int         `json:"profile_view_count"`
		TimeRead                 int         `json:"time_read"`
		RecentTimeRead           int         `json:"recent_time_read"`
		PrimaryGroupName         string      `json:"primary_group_name"`
		PrimaryGroupFlairURL     interface{} `json:"primary_group_flair_url"`
		PrimaryGroupFlairBgColor interface{} `json:"primary_group_flair_bg_color"`
		PrimaryGroupFlairColor   interface{} `json:"primary_group_flair_color"`
		CustomAvatarUploadID     int         `json:"custom_avatar_upload_id"`
		CustomAvatarTemplate     string      `json:"custom_avatar_template"`
		DateOfBirth              interface{} `json:"date_of_birth"`
		InvitedBy                interface{} `json:"invited_by"`
		Groups                   []struct {
			ID                        int         `json:"id"`
			Automatic                 bool        `json:"automatic"`
			Name                      string      `json:"name"`
			DisplayName               string      `json:"display_name,omitempty"`
			UserCount                 int         `json:"user_count"`
			MentionableLevel          int         `json:"mentionable_level"`
			MessageableLevel          int         `json:"messageable_level"`
			VisibilityLevel           int         `json:"visibility_level"`
			PrimaryGroup              bool        `json:"primary_group"`
			Title                     interface{} `json:"title"`
			GrantTrustLevel           interface{} `json:"grant_trust_level"`
			HasMessages               bool        `json:"has_messages,omitempty"`
			FlairURL                  interface{} `json:"flair_url"`
			FlairBgColor              interface{} `json:"flair_bg_color"`
			FlairColor                interface{} `json:"flair_color"`
			BioCooked                 interface{} `json:"bio_cooked"`
			PublicAdmission           bool        `json:"public_admission"`
			PublicExit                bool        `json:"public_exit"`
			AllowMembershipRequests   bool        `json:"allow_membership_requests"`
			FullName                  interface{} `json:"full_name"`
			DefaultNotificationLevel  int         `json:"default_notification_level"`
			MembershipRequestTemplate interface{} `json:"membership_request_template"`
		} `json:"groups"`
		GroupUsers []struct {
			GroupID           int `json:"group_id"`
			UserID            int `json:"user_id"`
			NotificationLevel int `json:"notification_level"`
		} `json:"group_users"`
		FeaturedUserBadgeIds []int `json:"featured_user_badge_ids"`
	} `json:"user"`
}

// Eagerly fetch a specific user by username.
func GetUser(username string) User {
	result := getUserResponse{}
	endpoint := fmt.Sprintf("/users/%s.json", username)
	err := makeRequest(endpoint, &result)

	if err != nil {
		return User{}
	}

	teamNumber, err := strconv.Atoi(result.User.UserFields.Num1)

	if err != nil {
		log.Printf("Bad team number from Chief Delphi on %s", endpoint)
		return User{result.User.ID, result.User.Username, -1}
	}

	return User{result.User.ID, result.User.Username, teamNumber}
}