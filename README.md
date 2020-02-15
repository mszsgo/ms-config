# ms-config

配置文件管理

提供websocket接口，实时更新配置

1. 监听mongodb的ms_config集合，有新增、更新、删除，推送
2. 提供ws服务，接收服务名，响应配置信息
 map[name]wsconn[]



轮询获取配置信息










