// Code generated by protoc-gen-go.
// source: reddit.proto
// DO NOT EDIT!

/*
Package redditproto is a generated protocol buffer package.

It is generated from these files:
	reddit.proto
	useragent.proto

It has these top-level messages:
	Comment
	Account
	Link
	LinkSet
	Message
	Subreddit
	UserAgent
*/
package redditproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Data type t1_
type Comment struct {
	ApprovedBy          *string `protobuf:"bytes,1,opt,name=approved_by" json:"approved_by,omitempty"`
	Author              *string `protobuf:"bytes,2,opt,name=author" json:"author,omitempty"`
	AuthorFlairCssClass *string `protobuf:"bytes,3,opt,name=author_flair_css_class" json:"author_flair_css_class,omitempty"`
	AuthorFlairText     *string `protobuf:"bytes,4,opt,name=author_flair_text" json:"author_flair_text,omitempty"`
	BannedBy            *string `protobuf:"bytes,5,opt,name=banned_by" json:"banned_by,omitempty"`
	Body                *string `protobuf:"bytes,6,opt,name=body" json:"body,omitempty"`
	BodyHtml            *string `protobuf:"bytes,7,opt,name=body_html" json:"body_html,omitempty"`
	// Field 8 is reserved for "edited".
	Gilded        *int32     `protobuf:"varint,9,opt,name=gilded" json:"gilded,omitempty"`
	LinkAuthor    *string    `protobuf:"bytes,10,opt,name=link_author" json:"link_author,omitempty"`
	LinkUrl       *string    `protobuf:"bytes,11,opt,name=link_url" json:"link_url,omitempty"`
	NumReports    *int32     `protobuf:"varint,12,opt,name=num_reports" json:"num_reports,omitempty"`
	ParentId      *string    `protobuf:"bytes,13,opt,name=parent_id" json:"parent_id,omitempty"`
	Replies       []*Comment `protobuf:"bytes,14,rep,name=replies" json:"replies,omitempty"`
	Subreddit     *string    `protobuf:"bytes,15,opt,name=subreddit" json:"subreddit,omitempty"`
	SubredditId   *string    `protobuf:"bytes,16,opt,name=subreddit_id" json:"subreddit_id,omitempty"`
	Distinguished *string    `protobuf:"bytes,17,opt,name=distinguished" json:"distinguished,omitempty"`
	// Implements Created
	Created    *float64 `protobuf:"fixed64,18,opt,name=created" json:"created,omitempty"`
	CreatedUtc *float64 `protobuf:"fixed64,19,opt,name=created_utc" json:"created_utc,omitempty"`
	// Implements Votable
	Ups   *int32 `protobuf:"varint,20,opt,name=ups" json:"ups,omitempty"`
	Downs *int32 `protobuf:"varint,21,opt,name=downs" json:"downs,omitempty"`
	Likes *bool  `protobuf:"varint,22,opt,name=likes" json:"likes,omitempty"`
	// Implements Thing
	Id   *string `protobuf:"bytes,23,opt,name=id" json:"id,omitempty"`
	Name *string `protobuf:"bytes,24,opt,name=name" json:"name,omitempty"`
	// Message-Comment hybrid fields; these are present when the comment is
	// actually a Message in the inbox, but Reddit still labels the object "t1"
	// (Comment) because that is its original type.
	Subject          *string                   `protobuf:"bytes,25,opt,name=subject" json:"subject,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}

var extRange_Comment = []proto.ExtensionRange{
	{100, 536870911},
}

func (*Comment) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Comment
}
func (m *Comment) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *Comment) GetApprovedBy() string {
	if m != nil && m.ApprovedBy != nil {
		return *m.ApprovedBy
	}
	return ""
}

func (m *Comment) GetAuthor() string {
	if m != nil && m.Author != nil {
		return *m.Author
	}
	return ""
}

func (m *Comment) GetAuthorFlairCssClass() string {
	if m != nil && m.AuthorFlairCssClass != nil {
		return *m.AuthorFlairCssClass
	}
	return ""
}

func (m *Comment) GetAuthorFlairText() string {
	if m != nil && m.AuthorFlairText != nil {
		return *m.AuthorFlairText
	}
	return ""
}

func (m *Comment) GetBannedBy() string {
	if m != nil && m.BannedBy != nil {
		return *m.BannedBy
	}
	return ""
}

func (m *Comment) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func (m *Comment) GetBodyHtml() string {
	if m != nil && m.BodyHtml != nil {
		return *m.BodyHtml
	}
	return ""
}

func (m *Comment) GetGilded() int32 {
	if m != nil && m.Gilded != nil {
		return *m.Gilded
	}
	return 0
}

func (m *Comment) GetLinkAuthor() string {
	if m != nil && m.LinkAuthor != nil {
		return *m.LinkAuthor
	}
	return ""
}

func (m *Comment) GetLinkUrl() string {
	if m != nil && m.LinkUrl != nil {
		return *m.LinkUrl
	}
	return ""
}

func (m *Comment) GetNumReports() int32 {
	if m != nil && m.NumReports != nil {
		return *m.NumReports
	}
	return 0
}

func (m *Comment) GetParentId() string {
	if m != nil && m.ParentId != nil {
		return *m.ParentId
	}
	return ""
}

func (m *Comment) GetReplies() []*Comment {
	if m != nil {
		return m.Replies
	}
	return nil
}

func (m *Comment) GetSubreddit() string {
	if m != nil && m.Subreddit != nil {
		return *m.Subreddit
	}
	return ""
}

func (m *Comment) GetSubredditId() string {
	if m != nil && m.SubredditId != nil {
		return *m.SubredditId
	}
	return ""
}

func (m *Comment) GetDistinguished() string {
	if m != nil && m.Distinguished != nil {
		return *m.Distinguished
	}
	return ""
}

func (m *Comment) GetCreated() float64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func (m *Comment) GetCreatedUtc() float64 {
	if m != nil && m.CreatedUtc != nil {
		return *m.CreatedUtc
	}
	return 0
}

func (m *Comment) GetUps() int32 {
	if m != nil && m.Ups != nil {
		return *m.Ups
	}
	return 0
}

func (m *Comment) GetDowns() int32 {
	if m != nil && m.Downs != nil {
		return *m.Downs
	}
	return 0
}

func (m *Comment) GetLikes() bool {
	if m != nil && m.Likes != nil {
		return *m.Likes
	}
	return false
}

func (m *Comment) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Comment) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Comment) GetSubject() string {
	if m != nil && m.Subject != nil {
		return *m.Subject
	}
	return ""
}

// Data type t2_
type Account struct {
	CommentKarma     *int32   `protobuf:"varint,1,opt,name=comment_karma" json:"comment_karma,omitempty"`
	HasMail          *bool    `protobuf:"varint,2,opt,name=has_mail" json:"has_mail,omitempty"`
	HasModMail       *bool    `protobuf:"varint,3,opt,name=has_mod_mail" json:"has_mod_mail,omitempty"`
	HasVerifiedEmail *bool    `protobuf:"varint,4,opt,name=has_verified_email" json:"has_verified_email,omitempty"`
	InboxCount       *int32   `protobuf:"varint,5,opt,name=inbox_count" json:"inbox_count,omitempty"`
	IsFriend         *bool    `protobuf:"varint,6,opt,name=is_friend" json:"is_friend,omitempty"`
	IsGold           *bool    `protobuf:"varint,7,opt,name=is_gold" json:"is_gold,omitempty"`
	IsMod            *bool    `protobuf:"varint,8,opt,name=is_mod" json:"is_mod,omitempty"`
	LinkKarma        *int32   `protobuf:"varint,9,opt,name=link_karma" json:"link_karma,omitempty"`
	Modhash          *string  `protobuf:"bytes,10,opt,name=modhash" json:"modhash,omitempty"`
	Over_18          *bool    `protobuf:"varint,11,opt,name=over_18" json:"over_18,omitempty"`
	GoldCredits      *int32   `protobuf:"varint,12,opt,name=gold_credits" json:"gold_credits,omitempty"`
	GoldExpiration   *float64 `protobuf:"fixed64,13,opt,name=gold_expiration" json:"gold_expiration,omitempty"`
	HideFromRobots   *bool    `protobuf:"varint,14,opt,name=hide_from_robots" json:"hide_from_robots,omitempty"`
	// Implements Created
	Created    *float64 `protobuf:"fixed64,15,opt,name=created" json:"created,omitempty"`
	CreatedUtc *float64 `protobuf:"fixed64,16,opt,name=created_utc" json:"created_utc,omitempty"`
	// Implements Thing
	Id               *string                   `protobuf:"bytes,17,opt,name=id" json:"id,omitempty"`
	Name             *string                   `protobuf:"bytes,18,opt,name=name" json:"name,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}

var extRange_Account = []proto.ExtensionRange{
	{100, 536870911},
}

func (*Account) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Account
}
func (m *Account) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *Account) GetCommentKarma() int32 {
	if m != nil && m.CommentKarma != nil {
		return *m.CommentKarma
	}
	return 0
}

