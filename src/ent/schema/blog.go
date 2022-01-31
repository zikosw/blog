package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blog holds the schema definition for the Blog entity.
type Blog struct {
	ent.Schema
}

// Fields of the Blog.
func (Blog) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("content"),
		field.Time("created_at").Default(time.Now),
		// field.String("slug").
		//     Unique(),
		// field.Int("age").Positive(),
		// field.String("name").Default("unknown"),
	}
}

// Edges of the Blog.
func (Blog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("slugs", Slug.Type),
	}
}
