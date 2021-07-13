# dubbo-go-config

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

### 命令行参数

默认解析yaml文件

> json文件解析

```shell
➜  samples git:(main) ✗ ./main --path=../conf/json --genre=json --name=application.json
start load config file
application name: dubbo-go
```

> toml 文件解析

```shell
➜  samples git:(main) ✗ ./main --path=../conf/toml --genre=toml --name=application.toml
start load config file
application name: dubbo-go
```

> yaml 文件解析

```shell
➜  samples git:(main) ✗ ./main --path=../conf                                          
start load config file
application name: dubbo-go
```

### 参数说明

* `WithGenre`: 文件类型yaml、json、toml 反正就是[viper](https://github.com/spf13/viper) 支持的类型
* `WithCache`: 是否需要缓存配置到本地
* `WithPath`: 配置文件地址
* `WithName`: 配置文件名称

### 没有配置`registries`情况下给默认`zk`

> 最终目标就这点配置

```yaml
dubbo:
  application:
    name: dubbo-go
    module: local
    version: 1.0.0
    owner: zhaoyunxing
```