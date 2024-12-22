// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// LikesColumns holds the columns for the "likes" table.
	LikesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "tweet_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// LikesTable holds the schema information for the "likes" table.
	LikesTable = &schema.Table{
		Name:       "likes",
		Columns:    LikesColumns,
		PrimaryKey: []*schema.Column{LikesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "likes_tweets_likes",
				Columns:    []*schema.Column{LikesColumns[3]},
				RefColumns: []*schema.Column{TweetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "likes_users_likes",
				Columns:    []*schema.Column{LikesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RetweetsColumns holds the columns for the "retweets" table.
	RetweetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "tweet_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
	}
	// RetweetsTable holds the schema information for the "retweets" table.
	RetweetsTable = &schema.Table{
		Name:       "retweets",
		Columns:    RetweetsColumns,
		PrimaryKey: []*schema.Column{RetweetsColumns[0]},
	}
	// TweetsColumns holds the columns for the "tweets" table.
	TweetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "content", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "parent_id", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "author_id", Type: field.TypeInt},
	}
	// TweetsTable holds the schema information for the "tweets" table.
	TweetsTable = &schema.Table{
		Name:       "tweets",
		Columns:    TweetsColumns,
		PrimaryKey: []*schema.Column{TweetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tweets_tweets_parent",
				Columns:    []*schema.Column{TweetsColumns[4]},
				RefColumns: []*schema.Column{TweetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tweets_users_tweets",
				Columns:    []*schema.Column{TweetsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "email_verified", Type: field.TypeBool, Default: false},
		{Name: "password", Type: field.TypeString},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "banner", Type: field.TypeString, Nullable: true},
		{Name: "bio", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// TweetRetweetsColumns holds the columns for the "tweet_retweets" table.
	TweetRetweetsColumns = []*schema.Column{
		{Name: "tweet_id", Type: field.TypeInt},
		{Name: "retweet_id", Type: field.TypeInt},
	}
	// TweetRetweetsTable holds the schema information for the "tweet_retweets" table.
	TweetRetweetsTable = &schema.Table{
		Name:       "tweet_retweets",
		Columns:    TweetRetweetsColumns,
		PrimaryKey: []*schema.Column{TweetRetweetsColumns[0], TweetRetweetsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tweet_retweets_tweet_id",
				Columns:    []*schema.Column{TweetRetweetsColumns[0]},
				RefColumns: []*schema.Column{TweetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tweet_retweets_retweet_id",
				Columns:    []*schema.Column{TweetRetweetsColumns[1]},
				RefColumns: []*schema.Column{RetweetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserRetweetsColumns holds the columns for the "user_retweets" table.
	UserRetweetsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "retweet_id", Type: field.TypeInt},
	}
	// UserRetweetsTable holds the schema information for the "user_retweets" table.
	UserRetweetsTable = &schema.Table{
		Name:       "user_retweets",
		Columns:    UserRetweetsColumns,
		PrimaryKey: []*schema.Column{UserRetweetsColumns[0], UserRetweetsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_retweets_user_id",
				Columns:    []*schema.Column{UserRetweetsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_retweets_retweet_id",
				Columns:    []*schema.Column{UserRetweetsColumns[1]},
				RefColumns: []*schema.Column{RetweetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		LikesTable,
		RetweetsTable,
		TweetsTable,
		UsersTable,
		TweetRetweetsTable,
		UserRetweetsTable,
	}
)

func init() {
	LikesTable.ForeignKeys[0].RefTable = TweetsTable
	LikesTable.ForeignKeys[1].RefTable = UsersTable
	TweetsTable.ForeignKeys[0].RefTable = TweetsTable
	TweetsTable.ForeignKeys[1].RefTable = UsersTable
	TweetRetweetsTable.ForeignKeys[0].RefTable = TweetsTable
	TweetRetweetsTable.ForeignKeys[1].RefTable = RetweetsTable
	TweetRetweetsTable.Annotation = &entsql.Annotation{}
	UserRetweetsTable.ForeignKeys[0].RefTable = UsersTable
	UserRetweetsTable.ForeignKeys[1].RefTable = RetweetsTable
	UserRetweetsTable.Annotation = &entsql.Annotation{}
}
