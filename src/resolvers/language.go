package resolvers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/kwcay/boateng-graph-service/src/generated"
)

type languageResolver struct{ *Resolver }

// Assert that todoResolver conforms to the generated.TodoResolver interface
// var _ generated.LanguageResolver = (*languageResolver)(nil)

func (resolver *queryResolver) Language(ctx context.Context, code string) (*generated.Language, error) {
	txn := resolver.Dgraph.NewReadOnlyTxn()
	defer txn.Discard(ctx)

	variables := make(map[string]string)
	variables["$code"] = code
	log.Printf("Querying Language with %v\n", variables)

	query := `
		query GetLanguage($code: string) {
			result(func: eq(<Language.code>, $code)) {
				expand(_all_)
			}
		}
	`

	res, err := txn.QueryWithVars(ctx, query, variables)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	// Convert Dgraph's JSON shape into the intended JSON shape so
	// we can unmarshal it properly.
	jsonResult := strings.ReplaceAll(string(res.Json), "Language.", "")
	log.Printf("%s\n%s\n", res.Json, jsonResult)

	type Root struct {
		Result []generated.Language `json:"result"`
	}

	var language *generated.Language

	var root Root
	err = json.Unmarshal([]byte(jsonResult), &root)
	if err != nil {
		return nil, err
	}

	if len(root.Result) == 1 {
		language = &root.Result[0]
	}

	return language, nil
}

// Languages ...
func (resolver *queryResolver) Languages(ctx context.Context) ([]*generated.Language, error) {
	txn := resolver.Dgraph.NewReadOnlyTxn()
	defer txn.Discard(ctx)

	res, err := txn.Query(ctx, `{
		result(func: type(Language)) {
			<Language.code>
		}
	}`)

	if err != nil {
		return nil, err
	}

	// Convert Dgraph's JSON shape into the intended JSON shape so
	// we can unmarshal it properly.
	jsonResult := strings.ReplaceAll(string(res.Json), "Language.", "")
	fmt.Printf("%s\n%s\n", res.Json, jsonResult)

	type Root struct {
		Result []generated.Language `json:"result"`
	}

	var root Root
	err = json.Unmarshal([]byte(jsonResult), &root)
	if err != nil {
		return nil, err
	}
	numLanguages := len(root.Result)

	out, _ := json.MarshalIndent(root.Result, "", "\t")
	fmt.Printf("%s\n", out)

	var languages []*generated.Language

	for i := 0; i < numLanguages; i++ {
		languages = append(languages, &root.Result[i])
	}

	return languages, nil
}

// func (r *languageResolver) CreatedBy(ctx context.Context, obj *models.Todo) (*models.User, error) {
// 	return ctx.Value(dataloaders.UserLoader).(*generated.UserLoader).Load(obj.CreatedBy)
// }

// func (r *languageResolver) UpdatedBy(ctx context.Context, obj *models.Todo) (*models.User, error) {
// 	return ctx.Value(dataloaders.UserLoader).(*generated.UserLoader).Load(obj.UpdatedBy)
// }

// func (r *mutationResolver) TodoCreate(ctx context.Context, todo models.TodoInput) (*models.Todo, error) {
// 	// Validate that createdby id actually exists
// 	err := r.DB.Select(&models.User{ID: todo.CreatedBy})
// 	if err != nil {
// 		return nil, err
// 	}

// 	t := models.Todo{
// 		Name: todo.Name,

// 		CreatedBy: todo.CreatedBy,
// 		UpdatedBy: todo.CreatedBy,

// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	err = r.DB.Insert(&t)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &t, nil
// }

// func (r *mutationResolver) TodoComplete(ctx context.Context, id int, updatedBy int) (*models.Todo, error) {
// 	// Validate that updatedBy id actually exists
// 	err := r.DB.Select(&models.User{ID: updatedBy})
// 	if err != nil {
// 		return nil, errors.New(fmt.Sprintf("user %d does not exist", updatedBy))
// 	}

// 	todo := models.Todo{
// 		ID: id,
// 	}

// 	err = r.DB.Select(&todo)
// 	if err != nil {
// 		return nil, errors.New(fmt.Sprintf("todo %d does not exist", updatedBy))
// 	}

// 	todo.UpdatedBy = updatedBy
// 	todo.IsComplete = true
// 	todo.UpdatedAt = time.Now()

// 	err = r.DB.Update(&todo)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &todo, nil
// }
