package resolvers

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/kwcay/boateng-graph-service/src/generated"
)

type languageResolver struct{ *Resolver }

// Assert that todoResolver conforms to the generated.TodoResolver interface
// var _ generated.LanguageResolver = (*languageResolver)(nil)

func (r *queryResolver) Language(ctx context.Context, code string) (*generated.Language, error) {
	transaction := r.Dgraph.NewReadOnlyTxn()
	defer transaction.Discard(ctx)

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

	response, err := transaction.QueryWithVars(ctx, query, variables)

	if err != nil {
		return nil, err
	}

	// Convert Dgraph's JSON shape into the intended JSON shape so
	// we can unmarshal it properly.
	responseJSON := strings.ReplaceAll(string(response.Json), "Language.", "")

	type ResponseObj struct {
		Result []generated.Language `json:"result"`
	}

	var responseObj ResponseObj
	err = json.Unmarshal([]byte(responseJSON), &responseObj)

	if err != nil {
		return nil, err
	}

	var language *generated.Language

	if len(responseObj.Result) == 1 {
		language = &responseObj.Result[0]
	}

	return language, nil
}

// Languages ...
func (r *queryResolver) Languages(ctx context.Context) ([]*generated.Language, error) {
	transaction := r.Dgraph.NewReadOnlyTxn()
	defer transaction.Discard(ctx)

	response, err := transaction.Query(ctx, `{
		result(func: type(Language)) {
			<Language.code>
		}
	}`)

	if err != nil {
		return nil, err
	}

	// Convert Dgraph's JSON shape into the intended JSON shape so
	// we can unmarshal it properly.
	responseJSON := strings.ReplaceAll(string(response.Json), "Language.", "")

	type ResponseObj struct {
		Result []generated.Language `json:"result"`
	}

	var responseObj ResponseObj
	err = json.Unmarshal([]byte(responseJSON), &responseObj)

	if err != nil {
		return nil, err
	}

	var languages []*generated.Language

	for i := 0; i < len(responseObj.Result); i++ {
		languages = append(languages, &responseObj.Result[i])
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
