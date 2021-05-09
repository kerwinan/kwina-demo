# kwina-demo
个人golang项目模板示例。

### TODO

1. 解析 config文件改善，解析指定目录下的所有yaml文件
2. 自己封装一套rpc服务注册与发现工具包，register, call, logout。通过idl定义服务并进行服务方法注册，不使用pb定义
3. 通过项目模板快速构建新项目

### 项目结构介绍

| 顶层目录/文件 | 内容说明                             |
| ------------- | ------------------------------------ |
| bootstrap/    | 启动过程                             |
| client/       | 调用客户端, 如rpc客户端              |
| config/       | 配置文件                             |
| dao/          | 数据库操作，如mysql，sds，tablestore |
| idl/          | 内部idl，不向其他程序提供调用；      |
| internal/     | 内部常量与结构                       |
| model/        | 为service/层提供查询业务             |
| option/       | 解析配置文件                         |
| provider/     | 启动服务                             |
| service/      | 提供的服务                           |

