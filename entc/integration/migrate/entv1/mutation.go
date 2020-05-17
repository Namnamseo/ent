// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv1

import (
	"context"
	"fmt"
	"sync"

	"github.com/facebookincubator/ent/entc/integration/migrate/entv1/car"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv1/user"

	"github.com/facebookincubator/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCar  = "Car"
	TypeUser = "User"
)

// CarMutation represents an operation that mutate the Cars
// nodes in the graph.
type CarMutation struct {
	config
	op            Op
	typ           string
	id            *int
	clearedFields map[string]struct{}
	owner         *int
	clearedowner  bool
	oldValue      func(context.Context) (*Car, error)
}

var _ ent.Mutation = (*CarMutation)(nil)

// carOption allows to manage the mutation configuration using functional options.
type carOption func(*CarMutation)

// newCarMutation creates new mutation for $n.Name.
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

// withCarID sets the id field of the mutation.
func withCarID(id int) carOption {
	return func(m *CarMutation) {
		var (
			err   error
			once  sync.Once
			value *Car
		)
		m.oldValue = func(ctx context.Context) (*Car, error) {
			once.Do(func() {
				value, err = m.Client().Car.Get(ctx, id)
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
		return nil, fmt.Errorf("entv1: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *CarMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetOwnerID sets the owner edge to User by id.
func (m *CarMutation) SetOwnerID(id int) {
	m.owner = &id
}

// ClearOwner clears the owner edge to User.
func (m *CarMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared returns if the edge owner was cleared.
func (m *CarMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerID returns the owner id in the mutation.
func (m *CarMutation) OwnerID() (id int, exists bool) {
	if m.owner != nil {
		return *m.owner, true
	}
	return
}

// OwnerIDs returns the owner ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// OwnerID instead. It exists only for internal usage by the builders.
func (m *CarMutation) OwnerIDs() (ids []int) {
	if id := m.owner; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwner reset all changes of the "owner" edge.
func (m *CarMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// Op returns the operation name.
func (m *CarMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Car).
func (m *CarMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *CarMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *CarMutation) Field(name string) (ent.Value, bool) {
	switch name {
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *CarMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	}
	return nil, fmt.Errorf("unknown Car field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *CarMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Car field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *CarMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *CarMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *CarMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Car numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *CarMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *CarMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *CarMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Car nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *CarMutation) ResetField(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Car field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *CarMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owner != nil {
		edges = append(edges, car.EdgeOwner)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *CarMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case car.EdgeOwner:
		if id := m.owner; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *CarMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *CarMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *CarMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowner {
		edges = append(edges, car.EdgeOwner)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *CarMutation) EdgeCleared(name string) bool {
	switch name {
	case car.EdgeOwner:
		return m.clearedowner
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *CarMutation) ClearEdge(name string) error {
	switch name {
	case car.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown Car unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *CarMutation) ResetEdge(name string) error {
	switch name {
	case car.EdgeOwner:
		m.ResetOwner()
		return nil
	}
	return fmt.Errorf("unknown Car edge %s", name)
}

// UserMutation represents an operation that mutate the Users
// nodes in the graph.
type UserMutation struct {
	config
	op              Op
	typ             string
	id              *int
	age             *int32
	addage          *int32
	name            *string
	nickname        *string
	address         *string
	renamed         *string
	blob            *[]byte
	state           *user.State
	clearedFields   map[string]struct{}
	parent          *int
	clearedparent   bool
	children        map[int]struct{}
	removedchildren map[int]struct{}
	spouse          *int
	clearedspouse   bool
	car             *int
	clearedcar      bool
	oldValue        func(context.Context) (*User, error)
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows to manage the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for $n.Name.
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

// withUserID sets the id field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				value, err = m.Client().User.Get(ctx, id)
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
		return nil, fmt.Errorf("entv1: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on User creation.
func (m *UserMutation) SetID(id int) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetAge sets the age field.
func (m *UserMutation) SetAge(i int32) {
	m.age = &i
	m.addage = nil
}

// Age returns the age value in the mutation.
func (m *UserMutation) Age() (r int32, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old age value, if exists.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldAge(ctx context.Context) (v int32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAge is allowed only on UpdateOne operations")
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

// AddAge adds i to age.
func (m *UserMutation) AddAge(i int32) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the age field in this mutation.
func (m *UserMutation) AddedAge() (r int32, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge reset all changes of the "age" field.
func (m *UserMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetName sets the name field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old name value, if exists.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is allowed only on UpdateOne operations")
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

// ResetName reset all changes of the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetNickname sets the nickname field.
func (m *UserMutation) SetNickname(s string) {
	m.nickname = &s
}

// Nickname returns the nickname value in the mutation.
func (m *UserMutation) Nickname() (r string, exists bool) {
	v := m.nickname
	if v == nil {
		return
	}
	return *v, true
}

// OldNickname returns the old nickname value, if exists.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldNickname(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldNickname is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldNickname requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNickname: %w", err)
	}
	return oldValue.Nickname, nil
}

// ResetNickname reset all changes of the "nickname" field.
func (m *UserMutation) ResetNickname() {
	m.nickname = nil
}

// SetAddress sets the address field.
func (m *UserMutation) SetAddress(s string) {
	m.address = &s
}

// Address returns the address value in the mutation.
func (m *UserMutation) Address() (r string, exists bool) {
	v := m.address
	if v == nil {
		return
	}
	return *v, true
}

// OldAddress returns the old address value, if exists.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldAddress(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAddress is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAddress requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAddress: %w", err)
	}
	return oldValue.Address, nil
}

// ClearAddress clears the value of address.
func (m *UserMutation) ClearAddress() {
	m.address = nil
	m.clearedFields[user.FieldAddress] = struct{}{}
}

// AddressCleared returns if the field address was cleared in this mutation.
func (m *UserMutation) AddressCleared() bool {
	_, ok := m.clearedFields[user.FieldAddress]
	return ok
}

// ResetAddress reset all changes of the "address" field.
func (m *UserMutation) ResetAddress() {
	m.address = nil
	delete(m.clearedFields, user.FieldAddress)
}

// SetRenamed sets the renamed field.
func (m *UserMutation) SetRenamed(s string) {
	m.renamed = &s
}

// Renamed returns the renamed value in the mutation.
func (m *UserMutation) Renamed() (r string, exists bool) {
	v := m.renamed
	if v == nil {
		return
	}
	return *v, true
}

// OldRenamed returns the old renamed value, if exists.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldRenamed(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRenamed is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRenamed requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRenamed: %w", err)
	}
	return oldValue.Renamed, nil
}

// ClearRenamed clears the value of renamed.
func (m *UserMutation) ClearRenamed() {
	m.renamed = nil
	m.clearedFields[user.FieldRenamed] = struct{}{}
}

// RenamedCleared returns if the field renamed was cleared in this mutation.
func (m *UserMutation) RenamedCleared() bool {
	_, ok := m.clearedFields[user.FieldRenamed]
	return ok
}

// ResetRenamed reset all changes of the "renamed" field.
func (m *UserMutation) ResetRenamed() {
	m.renamed = nil
	delete(m.clearedFields, user.FieldRenamed)
}

// SetBlob sets the blob field.
func (m *UserMutation) SetBlob(b []byte) {
	m.blob = &b
}

// Blob returns the blob value in the mutation.
func (m *UserMutation) Blob() (r []byte, exists bool) {
	v := m.blob
	if v == nil {
		return
	}
	return *v, true
}

// OldBlob returns the old blob value, if exists.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldBlob(ctx context.Context) (v []byte, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldBlob is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldBlob requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBlob: %w", err)
	}
	return oldValue.Blob, nil
}

// ClearBlob clears the value of blob.
func (m *UserMutation) ClearBlob() {
	m.blob = nil
	m.clearedFields[user.FieldBlob] = struct{}{}
}

// BlobCleared returns if the field blob was cleared in this mutation.
func (m *UserMutation) BlobCleared() bool {
	_, ok := m.clearedFields[user.FieldBlob]
	return ok
}

// ResetBlob reset all changes of the "blob" field.
func (m *UserMutation) ResetBlob() {
	m.blob = nil
	delete(m.clearedFields, user.FieldBlob)
}

// SetState sets the state field.
func (m *UserMutation) SetState(u user.State) {
	m.state = &u
}

// State returns the state value in the mutation.
func (m *UserMutation) State() (r user.State, exists bool) {
	v := m.state
	if v == nil {
		return
	}
	return *v, true
}

// OldState returns the old state value, if exists.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldState(ctx context.Context) (v user.State, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldState is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldState requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldState: %w", err)
	}
	return oldValue.State, nil
}

// ClearState clears the value of state.
func (m *UserMutation) ClearState() {
	m.state = nil
	m.clearedFields[user.FieldState] = struct{}{}
}

// StateCleared returns if the field state was cleared in this mutation.
func (m *UserMutation) StateCleared() bool {
	_, ok := m.clearedFields[user.FieldState]
	return ok
}

// ResetState reset all changes of the "state" field.
func (m *UserMutation) ResetState() {
	m.state = nil
	delete(m.clearedFields, user.FieldState)
}

// SetParentID sets the parent edge to User by id.
func (m *UserMutation) SetParentID(id int) {
	m.parent = &id
}

// ClearParent clears the parent edge to User.
func (m *UserMutation) ClearParent() {
	m.clearedparent = true
}

// ParentCleared returns if the edge parent was cleared.
func (m *UserMutation) ParentCleared() bool {
	return m.clearedparent
}

// ParentID returns the parent id in the mutation.
func (m *UserMutation) ParentID() (id int, exists bool) {
	if m.parent != nil {
		return *m.parent, true
	}
	return
}

// ParentIDs returns the parent ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// ParentID instead. It exists only for internal usage by the builders.
func (m *UserMutation) ParentIDs() (ids []int) {
	if id := m.parent; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetParent reset all changes of the "parent" edge.
func (m *UserMutation) ResetParent() {
	m.parent = nil
	m.clearedparent = false
}

// AddChildIDs adds the children edge to User by ids.
func (m *UserMutation) AddChildIDs(ids ...int) {
	if m.children == nil {
		m.children = make(map[int]struct{})
	}
	for i := range ids {
		m.children[ids[i]] = struct{}{}
	}
}

// RemoveChildIDs removes the children edge to User by ids.
func (m *UserMutation) RemoveChildIDs(ids ...int) {
	if m.removedchildren == nil {
		m.removedchildren = make(map[int]struct{})
	}
	for i := range ids {
		m.removedchildren[ids[i]] = struct{}{}
	}
}

// RemovedChildren returns the removed ids of children.
func (m *UserMutation) RemovedChildrenIDs() (ids []int) {
	for id := range m.removedchildren {
		ids = append(ids, id)
	}
	return
}

// ChildrenIDs returns the children ids in the mutation.
func (m *UserMutation) ChildrenIDs() (ids []int) {
	for id := range m.children {
		ids = append(ids, id)
	}
	return
}

// ResetChildren reset all changes of the "children" edge.
func (m *UserMutation) ResetChildren() {
	m.children = nil
	m.removedchildren = nil
}

// SetSpouseID sets the spouse edge to User by id.
func (m *UserMutation) SetSpouseID(id int) {
	m.spouse = &id
}

// ClearSpouse clears the spouse edge to User.
func (m *UserMutation) ClearSpouse() {
	m.clearedspouse = true
}

// SpouseCleared returns if the edge spouse was cleared.
func (m *UserMutation) SpouseCleared() bool {
	return m.clearedspouse
}

// SpouseID returns the spouse id in the mutation.
func (m *UserMutation) SpouseID() (id int, exists bool) {
	if m.spouse != nil {
		return *m.spouse, true
	}
	return
}

// SpouseIDs returns the spouse ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// SpouseID instead. It exists only for internal usage by the builders.
func (m *UserMutation) SpouseIDs() (ids []int) {
	if id := m.spouse; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetSpouse reset all changes of the "spouse" edge.
func (m *UserMutation) ResetSpouse() {
	m.spouse = nil
	m.clearedspouse = false
}

// SetCarID sets the car edge to Car by id.
func (m *UserMutation) SetCarID(id int) {
	m.car = &id
}

// ClearCar clears the car edge to Car.
func (m *UserMutation) ClearCar() {
	m.clearedcar = true
}

// CarCleared returns if the edge car was cleared.
func (m *UserMutation) CarCleared() bool {
	return m.clearedcar
}

// CarID returns the car id in the mutation.
func (m *UserMutation) CarID() (id int, exists bool) {
	if m.car != nil {
		return *m.car, true
	}
	return
}

// CarIDs returns the car ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// CarID instead. It exists only for internal usage by the builders.
func (m *UserMutation) CarIDs() (ids []int) {
	if id := m.car; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetCar reset all changes of the "car" edge.
func (m *UserMutation) ResetCar() {
	m.car = nil
	m.clearedcar = false
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.age != nil {
		fields = append(fields, user.FieldAge)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.nickname != nil {
		fields = append(fields, user.FieldNickname)
	}
	if m.address != nil {
		fields = append(fields, user.FieldAddress)
	}
	if m.renamed != nil {
		fields = append(fields, user.FieldRenamed)
	}
	if m.blob != nil {
		fields = append(fields, user.FieldBlob)
	}
	if m.state != nil {
		fields = append(fields, user.FieldState)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.Age()
	case user.FieldName:
		return m.Name()
	case user.FieldNickname:
		return m.Nickname()
	case user.FieldAddress:
		return m.Address()
	case user.FieldRenamed:
		return m.Renamed()
	case user.FieldBlob:
		return m.Blob()
	case user.FieldState:
		return m.State()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldAge:
		return m.OldAge(ctx)
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldNickname:
		return m.OldNickname(ctx)
	case user.FieldAddress:
		return m.OldAddress(ctx)
	case user.FieldRenamed:
		return m.OldRenamed(ctx)
	case user.FieldBlob:
		return m.OldBlob(ctx)
	case user.FieldState:
		return m.OldState(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int32)
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
	case user.FieldNickname:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNickname(v)
		return nil
	case user.FieldAddress:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAddress(v)
		return nil
	case user.FieldRenamed:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRenamed(v)
		return nil
	case user.FieldBlob:
		v, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBlob(v)
		return nil
	case user.FieldState:
		v, ok := value.(user.State)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetState(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, user.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldAddress) {
		fields = append(fields, user.FieldAddress)
	}
	if m.FieldCleared(user.FieldRenamed) {
		fields = append(fields, user.FieldRenamed)
	}
	if m.FieldCleared(user.FieldBlob) {
		fields = append(fields, user.FieldBlob)
	}
	if m.FieldCleared(user.FieldState) {
		fields = append(fields, user.FieldState)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldAddress:
		m.ClearAddress()
		return nil
	case user.FieldRenamed:
		m.ClearRenamed()
		return nil
	case user.FieldBlob:
		m.ClearBlob()
		return nil
	case user.FieldState:
		m.ClearState()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldAge:
		m.ResetAge()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldNickname:
		m.ResetNickname()
		return nil
	case user.FieldAddress:
		m.ResetAddress()
		return nil
	case user.FieldRenamed:
		m.ResetRenamed()
		return nil
	case user.FieldBlob:
		m.ResetBlob()
		return nil
	case user.FieldState:
		m.ResetState()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 4)
	if m.parent != nil {
		edges = append(edges, user.EdgeParent)
	}
	if m.children != nil {
		edges = append(edges, user.EdgeChildren)
	}
	if m.spouse != nil {
		edges = append(edges, user.EdgeSpouse)
	}
	if m.car != nil {
		edges = append(edges, user.EdgeCar)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeParent:
		if id := m.parent; id != nil {
			return []ent.Value{*id}
		}
	case user.EdgeChildren:
		ids := make([]ent.Value, 0, len(m.children))
		for id := range m.children {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeSpouse:
		if id := m.spouse; id != nil {
			return []ent.Value{*id}
		}
	case user.EdgeCar:
		if id := m.car; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 4)
	if m.removedchildren != nil {
		edges = append(edges, user.EdgeChildren)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeChildren:
		ids := make([]ent.Value, 0, len(m.removedchildren))
		for id := range m.removedchildren {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 4)
	if m.clearedparent {
		edges = append(edges, user.EdgeParent)
	}
	if m.clearedspouse {
		edges = append(edges, user.EdgeSpouse)
	}
	if m.clearedcar {
		edges = append(edges, user.EdgeCar)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeParent:
		return m.clearedparent
	case user.EdgeSpouse:
		return m.clearedspouse
	case user.EdgeCar:
		return m.clearedcar
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	case user.EdgeParent:
		m.ClearParent()
		return nil
	case user.EdgeSpouse:
		m.ClearSpouse()
		return nil
	case user.EdgeCar:
		m.ClearCar()
		return nil
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeParent:
		m.ResetParent()
		return nil
	case user.EdgeChildren:
		m.ResetChildren()
		return nil
	case user.EdgeSpouse:
		m.ResetSpouse()
		return nil
	case user.EdgeCar:
		m.ResetCar()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
