package mqserv

import (
	"os"
	"runtime"
	"strings"

	"github.com/jsuserapp/ju"
	"github.com/kardianos/service"
)

const (
	serviceName = "mqttserv"
)

type program struct {
	serv    service.Service
	logFile ju.LogDb
}

func (p *program) Start(serv service.Service) error {
	p.serv = serv
	go p.run()
	return nil
}
func (p *program) Stop(_ service.Service) error {
	ju.CloseFileLogDb(p.logFile)
	return nil
}

// 这里是由 Start 里的 go p.run() 调用的
func (p *program) run() {
	//很奇怪的一个问题, 就是在 Linux 服务里, 这个函数不会阻塞执行, 而是会返回, 但是程序功能正常
	//在windows下就会阻塞, Linux 普通应用里也是阻塞执行的.
	if !startMQTTServ() {
		if service.Interactive() {
			// 在调试模式下，直接以错误码 1 退出程序
			os.Exit(1)
		} else {
			err := p.serv.Stop()
			if err != nil {
				ju.LogRed(serviceName, "stoped error", err.Error())
			} else {
				ju.LogRed(serviceName, "stoped")
			}
		}
	}
}

// getServicePath 获取服务的可执行文件路径
func getServicePath(name string) string {
	switch runtime.GOOS {
	case "windows":
		return "."
	}
	fn := "/etc/systemd/system/" + name + ".service"
	data := ju.ReadFile(fn)
	if data == nil {
		return ""
	}
	str := string(data)
	flag := "ExecStart="
	pos := strings.Index(str, flag)
	if pos == -1 {
		return ""
	}
	str = str[pos+len(flag):]
	pos = strings.Index(str, "\n")
	if pos != -1 {
		str = str[:pos]
	}
	pos = strings.LastIndex(str, "/")
	if pos != -1 {
		str = str[:pos]
	}
	ju.LogYellow(str)
	return str
}

func RunService() {
	fileDb := ju.CreateFileLogDb("", 0, 0)
	ju.SetLogParam(fileDb, true, 0, 0)
	absPath := getServicePath(serviceName)
	//服务的配置信息
	cfg := &service.Config{
		Name:             serviceName,
		DisplayName:      "http or other service",
		Description:      "This is an service for jsuse",
		WorkingDirectory: absPath,
	}
	ju.LogGreen(serviceName, "start success")
	// Interface 接口
	prg := &program{
		logFile: fileDb,
	}
	// 构建服务对象
	s, err := service.New(prg, cfg)
	if ju.LogFail(err) {
		return
	}

	errs := make(chan error, 5)
	go func() {
		for {
			err = <-errs
			ju.LogError(err)
		}
	}()

	if len(os.Args) == 2 {
		//如果有命令则执行
		err = service.Control(s, os.Args[1])
		ju.LogError(err)
	} else {
		// 这里并不是调用的 program 的 run, 而是服务的 run, 它在内部会启动服务, 执行 Start
		err = s.Run()
		ju.LogError(err)
	}
}
