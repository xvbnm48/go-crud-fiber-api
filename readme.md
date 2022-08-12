# go rest api example with golang , gorm, and fiber


in fiber can use like gin.H for return message with object, like this

this gin example
```go
package bookController

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "pong",
    })
}

```

this is fiber example
```go
package bookController
import "github.com/gofiber/fiber"

func Index(c *fiber.Ctx) err {
    c.JSON(200, fiber.Map{
        "message": "pong",
    })
    return nil
} 
```