package subdir

import (
	"testing"

	"github.com/oscerai/gqlgen/_examples/embedding/subdir/gendir"
	"github.com/oscerai/gqlgen/client"
	"github.com/oscerai/gqlgen/graphql/handler"
	"github.com/stretchr/testify/require"
)

func TestEmbeddingWorks(t *testing.T) {
	c := client.New(handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{}})))
	var resp struct {
		InSchemadir string
		Parentdir   string
		Subdir      string
	}
	c.MustPost(`{
				inSchemadir
				parentdir
				subdir
			}
		`, &resp)

	require.Equal(t, "example", resp.InSchemadir)
	require.Equal(t, "example", resp.Parentdir)
	require.Equal(t, "example", resp.Subdir)
}

func TestEmbeddingWorksInGendir(t *testing.T) {
	c := client.New(handler.NewDefaultServer(gendir.NewExecutableSchema(gendir.Config{Resolvers: &GendirResolver{}})))
	var resp struct {
		InSchemadir string
		Parentdir   string
		Subdir      string
	}
	c.MustPost(`{
				inSchemadir
				parentdir
				subdir
			}
		`, &resp)

	require.Equal(t, "example", resp.InSchemadir)
	require.Equal(t, "example", resp.Parentdir)
	require.Equal(t, "example", resp.Subdir)
}
