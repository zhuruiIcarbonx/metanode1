我来展示如何使用 Gin 框架实现这个简单的 RESTful API：

## 1. 基础版本（最简洁）

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    
    r.Run() // 默认在 0.0.0.0:8080 启动
}
```

## 2. 使用结构体定义响应（推荐）

```go
package main

import "github.com/gin-gonic/gin"

// 定义响应结构体
type PingResponse struct {
    Message string `json:"message"`
}

func main() {
    r := gin.Default()
    
    r.GET("/ping", func(c *gin.Context) {
        response := PingResponse{
            Message: "pong",
        }
        c.JSON(200, response)
    })
    
    r.Run(":8080")
}
```

## 3. 完整版本（包含配置和健康检查）

```go
package main

import (
    "log"
    "time"
    
    "github.com/gin-gonic/gin"
)

type PingResponse struct {
    Message   string    `json:"message"`
    Timestamp time.Time `json:"timestamp"`
    Status    string    `json:"status"`
}

func main() {
    // 设置 Gin 模式（ReleaseMode 或 DebugMode）
    // gin.SetMode(gin.ReleaseMode)
    
    r := gin.Default()
    
    // 添加全局中间件（可选）
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    
    // 定义路由
    r.GET("/ping", pingHandler)
    r.GET("/", rootHandler)
    
    // 启动服务器
    log.Println("🚀 Server starting on http://localhost:8080")
    log.Println("📡 Try: curl http://localhost:8080/ping")
    r.Run(":8080")
}

func pingHandler(c *gin.Context) {
    response := PingResponse{
        Message:   "pong",
        Timestamp: time.Now(),
        Status:    "success",
    }
    c.JSON(200, response)
}

func rootHandler(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Welcome to Gin Server",
        "endpoints": gin.H{
            "ping": "GET /ping",
        },
    })
}
```

## 4. 带分组路由的版本

```go
package main

import "github.com/gin-gonic/gin"

type PingResponse struct {
    Message string `json:"message"`
}

func main() {
    r := gin.Default()
    
    // API 路由分组
    api := r.Group("/api")
    {
        api.GET("/ping", pingHandler)
        api.GET("/health", healthHandler)
    }
    
    r.Run(":8080")
}

func pingHandler(c *gin.Context) {
    response := PingResponse{Message: "pong"}
    c.JSON(200, response)
}

func healthHandler(c *gin.Context) {
    c.JSON(200, gin.H{
        "status": "healthy",
    })
}
```

## 运行步骤

1. **初始化项目**（如果还没有 go.mod 文件）：
```bash
go mod init myapp
```

2. **安装 Gin**：
```bash
go get -u github.com/gin-gonic/gin
```

3. **创建 main.go 文件**，复制上面的代码

4. **运行程序**：
```bash
go run main.go
```

5. **测试 API**：
```bash
curl http://localhost:8080/ping
```

预期输出：
```json
{"message":"pong"}
```

## 测试其他方式

**使用浏览器**：
- 直接访问 `http://localhost:8080/ping`

**使用其他工具**：
```bash
# 使用 httpie
http GET http://localhost:8080/ping

# 使用 wget
wget -q -O - http://localhost:8080/ping
```

## 项目结构示例

```
myapp/
├── go.mod
├── go.sum
└── main.go
```

## Gin 的优势

相比标准库，Gin 提供了：
- ✅ 更简洁的路由定义
- ✅ 自动的 JSON 序列化
- ✅ 内置中间件支持
- ✅ 更好的性能
- ✅ 丰富的上下文方法
- ✅ 参数绑定和验证

**推荐使用版本2（结构体版本）**，因为它代码清晰且易于扩展。当需要添加更多字段时，只需要修改结构体即可。
