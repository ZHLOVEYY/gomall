package dal

import (
	"github.com/ZHLOVEYY/gomall/app/frontend/biz/dal/mysql"
	"github.com/ZHLOVEYY/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
