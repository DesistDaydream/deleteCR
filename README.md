# 弃用
改名为 [k8s-resource-operate-job](https://github.com/DesistDaydream/k8s-resource-operate-job)，添加更多功能

# deleteCR
在 helm uninstall 时，先删除 CR 再删除 operator 的工具

## 构建
docker build --tag lchdzh/deletecr:v0.3 .

## 测试
deletecr --ns=rabbitmq --name=rabbitmq \
--crgroup=rabbitmq.com \
--crversion=v1beta1 \
--crname=rabbitmqclusters

helm template test . -s templates/rabbitmqcluster/job-delete-cr.yaml  | kubectl apply -f -