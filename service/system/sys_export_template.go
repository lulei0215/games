package system

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type SysExportTemplateService struct {
}

var SysExportTemplateServiceApp = new(SysExportTemplateService)

// CreateSysExportTemplate
// Author [piexlmax](https://github.com/piexlmax)
func (sysExportTemplateService *SysExportTemplateService) CreateSysExportTemplate(sysExportTemplate *system.SysExportTemplate) (err error) {
	err = global.GVA_DB.Create(sysExportTemplate).Error
	return err
}

// DeleteSysExportTemplate
// Author [piexlmax](https://github.com/piexlmax)
func (sysExportTemplateService *SysExportTemplateService) DeleteSysExportTemplate(sysExportTemplate system.SysExportTemplate) (err error) {
	err = global.GVA_DB.Delete(&sysExportTemplate).Error
	return err
}

// DeleteSysExportTemplateByIds
// Author [piexlmax](https://github.com/piexlmax)
func (sysExportTemplateService *SysExportTemplateService) DeleteSysExportTemplateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]system.SysExportTemplate{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSysExportTemplate
// Author [piexlmax](https://github.com/piexlmax)
func (sysExportTemplateService *SysExportTemplateService) UpdateSysExportTemplate(sysExportTemplate system.SysExportTemplate) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		conditions := sysExportTemplate.Conditions
		e := tx.Delete(&[]system.Condition{}, "template_id = ?", sysExportTemplate.TemplateID).Error
		if e != nil {
			return e
		}
		sysExportTemplate.Conditions = nil

		joins := sysExportTemplate.JoinTemplate
		e = tx.Delete(&[]system.JoinTemplate{}, "template_id = ?", sysExportTemplate.TemplateID).Error
		if e != nil {
			return e
		}
		sysExportTemplate.JoinTemplate = nil

		e = tx.Updates(&sysExportTemplate).Error
		if e != nil {
			return e
		}
		if len(conditions) > 0 {
			for i := range conditions {
				conditions[i].ID = 0
			}
			e = tx.Create(&conditions).Error
		}
		if len(joins) > 0 {
			for i := range joins {
				joins[i].ID = 0
			}
			e = tx.Create(&joins).Error
		}
		return e
	})
}

// GetSysExportTemplate id
// Author [piexlmax](https://github.com/piexlmax)
func (sysExportTemplateService *SysExportTemplateService) GetSysExportTemplate(id uint) (sysExportTemplate system.SysExportTemplate, err error) {
	err = global.GVA_DB.Where("id = ?", id).Preload("JoinTemplate").Preload("Conditions").First(&sysExportTemplate).Error
	return
}