func (m *Account) GetHasMail() bool {
	if m != nil && m.HasMail != nil {
		return *m.HasMail
	}
	return false
}

func (m *Account) GetHasModMail() bool {
	if m != nil && m.HasModMail != nil {
		return *m.HasModMail
	}
	return false
}

func (m *Account) GetHasVerifiedEmail() bool {
	if m != nil && m.HasVerifiedEmail != nil {
		return *m.HasVerifiedEmail
	}
	return false
}

func (m *Account) GetInboxCount() int32 {
	if m != nil && m.InboxCount != nil {
		return *m.InboxCount
	}
	return 0
}

func (m *Account) GetIsFriend() bool {
	if m != nil && m.IsFriend != nil {
		return *m.IsFriend
	}
	return false
}

func (m *Account) GetIsGold() bool {
	if m != nil && m.IsGold != nil {
		return *m.IsGold
	}
	return false
}

func (m *Account) GetIsMod() bool {
	if m != nil && m.IsMod != nil {
		return *m.IsMod
	}
	return false
}

func (m *Account) GetLinkKarma() int32 {
	if m != nil && m.LinkKarma != nil {
		return *m.LinkKarma
	}
	return 0
}

func (m *Account) GetModhash() string {
	if m != nil && m.Modhash != nil {
		return *m.Modhash
	}
	return ""
}

