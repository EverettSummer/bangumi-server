// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"github.com/bangumi/server/internal/dal/dao"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newIndexSubject(db *gorm.DB) indexSubject {
	_indexSubject := indexSubject{}

	_indexSubject.indexSubjectDo.UseDB(db)
	_indexSubject.indexSubjectDo.UseModel(&dao.IndexSubject{})

	tableName := _indexSubject.indexSubjectDo.TableName()
	_indexSubject.ALL = field.NewField(tableName, "*")
	_indexSubject.ID = field.NewUint32(tableName, "idx_rlt_id")
	_indexSubject.Cat = field.NewInt8(tableName, "idx_rlt_cat")
	_indexSubject.Rid = field.NewUint32(tableName, "idx_rlt_rid")
	_indexSubject.Type = field.NewUint8(tableName, "idx_rlt_type")
	_indexSubject.Sid = field.NewUint32(tableName, "idx_rlt_sid")
	_indexSubject.Order = field.NewUint32(tableName, "idx_rlt_order")
	_indexSubject.Comment = field.NewString(tableName, "idx_rlt_comment")
	_indexSubject.Dateline = field.NewInt32(tableName, "idx_rlt_dateline")
	_indexSubject.Subject = indexSubjectBelongsToSubject{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Subject", "dao.Subject"),
		Fields: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Subject.Fields", "dao.SubjectField"),
		},
	}

	_indexSubject.fillFieldMap()

	return _indexSubject
}

type indexSubject struct {
	indexSubjectDo indexSubjectDo

	ALL      field.Field
	ID       field.Uint32
	Cat      field.Int8
	Rid      field.Uint32
	Type     field.Uint8
	Sid      field.Uint32
	Order    field.Uint32
	Comment  field.String
	Dateline field.Int32
	Subject  indexSubjectBelongsToSubject

	fieldMap map[string]field.Expr
}

func (i indexSubject) Table(newTableName string) *indexSubject {
	i.indexSubjectDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i indexSubject) As(alias string) *indexSubject {
	i.indexSubjectDo.DO = *(i.indexSubjectDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *indexSubject) updateTableName(table string) *indexSubject {
	i.ALL = field.NewField(table, "*")
	i.ID = field.NewUint32(table, "idx_rlt_id")
	i.Cat = field.NewInt8(table, "idx_rlt_cat")
	i.Rid = field.NewUint32(table, "idx_rlt_rid")
	i.Type = field.NewUint8(table, "idx_rlt_type")
	i.Sid = field.NewUint32(table, "idx_rlt_sid")
	i.Order = field.NewUint32(table, "idx_rlt_order")
	i.Comment = field.NewString(table, "idx_rlt_comment")
	i.Dateline = field.NewInt32(table, "idx_rlt_dateline")

	i.fillFieldMap()

	return i
}

func (i *indexSubject) WithContext(ctx context.Context) *indexSubjectDo {
	return i.indexSubjectDo.WithContext(ctx)
}

func (i indexSubject) TableName() string { return i.indexSubjectDo.TableName() }

func (i *indexSubject) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *indexSubject) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 9)
	i.fieldMap["idx_rlt_id"] = i.ID
	i.fieldMap["idx_rlt_cat"] = i.Cat
	i.fieldMap["idx_rlt_rid"] = i.Rid
	i.fieldMap["idx_rlt_type"] = i.Type
	i.fieldMap["idx_rlt_sid"] = i.Sid
	i.fieldMap["idx_rlt_order"] = i.Order
	i.fieldMap["idx_rlt_comment"] = i.Comment
	i.fieldMap["idx_rlt_dateline"] = i.Dateline

}

func (i indexSubject) clone(db *gorm.DB) indexSubject {
	i.indexSubjectDo.ReplaceDB(db)
	return i
}

type indexSubjectBelongsToSubject struct {
	db *gorm.DB

	field.RelationField

	Fields struct {
		field.RelationField
	}
}

func (a indexSubjectBelongsToSubject) Where(conds ...field.Expr) *indexSubjectBelongsToSubject {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a indexSubjectBelongsToSubject) WithContext(ctx context.Context) *indexSubjectBelongsToSubject {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a indexSubjectBelongsToSubject) Model(m *dao.IndexSubject) *indexSubjectBelongsToSubjectTx {
	return &indexSubjectBelongsToSubjectTx{a.db.Model(m).Association(a.Name())}
}

type indexSubjectBelongsToSubjectTx struct{ tx *gorm.Association }

