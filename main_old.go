package main

import (
	"github.com/gin-gonic/gin"
	"paramboard_go/conf"
	"paramboard_go/dao"
	"paramboard_go/model"
	"time"
)

func main() {
	// Ek1+Ep1==Ek2+Ep2
	conf.Init()
	//r := routes.NewRouter()
	//_=r.Run(conf.HttpPort)
	ctx := gin.Context{}
	//cartDao := dao.NewCartDao(&ctx)
	//cart, _ := cartDao.ListCartByUserId(1)
	//print(cart[0].BossID)
	authorDao := dao.NewAuthorDaoDefault(&ctx)
	//k, _ := authorDao.FindOne(&ctx, "*", "id=? and model_name=?", 124, "2")
	//k, _ := authorDao.FindOne(&ctx, "*", "select * from t_author where id=? and model_name =?", 121, "1") 查找
	num := 1
	str := "2"
	ti := time.Now()
	author := model.Author{
		ID:          int64(num),
		AuthorToken: &str,
		ModelName:   &str,
		Describe:    &str,
		LabelA1Name: &str,
		LabelA2Name: &str,
		LabelA3Name: &str,
		LabelA4Name: &str,
		LabelA5Name: &str,
		LabelA6Name: &str,
		LabelB1Name: &str,
		LabelB2Name: &str,
		LabelB3Name: &str,

		CreateTime: &ti,
	}
	k := authorDao.CreateOne(&ctx, &author)
	//k := authorDao.Delete(&ctx, "id = ? AND model_name=?", 1, "2")
	//db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;
	//
	//db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	//
	//dbDriver.Where(strings.Join([]string{"id = ?", "model_name = ?"}, " AND "), []interface{}{121, "1"}).First(t)

	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;
	print(k)

}
