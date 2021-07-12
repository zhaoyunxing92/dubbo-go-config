# dubbo-go-config

dubbo-go 配置优化样例代码

## 配置使用
```go
config := config.Load(
		WithGenre("yaml"),
		WithPrefix("dubbo"),
		WithPath("../conf/yaml"),
		WithName("application.yaml"),
	)

application, _ := config.GetApplicationConfig()

registries, _ := config.GetRegistriesConfig()
```
### 没有配置`registries`情况下给默认

```yaml
dubbo:
  application:
    name: dubbo-go
    module: local
    version: 1.0.0
    owner: zhaoyunxing
  services:
    hello-service:
      interface: org.dubbo.service.HelloService
      registry: nacos,zk
    order-service:
      interface: org.dubbo.service.OrderService
      registry: nacos
```