package mmsg

import (
	"fmt"
	"github.com/github.com/mabetle/mcore"
	"github.com/github.com/mabetle/mcore/mmodel"
)

type ModelLabel struct {
	Field string
	Label string
}

// GetModelLabels
// returns model all labels
func GetModelLabels(locale string, model interface{}) (rows []ModelLabel) {
	fs := mcore.GetFields(model)
	for _, f := range fs {
		item := ModelLabel{}
		item.Field = f
		item.Label = GetModelFieldLabel(locale, model, f)
		rows = append(rows, item)
	}
	return
}

// GetModelFieldLabel
// returns model locale filed name message. Parameters:
// locale 
// model can be a slice? no. do as name.
// field
// Field name save as column name.
// table name should be follow some rules.
func GetModelFieldLabel(locale string, model interface{}, field string) string {
	//table := mcore.GetTypeName(model)
	table := mmodel.GetModelTableName(model)
	column:= field		
	return GetTableColumnLabel(locale,table,column)
}


// GetTableColumnLabel
// returns table column locale label message.
// if table is "", set table to common
// key format: common-UserName etc.
func GetTableColumnLabel(locale, table, column string) string{
	
	value := ""
	
	if table == ""{
		table = "common"	
	} 

	key := fmt.Sprintf("%s-%s-label", table, column)

	if Contains(locale, key) {
		logger.Debugf("Found model field label. Model:%s, Field:%s Key: %s", table, column, key)
		value = Message(locale, key)
	}

	// if not found, name to label name.
	if value == "" {
		logger.Warnf("Not found model field label. Model:%s, Field:%s", table, column)
		value = mcore.ToLabel(column)
	}

	return value
}
