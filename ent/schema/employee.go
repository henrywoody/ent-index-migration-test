package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Employee holds the schema definition for the Employee entity.
type Employee struct {
	ent.Schema
}

func (Employee) Config() ent.Config {
	return ent.Config{
		Table: "employee",
	}
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name"),
		field.String("middle_name"),
		field.String("last_name"),
	}
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return nil
}

// Indexes of the Employee.
func (Employee) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("first_name", "middle_name", "last_name").Unique(),
	}
}
