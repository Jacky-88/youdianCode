# KubeA平台部署



## 一、环境要求

- golang 1.16以上版本（helm sdk不支持更低版本）
- vue/cli 4.5的版本（不能使用5的版本）
- K8s 1.20以上 (版本不要差太多就行)



## 二、后端配置

### 1. 配置路径

- `k8s-demo/config/config.go`

### 2. 核心配置

#### 2.1 监听地址

- 服务端监听配置：

  ```go
  ListenAddr = "0.0.0.0:9090"
  ```

- websocket监听配置（用于webshell终端交互）：

  ```go
  WsAddr = "0.0.0.0:8081"
  ```

#### 2.2 K8S集群

- 支持跨集群，map格式，key为集群名（会在前端展示），value为config地址

  ```go
  //多集群
  Kubeconfigs = `{"TST-1":"/Users/adoo/.kube/config","TST-2":"/Users/adoo/.kube/config1"}`
  //单集群
  Kubeconfigs = `{"TST-1":"/Users/adoo/.kube/config"}`
  ```

#### 2.3 登录配置

- 管理员账号密码

  ```go
  AdminUser = "admin"
  AdminPwd = "123456"
  ```

#### 2.4 Helm Chart配置

- 上传路径

  ```go
  UploadPath = "/Users/adoo/chart"
  ```

#### 2.5 数据库配置

- 本平台会**自动**生成两张表，`k8s_event`和`helm_chart`

- 分别用于存储事件数据和chart信息

- 在启动程序前，先创建名为`k8s_demo`的数据库

  ```go
  DbType = "mysql"
  DbHost = "192.168.1.11"
  DbPort = 3306
  DbName = "k8s_demo"
  DbUser = "root"
  DbPwd = "Abc123@@@"
  ```

#### 2.6 Event监听任务开启

- k8s-demo/main.go

- 默认代码已被注销，不开启

- 支持跨集群监听，参数一定要与config中的集群名对齐

  ```go
  go func() {
  		service.Event.WatchEventTask("TST-1")
  	}()
  go func() {
      service.Event.WatchEventTask("TST-2")
  }()
  ```

#### 2.7 开启jwt token校验

- k8s-demo/main.go

- 默认关闭

- 若开启，不影响前端请求，但是使用postman或curl调试需要添加请求头`Authorization`

  ```go
  r.Use(middle.JWTAuth())
  ```

  

## 三、前端配置

### 1. 配置路径

- `kubea-fe/src/config/index.js`

### 2. 设置请求的后端Host

- 本地部署，则不用修改

- 在远程服务器上部署，则需要修改为真正的后端地址

  ```js
  const baseHost = 'http://localhost:9090'
  ```

  

## 四、本地部署

### 1. 前端

- 进入项目根目录

  ```shell
  cd kubea-fe
  ```

- 安装依赖，可提前配置npm淘宝源

  ```shell
  npm install
  ```

- 启动程序

  ```shell
  npm run serve
  ```

- 打开浏览器，输入：

  ```shell
  http://localhost:8080
  ```

- 默认账号密码：

  ```
  账号：Admin
  密码：123456
  ```

  

### 2. 后端

- 更改配置文件，详见第二步（后端配置）

- 进入项目根目录

  ```shell
  cd k8s-demo
  ```

- 安装依赖

  ```shell
  go mod tidy
  ```

- 运行程序

  ```shell
  go run main.go
  ```

- 测试接口

  ```shell
  adoodeMacBook-Pro:.kube adoo$ curl --location --request GET 'http://0.0.0.0:9090/api/k8s/clusters'
  {"data":["TST-1","TST-2"],"msg":"获取集群信息成功"}
  ```

  

## 五、远程服务器部署

### 1. 部署信息

- 服务器系统：Centos7.6

|      | 服务器       | 端口      |
| ---- | ------------ | --------- |
| 前端 | 192.168.1.11 | 8888      |
| 后端 | 192.168.1.11 | 8999/8082 |

### 2. 前端部署

#### 2.1 打包项目（本地操作）

- 修改后端host配置, `kubea-fe/src/config/index.js`中的baseHost和websocket地址

  ```js
  const baseHost = 'http://localhost:8999'
  k8sTerminalWs: 'ws://localhost:8082/ws'
  ```

- 编译

  ```shell
  cd kubea-fe/
  npm run build
  ```

- 生成`dist`目录，将目录压缩放到`192.168.1.11`服务器的`/data/`下,解压