func (a indexSubjectBelongsToSubjectTx) Find() (result *dao.Subject, err error) {
	return result, a.tx.Find(&result)
}

func (a indexSubjectBelongsToSubjectTx) Append(values ...*dao.Subject) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a indexSubjectBelongsToSubjectTx) Replace(values ...*dao.Subject) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a indexSubjectBelongsToSubjectTx) Delete(values ...*dao.Subject) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a indexSubjectBelongsToSubjectTx) Clear() error {
	return a.tx.Clear()
}

func (a indexSubjectBelongsToSubjectTx) Count() int64 {
	return a.tx.Count()
}

type indexSubjectDo struct{ gen.DO }

func (i indexSubjectDo) Debug() *indexSubjectDo {
	return i.withDO(i.DO.Debug())
}

func (i indexSubjectDo) WithContext(ctx context.Context) *indexSubjectDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i indexSubjectDo) Clauses(conds ...clause.Expression) *indexSubjectDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i indexSubjectDo) Not(conds ...gen.Condition) *indexSubjectDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i indexSubjectDo) Or(conds ...gen.Condition) *indexSubjectDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i indexSubjectDo) Select(conds ...field.Expr) *indexSubjectDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i indexSubjectDo) Where(conds ...gen.Condition) *indexSubjectDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i indexSubjectDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *indexSubjectDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i indexSubjectDo) Order(conds ...field.Expr) *indexSubjectDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i indexSubjectDo) Distinct(cols ...field.Expr) *indexSubjectDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i indexSubjectDo) Omit(cols ...field.Expr) *indexSubjectDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i indexSubjectDo) Join(table schema.Tabler, on ...field.Expr) *indexSubjectDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i indexSubjectDo) LeftJoin(table schema.Tabler, on ...field.Expr) *indexSubjectDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i indexSubjectDo) RightJoin(table schema.Tabler, on ...field.Expr) *indexSubjectDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i indexSubjectDo) Group(cols ...field.Expr) *indexSubjectDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i indexSubjectDo) Having(conds ...gen.Condition) *indexSubjectDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i indexSubjectDo) Limit(limit int) *indexSubjectDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i indexSubjectDo) Offset(offset int) *indexSubjectDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i indexSubjectDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *indexSubjectDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i indexSubjectDo) Unscoped() *indexSubjectDo {
	return i.withDO(i.DO.Unscoped())
}

func (i indexSubjectDo) Create(values ...*dao.IndexSubject) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i indexSubjectDo) CreateInBatches(values []*dao.IndexSubject, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i indexSubjectDo) Save(values ...*dao.IndexSubject) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i indexSubjectDo) First() (*dao.IndexSubject, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.IndexSubject), nil
	}
}

func (i indexSubjectDo) Take() (*dao.IndexSubject, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.IndexSubject), nil
	}
}

func (i indexSubjectDo) Last() (*dao.IndexSubject, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.IndexSubject), nil
	}
}

func (i indexSubjectDo) Find() ([]*dao.IndexSubject, error) {
	result, err := i.DO.Find()
	return result.([]*dao.IndexSubject), err
}

func (i indexSubjectDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.IndexSubject, err error) {
	buf := make([]*dao.IndexSubject, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i indexSubjectDo) FindInBatches(result *[]*dao.IndexSubject, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i indexSubjectDo) Attrs(attrs ...field.AssignExpr) *indexSubjectDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i indexSubjectDo) Assign(attrs ...field.AssignExpr) *indexSubjectDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i indexSubjectDo) Joins(field field.RelationField) *indexSubjectDo {
	return i.withDO(i.DO.Joins(field))
}

func (i indexSubjectDo) Preload(field field.RelationField) *indexSubjectDo {
	return i.withDO(i.DO.Preload(field))
}

func (i indexSubjectDo) FirstOrInit() (*dao.IndexSubject, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.IndexSubject), nil
	}
}

func (i indexSubjectDo) FirstOrCreate() (*dao.IndexSubject, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.IndexSubject), nil
	}
}

func (i indexSubjectDo) FindByPage(offset int, limit int) (result []*dao.IndexSubject, count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	if limit <= 0 {
		return
	}

	result, err = i.Offset(offset).Limit(limit).Find()
	return
}

func (i indexSubjectDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i *indexSubjectDo) withDO(do gen.Dao) *indexSubjectDo {
	i.DO = *do.(*gen.DO)
	return i
}