package apis

import (
	"ESClient/tools"
	"net/http"

	"github.com/labstack/echo"
)

// Search  s
func Search(c echo.Context) error {
	searchInput := c.FormValue("search_input")
	searchInput = tools.TranslateToPinyin(searchInput)
	b, _ := tools.Query(getQueryString(searchInput))

	return c.String(http.StatusOK, string(b))
}

func getQueryString(params string) []byte {
	return []byte(`{	
		"size" : 1500,
		"query": {
			"function_score": {
				"query": {
					"multi_match": {
						"query": "` + params + `",
						"fields": [
							"Patient's Name^3", 
							"Institution Name^2.5", 
							"Body Part Examined^2", 
							"Study Date^1.5",
							"SOP Instance UID^1",
							"Modality^1"
						],
						"fuzziness" : "AUTO",
						"prefix_length" : 3
					}
				}
			}
		}
	}`)
}

// // Search  s
// func Search(c echo.Context) error {
// 	searchInput := c.FormValue("search_input")
// 	// termQuery := elastic.NewTermQuery("Patient's Name", searchInput)
// 	multiMatchQuery := elastic.NewMultiMatchQuery(searchInput, "Patient's Name", "Institution Name")
// 	searchResult, err := tools.Client.Search().
// 		Index("dicom").
// 		Query(multiMatchQuery).
// 		// Sort("user", true). // sort by "user" field, ascending
// 		From(0).Size(1500). // take documents 0-9
// 		// Pretty(true).
// 		Do(tools.Ctx)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return c.JSON(http.StatusOK, searchResult)
// }
