package schema

import (
	"github.com/graph-gophers/graphql-go"
	"graphql-go-gin/resolver"
	"io/ioutil"
)

var Schema *graphql.Schema

func init() {

	s, err := getSchema("./schema/schema.graphql")
	if err != nil {
		panic(err)
	}
	Schema = graphql.MustParseSchema(s, &resolver.Resolver{}, graphql.UseStringDescriptions())
}

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
