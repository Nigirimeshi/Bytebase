package mysql

// Framework code is generated by the generator.

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pkg/errors"

	mysql "github.com/bytebase/mysql-parser"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	mysqlparser "github.com/bytebase/bytebase/backend/plugin/parser/mysql"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*TableNoDuplicateIndexAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MYSQL, advisor.MySQLTableNoDuplicateIndex, &TableNoDuplicateIndexAdvisor{})
}

// TableNoDuplicateIndexAdvisor is the advisor checking for no duplicate index in table.
type TableNoDuplicateIndexAdvisor struct {
}

// Check checks for no duplicate index in table.
func (*TableNoDuplicateIndexAdvisor) Check(ctx advisor.Context, _ string) ([]advisor.Advice, error) {
	stmtList, ok := ctx.AST.([]*mysqlparser.ParseResult)
	if !ok {
		return nil, errors.Errorf("failed to convert to mysql parser result")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}

	checker := &tableNoDuplicateIndexChecker{
		level:     level,
		title:     string(ctx.Rule.Type),
		indexList: []duplicateIndex{},
	}
	for _, stmt := range stmtList {
		checker.baseLine = stmt.BaseLine
		antlr.ParseTreeWalkerDefault.Walk(checker, stmt.Tree)
	}

	if len(checker.adviceList) == 0 {
		checker.adviceList = append(checker.adviceList, advisor.Advice{
			Status:  advisor.Success,
			Code:    advisor.Ok,
			Title:   "OK",
			Content: "",
		})
	}
	return checker.adviceList, nil
}

type tableNoDuplicateIndexChecker struct {
	*mysql.BaseMySQLParserListener

	baseLine   int
	adviceList []advisor.Advice
	level      advisor.Status
	title      string
	indexList  []duplicateIndex
}

type duplicateIndex struct {
	indexName string
	// BTREE, HASH, etc.
	indexType string
	unique    bool
	fulltext  bool
	spatial   bool
	table     string
	columns   []string
	// line is the line number of the index definition.
	line int
}

func (checker *tableNoDuplicateIndexChecker) EnterCreateTable(ctx *mysql.CreateTableContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	if ctx.TableName() == nil {
		return
	}
	if ctx.TableElementList() == nil {
		return
	}

	_, tableName := mysqlparser.NormalizeMySQLTableName(ctx.TableName())
	// Build suspect index list.
	for _, tableElement := range ctx.TableElementList().AllTableElement() {
		if tableElement == nil || tableElement.TableConstraintDef() == nil {
			continue
		}
		checker.handleConstraintDef(tableName, tableElement.TableConstraintDef())
	}
	// Check for duplicate index.
	if index := hasDuplicateIndexes(checker.indexList); index != nil {
		checker.adviceList = append(checker.adviceList, advisor.Advice{
			Status:  checker.level,
			Code:    advisor.DuplicateIndexInTable,
			Title:   checker.title,
			Content: fmt.Sprintf("`%s` has duplicate index `%s`", tableName, index.indexName),
			Line:    index.line,
		})
	}
}

func (checker *tableNoDuplicateIndexChecker) EnterAlterTable(ctx *mysql.AlterTableContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	if ctx.TableRef() == nil {
		return
	}
	if ctx.AlterTableActions() == nil || ctx.AlterTableActions().AlterCommandList() == nil || ctx.AlterTableActions().AlterCommandList().AlterList() == nil {
		return
	}

	_, tableName := mysqlparser.NormalizeMySQLTableRef(ctx.TableRef())
	for _, alterListItem := range ctx.AlterTableActions().AlterCommandList().AlterList().AllAlterListItem() {
		if alterListItem == nil {
			continue
		}
		if alterListItem.ADD_SYMBOL() != nil && alterListItem.TableConstraintDef() != nil {
			checker.handleConstraintDef(tableName, alterListItem.TableConstraintDef())
		}
	}
	if index := hasDuplicateIndexes(checker.indexList); index != nil {
		checker.adviceList = append(checker.adviceList, advisor.Advice{
			Status:  checker.level,
			Code:    advisor.DuplicateIndexInTable,
			Title:   checker.title,
			Content: fmt.Sprintf("`%s` has duplicate index `%s`", tableName, index.indexName),
			Line:    index.line,
		})
	}
}

