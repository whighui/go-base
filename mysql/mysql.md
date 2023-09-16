

# count(1)和count(*)和count(列)有什么区别
count(*)==count(1)==>count(主键值)>count(字段)
https://xiaolincoding.com/mysql/index/count.html

count(主键id) 
1. 如果表中只有主键id，innodb会遍历聚族索引，并将id返回给server层并判断是否为Null
2. 如果表中有一个二级索引，即便是count(主键) 那么也会是遍历二级索引B+树，因为二级索引比聚族索引记录主键ID占用更少的存储空间，二级索引比聚族索引占用空间小，所以IO成本更低一些
3. 但是表中如果存在多个二级索引，那么会选择二级索引中长度最短的那个索引来查询

count(1)
1. 基本上同count(主键)是一致的，但是会少一个步骤，那么就是不会判断主键是否为Null  所以相比效率会高一些

count(*)
1. count(*)其实基本上属于count(0)  select count(*) as count(0) from user_info 
2. count(*)和count(1)执行上基本上不会存在差异

count(字段)
1. 如果该字段不是索引的话，那么基本上就是全表查询 此效率是最慢的

如何优化count(*)
假如数据量上千万，即便存在二级索引查询耗时也是比较长的 差不多在5秒索引，需要看二级索引的长度
explain关键字，并不会正常的去查询而是进行估算 例如 explain select count(*) from user 

# limit越往后查越慢 是什么原因导致的
https://cloud.tencent.com/developer/article/1888828 

LIMIT 关键字可以指定查询结果从哪条记录开始显示，显示多少条记录。例如limit[M,N] 重记录M+1开始返回N条记录<br>

select * from table_name limit 10000,10 当offset偏移量大之后 Limit查询会变慢
1. 重表中读取第N条数据也即就是10到数据集当中
2. 重复第一步操作 直到取到第M+N 条数到数据集当中
3. 根据offset丢弃前M条数据
4. 返回剩余十条数据 <br>
会什么SQL会变,因为前1000条数据是无意义的<br>
**第一次优化：假如主键ID连续**<br>
如果根据主键ID查询 select * from user where id>=1000 limit 10  但是在实际场所当中主键ID可能是不连续的，所以该优化仅针对特殊场景

**第二次优化：覆盖索引+in操作**<br>
select * from user where id in (select id from user limit 1000,10) 

**第三次优化：覆盖索引+inner join操作**<br>

# mysql inner join 会比 in 效率操作高？
1. 优化器选择更好的执行计划：MySQL 的查询优化器会根据表的统计信息、索引情况和查询条件等因素，选择最优的执行计划。在使用 JOIN 时，优化器可以考虑多个表之间的关联条件和索引，选择合适的连接顺序和连接算法，以最小化数据读取和计算的成本。而使用 IN 时，优化器无法直接利用表之间的关联条件和索引，可能导致全表扫描或大量的随机访问。

2. 避免重复数据的处理：当使用 JOIN 连接多个表时，可以通过关联条件过滤掉不符合条件的行，避免了重复数据的处理。而使用 IN 时，可能会导致重复数据的出现，需要额外的去重操作。

3. 内存占用和磁盘访问：使用 JOIN 可以将相关的数据行放在内存中一次性处理，减少了对磁盘的访问。而使用 IN 时，每个 IN 子查询都需要单独执行，可能导致频繁的磁盘访问和内存交换。