func (m *Account) GetOver_18() bool {
	if m != nil && m.Over_18 != nil {
		return *m.Over_18
	}
	return false
}

func (m *Account) GetGoldCredits() int32 {
	if m != nil && m.GoldCredits != nil {
		return *m.GoldCredits
	}
	return 0
}

func (m *Account) GetGoldExpiration() float64 {
	if m != nil && m.GoldExpiration != nil {
		return *m.GoldExpiration
	}
	return 0
}

func (m *Account) GetHideFromRobots() bool {
	if m != nil && m.HideFromRobots != nil {
		return *m.HideFromRobots
	}
	return false
}

func (m *Account) GetCreated() float64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func (m *Account) GetCreatedUtc() float64 {
	if m != nil && m.CreatedUtc != nil {
		return *m.CreatedUtc
	}
	return 0
}

func (m *Account) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Account) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

// Data type t3_
type Link struct {
	Author              *string `protobuf:"bytes,1,opt,name=author" json:"author,omitempty"`
	AuthorFlairCssClass *string `protobuf:"bytes,2,opt,name=author_flair_css_class" json:"author_flair_css_class,omitempty"`
	AuthorFlairText     *string `protobuf:"bytes,3,opt,name=author_flair_text" json:"author_flair_text,omitempty"`
	Clicked             *bool   `protobuf:"varint,4,opt,name=clicked" json:"clicked,omitempty"`
	Domain              *string `protobuf:"bytes,5,opt,name=domain" json:"domain,omitempty"`
	Hidden              *bool   `protobuf:"varint,6,opt,name=hidden" json:"hidden,omitempty"`
	IsSelf              *bool   `protobuf:"varint,7,opt,name=is_self" json:"is_self,omitempty"`
	LinkFlairCssClass   *string `protobuf:"bytes,8,opt,name=link_flair_css_class" json:"link_flair_css_class,omitempty"`
	LinkFlairText       *string `protobuf:"bytes,9,opt,name=link_flair_text" json:"link_flair_text,omitempty"`
	// Field 10 is reserved for "media".
	// Field 11 is reserved for "media_embed".
	NumComments  *int32  `protobuf:"varint,12,opt,name=num_comments" json:"num_comments,omitempty"`
	Over_18      *bool   `protobuf:"varint,13,opt,name=over_18" json:"over_18,omitempty"`
	Permalink    *string `protobuf:"bytes,14,opt,name=permalink" json:"permalink,omitempty"`
	Saved        *bool   `protobuf:"varint,15,opt,name=saved" json:"saved,omitempty"`
	Score        *int32  `protobuf:"varint,16,opt,name=score" json:"score,omitempty"`
	Selftext     *string `protobuf:"bytes,17,opt,name=selftext" json:"selftext,omitempty"`
	SelftextHtml *string `protobuf:"bytes,18,opt,name=selftext_html" json:"selftext_html,omitempty"`
	Subreddit    *string `protobuf:"bytes,19,opt,name=subreddit" json:"subreddit,omitempty"`
	SubredditId  *string `protobuf:"bytes,20,opt,name=subreddit_id" json:"subreddit_id,omitempty"`
	Thumbnail    *string `protobuf:"bytes,21,opt,name=thumbnail" json:"thumbnail,omitempty"`
	Title        *string `protobuf:"bytes,22,opt,name=title" json:"title,omitempty"`
	Url          *string `protobuf:"bytes,23,opt,name=url" json:"url,omitempty"`
	// Field 24 is reserved for "edited".
	Distinguished *string `protobuf:"bytes,24,opt,name=distinguished" json:"distinguished,omitempty"`
	Stickied      *bool   `protobuf:"varint,25,opt,name=stickied" json:"stickied,omitempty"`
	// Implements Created
	Created    *float64 `protobuf:"fixed64,26,opt,name=created" json:"created,omitempty"`
	CreatedUtc *float64 `protobuf:"fixed64,27,opt,name=created_utc" json:"created_utc,omitempty"`
	// Implements Votable
	Ups   *int32 `protobuf:"varint,28,opt,name=ups" json:"ups,omitempty"`
	Downs *int32 `protobuf:"varint,29,opt,name=downs" json:"downs,omitempty"`
	Likes *bool  `protobuf:"varint,30,opt,name=likes" json:"likes,omitempty"`
	// Implements Thing
	Id   *string `protobuf:"bytes,31,opt,name=id" json:"id,omitempty"`
	Name *string `protobuf:"bytes,32,opt,name=name" json:"name,omitempty"`
	// Comment tree (not provided by Reddit).
	Comments         []*Comment                `protobuf:"bytes,33,rep,name=comments" json:"comments,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}

var extRange_Link = []proto.ExtensionRange{
	{100, 536870911},
}

func (*Link) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Link
}
func (m *Link) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *Link) GetAuthor() string {
	if m != nil && m.Author != nil {
		return *m.Author
	}
	return ""
}

func (m *Link) GetAuthorFlairCssClass() string {
	if m != nil && m.AuthorFlairCssClass != nil {
		return *m.AuthorFlairCssClass
	}
	return ""
}

func (m *Link) GetAuthorFlairText() string {
	if m != nil && m.AuthorFlairText != nil {
		return *m.AuthorFlairText
	}
	return ""
}

func (m *Link) GetClicked() bool {
	if m != nil && m.Clicked != nil {
		return *m.Clicked
	}
	return false
}

func (m *Link) GetDomain() string {
	if m != nil && m.Domain != nil {
		return *m.Domain
	}
	return ""
}

func (m *Link) GetHidden() bool {
	if m != nil && m.Hidden != nil {
		return *m.Hidden
	}
	return false
}

func (m *Link) GetIsSelf() bool {
	if m != nil && m.IsSelf != nil {
		return *m.IsSelf
	}
	return false
}

func (m *Link) GetLinkFlairCssClass() string {
	if m != nil && m.LinkFlairCssClass != nil {
		return *m.LinkFlairCssClass
	}
	return ""
}

func (m *Link) GetLinkFlairText() string {
	if m != nil && m.LinkFlairText != nil {
		return *m.LinkFlairText
	}
	return ""
}

func (m *Link) GetNumComments() int32 {
	if m != nil && m.NumComments != nil {
		return *m.NumComments
	}
	return 0
}

func (m *Link) GetOver_18() bool {
	if m != nil && m.Over_18 != nil {
		return *m.Over_18
	}
	return false
}

func (m *Link) GetPermalink() string {
	if m != nil && m.Permalink != nil {
		return *m.Permalink
	}
	return ""
}

func (m *Link) GetSaved() bool {
	if m != nil && m.Saved != nil {
		return *m.Saved
	}
	return false
}

func (m *Link) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *Link) GetSelftext() string {
	if m != nil && m.Selftext != nil {
		return *m.Selftext
	}
	return ""
}

func (m *Link) GetSelftextHtml() string {
	if m != nil && m.SelftextHtml != nil {
		return *m.SelftextHtml
	}
	return ""
}

func (m *Link) GetSubreddit() string {
	if m != nil && m.Subreddit != nil {
		return *m.Subreddit
	}
	return ""
}

func (m *Link) GetSubredditId() string {
	if m != nil && m.SubredditId != nil {
		return *m.SubredditId
	}
	return ""
}

func (m *Link) GetThumbnail() string {
	if m != nil && m.Thumbnail != nil {
		return *m.Thumbnail
	}
	return ""
}

func (m *Link) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *Link) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *Link) GetDistinguished() string {
	if m != nil && m.Distinguished != nil {
		return *m.Distinguished
	}
	return ""
}

func (m *Link) GetStickied() bool {
	if m != nil && m.Stickied != nil {
		return *m.Stickied
	}
	return false
}

func (m *Link) GetCreated() float64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func (m *Link) GetCreatedUtc() float64 {
	if m != nil && m.CreatedUtc != nil {
		return *m.CreatedUtc
	}
	return 0
}

func (m *Link) GetUps() int32 {
	if m != nil && m.Ups != nil {
		return *m.Ups
	}
	return 0
}

func (m *Link) GetDowns() int32 {
	if m != nil && m.Downs != nil {
		return *m.Downs
	}
	return 0
}

func (m *Link) GetLikes() bool {
	if m != nil && m.Likes != nil {
		return *m.Likes
	}
	return false
}

func (m *Link) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Link) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Link) GetComments() []*Comment {
	if m != nil {
		return m.Comments
	}
	return nil
}

// LinkSet holds links and data set annotations.
type LinkSet struct {
	Links            []*Link                   `protobuf:"bytes,1,rep,name=links" json:"links,omitempty"`
	Data             *string                   `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *LinkSet) Reset()         { *m = LinkSet{} }