func (checker *tableNoDuplicateIndexChecker) handleConstraintDef(tableName string, ctx mysql.ITableConstraintDefContext) {
	switch ctx.GetType_().GetTokenType() {
	case mysql.MySQLParserINDEX_SYMBOL, mysql.MySQLParserKEY_SYMBOL, mysql.MySQLParserPRIMARY_SYMBOL, mysql.MySQLParserUNIQUE_SYMBOL, mysql.MySQLParserFULLTEXT_SYMBOL, mysql.MySQLParserSPATIAL_SYMBOL:
	default:
		return
	}

	index := duplicateIndex{
		indexType: "BTREE",
		line:      checker.baseLine + ctx.GetStart().GetLine(),
		table:     tableName,
	}
	if ctx.KeyListVariants() != nil {
		index.columns = mysqlparser.NormalizeKeyListVariants(ctx.KeyListVariants())
	}
	if ctx.UNIQUE_SYMBOL() != nil {
		index.unique = true
	} else if ctx.FULLTEXT_SYMBOL() != nil {
		index.fulltext = true
	} else if ctx.SPATIAL_SYMBOL() != nil {
		index.spatial = true
	}
	if ctx.IndexNameAndType() != nil {
		if ctx.IndexNameAndType().IndexName() != nil {
			index.indexName = mysqlparser.NormalizeIndexName(ctx.IndexNameAndType().IndexName())
		}
		if ctx.IndexNameAndType().IndexType() != nil {
			index.indexType = ctx.IndexNameAndType().IndexType().GetText()
		}
	}
	checker.indexList = append(checker.indexList, index)
}

func (checker *tableNoDuplicateIndexChecker) EnterCreateIndex(ctx *mysql.CreateIndexContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	if ctx.CreateIndexTarget() == nil || ctx.CreateIndexTarget().TableRef() == nil || ctx.CreateIndexTarget().KeyListVariants() == nil {
		return
	}

	_, tableName := mysqlparser.NormalizeMySQLTableRef(ctx.CreateIndexTarget().TableRef())
	index := duplicateIndex{
		indexType: "BTREE",
		line:      checker.baseLine + ctx.GetStart().GetLine(),
		table:     tableName,
	}
	if ctx.UNIQUE_SYMBOL() != nil {
		index.unique = true
	} else if ctx.FULLTEXT_SYMBOL() != nil {
		index.fulltext = true
	} else if ctx.SPATIAL_SYMBOL() != nil {
		index.spatial = true
	}
	if ctx.IndexName() != nil {
		index.indexName = mysqlparser.NormalizeIndexName(ctx.IndexName())
	}
	if ctx.IndexNameAndType() != nil {
		if ctx.IndexNameAndType().IndexName() != nil {
			index.indexName = mysqlparser.NormalizeIndexName(ctx.IndexNameAndType().IndexName())
		}
		if ctx.IndexNameAndType().IndexType() != nil {
			index.indexType = ctx.IndexNameAndType().IndexType().GetText()
		}
	}
	if ctx.IndexTypeClause() != nil && ctx.IndexTypeClause().IndexType() != nil {
		index.indexType = ctx.IndexTypeClause().IndexType().GetText()
	}

	index.columns = mysqlparser.NormalizeKeyListVariants(ctx.CreateIndexTarget().KeyListVariants())
	checker.indexList = append(checker.indexList, index)
	if index := hasDuplicateIndexes(checker.indexList); index != nil {
		checker.adviceList = append(checker.adviceList, advisor.Advice{
			Status:  checker.level,
			Code:    advisor.DuplicateIndexInTable,
			Title:   checker.title,
			Content: fmt.Sprintf("`%s` has duplicate index `%s`", tableName, index.indexName),
			Line:    index.line,
		})
	}
}

// hasDuplicateIndexes returns the first duplicate index if found, otherwise nil.
func hasDuplicateIndexes(indexList []duplicateIndex) *duplicateIndex {
	seen := make(map[string]struct{})
	for _, index := range indexList {
		key := indexKey(index)
		if _, exists := seen[key]; exists {
			return &index
		}
		seen[key] = struct{}{}
	}
	return nil
}

// indexKey returns a string key for the index with the index type and columns.
func indexKey(index duplicateIndex) string {
	parts := []string{}
	if index.unique {
		parts = append(parts, "unique")
	}
	if index.fulltext {
		parts = append(parts, "fulltext")
	}
	if index.spatial {
		parts = append(parts, "spatial")
	}
	parts = append(parts, index.indexType)
	parts = append(parts, index.table)
	parts = append(parts, index.columns...)
	return strings.Join(parts, "-")
}