// GetSysExportTemplateInfoList
// Author [piexlmax](https://github.com/piexlmax)
func (sysExportTemplateService *SysExportTemplateService) GetSysExportTemplateInfoList(info systemReq.SysExportTemplateSearch) (list []system.SysExportTemplate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// db
	db := global.GVA_DB.Model(&system.SysExportTemplate{})
	var sysExportTemplates []system.SysExportTemplate
	//
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.TableName != "" {
		db = db.Where("table_name = ?", info.TableName)
	}
	if info.TemplateID != "" {
		db = db.Where("template_id = ?", info.TemplateID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sysExportTemplates).Error
	return sysExportTemplates, total, err
}

// ExportExcel Excel
// Author [piexlmax](https://github.com/piexlmax)
func (sysExportTemplateService *SysExportTemplateService) ExportExcel(templateID string, values url.Values) (file *bytes.Buffer, name string, err error) {
	var params = values.Get("params")
	paramsValues, err := url.ParseQuery(params)
	if err != nil {
		return nil, "", fmt.Errorf(" params : %v", err)
	}
	var template system.SysExportTemplate
	err = global.GVA_DB.Preload("Conditions").Preload("JoinTemplate").First(&template, "template_id = ?", templateID).Error
	if err != nil {
		return nil, "", err
	}
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	var templateInfoMap = make(map[string]string)
	columns, err := utils.GetJSONKeys(template.TemplateInfo)
	if err != nil {
		return nil, "", err
	}
	err = json.Unmarshal([]byte(template.TemplateInfo), &templateInfoMap)
	if err != nil {
		return nil, "", err
	}
	var tableTitle []string
	var selectKeyFmt []string
	for _, key := range columns {
		selectKeyFmt = append(selectKeyFmt, key)
		tableTitle = append(tableTitle, templateInfoMap[key])
	}

	selects := strings.Join(selectKeyFmt, ", ")
	var tableMap []map[string]interface{}
	db := global.GVA_DB
	if template.DBName != "" {
		db = global.MustGetGlobalDBByDBName(template.DBName)
	}

	if len(template.JoinTemplate) > 0 {
		for _, join := range template.JoinTemplate {
			db = db.Joins(join.JOINS + " " + join.Table + " ON " + join.ON)
		}
	}

	db = db.Select(selects).Table(template.TableName)

	filterDeleted := false

	filterParam := paramsValues.Get("filterDeleted")
	if filterParam == "true" {
		filterDeleted = true
	}

	if filterDeleted {
		//
		db = db.Where(fmt.Sprintf("%s.deleted_at IS NULL", template.TableName))

		// ()
		if len(template.JoinTemplate) > 0 {
			for _, join := range template.JoinTemplate {
				// deleted_at
				hasDeletedAt := sysExportTemplateService.hasDeletedAtColumn(join.Table)
				if hasDeletedAt {
					db = db.Where(fmt.Sprintf("%s.deleted_at IS NULL", join.Table))
				}
			}
		}
	}

	if len(template.Conditions) > 0 {
		for _, condition := range template.Conditions {
			sql := fmt.Sprintf("%s %s ?", condition.Column, condition.Operator)
			value := paramsValues.Get(condition.From)

			if condition.Operator == "IN" || condition.Operator == "NOT IN" {
				sql = fmt.Sprintf("%s %s (?)", condition.Column, condition.Operator)
			}

			if value != "" {
				if condition.Operator == "LIKE" {
					value = "%" + value + "%"
				}
				db = db.Where(sql, value)
			}
		}
	}
	// limit
	limit := paramsValues.Get("limit")
	if limit != "" {
		l, e := strconv.Atoi(limit)
		if e == nil {
			db = db.Limit(l)
		}
	}
	// limit
	if limit == "" && template.Limit != nil && *template.Limit != 0 {
		db = db.Limit(*template.Limit)
	}

	// offset
	offset := paramsValues.Get("offset")
	if offset != "" {
		o, e := strconv.Atoi(offset)
		if e == nil {
			db = db.Offset(o)
		}
	}

	//
	table := template.TableName
	orderColumns, err := db.Migrator().ColumnTypes(table)
	if err != nil {
		return nil, "", err
	}

	//  map
	fields := make(map[string]bool)

	for _, column := range orderColumns {
		fields[column.Name()] = true
	}

	// order
	order := paramsValues.Get("order")

	if order == "" && template.Order != "" {
		// order，
		order = template.Order
	}

	if order != "" {
		checkOrderArr := strings.Split(order, " ")
		orderStr := ""
		//
		if _, ok := fields[checkOrderArr[0]]; !ok {
			return nil, "", fmt.Errorf("order by %s is not in the fields", order)
		}
		orderStr = checkOrderArr[0]
		if len(checkOrderArr) > 1 {
			if checkOrderArr[1] != "asc" && checkOrderArr[1] != "desc" {
				return nil, "", fmt.Errorf("order by %s is not secure", order)
			}
			orderStr = orderStr + " " + checkOrderArr[1]
		}
		db = db.Order(orderStr)
	}

	err = db.Debug().Find(&tableMap).Error
	if err != nil {
		return nil, "", err
	}
	var rows [][]string
	rows = append(rows, tableTitle)
	for _, exTable := range tableMap {
		var row []string
		for _, column := range columns {
			column = strings.ReplaceAll(column, "\"", "")
			column = strings.ReplaceAll(column, "`", "")
			if len(template.JoinTemplate) > 0 {
				columnAs := strings.Split(column, " as ")
				if len(columnAs) > 1 {
					column = strings.TrimSpace(strings.Split(column, " as ")[1])
				} else {
					columnArr := strings.Split(column, ".")
					if len(columnArr) > 1 {
						column = strings.Split(column, ".")[1]
					}
				}
			}
			//
			if t, ok := exTable[column].(time.Time); ok {
				row = append(row, t.Format("2006-01-02 15:04:05"))
			} else {
				row = append(row, fmt.Sprintf("%v", exTable[column]))
			}
		}
		rows = append(rows, row)
	}
	for i, row := range rows {
		for j, colCell := range row {
			sErr := f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", getColumnName(j+1), i+1), colCell)
			if sErr != nil {
				return nil, "", sErr
			}
		}
	}
	f.SetActiveSheet(index)
	file, err = f.WriteToBuffer()
	if err != nil {
		return nil, "", err
	}

	return file, template.Name, nil
}

// ExportTemplate Excel
// Author [piexlmax](https://github.com/piexlmax)
func (sysExportTemplateService *SysExportTemplateService) ExportTemplate(templateID string) (file *bytes.Buffer, name string, err error) {
	var template system.SysExportTemplate
	err = global.GVA_DB.First(&template, "template_id = ?", templateID).Error
	if err != nil {
		return nil, "", err
	}
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	var templateInfoMap = make(map[string]string)

	columns, err := utils.GetJSONKeys(template.TemplateInfo)

	err = json.Unmarshal([]byte(template.TemplateInfo), &templateInfoMap)
	if err != nil {
		return nil, "", err
	}
	var tableTitle []string
	for _, key := range columns {
		tableTitle = append(tableTitle, templateInfoMap[key])
	}

	for i := range tableTitle {
		fErr := f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", getColumnName(i+1), 1), tableTitle[i])
		if fErr != nil {
			return nil, "", fErr
		}
	}
	f.SetActiveSheet(index)
	file, err = f.WriteToBuffer()
	if err != nil {
		return nil, "", err
	}

	return file, template.Name, nil
}

// ：deleted_at
func (s *SysExportTemplateService) hasDeletedAtColumn(tableName string) bool {
	var count int64
	global.GVA_DB.Raw("SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = ? AND COLUMN_NAME = 'deleted_at'", tableName).Count(&count)
	return count > 0
}

// ImportExcel Excel
// Author [piexlmax](https://github.com/piexlmax)
func (sysExportTemplateService *SysExportTemplateService) ImportExcel(templateID string, file *multipart.FileHeader) (err error) {
	var template system.SysExportTemplate
	err = global.GVA_DB.First(&template, "template_id = ?", templateID).Error
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	f, err := excelize.OpenReader(src)
	if err != nil {
		return err
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return err
	}
	if len(rows) < 2 {
		return errors.New("Excel data is not enough.\nIt should contain title row and data")
	}

	var templateInfoMap = make(map[string]string)
	err = json.Unmarshal([]byte(template.TemplateInfo), &templateInfoMap)
	if err != nil {
		return err
	}

	var titleKeyMap = make(map[string]string)
	for key, title := range templateInfoMap {
		titleKeyMap[title] = key
	}

	db := global.GVA_DB
	if template.DBName != "" {
		db = global.MustGetGlobalDBByDBName(template.DBName)
	}

	return db.Transaction(func(tx *gorm.DB) error {
		excelTitle := rows[0]
		for i, str := range excelTitle {
			excelTitle[i] = strings.TrimSpace(str)
		}
		values := rows[1:]
		items := make([]map[string]interface{}, 0, len(values))
		for _, row := range values {
			var item = make(map[string]interface{})
			for ii, value := range row {
				if _, ok := titleKeyMap[excelTitle[ii]]; !ok {
					continue // excel，，key，
				}
				key := titleKeyMap[excelTitle[ii]]
				item[key] = value
			}

			needCreated := tx.Migrator().HasColumn(template.TableName, "created_at")
			needUpdated := tx.Migrator().HasColumn(template.TableName, "updated_at")

			if item["created_at"] == nil && needCreated {
				item["created_at"] = time.Now()
			}
			if item["updated_at"] == nil && needUpdated {
				item["updated_at"] = time.Now()
			}

			items = append(items, item)
		}
		cErr := tx.Table(template.TableName).CreateInBatches(&items, 1000).Error
		return cErr
	})
}

func getColumnName(n int) string {
	columnName := ""
	for n > 0 {
		n--
		columnName = string(rune('A'+n%26)) + columnName
		n /= 26
	}
	return columnName
}
