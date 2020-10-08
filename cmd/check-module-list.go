package cmd

import (
	"wait4it/http"
	"wait4it/model"
	"wait4it/mongo"
	"wait4it/mysql"
	"wait4it/postgres"
	"wait4it/redis"
	"wait4it/tcp"
)

var cm = map[string]model.Checker{
	"tcp":      &tcp.Tcp{},
	"mysql":    &mysql.MySQLConnection{},
	"postgres": &postgres.PostgresSQLConnection{},
	"http":     &http.HttpCheck{},
	"mongo":    &mongo.MongoDbConnection{},
	"redis":    &redis.RedisConnection{},
}
