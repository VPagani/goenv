Environment variables loading utility

# Usage

### **`env.go`**

```go
var env = goenv.SetupEnv(".env")

var (
    AppEnv  = env("APP_ENV", "development")
    AppPort = env("APP_PORT", env("PORT"))
    AppAddr = env("APP_ADDR", ":"+AppPort)
)
```