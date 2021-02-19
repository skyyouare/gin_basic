# docker上线步骤

#1、部署相关

##1.1、测试环境：

①、瓦力后置执行脚本：
```shell
chmod 777 /home/www/gin_basic/test.sh
/home/www/gin_basic/test.sh >/tmp/gin_basic.log 2>&1
```
②test.sh主要内容
```shell
#!/bin/bash
echo 'start'
st=$(date +%Y-%m-%d\ %H:%M:%S)
echo "$st"
cd /home/www/gin_basic || return
#build
docker-compose -f docker-compose-test.yml build
#停止容器 有bug
#docker-compose -f docker-compose-test.yml down
#移除当前容器
for a in $(docker ps -a |grep gin_basic | awk '{print $1}')
do
        docker stop "$a"
        docker rm "$a"
done
#启动容器
docker-compose -f docker-compose-test.yml up -d
et=$(date +%Y-%m-%d\ %H:%M:%S)
echo "$et"
echo 'end'
```
##1.2、生成环境：

①瓦力后置执行脚本：
```shell
chmod 777 /home/www/gin_basic/prod.sh
/home/www/gin_basic/prod.sh >/tmp/gin_basic.log 2>&1
```   

②prod.sh主要内容
```shell
#!/bin/bash
echo 'start'
st=$(date +%Y-%m-%d\ %H:%M:%S)
echo "$st"
cd /home/www/gin_basic || return
#build
docker-compose -f docker-compose-prod.yml build
#down 有bug
#docker-compose -f docker-compose-prod.yml down
#移除当前容器
for a in $(docker ps -a |grep gin_basic | awk '{print $1}')
do
        docker stop "$a"
        docker rm "$a"
done
#启动容器
docker-compose -f docker-compose-prod.yml up -d
et=$(date +%Y-%m-%d\ %H:%M:%S)
echo "$et"
echo 'end'
```



#2、golang镜像相关 --- 内部docker-hub

##2.1、打标签:
```shell
docker tag golang:1.14.1 docker-hub.shiqutech.com/oa/golang:1.14.1
```

##2.2、推到仓库:
```shell
docker push docker-hub.shiqutech.com/oa/golang:1.14.1
```

##2.3、仓库查看地址:
```shell
docker-hub.shiqutech.com
```

#3、开发相关

##3.1、开发工具

```shell
推荐使用goland
```
    

##3.2、开发环境
开发机：192.168.8.32

goland sftp部署上传如 /home/doujinya/www/gin_basic

执行：

```shell
cd /home/doujinya/www/gin_basic
go run main.go dev
```

访问:192.168.8.32:9090

代码修改完后，需要ctrl+c退出，并且重新执行

```shell
go run main.go dev
```

#4、框架相关

##4.1、配置文件 redis暂不可用

开发环境配置；conf/dev下 app.toml等

测试环境配置: conf/test下 app.toml等

生成环节配置：conf/prod下 prod.toml等

##3.4、路由

routes/route.go中

```go
router.GET("/", func(c *gin.Context) {
c.String(http.StatusOK, "Welcome Gin Server5")
})
```
定义入口路口
访问地址：ip:9090
```go
v1 := router.Group("/v1")

v1.Use(
//middleware.RecoveryMiddleware(),
)
{
    controller.ReportRegister(v1)
}
```
中 controller.ReportRegister(v1)调用controller/ReportController

```go
func ReportRegister(router *gin.RouterGroup) {
    controller := new(ReportController)
    router.POST("/report/index", controller.index)
}
```
方法注册路由

访问地址：ip:9090/v1//report/index

##3.5、控制器
参考 controller/ReportController





