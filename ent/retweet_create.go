// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mayocream/twitter2/ent/retweet"
	"github.com/mayocream/twitter2/ent/tweet"
	"github.com/mayocream/twitter2/ent/user"
)

// RetweetCreate is the builder for creating a Retweet entity.
type RetweetCreate struct {
	config
	mutation *RetweetMutation
	hooks    []Hook
}

// SetTweetID sets the "tweet_id" field.
func (rc *RetweetCreate) SetTweetID(i int) *RetweetCreate {
	rc.mutation.SetTweetID(i)
	return rc
}

// SetUserID sets the "user_id" field.
func (rc *RetweetCreate) SetUserID(i int) *RetweetCreate {
	rc.mutation.SetUserID(i)
	return rc
}

// SetCreatedAt sets the "created_at" field.
func (rc *RetweetCreate) SetCreatedAt(t time.Time) *RetweetCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetTweet sets the "tweet" edge to the Tweet entity.
func (rc *RetweetCreate) SetTweet(t *Tweet) *RetweetCreate {
	return rc.SetTweetID(t.ID)
}

// SetUser sets the "user" edge to the User entity.
func (rc *RetweetCreate) SetUser(u *User) *RetweetCreate {
	return rc.SetUserID(u.ID)
}

// Mutation returns the RetweetMutation object of the builder.
func (rc *RetweetCreate) Mutation() *RetweetMutation {
	return rc.mutation
}

// Save creates the Retweet in the database.
func (rc *RetweetCreate) Save(ctx context.Context) (*Retweet, error) {
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RetweetCreate) SaveX(ctx context.Context) *Retweet {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RetweetCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RetweetCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RetweetCreate) check() error {
	if _, ok := rc.mutation.TweetID(); !ok {
		return &ValidationError{Name: "tweet_id", err: errors.New(`ent: missing required field "Retweet.tweet_id"`)}
	}
	if _, ok := rc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Retweet.user_id"`)}
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Retweet.created_at"`)}
	}
	if len(rc.mutation.TweetIDs()) == 0 {
		return &ValidationError{Name: "tweet", err: errors.New(`ent: missing required edge "Retweet.tweet"`)}
	}
	if len(rc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Retweet.user"`)}
	}
	return nil
}

func (rc *RetweetCreate) sqlSave(ctx context.Context) (*Retweet, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RetweetCreate) createSpec() (*Retweet, *sqlgraph.CreateSpec) {
	var (
		_node = &Retweet{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(retweet.Table, sqlgraph.NewFieldSpec(retweet.FieldID, field.TypeInt))
	)
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(retweet.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := rc.mutation.TweetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   retweet.TweetTable,
			Columns: []string{retweet.TweetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TweetID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   retweet.UserTable,
			Columns: []string{retweet.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RetweetCreateBulk is the builder for creating many Retweet entities in bulk.
type RetweetCreateBulk struct {
	config
	err      error
	builders []*RetweetCreate
}

// Save creates the Retweet entities in the database.
func (rcb *RetweetCreateBulk) Save(ctx context.Context) ([]*Retweet, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Retweet, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RetweetMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RetweetCreateBulk) SaveX(ctx context.Context) []*Retweet {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RetweetCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RetweetCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
