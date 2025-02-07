// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mayocream/twitter2/ent/like"
	"github.com/mayocream/twitter2/ent/predicate"
	"github.com/mayocream/twitter2/ent/tweet"
	"github.com/mayocream/twitter2/ent/user"
)

// LikeUpdate is the builder for updating Like entities.
type LikeUpdate struct {
	config
	hooks    []Hook
	mutation *LikeMutation
}

// Where appends a list predicates to the LikeUpdate builder.
func (lu *LikeUpdate) Where(ps ...predicate.Like) *LikeUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetTweetID sets the "tweet_id" field.
func (lu *LikeUpdate) SetTweetID(i int) *LikeUpdate {
	lu.mutation.SetTweetID(i)
	return lu
}

// SetNillableTweetID sets the "tweet_id" field if the given value is not nil.
func (lu *LikeUpdate) SetNillableTweetID(i *int) *LikeUpdate {
	if i != nil {
		lu.SetTweetID(*i)
	}
	return lu
}

// SetUserID sets the "user_id" field.
func (lu *LikeUpdate) SetUserID(i int) *LikeUpdate {
	lu.mutation.SetUserID(i)
	return lu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (lu *LikeUpdate) SetNillableUserID(i *int) *LikeUpdate {
	if i != nil {
		lu.SetUserID(*i)
	}
	return lu
}

// SetCreatedAt sets the "created_at" field.
func (lu *LikeUpdate) SetCreatedAt(t time.Time) *LikeUpdate {
	lu.mutation.SetCreatedAt(t)
	return lu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lu *LikeUpdate) SetNillableCreatedAt(t *time.Time) *LikeUpdate {
	if t != nil {
		lu.SetCreatedAt(*t)
	}
	return lu
}

// SetUpdatedAt sets the "updated_at" field.
func (lu *LikeUpdate) SetUpdatedAt(t time.Time) *LikeUpdate {
	lu.mutation.SetUpdatedAt(t)
	return lu
}

// SetTweet sets the "tweet" edge to the Tweet entity.
func (lu *LikeUpdate) SetTweet(t *Tweet) *LikeUpdate {
	return lu.SetTweetID(t.ID)
}

// SetUser sets the "user" edge to the User entity.
func (lu *LikeUpdate) SetUser(u *User) *LikeUpdate {
	return lu.SetUserID(u.ID)
}

// Mutation returns the LikeMutation object of the builder.
func (lu *LikeUpdate) Mutation() *LikeMutation {
	return lu.mutation
}

// ClearTweet clears the "tweet" edge to the Tweet entity.
func (lu *LikeUpdate) ClearTweet() *LikeUpdate {
	lu.mutation.ClearTweet()
	return lu
}

// ClearUser clears the "user" edge to the User entity.
func (lu *LikeUpdate) ClearUser() *LikeUpdate {
	lu.mutation.ClearUser()
	return lu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LikeUpdate) Save(ctx context.Context) (int, error) {
	lu.defaults()
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LikeUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LikeUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LikeUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lu *LikeUpdate) defaults() {
	if _, ok := lu.mutation.UpdatedAt(); !ok {
		v := like.UpdateDefaultUpdatedAt()
		lu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LikeUpdate) check() error {
	if lu.mutation.TweetCleared() && len(lu.mutation.TweetIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Like.tweet"`)
	}
	if lu.mutation.UserCleared() && len(lu.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Like.user"`)
	}
	return nil
}

func (lu *LikeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(like.Table, like.Columns, sqlgraph.NewFieldSpec(like.FieldID, field.TypeInt))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.CreatedAt(); ok {
		_spec.SetField(like.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := lu.mutation.UpdatedAt(); ok {
		_spec.SetField(like.FieldUpdatedAt, field.TypeTime, value)
	}
	if lu.mutation.TweetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   like.TweetTable,
			Columns: []string{like.TweetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.TweetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   like.TweetTable,
			Columns: []string{like.TweetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   like.UserTable,
			Columns: []string{like.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   like.UserTable,
			Columns: []string{like.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{like.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LikeUpdateOne is the builder for updating a single Like entity.
type LikeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LikeMutation
}

// SetTweetID sets the "tweet_id" field.
func (luo *LikeUpdateOne) SetTweetID(i int) *LikeUpdateOne {
	luo.mutation.SetTweetID(i)
	return luo
}

// SetNillableTweetID sets the "tweet_id" field if the given value is not nil.
func (luo *LikeUpdateOne) SetNillableTweetID(i *int) *LikeUpdateOne {
	if i != nil {
		luo.SetTweetID(*i)
	}
	return luo
}

// SetUserID sets the "user_id" field.
func (luo *LikeUpdateOne) SetUserID(i int) *LikeUpdateOne {
	luo.mutation.SetUserID(i)
	return luo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (luo *LikeUpdateOne) SetNillableUserID(i *int) *LikeUpdateOne {
	if i != nil {
		luo.SetUserID(*i)
	}
	return luo
}

// SetCreatedAt sets the "created_at" field.
func (luo *LikeUpdateOne) SetCreatedAt(t time.Time) *LikeUpdateOne {
	luo.mutation.SetCreatedAt(t)
	return luo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (luo *LikeUpdateOne) SetNillableCreatedAt(t *time.Time) *LikeUpdateOne {
	if t != nil {
		luo.SetCreatedAt(*t)
	}
	return luo
}

// SetUpdatedAt sets the "updated_at" field.
func (luo *LikeUpdateOne) SetUpdatedAt(t time.Time) *LikeUpdateOne {
	luo.mutation.SetUpdatedAt(t)
	return luo
}

// SetTweet sets the "tweet" edge to the Tweet entity.
func (luo *LikeUpdateOne) SetTweet(t *Tweet) *LikeUpdateOne {
	return luo.SetTweetID(t.ID)
}

// SetUser sets the "user" edge to the User entity.
func (luo *LikeUpdateOne) SetUser(u *User) *LikeUpdateOne {
	return luo.SetUserID(u.ID)
}

// Mutation returns the LikeMutation object of the builder.
func (luo *LikeUpdateOne) Mutation() *LikeMutation {
	return luo.mutation
}

// ClearTweet clears the "tweet" edge to the Tweet entity.
func (luo *LikeUpdateOne) ClearTweet() *LikeUpdateOne {
	luo.mutation.ClearTweet()
	return luo
}

// ClearUser clears the "user" edge to the User entity.
func (luo *LikeUpdateOne) ClearUser() *LikeUpdateOne {
	luo.mutation.ClearUser()
	return luo
}

// Where appends a list predicates to the LikeUpdate builder.
func (luo *LikeUpdateOne) Where(ps ...predicate.Like) *LikeUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LikeUpdateOne) Select(field string, fields ...string) *LikeUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Like entity.
func (luo *LikeUpdateOne) Save(ctx context.Context) (*Like, error) {
	luo.defaults()
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LikeUpdateOne) SaveX(ctx context.Context) *Like {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LikeUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LikeUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luo *LikeUpdateOne) defaults() {
	if _, ok := luo.mutation.UpdatedAt(); !ok {
		v := like.UpdateDefaultUpdatedAt()
		luo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LikeUpdateOne) check() error {
	if luo.mutation.TweetCleared() && len(luo.mutation.TweetIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Like.tweet"`)
	}
	if luo.mutation.UserCleared() && len(luo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Like.user"`)
	}
	return nil
}

func (luo *LikeUpdateOne) sqlSave(ctx context.Context) (_node *Like, err error) {
	if err := luo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(like.Table, like.Columns, sqlgraph.NewFieldSpec(like.FieldID, field.TypeInt))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Like.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, like.FieldID)
		for _, f := range fields {
			if !like.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != like.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.CreatedAt(); ok {
		_spec.SetField(like.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := luo.mutation.UpdatedAt(); ok {
		_spec.SetField(like.FieldUpdatedAt, field.TypeTime, value)
	}
	if luo.mutation.TweetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   like.TweetTable,
			Columns: []string{like.TweetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.TweetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   like.TweetTable,
			Columns: []string{like.TweetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   like.UserTable,
			Columns: []string{like.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   like.UserTable,
			Columns: []string{like.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Like{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{like.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
