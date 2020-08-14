#### 编译
```bash
# 使用 go.mod

# 安装go依赖包
go list (go mod tidy)

# 编译
go build

#### Run
```bash
skaffold dev
```


#### 目录结构

```
    ├─server  	     （后端文件夹）
    │  ├─api            （API）
    │  ├─config         （配置包）
    │  ├─pkg            （业务代码）
    │      ├─core  	        （內核）
    │      ├─db             （数据库脚本）
    │      ├─router         （路由）
    │      ├─service         (服务)
    │      ├─middleware     （中间件）
    │      ├─global         （全局对象）
    │      ├─model          （结构体层）
    │  ├─docs  	        （swagger文档目录）
    │  ├─initialiaze    （初始化）
    │  ├─resource       （资源）
    │  ├─log            （日志）
    │  └─tools	        （公共功能）
    

```

## Dockerfile

## k8s-pod.yaml

## Makefile

## skaffold.yaml

