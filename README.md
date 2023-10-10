## 目录介绍
├─build  编译文件
│      .gitignore
│      .gitkeep
│      
├─cmd  启动目录
│      app.go 
│      wire.go
│      wire_gen.go
│      
├─conf 配置文件
│  ├─dev
│  │      config.yaml
│  │      dubbogo.yaml
│  │      
│  ├─prod
│  └─test
├─internal  内部代码
│  ├─controller
│  │      video.go
│  │      
│  ├─dao  数据操作
│  ├─handler
│  │      default_handler.go
│  │      
│  ├─logic 逻辑处理
│  │      video.go
│  │      
│  └─model 模板定义
├─logs  日志目录
│      .gitignore
│      .gitkeep
│      
├─pkg 核心包
│  ├─amqp amqp
│  │      rabbitmq.go
│  │      
│  ├─cache 缓存
│  │      redis.go
│  │      
│  ├─config  配置
│  │      config.go
│  │      
│  ├─db 数据库
│  │      mysql.go
│  │      [DockerBuild.bat](..%2F..%2F..%2Fproject%2Fjava%2Fdubbo-chain%2Fapproval-provider%2FDockerBuild.bat)
│  ├─logger 日志
│  │      zap_logger.go
│  │      
│  ├─msg 消息
│  │      msg.go
│  │      
│  ├─service 服务
│  │      init.go
│  │      
│  ├─snowf 雪花算法
│  │      snowflake.go
│  │      
│  ├─tracing 链路追踪
│  │      jaeger.go
│  │      
│  └─utils 工具库
│          http.go
│          msg.go
│          time.go
│          
└─runtime
│  .gitignore
│  .gitkeep
│  
└─logs