#### 2.2 登录服务器

```shell
ssh root@192.168.1.11
```

#### 2.3 安装Nginx

- 安装并启动

  ```shell
  yum install -y nginx
  systemctl start nginx
  ```

- 修改nginx配置文件，打开`/etc/nginx/conf.d/kubea-fe.conf`

  ```shell
  server {
      listen    8888;
      server_name _;
      access_log /data/nginx_access.log;
      location / {
          root /data/dist;
          index index.html index.htm;
          try_files $uri $uri/ /index.html;
      }
  }
  ```

- Reload Nginx

  ```
  systemctl reload nginx
  ```

  

### 3. 后端部署

#### 3.1 打包项目（本地操作）

- 更改配置文件，详见第二步（后端配置）

- 进入项目根目录

  ```shell
  cd k8s-demo/
  ```

- 跨平台编译为linux二进制文件

  ```shell
  #windows编译linux
  SET CGO_ENABLED=0
  SET GOOS=linux
  SET GOARCH=amd64
  go build
  #mac编译linux
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
  ```

- 生成二进制文件`main`，可随意改名字
- 将二进制文件放到`192.168.1.11`服务器的`/data/`下

#### 3.2 登录服务器

```shell
ssh root@192.168.1.11
```

### 3.3 启动后端程序

- 进入`/data`目录下

  ```shell
  #后台执行，日志会打印在nohup.out
  nohup main &
  ```

### 3.4 配置Nginx

- 修改nginx配置文件，打开`/etc/nginx/conf.d/k8s-demo.conf`

  ```shell
  server {
      listen    8999;
      server_name _;
      access_log /data/nginx_access.log;
      location / {
          proxy_pass: 0.0.0.0:9090; #这里是后端程序的监听端口
      }
  }
  
  server {
      listen    8082;
      server_name _;
      access_log /data/nginx_access.log;
      location / {
          proxy_pass: 0.0.0.0:8081; #这里是后端程序的websocket端口
          #websocket代理我还没试过，若上面配置异常就取消下面注释
          #proxy_http_version 1.1;
          #proxy_set_header Upgrade $http_upgrade;
          #proxy_set_header Connection "$connection_upgrade";
      }
  }
  ```

- Reload Nginx

  ```
  systemctl reload nginx
  ```

### 4. 验证

- 浏览器打开`http://192.168.1.11:8888`



## 六、功能使用注意事项

### 1. Helm应用市场

- 进入`http://192.168.1.11:8888/helmstore/app`

- 点击 “新增”，填写新增信息，其中：

  - Icon图标：仅支持网络图片，需要填写url，可参考一下url：

    ```shell
    my-oss-testing.oss-cn-beijing.aliyuncs.com/course/gotrain/assets/app-icon/tomcat.svg
    my-oss-testing.oss-cn-beijing.aliyuncs.com/course/gotrain/assets/app-icon/rabbitmq.svg
    my-oss-testing.oss-cn-beijing.aliyuncs.com/course/gotrain/assets/app-icon/nacos.jpeg
    my-oss-testing.oss-cn-beijing.aliyuncs.com/course/gotrain/assets/app-icon/logstash.svg
    my-oss-testing.oss-cn-beijing.aliyuncs.com/course/gotrain/assets/app-icon/kibana.svg
    my-oss-testing.oss-cn-beijing.aliyuncs.com/course/gotrain/assets/app-icon/kafka.svg
    my-oss-testing.oss-cn-beijing.aliyuncs.com/course/gotrain/assets/app-icon/jenkins.svg
    my-oss-testing.oss-cn-beijing.aliyuncs.com/course/gotrain/assets/app-icon/istio.png
    my-oss-testing.oss-cn-beijing.aliyuncs.com/course/gotrain/assets/app-icon/etcd.svg
    ```

  - 上传Chart：仅支持上传一个文件，`.tgz`结尾的chart包，会上传到后端配置中的`UploadPath`中

  - ps : 前期可以随便那个tgz文件上传

### 2. Event事件统计

- 概览中event表格在后端不开启以下代码时，无数据

- `k8s-demo/main.go`

  ```go
  //event任务,用于监听event并写入数据库,这里的传参是集群名，一定要与config中的集群名对齐
  //go func() {
  //	service.Event.WatchEventTask("TST-1")
  //}()
  //go func() {
  //	service.Event.WatchEventTask("TST-2")
  //}()
  ```

  