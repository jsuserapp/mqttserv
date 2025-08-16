# mqttserv
这是一个旧版的mqtt服务器，因为其特殊性，某些客户端设备是基于其实现的，不兼容主流的mqtt服务器，故保留

运行下面的命令安装服务
```shell
./mqttserv install
#成功 install 之后，服务就会启动
# install 不能重复执行，重新安装需要先 uninstall
./mqttserv start
#已经运行时重复 start 操作不会报错，系统会自动忽略
./mqttserv restart
./mqttserv stop
./mqttserv uninstall

```
`/etc/systemd/system/mqttserv.service`文件是服务的配置文件，其中 `RestartSec=120` 指示服务异常退出后的重启时间，重新配置这个值只能通过手动修改。