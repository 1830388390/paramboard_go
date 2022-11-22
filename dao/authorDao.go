package dao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"paramboard_go/model"
)

type AuthorDao struct {
	sourceDB  *gorm.DB
	replicaDB []*gorm.DB
	m         *model.Author
}

func NewAuthorDaoDefault(ctx context.Context) *AuthorDao {
	return CreatAuthorDao(NewDBClient(ctx))
}

func NewAuthorDaoByDB(dbs ...*gorm.DB) *AuthorDao {
	return CreatAuthorDao(dbs...)
}

func CreatAuthorDao(dbs ...*gorm.DB) *AuthorDao {
	switch len(dbs) {
	case 0:
		panic("database connection required")
	case 1:
		return &AuthorDao{dbs[0], []*gorm.DB{dbs[0]}, &model.Author{}}
	default:
		return &AuthorDao{dbs[0], dbs[1:], &model.Author{}}
	}
	return &AuthorDao{}
}

func (d *AuthorDao) CreateOne(ctx context.Context, obj *model.Author) error {
	//err := d.Create(ctx, obj)
	err := d.createRaw(ctx, "insert into t_author(author_token,model_name,`describe`,label_a1_name,label_a2_name,label_a3_name,label_a4_name,label_a5_name,label_a6_name,label_b1_name,label_b2_name,label_b3_name)values(?,?,?,?,?,?,?,?,?,?,?,?)",
		obj.AuthorToken, obj.ModelName, obj.Describe, obj.LabelA1Name, obj.LabelA2Name, obj.LabelA3Name, obj.LabelA4Name, obj.LabelA5Name, obj.LabelA6Name, obj.LabelB1Name, obj.LabelB2Name, obj.LabelB3Name)
	if err != nil {
		return fmt.Errorf("AuthorDao: %w", err)
	}
	return nil
}

func (d *AuthorDao) createRaw(ctx context.Context, sql string, args ...interface{}) error {
	err := d.sourceDB.Raw(sql, args...).Scan(1).Error
	if err != nil {
		return fmt.Errorf("AuthorDao: %w", err)
	}
	return nil
}

func (d *AuthorDao) create(ctx context.Context, obj *model.Author) error {
	err := d.sourceDB.Model(d.m).Create(&obj).Error
	if err != nil {
		return fmt.Errorf("AuthorDao: %w", err)
	}
	return nil
}

func (d *AuthorDao) FindAll(ctx context.Context, sql string, args ...interface{}) (objs []*model.Author, err error) {
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

// k, _ := demo.FindOne(&ctx, "*", "select * from t_demo where id=? and model_name =?", 1, "1") ListRaw用例
// k, _ := demo.FindOne(&ctx, "*", "id=? and model_name=?", 121, "1") List用例
func (d *AuthorDao) FindOne(ctx context.Context, fields string, whereOrSql string, args ...interface{}) (author *model.Author, err error) {
	authors := []*model.Author{}
	//authors, err = d.listRaw(ctx, whereOrSql, args...)
	authors, err = d.list(ctx, fields, whereOrSql, 0, 1, args...)
	if err != nil {
		err = fmt.Errorf("AuthorDao: Get where=%s: %w", whereOrSql, err)
		return
	}
	if len(authors) == 0 {
		err = gorm.ErrRecordNotFound
		return
	}
	author = authors[0]
	return
}

func (d *AuthorDao) listRaw(ctx context.Context, sql string, args ...interface{}) (authors []*model.Author, err error) {
	err = d.replicaDB[rand.Intn(len(d.replicaDB))].Raw(sql, args...).Scan(&authors).Error
	if err != nil {
		err = fmt.Errorf("AuthorDao: ListRaw %s: %w", sql, err)
		return
	}
	return
}

func (d *AuthorDao) list(ctx context.Context, fields string, where string, offset int, limit int, args ...interface{}) (authors []*model.Author, err error) {
	err = d.replicaDB[rand.Intn(len(d.replicaDB))].Model(&model.Author{}).
		Select(fields).Where(where, args...).Offset(offset).Limit(limit).Find(&authors).Error
	if err != nil {
		err = fmt.Errorf("AuthorDao: List where=%s: %w", where, err)
		return
	}
	return
}

func (d *AuthorDao) Update(ctx context.Context, where string, update map[string]interface{}, args ...interface{}) (err error) {
	err = d.sourceDB.Model(d.m).Where(where, args...).
		Updates(update).Error
	if err != nil {
		err = fmt.Errorf("AuthorDao:Update where=%s: %w", where, err)
		return
	}
	return
}

// Delete
// k := demo.Delete(&ctx, "id = ? AND model_name=?", 1, "2")
func (d *AuthorDao) Delete(ctx context.Context, where string, args ...interface{}) (err error) {
	if len(where) == 0 {
		err = gorm.ErrInvalidData
		return
	}
	if err := d.sourceDB.Where(where, args...).Delete(d.m).Error; err != nil {
		return fmt.Errorf("AuthorDao: Delete where=%s: %w", where, err)
	}
	return
}
