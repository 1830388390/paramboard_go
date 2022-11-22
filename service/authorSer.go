package service

import (
	"context"
	"paramboard_go/dao"
	e "paramboard_go/enums"
	"paramboard_go/model"
	"paramboard_go/serializer"
	"paramboard_go/utilss"
	"strconv"
	"strings"
	"time"
)

type AuthorSer struct {
	result      string
	CaptchaCode string `json:"captcha" form:"captcha"`
	model.Author
}

func (a *AuthorSer) AddAuthor(ctx context.Context) serializer.Response {
	rd := model.NewRedisClient(ctx)
	//检查验证码
	if a.CaptchaCode != rd.Get(e.CAPTCHA_CODE).Val() {
		return serializer.Response{Code: e.ErrorCaptchaCode, Msg: e.GetMsg(e.ErrorCaptchaCode), Error: e.GetMsg(e.ErrorCaptchaCode)}
	}
	//限制注册数量
	//加锁
	isUnused, err := rd.SetNX(e.Registering, 1, 30*time.Second).Result()
	if !isUnused {
		_ = rd.Del(e.Registering)
		return serializer.Response{Code: e.ErrorCloseRegistration, Msg: e.GetMsg(e.ErrorCloseRegistration), Error: err.Error()}
	}
	//查找剩余注册数量
	t, err := rd.Get(e.CAPTCHA_NUM).Result()
	//t, err := rd.Set(e.CAPTCHA_NUM, "10", -1).Result()
	if err != nil {
		_ = rd.Del(e.Registering)
		return serializer.Response{Code: e.ErrorCloseRegistration,
			Msg: e.GetMsg(e.ErrorCloseRegistration), Error: err.Error(),
		}
	}
	//剩余数量-1
	i, err := strconv.Atoi(t)
	if err != nil || i <= 0 {
		_ = rd.Del(e.Registering)
		return serializer.Response{Code: e.ErrorCloseRegistration, Msg: e.GetMsg(e.ErrorCloseRegistration), Error: err.Error()}
	}
	i = i - 1
	_ = rd.Set(e.CAPTCHA_NUM, i, -1)
	//redis添加author_token
	authorToken := utilss.RandStr(10, utilss.LettersNums)
	_ = rd.Set(e.AUTOHR_TOKEN_PREFIX+authorToken, i, -1)

	//释放锁
	_ = rd.Del(e.Registering)
	authorDao := dao.NewAuthorDaoDefault(ctx)
	author := a.Author
	author.SetAuthorToken(&authorToken)
	err = authorDao.CreateOne(ctx, &author)

	var result strings.Builder
	nowTime := time.Now().String()[0:19]
	result.WriteString("# 结果查看网址( 创建时间: " + nowTime + " 有效期:永久): http://board.权.xyz/#" + authorToken + "<br/>")
	result.WriteString("from urllib import request,parse<br/>")
	result.WriteString("import time<br/>")
	result.WriteString("try:<br/>")
	result.WriteString("    pbe3T='http://board.xn--wqv.xyz/aP?'<br/>")
	result.WriteString("    sjqY9 = {")
	result.WriteString("'author_token': '" + authorToken + "'")
	result.WriteString(", 'train_num': epoch")
	result.WriteString(", 's_t': int(time.time()*1000)")
	if author.LabelA1Name != nil {
		result.WriteString(", 'la1': " + *author.LabelA1Name)
	}
	if author.LabelA2Name != nil {
		result.WriteString(", 'la2': " + *author.LabelA2Name)
	}
	if author.LabelA3Name != nil {
		result.WriteString(", 'la3': " + *author.LabelA3Name)
	}
	if author.LabelA4Name != nil {
		result.WriteString(", 'la4': " + *author.LabelA4Name)
	}
	if author.LabelA5Name != nil {
		result.WriteString(", 'la5': " + *author.LabelA5Name)
	}
	if author.LabelA6Name != nil {
		result.WriteString(", 'la6': " + *author.LabelA6Name)
	}
	if author.LabelB1Name != nil {
		result.WriteString(", 'lb1': " + *author.LabelB1Name)
	}
	if author.LabelB2Name != nil {
		result.WriteString(", 'lb2': " + *author.LabelB2Name)
	}
	if author.LabelB3Name != nil {
		result.WriteString(", 'lb3': " + *author.LabelB3Name)
	}
	if author.LabelB4Name != nil {
		result.WriteString(", 'lb4': " + *author.LabelB4Name)
	}
	if author.LabelB5Name != nil {
		result.WriteString(", 'lb5': " + *author.LabelB5Name)
	}
	if author.LabelB6Name != nil {
		result.WriteString(", 'lb6': " + *author.LabelB6Name)
	}
	result.WriteString("}<br/>")
	result.WriteString("    pbe3T = pbe3T + parse.urlencode(sjqY9)<br/>")
	result.WriteString("    request.urlopen(pbe3T)<br/>")
	result.WriteString("except Exception as E32hK:<br/>")
	result.WriteString("    print(E32hK)<br/>")
	return serializer.Response{Code: e.SUCCESS, Data: result.String(), Msg: e.GetMsg(e.SUCCESS), Error: ""}
}
