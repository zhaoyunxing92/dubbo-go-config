# dubbo-go-config

dubbo-go 配置优化样例代码

考虑到目前用户配置dubbo-go文件很繁琐，且没有层级，最终的的目标是简化到`spring boot`的那种程度

## 配置使用

```go
config := config.Load(
    WithGenre("yaml"),
    WithCache(false),
    WithPath("../conf/yaml"),
    WithName("application.yaml"),
)

application, _ := config.GetApplicationConfig()

registries, _ := config.GetRegistriesConfig()
```

### 已经支持命令行参数

### 参数说明

* `WithGenre`: 文件类型yaml、json、toml 反正就是[viper](https://github.com/spf13/viper) 支持的类型
* `WithCache`: 是否需要缓存配置到本地
* `WithPath`: 配置文件地址
* `WithName`: 配置文件名称

### 没有配置`registries`情况下给默认`zk`

```yaml
dubbo:
  application:
    name: dubbo-go
    module: local
    version: 1.0.0
    owner: zhaoyunxing
```