package dgraph

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
)

// GetClient ...
func GetClient() (*dgo.Dgraph, context.CancelFunc) {
	// Open gRPC connections to Dgraph Alpha nodes.
	conn, err := grpc.Dial("alpha:9080", grpc.WithInsecure())

	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}

	alpha1 := api.NewDgraphClient(conn)
	client := dgo.NewDgraphClient(alpha1)

	return client, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection:%v", err)
		}
	}
}

// LoadSchema ...
// TODO: The schema should be embedded with the binary file and loaded statically.
func LoadSchema(client *dgo.Dgraph) {
	indicesByteStr, err := ioutil.ReadFile("./src/schema/indices.dgraph")
	if err != nil {
		log.Fatal(err)
	}

	schemaByteStr, err := ioutil.ReadFile("./src/schema/graph.gql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = http.Post(
		"http://alpha:8080/admin/schema",
		"text/plain",
		bytes.NewBuffer(schemaByteStr),
	)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Alter(context.Background(), &api.Operation{
		RunInBackground: true,
		Schema:          string(indicesByteStr),
	})

	if err != nil {
		log.Fatal(err)
	}
}
