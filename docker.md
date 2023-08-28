
# docker启动MySQL
docker pull mysql/mysql-server  拉去最新mysql-server适配mac m1版本
1. 启动 docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql/mysql-server
2. -v /Users/byte/Desktop/docker/mysql/data:/var/lib/mysql  -v /Users/byte/Desktop/docker/mysql/log:/logs   -v /Users/byte/Desktop/docker/mysql/conf:/etc/mysql/conf.d
3. 上述 -v 是Mysql数据映射到主机目录上 可以这么理解呗
2. 进入到mysql容器内部 docker exec -it mysql bash
   https://www.jianshu.com/p/eb3d9129d880  参考这个博客 基本上就是没有错误

# docker启动redis
1. 启动redis: docker run -p 6379:6379 --name redis -v /Users/bytedance/Desktop/whig/data/redis/data:/data -v /Users/bytedance/Desktop/whig/data/redis/redis.conf:/etc/redis/redis.conf -d redis redis-server /etc/redis/redis.conf
2. 参考博客 https://blog.csdn.net/u013868195/article/details/119778589