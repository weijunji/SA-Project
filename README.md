# SA-Project
[![Go Report Card](https://goreportcard.com/badge/github.com/weijunji/SA-Project)](https://goreportcard.com/report/github.com/weijunji/SA-Project)
[![Build Status](https://travis-ci.org/weijunji/SA-Project.svg?branch=main)](https://travis-ci.org/weijunji/SA-Project)
[![codecov](https://codecov.io/gh/weijunji/SA-Project/branch/main/graph/badge.svg)](https://codecov.io/gh/weijunji/SA-Project)

Project of Software Architecture

## 任务进度
### 服务器
- [x] 客户端接入 验证客户端并修改数据库中的数据
- [x] 心跳 接收基站控制器心跳数据信息，确保其是否在线
- [x] 状态接受 接受client上传的状态信息并保存
- [x] 控制命令传输 接受web服务器的请求，将操作分发到对应的客户端上
- [x] 客户端状态 记录客户端状态并保存在数据库中

### 分布式机器
- [x] 控制命令 基于标准协议接收服务器发送的指令，实现基本控制
- [x] 状态上传 基于标准协议实现各种机器状态的上传
- [ ] 远程配置 实现远端对IP地址、以及各种管理数据的配置
- [x] 心跳 实现对服务器的心跳信息控制

### web服务器
基于go-gin框架
- [ ] 浏览器信息推送 有状态变化时向浏览器发送状态改变的信息
- [x] 控制命令传输 接受前端请求并发送控制命令到服务器上
- [x] 机器状态发送 从数据库中读取机器状态并返回给前端
- [ ] 管理配置 接收请求并转发到服务器上
- [x] 数据查询 从数据库中查询对应数据并返回前端

### web前端
基于vue、vuetify等构建，前后端分离
- [ ] 浏览器界面 采用JavaScript实现富客户端界面，如站点树、圆饼图、菜单等
- [ ] 控制命令传输 基于标准Http协议，实现控制命令发送，如开锁、上锁、停机
- [ ] 机器状态接收 基于标准Http协议，实现门禁状态信息的接收，如在线离线、锁的开启关闭
- [ ] 管理配置 实现同步信息的更新频率等各种设置信息的远程配置
- [ ] 数据查询 实现各种状态数据的查询汇总等
- [ ] 报表 实现以EXCEL、PDF等格式保存打印报表等