func (m *LinkSet) String() string { return proto.CompactTextString(m) }
func (*LinkSet) ProtoMessage()    {}

var extRange_LinkSet = []proto.ExtensionRange{
	{100, 536870911},
}

func (*LinkSet) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_LinkSet
}
func (m *LinkSet) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *LinkSet) GetLinks() []*Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *LinkSet) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

// Data type t4_
type Message struct {
	Author   *string `protobuf:"bytes,1,opt,name=author" json:"author,omitempty"`
	Body     *string `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
	BodyHtml *string `protobuf:"bytes,3,opt,name=body_html" json:"body_html,omitempty"`
	Context  *string `protobuf:"bytes,4,opt,name=context" json:"context,omitempty"`
	// Field 5 reserved for mystic "first_message".
	FirstMessageName *string `protobuf:"bytes,6,opt,name=first_message_name" json:"first_message_name,omitempty"`
	Likes            *bool   `protobuf:"varint,7,opt,name=likes" json:"likes,omitempty"`
	LinkTitle        *string `protobuf:"bytes,8,opt,name=link_title" json:"link_title,omitempty"`
	New              *bool   `protobuf:"varint,9,opt,name=new" json:"new,omitempty"`
	ParentId         *string `protobuf:"bytes,10,opt,name=parent_id" json:"parent_id,omitempty"`
	Replies          *string `protobuf:"bytes,11,opt,name=replies" json:"replies,omitempty"`
	Subject          *string `protobuf:"bytes,12,opt,name=subject" json:"subject,omitempty"`
	Subreddit        *string `protobuf:"bytes,13,opt,name=subreddit" json:"subreddit,omitempty"`
	WasComment       *bool   `protobuf:"varint,14,opt,name=was_comment" json:"was_comment,omitempty"`
	// Implements Created
	Created    *float64 `protobuf:"fixed64,15,opt,name=created" json:"created,omitempty"`
	CreatedUtc *float64 `protobuf:"fixed64,16,opt,name=created_utc" json:"created_utc,omitempty"`
	// Implements Thing
	Id   *string `protobuf:"bytes,17,opt,name=id" json:"id,omitempty"`
	Name *string `protobuf:"bytes,18,opt,name=name" json:"name,omitempty"`
	// This field contains the chronological sequence of messages following this
	// one.
	Messages         []*Message                `protobuf:"bytes,19,rep,name=messages" json:"messages,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}

var extRange_Message = []proto.ExtensionRange{
	{100, 536870911},
}

func (*Message) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Message
}
func (m *Message) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *Message) GetAuthor() string {
	if m != nil && m.Author != nil {
		return *m.Author
	}
	return ""
}

