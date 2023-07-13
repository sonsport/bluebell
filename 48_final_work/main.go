package main

import (
	"context"
	"fmt"
	"go_web_demo/48_final_work/controller"
	"go_web_demo/48_final_work/dao/mysqlx"
	"go_web_demo/48_final_work/dao/redisx"
	"go_web_demo/48_final_work/logger"
	"go_web_demo/48_final_work/pkg/snowflake"
	"go_web_demo/48_final_work/routers"
	"go_web_demo/48_final_work/settings"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	// 1. viper 读取配置文件
	if err := settings.Init(); err != nil {
		fmt.Println("viper init failed,err:", err)
		return
	}
	// 2. 初始化zap
	logger.Init("dev")
	// 3. 初始化数据库
	if err := mysqlx.Init(); err != nil {
		fmt.Println("init mysqlx failed,err:", err)
		return
	}
	// 4. 初始化redis
	if err := redisx.Init(); err != nil {
		fmt.Println("init redisx failed,err:", err)
		return
	}
	if err := snowflake.Init("2023-04-01", 1); err != nil {
		fmt.Println("init snowflake failed,err:", err)
		return
	}
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Println("init trans failed,err:", err)
		return
	}
	// 5. 路由设置
	r := routers.Setup("release")
	// 6. 优雅关机

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.ViperConfig.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: %s\n", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}
