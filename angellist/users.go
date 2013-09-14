package angellist

import (
	"crypto/md5"
	"fmt"
	"io"
)

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Followers int    `json:"follower_count"`
	Investor  bool   `json:"investor"`

	// Links to social websites where the
	// user has a professional presence.
	Blog      string `json:"blog_url"`
	OnlineBio string `json:"online_bio_url"`
	AngelList string `json:"angellist_url"`
	Twitter   string `json:"twitter_url"`
	Facebook  string `json:"facebook_url"`
	LinkedIn  string `json:"linkedin_url"`
	AboutMe   string `json:"aboutme_url"`
	GitHub    string `json:"github_url"`
	Dribbble  string `json:"dribbble_url"`
	Behance   string `json:"behance_url"`

	// Child collections
	Locations []*Location `json:"locations"`
	Roles     []*Role     `json:"roles"`
	Skills    []*Skill    `json:"skills"`

	// Optional Investor Detail
	InvestorDetail *InvestorDetail `json:"investor_details"`

}

type Location struct {
	Id      int    `json:"id"`
    Type    string `json:"tag_type"`
    Name    string `json:"name"`
    Display string `json:"display_name"`
    Url     string `json:"angellist_url"`
}

type Role struct {
	Id      int    `json:"id"`
    Type    string `json:"tag_type"`
    Name    string `json:"name"`
    Display string `json:"display_name"`
    Url     string `json:"angellist_url"`
}

type Skill struct {
	Id      int    `json:"id"`
    Type    string `json:"tag_type"`
    Name    string `json:"name"`
    Display string `json:"display_name"`
    Url     string `json:"angellist_url"`
}

type Market struct {
	Id      int    `json:"id"`
    Type    string `json:"tag_type"`
    Name    string `json:"name"`
    Display string `json:"display_name"`
    Url     string `json:"angellist_url"`
}

type Investment struct {
	Id      int    `json:"id"`
    Name    string `json:"name"`
	Quality int    `json:"quality"`
}

type InvestorDetail struct {
	Accreditation string `json:"accreditation"`

	// Child collections
	Locations   []*Location   `json:"locations"`
	Investments []*Investment `json:"investment"`
	Markets     []*Market     `json:"markets"`
}


type UserResource struct {
	client *Client
}

// Get a user's information given an id.
func (u *UserResource) Get(id int) (*User, error) {
	path := fmt.Sprintf("/1/users/%d?include_details=investor", id)
	user := User{}
	if err := u.client.do("GET", path, nil, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// Get a user's information given an id.
func (u *UserResource) GetMulti(ids ...int) ([]*User, error) {

	return nil, nil
}

// Search for a user given a URL slug.
func (u *UserResource) GetSlug(slug string) (*User, error) {
	path := fmt.Sprintf("/1/users/search?include_details=investor&slug=%s", slug)
	user := User{}
	if err := u.client.do("GET", path, nil, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// Search for a user given an MD5 email hash.
func (u *UserResource) GetEmail(email string) (*User, error) {
	// calculate emails md5 hash
	h := md5.New()
    io.WriteString(h, email)
    hash := h.Sum(nil)

	path := fmt.Sprintf("/1/users/search?include_details=investor&md5=%x", hash)
	user := User{}
	if err := u.client.do("GET", path, nil, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
