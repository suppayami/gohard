package main

import (
	"io/ioutil"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	relay "github.com/graph-gophers/graphql-go/relay"
	"github.com/suppayami/gohard/models"
	"github.com/suppayami/gohard/repositories"
	"github.com/suppayami/gohard/resolver"
)

var posts = []*models.Post{
	{
		ID:    "1",
		Title: "Unlimited Blade Works",
	},

	{
		ID:    "2",
		Title: "We don't need a title",
	},
}

var schema *graphql.Schema
var postRepo *repositories.Posts

func init() {
	postData := make(map[graphql.ID]*models.Post)
	for _, post := range posts {
		postData[post.ID] = post
	}
	postRepo = repositories.NewPosts(postData)

	schemaText, err := getSchema("./resolver/schema.graphql")
	if err != nil {
		panic(err)
	}
	schema = graphql.MustParseSchema(schemaText, &resolver.Resolver{PostRepository: postRepo})
}

func main() {
	// hax render graphiql
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))

	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var page = []byte(`
    <!DOCTYPE html>
    <html>
        <head>
            <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
            <script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
            <script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
            <script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
            <script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
        </head>
        <body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
            <div id="graphiql" style="height: 100vh;">Loading...</div>
            <script>
                function graphQLFetcher(graphQLParams) {
                    return fetch("/query", {
                        method: "post",
                        body: JSON.stringify(graphQLParams),
                        credentials: "include",
                    }).then(function (response) {
                        return response.text();
                    }).then(function (responseBody) {
                        try {
                            return JSON.parse(responseBody);
                        } catch (error) {
                            return responseBody;
                        }
                    });
                }
                ReactDOM.render(
                    React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
                    document.getElementById("graphiql")
                );
            </script>
        </body>
    </html>
    `)

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
