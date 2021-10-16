// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-ent-sample/ent/car"
	"go-ent-sample/ent/group"
	"go-ent-sample/ent/predicate"
	"go-ent-sample/ent/user"
	"sync"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCar   = "Car"
	TypeGroup = "Group"
	TypeUser  = "User"
)

// CarMutation represents an operation that mutates the Car nodes in the graph.
type CarMutation struct {
	config
	op            Op
	typ           string
	id            *int
	model         *string
	registered_at *time.Time
	clearedFields map[string]struct{}
	owner         *int
	clearedowner  bool
	done          bool
	oldValue      func(context.Context) (*Car, error)
	predicates    []predicate.Car
}

var _ ent.Mutation = (*CarMutation)(nil)

// carOption allows management of the mutation configuration using functional options.
type carOption func(*CarMutation)

// newCarMutation creates new mutation for the Car entity.
func newCarMutation(c config, op Op, opts ...carOption) *CarMutation {
	m := &CarMutation{
		config:        c,
		op:            op,
		typ:           TypeCar,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCarID sets the ID field of the mutation.
func withCarID(id int) carOption {
	return func(m *CarMutation) {
		var (
			err   error
			once  sync.Once
			value *Car
		)
		m.oldValue = func(ctx context.Context) (*Car, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Car.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCar sets the old Car of the mutation.
func withCar(node *Car) carOption {
	return func(m *CarMutation) {
		m.oldValue = func(context.Context) (*Car, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CarMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CarMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CarMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetModel sets the "model" field.
func (m *CarMutation) SetModel(s string) {
	m.model = &s
}

// Model returns the value of the "model" field in the mutation.
func (m *CarMutation) Model() (r string, exists bool) {
	v := m.model
	if v == nil {
		return
	}
	return *v, true
}

// OldModel returns the old "model" field's value of the Car entity.
// If the Car object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CarMutation) OldModel(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldModel is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldModel requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldModel: %w", err)
	}
	return oldValue.Model, nil
}

// ResetModel resets all changes to the "model" field.
func (m *CarMutation) ResetModel() {
	m.model = nil
}

// SetRegisteredAt sets the "registered_at" field.
func (m *CarMutation) SetRegisteredAt(t time.Time) {
	m.registered_at = &t
}

// RegisteredAt returns the value of the "registered_at" field in the mutation.
func (m *CarMutation) RegisteredAt() (r time.Time, exists bool) {
	v := m.registered_at
	if v == nil {
		return
	}
	return *v, true
}

// OldRegisteredAt returns the old "registered_at" field's value of the Car entity.
// If the Car object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CarMutation) OldRegisteredAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRegisteredAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRegisteredAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRegisteredAt: %w", err)
	}
	return oldValue.RegisteredAt, nil
}

// ResetRegisteredAt resets all changes to the "registered_at" field.
func (m *CarMutation) ResetRegisteredAt() {
	m.registered_at = nil
}

// SetOwnerID sets the "owner" edge to the User entity by id.
func (m *CarMutation) SetOwnerID(id int) {
	m.owner = &id
}

// ClearOwner clears the "owner" edge to the User entity.
func (m *CarMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared reports if the "owner" edge to the User entity was cleared.
func (m *CarMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerID returns the "owner" edge ID in the mutation.
func (m *CarMutation) OwnerID() (id int, exists bool) {
	if m.owner != nil {
		return *m.owner, true
	}
	return
}

// OwnerIDs returns the "owner" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// OwnerID instead. It exists only for internal usage by the builders.
func (m *CarMutation) OwnerIDs() (ids []int) {
	if id := m.owner; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwner resets all changes to the "owner" edge.
func (m *CarMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// Where appends a list predicates to the CarMutation builder.
func (m *CarMutation) Where(ps ...predicate.Car) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *CarMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Car).
func (m *CarMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CarMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.model != nil {
		fields = append(fields, car.FieldModel)
	}
	if m.registered_at != nil {
		fields = append(fields, car.FieldRegisteredAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CarMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case car.FieldModel:
		return m.Model()
	case car.FieldRegisteredAt:
		return m.RegisteredAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CarMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case car.FieldModel:
		return m.OldModel(ctx)
	case car.FieldRegisteredAt:
		return m.OldRegisteredAt(ctx)
	}
	return nil, fmt.Errorf("unknown Car field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CarMutation) SetField(name string, value ent.Value) error {
	switch name {
	case car.FieldModel:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetModel(v)
		return nil
	case car.FieldRegisteredAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRegisteredAt(v)
		return nil
	}
	return fmt.Errorf("unknown Car field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CarMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CarMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CarMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Car numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CarMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CarMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CarMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Car nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CarMutation) ResetField(name string) error {
	switch name {
	case car.FieldModel:
		m.ResetModel()
		return nil
	case car.FieldRegisteredAt:
		m.ResetRegisteredAt()
		return nil
	}
	return fmt.Errorf("unknown Car field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CarMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owner != nil {
		edges = append(edges, car.EdgeOwner)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CarMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case car.EdgeOwner:
		if id := m.owner; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CarMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CarMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CarMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowner {
		edges = append(edges, car.EdgeOwner)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CarMutation) EdgeCleared(name string) bool {
	switch name {
	case car.EdgeOwner:
		return m.clearedowner
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CarMutation) ClearEdge(name string) error {
	switch name {
	case car.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown Car unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CarMutation) ResetEdge(name string) error {
	switch name {
	case car.EdgeOwner:
		m.ResetOwner()
		return nil
	}
	return fmt.Errorf("unknown Car edge %s", name)
}

// GroupMutation represents an operation that mutates the Group nodes in the graph.
type GroupMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	users         map[int]struct{}
	removedusers  map[int]struct{}
	clearedusers  bool
	done          bool
	oldValue      func(context.Context) (*Group, error)
	predicates    []predicate.Group
}

var _ ent.Mutation = (*GroupMutation)(nil)

// groupOption allows management of the mutation configuration using functional options.
type groupOption func(*GroupMutation)

// newGroupMutation creates new mutation for the Group entity.
func newGroupMutation(c config, op Op, opts ...groupOption) *GroupMutation {
	m := &GroupMutation{
		config:        c,
		op:            op,
		typ:           TypeGroup,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withGroupID sets the ID field of the mutation.
func withGroupID(id int) groupOption {
	return func(m *GroupMutation) {
		var (
			err   error
			once  sync.Once
			value *Group
		)
		m.oldValue = func(ctx context.Context) (*Group, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Group.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withGroup sets the old Group of the mutation.
func withGroup(node *Group) groupOption {
	return func(m *GroupMutation) {
		m.oldValue = func(context.Context) (*Group, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m GroupMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m GroupMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *GroupMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *GroupMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *GroupMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Group entity.
// If the Group object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *GroupMutation) ResetName() {
	m.name = nil
}

// AddUserIDs adds the "users" edge to the User entity by ids.
func (m *GroupMutation) AddUserIDs(ids ...int) {
	if m.users == nil {
		m.users = make(map[int]struct{})
	}
	for i := range ids {
		m.users[ids[i]] = struct{}{}
	}
}

// ClearUsers clears the "users" edge to the User entity.
func (m *GroupMutation) ClearUsers() {
	m.clearedusers = true
}

// UsersCleared reports if the "users" edge to the User entity was cleared.
func (m *GroupMutation) UsersCleared() bool {
	return m.clearedusers
}

// RemoveUserIDs removes the "users" edge to the User entity by IDs.
func (m *GroupMutation) RemoveUserIDs(ids ...int) {
	if m.removedusers == nil {
		m.removedusers = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.users, ids[i])
		m.removedusers[ids[i]] = struct{}{}
	}
}

// RemovedUsers returns the removed IDs of the "users" edge to the User entity.
func (m *GroupMutation) RemovedUsersIDs() (ids []int) {
	for id := range m.removedusers {
		ids = append(ids, id)
	}
	return
}

// UsersIDs returns the "users" edge IDs in the mutation.
func (m *GroupMutation) UsersIDs() (ids []int) {
	for id := range m.users {
		ids = append(ids, id)
	}
	return
}

// ResetUsers resets all changes to the "users" edge.
func (m *GroupMutation) ResetUsers() {
	m.users = nil
	m.clearedusers = false
	m.removedusers = nil
}

// Where appends a list predicates to the GroupMutation builder.
func (m *GroupMutation) Where(ps ...predicate.Group) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *GroupMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Group).
func (m *GroupMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *GroupMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, group.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *GroupMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case group.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *GroupMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case group.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Group field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *GroupMutation) SetField(name string, value ent.Value) error {
	switch name {
	case group.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Group field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *GroupMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *GroupMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *GroupMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Group numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *GroupMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *GroupMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *GroupMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Group nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *GroupMutation) ResetField(name string) error {
	switch name {
	case group.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Group field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *GroupMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.users != nil {
		edges = append(edges, group.EdgeUsers)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *GroupMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case group.EdgeUsers:
		ids := make([]ent.Value, 0, len(m.users))
		for id := range m.users {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *GroupMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedusers != nil {
		edges = append(edges, group.EdgeUsers)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *GroupMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case group.EdgeUsers:
		ids := make([]ent.Value, 0, len(m.removedusers))
		for id := range m.removedusers {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *GroupMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedusers {
		edges = append(edges, group.EdgeUsers)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *GroupMutation) EdgeCleared(name string) bool {
	switch name {
	case group.EdgeUsers:
		return m.clearedusers
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *GroupMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Group unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *GroupMutation) ResetEdge(name string) error {
	switch name {
	case group.EdgeUsers:
		m.ResetUsers()
		return nil
	}
	return fmt.Errorf("unknown Group edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	age           *int
	addage        *int
	name          *string
	clearedFields map[string]struct{}
	cars          map[int]struct{}
	removedcars   map[int]struct{}
	clearedcars   bool
	groups        map[int]struct{}
	removedgroups map[int]struct{}
	clearedgroups bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetAge sets the "age" field.
func (m *UserMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the value of the "age" field in the mutation.
func (m *UserMutation) Age() (r int, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old "age" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAge(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds i to the "age" field.
func (m *UserMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the "age" field in this mutation.
func (m *UserMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge resets all changes to the "age" field.
func (m *UserMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// AddCarIDs adds the "cars" edge to the Car entity by ids.
func (m *UserMutation) AddCarIDs(ids ...int) {
	if m.cars == nil {
		m.cars = make(map[int]struct{})
	}
	for i := range ids {
		m.cars[ids[i]] = struct{}{}
	}
}

// ClearCars clears the "cars" edge to the Car entity.
func (m *UserMutation) ClearCars() {
	m.clearedcars = true
}

// CarsCleared reports if the "cars" edge to the Car entity was cleared.
func (m *UserMutation) CarsCleared() bool {
	return m.clearedcars
}

// RemoveCarIDs removes the "cars" edge to the Car entity by IDs.
func (m *UserMutation) RemoveCarIDs(ids ...int) {
	if m.removedcars == nil {
		m.removedcars = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.cars, ids[i])
		m.removedcars[ids[i]] = struct{}{}
	}
}

// RemovedCars returns the removed IDs of the "cars" edge to the Car entity.
func (m *UserMutation) RemovedCarsIDs() (ids []int) {
	for id := range m.removedcars {
		ids = append(ids, id)
	}
	return
}

// CarsIDs returns the "cars" edge IDs in the mutation.
func (m *UserMutation) CarsIDs() (ids []int) {
	for id := range m.cars {
		ids = append(ids, id)
	}
	return
}

// ResetCars resets all changes to the "cars" edge.
func (m *UserMutation) ResetCars() {
	m.cars = nil
	m.clearedcars = false
	m.removedcars = nil
}

// AddGroupIDs adds the "groups" edge to the Group entity by ids.
func (m *UserMutation) AddGroupIDs(ids ...int) {
	if m.groups == nil {
		m.groups = make(map[int]struct{})
	}
	for i := range ids {
		m.groups[ids[i]] = struct{}{}
	}
}

// ClearGroups clears the "groups" edge to the Group entity.
func (m *UserMutation) ClearGroups() {
	m.clearedgroups = true
}

// GroupsCleared reports if the "groups" edge to the Group entity was cleared.
func (m *UserMutation) GroupsCleared() bool {
	return m.clearedgroups
}

// RemoveGroupIDs removes the "groups" edge to the Group entity by IDs.
func (m *UserMutation) RemoveGroupIDs(ids ...int) {
	if m.removedgroups == nil {
		m.removedgroups = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.groups, ids[i])
		m.removedgroups[ids[i]] = struct{}{}
	}
}

// RemovedGroups returns the removed IDs of the "groups" edge to the Group entity.
func (m *UserMutation) RemovedGroupsIDs() (ids []int) {
	for id := range m.removedgroups {
		ids = append(ids, id)
	}
	return
}

// GroupsIDs returns the "groups" edge IDs in the mutation.
func (m *UserMutation) GroupsIDs() (ids []int) {
	for id := range m.groups {
		ids = append(ids, id)
	}
	return
}

// ResetGroups resets all changes to the "groups" edge.
func (m *UserMutation) ResetGroups() {
	m.groups = nil
	m.clearedgroups = false
	m.removedgroups = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.age != nil {
		fields = append(fields, user.FieldAge)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.Age()
	case user.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldAge:
		return m.OldAge(ctx)
	case user.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, user.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldAge:
		m.ResetAge()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.cars != nil {
		edges = append(edges, user.EdgeCars)
	}
	if m.groups != nil {
		edges = append(edges, user.EdgeGroups)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeCars:
		ids := make([]ent.Value, 0, len(m.cars))
		for id := range m.cars {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeGroups:
		ids := make([]ent.Value, 0, len(m.groups))
		for id := range m.groups {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedcars != nil {
		edges = append(edges, user.EdgeCars)
	}
	if m.removedgroups != nil {
		edges = append(edges, user.EdgeGroups)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeCars:
		ids := make([]ent.Value, 0, len(m.removedcars))
		for id := range m.removedcars {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeGroups:
		ids := make([]ent.Value, 0, len(m.removedgroups))
		for id := range m.removedgroups {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedcars {
		edges = append(edges, user.EdgeCars)
	}
	if m.clearedgroups {
		edges = append(edges, user.EdgeGroups)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeCars:
		return m.clearedcars
	case user.EdgeGroups:
		return m.clearedgroups
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeCars:
		m.ResetCars()
		return nil
	case user.EdgeGroups:
		m.ResetGroups()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
