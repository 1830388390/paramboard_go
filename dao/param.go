package dao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type Param struct {
	ID          *int64     `json:"id" gorm:"column:id" form:"id"`
	AuthorToken *string    `json:"author_token" gorm:"column:author_token" form:"author_token"`
	TrainNum    *int       `json:"train_num" gorm:"column:train_num" form:"train_num"`
	La1         *float64   `json:"la1" gorm:"column:la1" form:"la1"`
	La2         *float64   `json:"la2" gorm:"column:la2" form:"la2"`
	La3         *float64   `json:"la3" gorm:"column:la3" form:"la3"`
	La4         *float64   `json:"la4" gorm:"column:la4" form:"la4"`
	La5         *float64   `json:"la5" gorm:"column:la5" form:"la5"`
	La6         *float64   `json:"la6" gorm:"column:la6" form:"la6"`
	Lb1         *float64   `json:"lb1" gorm:"column:lb1" form:"lb1"`
	Lb2         *float64   `json:"lb2" gorm:"column:lb2" form:"lb2"`
	Lb3         *float64   `json:"lb3" gorm:"column:lb3" form:"lb3"`
	Lb4         *float64   `json:"lb4" gorm:"column:lb4" form:"lb4"`
	Lb5         *float64   `json:"lb5" gorm:"column:lb5" form:"lb5"`
	Lb6         *float64   `json:"lb6" gorm:"column:lb6" form:"lb6"`
	CreateTime  *time.Time `json:"create_time" gorm:"column:create_time" form:"create_time"`
}

func (m *Param) TableName() string {
	return "t_param"
}

func (m *Param) SetCreateTime(time *time.Time) {
	m.CreateTime = time
	return
}

type ParamDao struct {
	sourceDB  *gorm.DB
	replicaDB []*gorm.DB
	m         *Param
}

func NewParamDaoDefault(ctx context.Context) *ParamDao {
	return CreatParamDao(NewDBClient(ctx))
}

func NewParamDaoByDB(dbs ...*gorm.DB) *ParamDao {
	return CreatParamDao(dbs...)
}

func CreatParamDao(dbs ...*gorm.DB) *ParamDao {
	switch len(dbs) {
	case 0:
		panic("database connection required")
	case 1:
		return &ParamDao{dbs[0], []*gorm.DB{dbs[0]}, &Param{}}
	default:
		return &ParamDao{dbs[0], dbs[1:], &Param{}}
	}
	return &ParamDao{}
}

func (d *ParamDao) CreateOne(ctx context.Context, obj *Param) error {
	err := d.create(ctx, obj)
	//err := d.createRaw(ctx, "insert into t_user(user_name)values(?)",obj.UserName)
	if err != nil {
		return fmt.Errorf("ParamDao: %w", err)
	}
	return nil
}

func (d *ParamDao) createRaw(ctx context.Context, sql string, args ...interface{}) error {
	err := d.sourceDB.Raw(sql, args...).Scan(1).Error
	if err != nil {
		return fmt.Errorf("ParamDao: %w", err)
	}
	return nil
}

func (d *ParamDao) create(ctx context.Context, obj *Param) error {
	err := d.sourceDB.Model(d.m).Create(&obj).Error
	if err != nil {
		return fmt.Errorf("ParamDao: %w", err)
	}
	return nil
}

// k, _ := demo.FindOne(&ctx, "*", "select * from t_demo where id=? and model_name =?", 1, "1") ListRaw用例
// k, _ := demo.FindOne(&ctx, "*", "id=? and model_name=?", 121, "1") List用例
func (d *ParamDao) FindOne(ctx context.Context, fields string, whereOrSql string, args ...interface{}) (obj *Param, err error) {
	objs := []*Param{}
	objs, err = d.listPage(ctx, fields, whereOrSql, 0, 1, args...)
	if err != nil {
		err = fmt.Errorf("ParamDao: Get where=%s: %w", whereOrSql, err)
		return
	}
	if len(objs) == 0 {
		err = gorm.ErrRecordNotFound
		return
	}
	obj = objs[0]
	return
}

func (d *ParamDao) FindAll(ctx context.Context, sql string, args ...interface{}) (objs []*Param, err error) {
	objs, err = d.listRaw(ctx, sql, args...)
	if err != nil {
		err = fmt.Errorf("ParamDao: Get where=%s: %w", sql, err)
		return
	}
	if len(objs) == 0 {
		err = gorm.ErrRecordNotFound
		return
	}
	return
}

func (d *ParamDao) listRaw(ctx context.Context, sql string, args ...interface{}) (objs []*Param, err error) {
	err = d.replicaDB[rand.Intn(len(d.replicaDB))].Raw(sql, args...).Scan(&objs).Error
	if err != nil {
		err = fmt.Errorf("ParamDao: ListRaw %s: %w", sql, err)
		return
	}
	return
}

func (d *ParamDao) listPage(ctx context.Context, fields string, where string, offset int, limit int, args ...interface{}) (objs []*Param, err error) {
	err = d.replicaDB[rand.Intn(len(d.replicaDB))].Model(&Param{}).
		Select(fields).Where(where, args...).Offset(offset).Limit(limit).Find(&objs).Error
	if err != nil {
		err = fmt.Errorf("ParamDao: List where=%s: %w", where, err)
		return
	}
	return
}

func (d *ParamDao) Update(ctx context.Context, where string, update map[string]interface{}, args ...interface{}) error {
	err := d.sourceDB.Model(d.m).Where(where, args...).
		Updates(update).Error
	if err != nil {
		return fmt.Errorf("ParamDao:Update where=%s: %w", where, err)
	}
	return nil
}

// k := demo.Delete(&ctx, "id = ? AND model_name=?", 1, "2")
func (d *ParamDao) Delete(ctx context.Context, where string, args ...interface{}) error {
	if len(where) == 0 {
		return gorm.ErrInvalidData
	}
	if err := d.sourceDB.Where(where, args...).Delete(d.m).Error; err != nil {
		return fmt.Errorf("ParamDao: Delete where=%s: %w", where, err)
	}
	return nil
}
