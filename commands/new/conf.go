package commands

var conf = `gormlog: true
app_name: "irisTool"
mysql:
  username: "root"
  password: "root"
  addr: "127.0.0.1:3306"
  name: "test"  #数据库名
redis:
  addr: "127.0.0.1:6379"
  maxActive: 10
  maxIdle: 10
  wait: true
  count: 10 #超时时间（单位是：秒)
  db: 1  #连接数据库的数量
mongodb:
  uri: "mongodb://localhost:27017"
  database: "class" 
  coll: "user" #连接名
log:
  logger_level: DEBUG #
  logger_file: log
  with_max_age: 7*24h #文件最大保存时间
  with_rotation_count: 5 #设置文件清理前最多保存的个数。
  with_rotation_time: 24h #日志切割时间间隔
addr: ":8080"
`