func (m *Message) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func (m *Message) GetBodyHtml() string {
	if m != nil && m.BodyHtml != nil {
		return *m.BodyHtml
	}
	return ""
}

func (m *Message) GetContext() string {
	if m != nil && m.Context != nil {
		return *m.Context
	}
	return ""
}

func (m *Message) GetFirstMessageName() string {
	if m != nil && m.FirstMessageName != nil {
		return *m.FirstMessageName
	}
	return ""
}

func (m *Message) GetLikes() bool {
	if m != nil && m.Likes != nil {
		return *m.Likes
	}
	return false
}

func (m *Message) GetLinkTitle() string {
	if m != nil && m.LinkTitle != nil {
		return *m.LinkTitle
	}
	return ""
}

func (m *Message) GetNew() bool {
	if m != nil && m.New != nil {
		return *m.New
	}
	return false
}

func (m *Message) GetParentId() string {
	if m != nil && m.ParentId != nil {
		return *m.ParentId
	}
	return ""
}

func (m *Message) GetReplies() string {
	if m != nil && m.Replies != nil {
		return *m.Replies
	}
	return ""
}

func (m *Message) GetSubject() string {
	if m != nil && m.Subject != nil {
		return *m.Subject
	}
	return ""
}

func (m *Message) GetSubreddit() string {
	if m != nil && m.Subreddit != nil {
		return *m.Subreddit
	}
	return ""
}

