package model

import (
	"database/sql"
	"fmt"
)

//根据上月数据合并
func getMergeRows(rows *sql.Rows, last map[string][]string, title []string) ([][]string, error) {
	//返回列名称组成的slince,也就是字段名的合集
	columns, _ := rows.Columns()
	//vals用来存放取出来的数据结果，表示一行所有列的值，后面的长度表示行数
	vals := make([][]byte, len(columns))
	//这个切片用来做rows,scan参数，将扫描后的数据存储在scan中。（将数据库的值复制到scan中）
	scans := make([]interface{}, len(columns))
	//将每一行数据填充到[][]byte中
	for k := range vals {
		//因为rows.scan是指针类型，这里遍历将指针变量保存到切片中
		scans[k] = &vals[k]
	}
	//key为column的中的字段名，值为字段名和记录
	result := make([][]string, 1)
	result = append(result, title)
	i := 0
	for rows.Next() {
		//将查询的结果放入scan中，也就是放到vals变量中
		if err := rows.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			fmt.Println(err)
			return result, err
		}
		//获取上月key
		var lastKey string
		for k, v := range vals {
			key := columns[k] //字段名
			if key == "clue_num" || key == "ver" {
				if key == "clue_num" {
					lastKey = string(v) + "_" + lastKey
				} else {
					lastKey = string(v) + lastKey
				}
			}
		}
		//因为是byte切片，所以循环取出转成string
		row := make([]string, len(columns)*2)
		for k, v := range vals {
			if _, ok := last[lastKey]; ok {
				row[k*2] = last[lastKey][k]
			} else {
				row[k*2] = ""
			}
			row[k*2+1] = string(v) //原来是byte类型，现在转成string类型
		}
		//加入切片中
		result = append(result, row)
		i++
	}
	return result, nil
}

//组装数据 map slice
func getAllRowsMapSlice(rows *sql.Rows) (map[string][]string, error) {
	//返回列名称组成的slice,也就是字段名的合集
	columns, _ := rows.Columns()
	//vals用来存放取出来的数据结果，表示一行所有列的值，后面的长度表示行数
	vals := make([][]byte, len(columns))
	//这个切片用来做rows,scan参数，将扫描后的数据存储在scan中。（将数据库的值复制到scan中）
	scans := make([]interface{}, len(columns))
	//将每一行数据填充到[][]byte中
	for k := range vals {
		//因为rows.scan是指针类型，这里遍历将指针变量保存到切片中
		scans[k] = &vals[k]
	}
	//key为column的中的字段名，值为字段名和记录
	result := make(map[string][]string)
	for rows.Next() {
		//将查询的结果放入scan中，也就是放到vals变量中
		if err := rows.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			// fmt.Println(err)
			return result, err
		}
		//因为是byte切片，所以循环取出转成string
		row := make([]string, len(columns))
		var i string
		for k, v := range vals {
			key := columns[k] //字段名
			if key == "clue_num" || key == "ver" {
				if key == "clue_num" {
					i = string(v) + "_" + i
				} else {
					i = string(v) + i
				}
			}
			row[k] = string(v) //原来是byte类型，现在转成string类型
		}
		//放入总结果集中，i用来记录读取的条数
		result[i] = row
	}
	return result, nil
}

//获取备注
func getComments(rows *sql.Rows) ([]string, error) {
	//返回列名称组成的slice,也就是字段名的合集
	columns, _ := rows.Columns()
	//vals用来存放取出来的数据结果，表示一行所有列的值，后面的长度表示行数
	vals := make([][]byte, len(columns))
	//这个切片用来做rows,scan参数，将扫描后的数据存储在scan中。（将数据库的值复制到scan中）
	scans := make([]interface{}, len(columns))
	//将每一行数据填充到[][]byte中
	for k := range vals {
		//因为rows.scan是指针类型，这里遍历将指针变量保存到切片中
		scans[k] = &vals[k]
	}
	//key为column的中的字段名，值为字段名和记录
	result := make([]string, 0)
	i := 0
	for rows.Next() {
		//将查询的结果放入scan中，也就是放到vals变量中
		if err := rows.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			// fmt.Println(err)
			return result, err
		}
		for k, v := range vals {
			key := columns[k] //字段名
			if key == "Comment" {
				result = append(result, "上月"+string(v))
				result = append(result, string(v))
			}
		}
		//放入总结果集中，i用来记录读取的条数
		i++
	}
	return result, nil
}

//组装数据切片
func getAllRowsSlice(rows *sql.Rows) ([][]string, error) {
	//返回列名称组成的slince,也就是字段名的合集
	columns, _ := rows.Columns()
	//vals用来存放取出来的数据结果，表示一行所有列的值，后面的长度表示行数
	vals := make([][]byte, len(columns))
	//这个切片用来做rows,scan参数，将扫描后的数据存储在scan中。（将数据库的值复制到scan中）
	scans := make([]interface{}, len(columns))
	//将每一行数据填充到[][]byte中
	for k := range vals {
		//因为rows.scan是指针类型，这里遍历将指针变量保存到切片中
		scans[k] = &vals[k]
	}
	//key为column的中的字段名，值为字段名和记录
	result := make([][]string, 0)
	i := 0
	for rows.Next() {
		//将查询的结果放入scan中，也就是放到vals变量中
		if err := rows.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			fmt.Println(err)
			return result, err
		}
		//因为是byte切片，所以循环取出转成string
		row := make([]string, len(columns))
		for k, v := range vals {
			row[k] = string(v) //原来是byte类型，现在转成string类型
		}
		//加入切片中
		result = append(result, row)
		i++
	}
	return result, nil
}

//组装数据map
func getAllRowsMap(rows *sql.Rows) map[int]map[string]string {
	//返回列名称组成的slice,也就是字段名的合集
	columns, _ := rows.Columns()
	//vals用来存放取出来的数据结果，表示一行所有列的值，后面的长度表示行数
	vals := make([][]byte, len(columns))
	//这个切片用来做rows,scan参数，将扫描后的数据存储在scan中。（将数据库的值复制到scan中）
	scans := make([]interface{}, len(columns))
	//将每一行数据填充到[][]byte中
	for k := range vals {
		//因为rows.scan是指针类型，这里遍历将指针变量保存到切片中
		scans[k] = &vals[k]
	}
	//key为column的中的字段名，值为字段名和记录
	result := make(map[int]map[string]string)
	i := 0
	for rows.Next() {
		//将查询的结果放入scan中，也就是放到vals变量中
		if err := rows.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			// fmt.Println(err)
			return result
		}
		//因为是byte切片，所以循环取出转成string
		row := make(map[string]string)
		for k, v := range vals {
			key := columns[k]    //字段名
			row[key] = string(v) //原来是byte类型，现在转成string类型
		}
		//放入总结果集中，i用来记录读取的条数
		result[i] = row
		i++
	}
	fmt.Println(i)
	return result
}
