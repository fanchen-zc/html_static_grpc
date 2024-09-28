# golang 静态页生成服务htmlstatic-web-grpc

使用gin框架  调用grpc动态html抓取程序，返回静态页内容， 在本地保存为html

## 本地运行
```bash
go mod init htmlstatic-web-grpc
go mod tidy
```

## 打包与压缩
```bash
git checkout master
git pull origin master
SET CGO_ENABLED=0
SET GOARCH=amd64
SET GOOS=linux
go build -ldflags "-w -s" -o htmlstatic-web-grpc main.go
upx htmlstatic-web-grpc
#可将这些命令保存为build.bat文件, 执行./build.bat即可
```

## 服务器部署
* 主目录上传编译后的htmlstatic-web-grpc 和 config.ini正式服配置文件
* 执行以下命令安装和启动
```bash
#安装
./htmlstatic-web-grpc install
#启动
./htmlstatic-web-grpc start
#重启
./htmlstatic-web-grpc restart
#停止
./htmlstatic-web-grpc stop
#卸载
./htmlstatic-web-grpc uninstall
```
## Git提交规范
```bash
#提交
git add .
git commit -m "Add new feature X"
#打标签
git tag -a v1.0.0 -m "Release version 1.0.0(版本号)"
#推送
git push origin v1.0.0
```

## 版本迭代说明
### [v1.1.1] - 2024/09/01
- README完善