func (m *Message) GetWasComment() bool {
	if m != nil && m.WasComment != nil {
		return *m.WasComment
	}
	return false
}

func (m *Message) GetCreated() float64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func (m *Message) GetCreatedUtc() float64 {
	if m != nil && m.CreatedUtc != nil {
		return *m.CreatedUtc
	}
	return 0
}

func (m *Message) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Message) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Message) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

// Data type t5_
type Subreddit struct {
	AccountsActive  *int32  `protobuf:"varint,1,opt,name=accounts_active" json:"accounts_active,omitempty"`
	CommentScore    *int32  `protobuf:"varint,2,opt,name=comment_score" json:"comment_score,omitempty"`
	Description     *string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	DescriptionHtml *string `protobuf:"bytes,4,opt,name=description_html" json:"description_html,omitempty"`
	DisplayName     *string `protobuf:"bytes,5,opt,name=display_name" json:"display_name,omitempty"`
	HeaderImg       *string `protobuf:"bytes,6,opt,name=header_img" json:"header_img,omitempty"`
	// Field 7 is reserved for "header_size".
	HeaderTitle       *string `protobuf:"bytes,7,opt,name=header_title" json:"header_title,omitempty"`
	Over18            *bool   `protobuf:"varint,8,opt,name=over18" json:"over18,omitempty"`
	PublicDescription *string `protobuf:"bytes,9,opt,name=public_description" json:"public_description,omitempty"`
	PublicTraffic     *bool   `protobuf:"varint,10,opt,name=public_traffic" json:"public_traffic,omitempty"`
	Subscribers       *int64  `protobuf:"varint,11,opt,name=subscribers" json:"subscribers,omitempty"`
	SubmissionType    *string `protobuf:"bytes,12,opt,name=submission_type" json:"submission_type,omitempty"`
	SubmitLinkLabel   *string `protobuf:"bytes,13,opt,name=submit_link_label" json:"submit_link_label,omitempty"`
	SubmitTextLabel   *string `protobuf:"bytes,14,opt,name=submit_text_label" json:"submit_text_label,omitempty"`
	SubredditType     *string `protobuf:"bytes,15,opt,name=subreddit_type" json:"subreddit_type,omitempty"`
	Title             *string `protobuf:"bytes,16,opt,name=title" json:"title,omitempty"`
	Url               *string `protobuf:"bytes,17,opt,name=url" json:"url,omitempty"`
	UserIsBanned      *bool   `protobuf:"varint,18,opt,name=user_is_banned" json:"user_is_banned,omitempty"`
	UserIsContributor *bool   `protobuf:"varint,19,opt,name=user_is_contributor" json:"user_is_contributor,omitempty"`
	UserIsModerator   *bool   `protobuf:"varint,20,opt,name=user_is_moderator" json:"user_is_moderator,omitempty"`
	UserIsSubscriber  *bool   `protobuf:"varint,21,opt,name=user_is_subscriber" json:"user_is_subscriber,omitempty"`
	// Implements Thing
	Id               *string                   `protobuf:"bytes,22,opt,name=id" json:"id,omitempty"`
	Name             *string                   `protobuf:"bytes,24,opt,name=name" json:"name,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Subreddit) Reset()         { *m = Subreddit{} }
func (m *Subreddit) String() string { return proto.CompactTextString(m) }
func (*Subreddit) ProtoMessage()    {}

var extRange_Subreddit = []proto.ExtensionRange{
	{100, 536870911},
}

func (*Subreddit) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Subreddit
}
func (m *Subreddit) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *Subreddit) GetAccountsActive() int32 {
	if m != nil && m.AccountsActive != nil {
		return *m.AccountsActive
	}
	return 0
}

func (m *Subreddit) GetCommentScore() int32 {
	if m != nil && m.CommentScore != nil {
		return *m.CommentScore
	}
	return 0
}

func (m *Subreddit) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *Subreddit) GetDescriptionHtml() string {
	if m != nil && m.DescriptionHtml != nil {
		return *m.DescriptionHtml
	}
	return ""
}

func (m *Subreddit) GetDisplayName() string {
	if m != nil && m.DisplayName != nil {
		return *m.DisplayName
	}
	return ""
}

func (m *Subreddit) GetHeaderImg() string {
	if m != nil && m.HeaderImg != nil {
		return *m.HeaderImg
	}
	return ""
}

func (m *Subreddit) GetHeaderTitle() string {
	if m != nil && m.HeaderTitle != nil {
		return *m.HeaderTitle
	}
	return ""
}

func (m *Subreddit) GetOver18() bool {
	if m != nil && m.Over18 != nil {
		return *m.Over18
	}
	return false
}

func (m *Subreddit) GetPublicDescription() string {
	if m != nil && m.PublicDescription != nil {
		return *m.PublicDescription
	}
	return ""
}

func (m *Subreddit) GetPublicTraffic() bool {
	if m != nil && m.PublicTraffic != nil {
		return *m.PublicTraffic
	}
	return false
}

func (m *Subreddit) GetSubscribers() int64 {
	if m != nil && m.Subscribers != nil {
		return *m.Subscribers
	}
	return 0
}

func (m *Subreddit) GetSubmissionType() string {
	if m != nil && m.SubmissionType != nil {
		return *m.SubmissionType
	}
	return ""
}

func (m *Subreddit) GetSubmitLinkLabel() string {
	if m != nil && m.SubmitLinkLabel != nil {
		return *m.SubmitLinkLabel
	}
	return ""
}

func (m *Subreddit) GetSubmitTextLabel() string {
	if m != nil && m.SubmitTextLabel != nil {
		return *m.SubmitTextLabel
	}
	return ""
}

func (m *Subreddit) GetSubredditType() string {
	if m != nil && m.SubredditType != nil {
		return *m.SubredditType
	}
	return ""
}

func (m *Subreddit) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *Subreddit) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *Subreddit) GetUserIsBanned() bool {
	if m != nil && m.UserIsBanned != nil {
		return *m.UserIsBanned
	}
	return false
}

func (m *Subreddit) GetUserIsContributor() bool {
	if m != nil && m.UserIsContributor != nil {
		return *m.UserIsContributor
	}
	return false
}

func (m *Subreddit) GetUserIsModerator() bool {
	if m != nil && m.UserIsModerator != nil {
		return *m.UserIsModerator
	}
	return false
}

func (m *Subreddit) GetUserIsSubscriber() bool {
	if m != nil && m.UserIsSubscriber != nil {
		return *m.UserIsSubscriber
	}
	return false
}

func (m *Subreddit) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Subreddit) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Comment)(nil), "redditproto.Comment")
	proto.RegisterType((*Account)(nil), "redditproto.Account")
	proto.RegisterType((*Link)(nil), "redditproto.Link")
	proto.RegisterType((*LinkSet)(nil), "redditproto.LinkSet")
	proto.RegisterType((*Message)(nil), "redditproto.Message")
	proto.RegisterType((*Subreddit)(nil), "redditproto.Subreddit")
}
