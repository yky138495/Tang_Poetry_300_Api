# Tang_Poetry_300_Api

唐诗300首+API，基于beego实现的Restful api

## go 环境安装
```
brew install go
```

## IDE 
```
可以使用intelliJ作为go的IDE:http://www.jianshu.com/p/943870134593
也可以使用Atom作用go的IDE:http://www.jianshu.com/p/c1d8cf274ec7
```

查看go是否安装成功：
```
go version
```

## 配置Go环境变量
如果.bash_profile不存在，则执行vi语句：
```
vi ~/.bash_profile
```

在 .bash_profile 中加入如下（下面1.11.4为go的版本号）：
```
export GOPATH=/usr/local/Cellar/go/1.12.4
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
export BeeBIN=/Users/***/go/bin
export PATH=$PATH:$BeeBIN
```

让配置生效
```
source ~/.bash_profile
```

查看配置结果：
```
go env
```

eg：新建一个main.go文件
```
package main
import ("fmt")
func main() {
fmt.Println("hello world!")
}
```

build：
```
go build main.go
```

run：
```
go run main.go
```

## go 安装sqllite支持

```
go get github.com/mattn/go-sqlite3
go get github.com/astaxie/beego
go get github.com/beego/bee
```

## 安装后引用三方 （相对路径）

```
_ "./github.com/mattn/go-sqlite3"
```



# Beego

## 生成项目
```
bee api snmpcheck
```
创建一个api项目，生成的目录结构为
```
├── conf
│   └── app.conf
├── controllers
│   ├── object.go
│   └── user.go
├── main.go
├── models
│   ├── object.go
│   └── user.go
├── routers
│   └── router.go
└── tests
└── default_test.go
```

beego生成的model文件中
```
bee generate model user -fields="name:string,age:int"
```

运行
```
bee run
##-gendoc=true表示每次会自动化的build文档，-downdoc=true表示会自动下载swagger文档查看
bee run -gendoc=true -downdoc=true 
```

打包
```
bee pack
修改app.conf：runmode=prod
```

浏览器访问：
```
http://localhost:8080/swagger/
```

### Beego 架构
Beego架构图如(Beegode 八大独立的模块):

```
997599-20180321233004672-238888081.png
```
Beego的执行逻辑，beego 是一个典型的 MVC 架构
Beego架构图如(Beegode 八大独立的模块):
```
997599-20180321233021296-480972597.png
```
