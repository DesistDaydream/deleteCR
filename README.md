# deleteCR
在 helm uninstall 时，先删除 CR 再删除 operator 的工具

## 构建
docker build . -tag lchdzh/deletecr:v0.2

## 测试
deletecr --ns=rabbitmq --name=rabbitmq \
--crgroup=rabbitmq.com \
--crversion=v1beta1 \
--crname=rabbitmqcluster

helm template test . -s templates/rabbitmqcluster/job-delete-cr.yaml  | kubectl apply -f -