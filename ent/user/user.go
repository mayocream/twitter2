// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldEmailVerified holds the string denoting the email_verified field in the database.
	FieldEmailVerified = "email_verified"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldBanner holds the string denoting the banner field in the database.
	FieldBanner = "banner"
	// FieldBio holds the string denoting the bio field in the database.
	FieldBio = "bio"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTweets holds the string denoting the tweets edge name in mutations.
	EdgeTweets = "tweets"
	// EdgeLikes holds the string denoting the likes edge name in mutations.
	EdgeLikes = "likes"
	// Table holds the table name of the user in the database.
	Table = "users"
	// TweetsTable is the table that holds the tweets relation/edge.
	TweetsTable = "tweets"
	// TweetsInverseTable is the table name for the Tweet entity.
	// It exists in this package in order to avoid circular dependency with the "tweet" package.
	TweetsInverseTable = "tweets"
	// TweetsColumn is the table column denoting the tweets relation/edge.
	TweetsColumn = "author_id"
	// LikesTable is the table that holds the likes relation/edge.
	LikesTable = "likes"
	// LikesInverseTable is the table name for the Like entity.
	// It exists in this package in order to avoid circular dependency with the "like" package.
	LikesInverseTable = "likes"
	// LikesColumn is the table column denoting the likes relation/edge.
	LikesColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldUsername,
	FieldEmail,
	FieldEmailVerified,
	FieldPassword,
	FieldAvatar,
	FieldBanner,
	FieldBio,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultEmailVerified holds the default value on creation for the "email_verified" field.
	DefaultEmailVerified bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByEmailVerified orders the results by the email_verified field.
func ByEmailVerified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmailVerified, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByBanner orders the results by the banner field.
func ByBanner(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBanner, opts...).ToFunc()
}

// ByBio orders the results by the bio field.
func ByBio(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBio, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByTweetsCount orders the results by tweets count.
func ByTweetsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTweetsStep(), opts...)
	}
}

// ByTweets orders the results by tweets terms.
func ByTweets(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTweetsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLikesCount orders the results by likes count.
func ByLikesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLikesStep(), opts...)
	}
}

// ByLikes orders the results by likes terms.
func ByLikes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLikesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTweetsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TweetsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TweetsTable, TweetsColumn),
	)
}
func newLikesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LikesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LikesTable, LikesColumn),
	)
}
