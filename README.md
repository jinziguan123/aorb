# 👐 Welcome to Aorb

## 💖 简介

- 这里是一个社交应用的demo，我们在这个项目中探索最佳实践
- 项目的主要目标是提供一个社交平台，用户可以在这里发布自己的动态，参与投票，评论等
- 与此同时，在客户端上将尝试一些动画的制作

## 🔨 技术栈

- 使用Flutter进行前端开发
- 使用gRPC进行微服务之间的通信
- 使用Consul进行服务注册和发现
- 使用RabbitMQ进行消息队列
- 使用Redis进行缓存
- 使用MongoDB进行数据存储
- 使用Kubernetes进行容器编排
- 使用Prometheus+Grafana进行监控

## 📋 项目结构
```
aorb
├── backend
│   ├── api-gateway
│   ├── go-services
│   │   ├── auth
│   │   ├── comment
│   │   ├── event
│   │   ├── poll
│   │   ├── user
│   │   └── vote
│   ├── rpc
│   └── utils
├── build
├── frontend
│   ├── lib
│   │   ├── conf
│   │   ├── generated
│   │   ├── models
│   │   ├── routes
│   │   ├── screens
│   │   ├── services
│   │   ├── utils
│   │   └── widgets
│   └── test
├── idl
├── k8s
├── logs
├── monitoring
│   ├── grafana
│   └── prometheus
└── scripts
```


## 🚀 快速开始

推荐使用vscode进行开发，安装flutter插件，以及dart插件

### 将项目克隆到本地

```bash
git clone https://github.com/BigNoseCattyHome/aorb.git
```

### 需要安装的工具

在这个项目中需要用到的工具有：

- flutter
- go
- protoc
- consul
- rabbitMQ
- redis
- mongodb



在项目拉取到本地之后需要执行`make proto`生成前后端中所需要的一些代码

### 前端开发 

开发和测试flutter应用，在项目根目录下执行

```shell
make run frontend
```

或者是尝试进入到frontend目录下执行

```shell
flutter run
```

Flutter 会自动编译 `fronted/lib/main.dart` 文件并运行，选择一个合适的平台进行查看就好，不同平台需要满足特定的工具包。


figma原型设计共享链接：[Aorb原型设计](https://www.figma.com/design/roDqwgrlbQo29vpSqeCVFw/Aorb?node-id=0-1&t=SOBamnPsEXegjKDF-1)

### 数据库初始化

这里是一篇[MongoDB安装和简单上手](https://obyi4vacom.feishu.cn/file/DTTWb1DMjoGynkxmgOBc0qgInWd)文档，可以参考一下

确保在本地安装好MongoDB后，进行数据库初始化：

```shell    
mongosh
```

进入到mongodb shell之后输入命令：
```shell
load("scripts/init_db.js")
```

### 后台各个服务的开启

RabbitMQ:
```shell
systemctl start rabbitmq-server     # Linux
brew services start rabbitmq        # MacOS
```

Consul:
```shell
consul agent -dev
```

Redis:
```shell
redis-server
```

### 微服务的启动

可以执行以下命令来启动后端各个微服务

```shell
make run backend
```

### 客户端启动

执行以下命令来启动客户端，因为项目中运用了 gRPC 进行通讯，浏览器目前不支持

推荐使用各个平台的客户端，比如Linux、Windows、MacOS等

```shell
make run frontend
```

#### 安卓设备

对于在手机上进行真机测试，需要手机打开开发者模式，并且使用USB连接到电脑上，并将连接方式设置为文件传输。

执行以下命令检查是否连接成功

```shell
adb devices
```

如果连接成功，执行以下命令启动客户端

```shell
make run frontend   # 在项目根目录下
flutter run         # 在frontend目录下
```

## 📝 开发文档

[Flutter开发过程用到组件指南](https://obyi4vacom.feishu.cn/file/E9vdbu0RBocg4yxfV0NcS1kHnwe)

[Git使用指南](http://sirius1y.top/posts/notes/dev/%E6%8C%87%E5%8D%97%E5%9B%A2%E9%98%9Fgit%E5%8D%8F%E4%BD%9C/)

[开发踩坑记录](http://sirius1y.top/posts/notes/dev/dev-aorb-grpc/)

## 踩坑记录补充

1. 使用Apifox测试的时候返回了```invalid wire type[13 INTERNAL]```错误
- 原因：本质上是因为客户端(Apifox)与服务端(项目后端)所使用的pb类型定义不一致
- 解决方法：检查后端的proto文件，并且重新上传到Apifox，参考链接是[这篇博客](https://loesspie.com/2021/09/14/grpc-did-not-read-entire-message/)
2. Consul报错：```too many colons in address```
- 原因：grpc的包里面没有针对consul的解析器，无法讲请求解析到正确的微服务端口
- 解决方法：在每个微服务的main.go中引入```import _ "github.com/mbobakov/grpc-consul-resolver"```，[参考链接](https://blog.csdn.net/dorlolo/article/details/123416857)

