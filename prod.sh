#!/bin/bash
echo 'start'
st=$(date +%Y-%m-%d\ %H:%M:%S)
echo "$st"
cd /home/www/gin_basic || return
# build
docker-compose -f docker-compose-prod.yml build
# down 有bug
# docker-compose -f docker-compose-prod.yml down
# 移除当前容器
for a in $(docker ps -a |grep gin_basic | awk '{print $1}')
do
        docker stop "$a"
        docker rm "$a"
done
# 启动容器
docker-compose -f docker-compose-prod.yml up -d
et=$(date +%Y-%m-%d\ %H:%M:%S)
echo "$et"
echo 'end'