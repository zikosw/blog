package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Slug holds the schema definition for the Slug entity.
type Slug struct {
	ent.Schema
}

// Fields of the Slug.
func (Slug) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").Unique().Immutable(),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

// Edges of the Slug.
func (Slug) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("link", Blog.Type).
			Ref("slugs").
			Unique(),
	}
}
