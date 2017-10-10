package testSchemata

import (
	"github.com/graphql-go/graphql"
)

// Root struct for RootType
type Root struct {
	Field1 string `json:"Field1"`
}

// RootType is the Object on root
var RootType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootType",
	Fields: graphql.Fields{
		"Field1": &graphql.Field{Type: graphql.String, Description: "Just a string"},
	},
})

// RootSchema the root-Schema
var RootSchema = graphql.SchemaConfig{
	Query: graphql.NewObject(
		graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"RootType": &graphql.Field{
					Type: RootType,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						root := Root{Field1: "value"}
						return root, nil
					},
				},
			},
		},
	),
}
