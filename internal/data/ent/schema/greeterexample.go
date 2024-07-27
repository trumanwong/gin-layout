package schema

import "entgo.io/ent"

// GreeterExample holds the schema definition for the GreeterExample entity.
type GreeterExample struct {
	ent.Schema
}

// Fields of the GreeterExample.
func (GreeterExample) Fields() []ent.Field {
	return nil
}

// Edges of the GreeterExample.
func (GreeterExample) Edges() []ent.Edge {
	return nil
}

func (GreeterExample) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}
