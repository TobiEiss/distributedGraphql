package distributedGraphql_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Jeffail/gabs"
	"github.com/fino-digital/distributedGraphql/testSchemata"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func TestQuerySchema(t *testing.T) {
	// GIVEN
	schema, err := graphql.NewSchema(testSchemata.RootSchema)
	if err != nil {
		t.Fail()
	}
	params := graphql.Params{Schema: schema, RequestString: `{RootType {Field1}}`}

	// TEST
	result := graphql.Do(params)

	// CHECK
	if len(result.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", result.Errors)
	}

	resultDataByte, _ := json.Marshal(result.Data)
	jsonParsed, _ := gabs.ParseJSON(resultDataByte)
	var root testSchemata.Root
	if err := json.Unmarshal(jsonParsed.Path("RootType").Bytes(), &root); err != nil {
		rJSON, _ := json.Marshal(result.Data)
		fmt.Printf("%s \n", rJSON)
		t.Fail()
	}
}

func Test1QuerySchemaThroughHandler(t *testing.T) {
	// GIVEN
	schema, err := graphql.NewSchema(testSchemata.RootSchema)
	if err != nil {
		t.Fail()
	}
	testServer := httptest.NewServer(handler.New(&handler.Config{Schema: &schema, Pretty: true}))
	defer testServer.Close()

	testQuery := `{RootType {Field1}}`

	// TEST
	response, err := http.Post(testServer.URL, handler.ContentTypeGraphQL, strings.NewReader(testQuery))
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	var result graphql.Result
	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)

	if err = json.Unmarshal(bytes, &result); err != nil {
		log.Println(err)
		t.Fail()
	}

	// CHECK
	if len(result.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", result.Errors)
	}
	jsonParsed, _ := gabs.ParseJSON(bytes)
	var root testSchemata.Root
	if err := json.Unmarshal(jsonParsed.Path("data.RootType").Bytes(), &root); err != nil {
		rJSON, _ := json.Marshal(result.Data)
		fmt.Printf("%s \n", rJSON)
		t.Fail()
	}
}
