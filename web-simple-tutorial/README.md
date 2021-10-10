golang的简单web demo
# 构建docker镜像
```
docker build -t web-app:1.0 .
```
# docker运行镜像
```
docker run -d -p 8080:8080 web-app:1.0
```
# 参考
https://dev.to/koddr/let-s-write-config-for-your-golang-web-app-on-right-way-yaml-5ggp