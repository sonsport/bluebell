package routers

import (
	"go_web_demo/48_final_work/controller"
	"go_web_demo/48_final_work/middlewares"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Setup(mode string) *gin.Engine {
	gin.SetMode(mode)
	r := gin.New()
	r.Use(GinLogger(), GinRecovery(true))
	r.LoadHTMLGlob("./templates/*")
	r.Static("/static", "./static")
	// 映射静态文件
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//路由组
	v1 := r.Group("/api/v1")
	//注册&
	v1.POST("/signup", controller.SignUpHandler)
	//登录&
	v1.POST("/login", controller.LoginHandler)
	// 列表显示所有社区&
	v1.GET("/community", controller.CommunityHandler)
	// 显示单个社区
	v1.GET("/community/:id", controller.CommunityDetailHandler)
	// 显示指定哪个帖子&
	v1.GET("/post/:id", controller.GetDetailPost)
	// 显示所有帖子&
	v1.GET("/post", controller.GetPost)
	// 临时2&
	//v1.GET("/posts2", controller.GetPost)
	// 获取排序帖子&
	// .GET posts2
	// 投票&
	// .POST vote

	// 分类 帖子排序
	//v1.GET("/posts2", controller.GetPostOrder)
	// 分类 帖子排序 + 得到的赞
	//v1.GET("/getpostorder2", controller.GetPostOrder2)
	//v1.GET("/getcommunitypostorder", controller.GetCommunityPostOrder)

	// 使用认证服务
	v1.Use(middlewares.JWTAuthMiddleware())

	// 获取信息
	v1.GET("/getInfo", controller.GetDetail)

	// 发表帖子&
	v1.POST("/post2w", controller.CreatePostHandler)

	// 投票
	//v1.POST("/vote", controller.PostVote)

	return r
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
