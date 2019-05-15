package controllers

import (
	"net/http"
	"strings"
	"time"

	"../models"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

type IndexController struct {
	BaseController
}

func (index *IndexController) Hello(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Info("execute hello function")

	index.sendOk(w, "test")
}

func (index *IndexController) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	r.ParseForm()

	username := strings.Trim(r.Form.Get("username"), " ")
	password := r.Form.Get("password")

	if username == "" {
		index.sendError(w, 101, "用户名不能为空")
		return
	}

	if password == "" {
		index.sendError(w, 102, "请输入密码")
		return
	}

	dbDefaut := models.BaseModel.ConnectDB("default")
	defer dbDefaut.Close()

	dbMarket := models.BaseModel.ConnectDB("market")
	defer dbMarket.Close()

	var banner models.BannerModel
	dbMarket.First(&banner)

	var user models.UserModel

	user.GetFirstByUsername(dbDefaut, username)
	if user.ID == 0 {
		index.sendError(w, 103, "用户名或密码错误")
		return
	}

	if (time.Now().Unix()-int64(user.LastLogin)) < 3600 && user.TryTime > 5 {
		index.sendError(w, 104, "输错密码次数太多，请一小时后再试！")
		return
	}

	if !user.CheckPassword(password) {
		user.TryTime += 1
		user.LastLogin = float64(time.Now().Unix())
		dbDefaut.Save(&user)
		index.sendError(w, 105, "用户名或密码错误")
		return
	}

	user.LastLogin = float64(time.Now().Unix())
	user.TryTime = 0
	user.LastIp = r.RemoteAddr
	dbDefaut.Save(&user)

	type res struct {
		Id       int                `json:"id"`
		Username string             `json:"username"`
		Banner   models.BannerModel `json:"banner"`
	}

	index.sendOk(w, &res{Id: user.ID, Username: user.Username, Banner: banner})
}
