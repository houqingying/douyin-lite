# douyin-lite

使用gorm + gin + mysql + redis + rabbitmq 编写的极简版抖音，实现基础接口，社交接口和互动接口。

视频展示网站：https://www.bilibili.com/video/BV1334y1T73z/?share_source=copy_web&vd_source=21cac58a94f918e40674f147fb0b2fc2

项目文档：https://l0wshqly3u6.feishu.cn/docx/JMRfdJyRkoDelaxOLiJcC0qcnQf?from=from_copylink

## 技术栈

- Gin：Http框架，提供路由和 HTTP 服务；
- Gorm：ORM框架，用面向对象的方式对数据库进行 增删查找 操作；
- MySQL: 关系数据库，存储数据；
- Redis：存储热点数据。对关注/取关操作进行缓存，按照一定策略使键过期，并定时同步数据到数据库；
- RabbitMQ：对 Redis 的异步操作、流量削峰；
- Viper：用文件的方式存储MySQL, Redis的配置；
- SnowFlake: 唯一分布式ID生成算法；
- JWT：用户鉴权，token 的生成与校验；
- ffmpeg：获取封面截图
- sha1：将密码加密，以密文的形式存入数据库。
