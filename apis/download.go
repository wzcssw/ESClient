package apis

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
)

// Download  Download
func Download(c echo.Context) error {
	filePath := c.FormValue("file_path")
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return c.HTML(404, "文件不存在")
	}
	// 获得文件名
	strs := strings.Split(filePath, "/")
	fileName := strs[len(strs)-1]
	header := c.Response().Header()
	header.Set("Content-Disposition", "attachment; filename="+fileName)
	return c.Stream(http.StatusOK, "application/dicom", file)
}
