package service

import (
	"context"
	"paramboard_go/dao"
	e "paramboard_go/enums"
	"paramboard_go/model"
	"paramboard_go/serializer"
	"strconv"
	"time"
)

type ParamSer struct {
	dao.Param
	*model.Author
	params      []*dao.Param `json:"params" form:"params"`
	St          int64        `json:"s_t" form:"s_t"`
	AuthorToken string       `json:"author_token" gorm:"column:author_token" form:"author_token"`
}

//    AuthorParamVO findByAuthorToken(String author_token);
//
//    int addParam(Param param,Long s_t);
//
//    int delParam(String author_token);

func (p *ParamSer) FindByAuthorToken(ctx context.Context) serializer.Response {
	paramDao := dao.NewParamDaoDefault(ctx)
	authorDao := dao.NewAuthorDaoDefault(ctx)
	var err error
	p.params, err = paramDao.FindAll(ctx, "select * from t_param where author_token=?", p.AuthorToken)
	p.Author, err = authorDao.FindOne(ctx, "*", "author_token=?", p.AuthorToken)
	if err != nil {
		return serializer.Response{Code: e.AuthorTokenFail, Msg: e.GetMsg(e.AuthorTokenFail), Error: err.Error()}
	}
	return serializer.Response{Code: e.SUCCESS, Data: map[string]interface{}{"params": p.params, "author": p.Author}, Msg: e.GetMsg(e.SUCCESS)}
}

func (p *ParamSer) AddParam(ctx context.Context) serializer.Response {

	rd := model.NewRedisClient(ctx)
	_, err := rd.Get(e.AUTOHR_TOKEN_PREFIX + p.AuthorToken).Result()
	if err != nil {
		return serializer.Response{Code: e.AuthorTokenFail, Msg: e.GetMsg(e.AuthorTokenFail), Error: err.Error()}
	}
	isUnused, err := rd.SetNX(e.PARAM_PREFIX+p.AuthorToken+strconv.FormatInt(p.St, 10), 1, 24*time.Hour).Result()
	if !isUnused {
		return serializer.Response{Code: e.ErrorCrowding, Msg: e.GetMsg(e.ErrorCrowding), Error: err.Error()}
	}
	Stime := time.Unix(p.St/1000, 0)
	p.SetCreateTime(&Stime)
	paramDao := dao.NewParamDaoDefault(ctx)
	err = paramDao.CreateOne(ctx, &p.Param)
	if err != nil {
		return serializer.Response{Code: e.ERROR, Msg: e.GetMsg(e.ERROR), Error: err.Error()}
	}
	return serializer.Response{Code: e.SUCCESS, Msg: e.GetMsg(e.SUCCESS), Error: ""}
}

func (p *ParamSer) DelParam(ctx context.Context) serializer.Response {
	rd := model.NewRedisClient(ctx)
	isUnused, err := rd.SetNX(e.PARAM_PREFIX+p.AuthorToken+strconv.FormatInt(p.St, 10), 1, 24*time.Hour).Result()
	if !isUnused {
		return serializer.Response{Code: e.ErrorCrowding, Msg: e.GetMsg(e.ErrorCrowding), Error: err.Error()}
	}
	paramDao := dao.NewParamDaoDefault(ctx)
	err = paramDao.Delete(ctx, "author_token=?", p.AuthorToken)
	if err != nil {
		return serializer.Response{Code: e.AuthorTokenFail, Msg: e.GetMsg(e.AuthorTokenFail), Error: err.Error()}
	}
	return serializer.Response{Code: e.SUCCESS, Msg: e.GetMsg(e.SUCCESS), Error: err.Error()}
}
