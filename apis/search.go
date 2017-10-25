package apis

import (
	"ESClient/tools"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Search 搜索
func Search(c echo.Context) error {
	searchInput := c.FormValue("search_input")
	searchInput = tools.TranslateToPinyin(searchInput)
	b, _ := tools.Query(getQueryString(searchInput, 3))

	return c.String(http.StatusOK, string(b))
}

func getQueryString(params string, length int) []byte {
	return []byte(`{	
		"size" : ` + strconv.Itoa(length) + `,
		"query": {
			"function_score": {
				"query": {
					"multi_match": {
						"query": "` + params + `",
						"fields": [
							"Patient's Name^8", 
							"Institution Name^2.5", 
							"Body Part Examined^2", 
							"Study Date^1.5",
							"SOP Instance UID^1",
							"Modality^1"
						],
						"fuzziness" : "AUTO",
						"prefix_length" : 1
					}
				}
			}
		}
	}`)
}

// SearchSimple 搜索(可选择字段)
func SearchSimple(c echo.Context) error {
	fields := c.FormValue("fields")
	searchInput := c.FormValue("search_input")
	length := c.FormValue("length")
	lengthNum := 10
	if fields == "" {
		return c.String(http.StatusOK, "缺少[fields]参数")
	}
	if searchInput == "" {
		return c.String(http.StatusOK, "缺少[searchInput]参数")
	}
	if length != "" {
		n, err := strconv.Atoi(length)
		if err == nil {
			lengthNum = n
		}
	}
	searchInput = tools.TranslateToPinyin(searchInput)
	byteResult, _ := tools.Query(getQueryString(searchInput, lengthNum))
	result, _ := tools.ProcessResult(fields, byteResult)

	return c.String(http.StatusOK, string(result))
}
