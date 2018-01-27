从零开始写一个运行在Kubernetes上的服务程序
=====================================

## 参考资料
http://dockone.io/article/2980
https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/

## 备忘
1、通过make push命令推送docker镜像前需要登录hub.docker.com
docker login -u [用户名] -p [密码]
2、结束时退出登录
docker logout
3、启动minikube时使用代理
minikube start --docker-env HTTP_PROXY=http://192.168.99.1:1087 --docker-env HTTPS_PROXY=http://192.168.99.1:1087

