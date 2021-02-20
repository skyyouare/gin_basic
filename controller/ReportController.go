package controller

import (
	"errors"
	"fmt"
	"gin_basic/middleware"
	"gin_basic/pkg/tools"
	"net/http"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
)

//ReqBody 请求参数
type ReqBody struct {
	Offset   []int        `json:"offset"`
	Title    []string     `json:"title"`
	Data     [][][]string `json:"data"`
	FileName string       `json:"file_name"`
	Muban    string       `json:"muban"`
}

//ReportController 控制器
type ReportController struct {
}

//ReportRegister 路由注册
func ReportRegister(router *gin.RouterGroup) {
	controller := new(ReportController)
	router.POST("/report/index", controller.index)
	router.GET("/report/test", controller.test)

}

//测试
func (t *ReportController) test(c *gin.Context) {
	c.String(http.StatusOK, "Welcome Gin Server14")
}

//商机基础报表
func (t *ReportController) index(c *gin.Context) {
	var reqInfo ReqBody
	err := c.BindJSON(&reqInfo)
	if err != nil {
		middleware.ResponseError(c, middleware.ErrorCode, err)
		return
	}
	offset := reqInfo.Offset
	titles := reqInfo.Title
	lists := reqInfo.Data
	file := reqInfo.FileName
	muban := reqInfo.Muban
	if len(offset) == 0 {
		middleware.ResponseError(c, middleware.ErrorCode, errors.New("offset格式不正确"))
		return
	}
	if len(titles) == 0 {
		middleware.ResponseError(c, middleware.ErrorCode, errors.New("title格式不正确"))
		return
	}
	if len(lists) == 0 {
		middleware.ResponseError(c, middleware.ErrorCode, errors.New("data格式不正确"))
		return
	}
	if file == "" {
		middleware.ResponseError(c, middleware.ErrorCode, errors.New("file不能为空"))
		return
	}
	//判断是否有模板参数，生成 *excelize.File
	var f *excelize.File
	if muban != "" {
		res, err := http.Get(muban)
		if err != nil {
			middleware.ResponseError(c, middleware.ErrorCode, err)
			return
		}
		//读取excel
		f, err = excelize.OpenReader(res.Body, excelize.Options{})
		if err != nil {
			middleware.ResponseError(c, middleware.ErrorCode, err)
			return
		}
	} else {
		f = excelize.NewFile()
	}
	//设置新增sheet，title
	for k, title := range titles {
		if k == 0 {
			f.SetSheetName(f.GetSheetName(0), title)
		} else {
			f.NewSheet(title)
		}
	}
	//遍历写入
	for index, list := range lists {
		//设置当前sheet
		f.SetActiveSheet(index)
		for lineNum, v := range list {
			// Set value of a cell.
			colNum := 0
			for _, vv := range v {
				colNum++
				sheetPosition := tools.Num2str(colNum) + strconv.Itoa(lineNum+offset[index])
				//格式转换
				fvv, err := strconv.ParseFloat(vv, 64)
				if err == nil {
					err = f.SetCellValue(titles[index], sheetPosition, fvv)
				} else {
					err = f.SetCellValue(titles[index], sheetPosition, vv)
				}
				if err != nil {
					middleware.ResponseError(c, middleware.ErrorCode, err)
					return
				}
			}
		}
	}
	fmt.Println(file)
	//输出二进制流
	c.Header("response-type", "blob")
	data, _ := f.WriteToBuffer()
	c.Data(http.StatusOK, "application/vnd.ms-excel", data.Bytes())
	return
}
