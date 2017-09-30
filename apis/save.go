package apis

import (
	"ESClient/tools"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

// Save  添加DICOM数据
func Save(c echo.Context) error {
	b, _ := ioutil.ReadAll(c.Request().Body)
	result, _ := tools.Save(b)
	return c.String(http.StatusOK, string(result))
}
