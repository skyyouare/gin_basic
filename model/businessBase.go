package model

import (
	"fmt"
	"gin_basic/pkg/db"
)

// 组装数据
func BusinessBasic(day string, endDayLastMonth string, oids string) (lists [][][]string, titles []string, err error) {
	// 组装数据切片
	lists = make([][][]string, 3)
	lists[0], err = getBusinessBaseData("operating_base_business_exec", day, endDayLastMonth, oids)
	if err != nil {
		return lists, titles, nil
	}
	lists[1], err = getBusinessBaseData("operating_base_business_cash", day, endDayLastMonth, oids)
	if err != nil {
		return lists, titles, nil
	}
	lists[2], err = getBusinessBaseData("operating_base_business_invoice", day, endDayLastMonth, oids)
	if err != nil {
		return lists, titles, nil
	}
	// 标题切片
	titles = []string{"商机基础执行", "商机基础现金流", "商机基础开票"}
	// 文件名
	return lists, titles, nil
}

// 查询数据
func getBusinessBaseData(table string, day string, lastDay string, oids string) (res [][]string, err error) {
	// 获取表头
	title, err := getBusinessBaseTitle(table)
	if err != nil {
		fmt.Printf("getBusinessBaseTitle faied, %v\n", err)
		return res, fmt.Errorf("getBusinessBaseTitle faied2, %v", err)
	}
	last, err := getBusinessBaseDataByLastDay(table, lastDay, oids)
	if err != nil {
		fmt.Printf("getBusinessBaseDataByLastDay faied, %v\n", err)
		return res, fmt.Errorf("getBusinessBaseDataByLastDay faied2, %v", err)
	}
	res, err = getBusinessBaseDataByDay(table, day, last, title, oids)
	if err != nil {
		fmt.Printf("getBusinessBaseDataByDay faied, %v\n", err)
		return res, fmt.Errorf("getBusinessBaseDataByDay faied2, %v", err)
	}
	return res, nil
}

// 获取表头
func getBusinessBaseTitle(table string) (res []string, err error) {
	// 获取当月
	rows, err := db.Conn.Query("show full columns from " + table)
	if err != nil {
		fmt.Printf("query faied, %v\n", err)
		return res, fmt.Errorf("query faied, %v", err)
	}
	res, err = getComments(rows)
	if err != nil {
		return res, fmt.Errorf("getComments failed, %v", err)
	}
	// 关闭结果集（释放连接）
	err = rows.Close()
	if err != nil {
		fmt.Printf("rows close faied, %v\n", err)
		return res, fmt.Errorf("rows close faied, %v", err)
	}
	return res, nil
}

// 获取当月数据
func getBusinessBaseDataByDay(table string, day string, last map[string][]string, title []string, oids string) (res [][]string, err error) {
	// 获取当月
	rows, err := db.Conn.Query("select * from "+table+" where day=? and accounting_department_id in ("+oids+")", day)
	if err != nil {
		fmt.Printf("query faied, %v\n", err)
		return res, fmt.Errorf("query faied, %v", err)
	}
	res, err = getMergeRows(rows, last, title)
	if err != nil {
		return res, fmt.Errorf("getAllRowsSlice failed, %v", err)
	}
	// fmt.Println(res)
	// 关闭结果集（释放连接）
	err = rows.Close()
	if err != nil {
		fmt.Printf("rows close faied, %v\n", err)
		return res, fmt.Errorf("rows close faied, %v", err)
	}
	return res, nil
}

// 获取上月数据
func getBusinessBaseDataByLastDay(table string, lastDay string, oids string) (res map[string][]string, err error) {
	// 获取上月最后一天
	rows, err := db.Conn.Query("select * from "+table+" where day=? and accounting_department_id in ("+oids+")", lastDay)
	if err != nil {
		fmt.Printf("query faied, %v\n", err)
		return res, fmt.Errorf("query faied, %v", err)
	}
	res, err = getAllRowsMapSlice(rows)
	if err != nil {
		return res, fmt.Errorf("getAllRowsMapSlice failed, %v", err)
	}
	// 关闭结果集（释放连接）
	err = rows.Close()
	if err != nil {
		fmt.Printf("rows close faied, %v\n", err)
		return res, fmt.Errorf("rows close faied, %v", err)
	}
	return res, nil
}
