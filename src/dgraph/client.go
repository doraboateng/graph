package dgraph

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
)

// GetClient connects to all Dgraph Alpha instances.
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
			log.Printf("Error while closing Dgraph connection: %v", err)
		}
	}
}

// LoadSchema updates the Dgraph schema and indices.
func LoadSchema(client *dgo.Dgraph) error {
	log.Println("Refreshing Graph schema...")

	schemaByteStr, err := readFile("GRAPH_SCHEMA_PATH", "graph.gql")
	if err != nil {
		return err
	}

	indicesByteStr, err := readFile("GRAPH_INDICES_PATH", "indices.dgraph")
	if err != nil {
		return err
	}

	_, err = http.Post(
		"http://alpha:8080/admin/schema",
		"text/plain",
		bytes.NewBuffer(schemaByteStr),
	)

	if err != nil {
		return err
	}

	err = client.Alter(context.Background(), &api.Operation{
		RunInBackground: true,
		Schema:          string(indicesByteStr),
	})

	if err != nil {
		return err
	}

	return err
}

// RefreshSchema ...
func RefreshSchema() {
	client, close := GetClient()
	defer close()

	LoadSchema(client)
}

func readFile(envKey string, filename string) ([]byte, error) {
	filepath := "./src/schema/" + filename

	if envFilePath, ok := os.LookupEnv(envKey); ok {
		filepath = envFilePath
	}

	byteStr, err := ioutil.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	return byteStr, nil
}
