# 时序数据库性能基准测试说明
支持以下的数据库，或者基于以下数据库改写的数据库 

+ InfluxDB
+ Elasticsearch ([announcement blog here](https://influxdata.com/blog/influxdb-markedly-elasticsearch-in-time-series-data-metrics-benchmark/))
+ Cassandra ([InfluxDB Tops Cassandra in Time-Series Data & Metrics Benchmark](https://www.influxdata.com/influxdb-vs-cassandra-benchmark-time-series-metrics/))
+ MongoDB ([InfluxDB is 27x Faster vs MongoDB for Time-Series Workloads](https://www.influxdata.com/influxdb-is-27x-faster-vs-mongodb-for-time-series-workloads/))
+ OpenTSDB

## 准备
需要安装golang和安装测试工程

### Phase 1: 安装golang

```powershell
yum install golang
```
配置GOROOT，修改~/.bash_profile，添加以下语句
```powershell
export GOROOT=/usr/local/go # 你安装的路径
```
执行
```powershell
source ~/.bash_profile
```

### Phase 2: 安装测试工程

测试工程包括生成数据工具、导入工具、生成查询语句工具和查询测试工具

#### 如果没有指定GOPATH，需要指定GOPATH

```powershell
export GOPATH=/usr/local/go  # 指定工程目录，随便指定一个地方
```

#### 安装生成数据工具

```powershell
go get github.com/gdchaochao/influxdb-comparisons/cmd/bulk_data_gen
```

#### 安装导入数据工具

导入数据需要根据你基于的数据库不同，安装不同的导入工具
```powershell
# influx
go get github.com/gdchaochao/influxdb-comparisons/cmd/bulk_load_influx

# ES
go get github.com/gdchaochao/influxdb-comparisons/cmd/bulk_load_es

# OPENTSDB
go get github.com/gdchaochao/influxdb-comparisons/cmd/bulk_load_opentsdb
```


## 测试工具的使用
在安装完golang和测试工具以后，就可以开始测试了

### 启动数据库
测试之前，需要把你的时序数据库启动。测试工具默认去读写的是localhost:8086，如果你的数据库不是在这个地址，请在使用工具的时候指定参数-url

### 生成数据

生成数据有两种方式，分别是生成数据到文件和边生成数据边写入数据库

#### 生成数据到文件
以influx为例（其他数据库替换工具名即可）
```powershell
$GOPATH/bin/bulk_data_gen --seed=123 --use-case=vehicle --scale-var=1 --format=influx-bulk | gzip > influx_bulk_records__usecase_vehicle__scalevar_1__seed_123.gz
```
--use-case  必须为vehicle

#### 边生成数据边导入数据库
以influx为例（其他数据库替换工具名即可）
```powershell
$GOPATH/bin/bulk_data_gen --seed=123 --use-case=vehicle --scale-var=1 --format=influx-bulk | $GOPATH/bin/bulk_load_influx  -workers 10 
```

### 导入数据
以influx为例（其他数据库替换工具名即可）
```powershell
cat influx_bulk_records__usecase_vehicle__scalevar_1__seed_123.gz | gunzip | ./bulk_load_influx --batch-size=5000 --workers=2
```

### 生成查询语句
TODO

### 执行查询
TODO

### 测试结束后清理数据
```powershell
curl 'http://localhost:8086/query?q=drop%20database%20benchmark_db'
```

### 工具参数说明
TODO   
具体参数说明后续会补充完整

## 时序数据库基准测试案例

### 实时数据集
测点数：60个指标*20,000辆车=1,200,000个测点  
数据生成间隔：1s（每个测点每隔1s产生一条数据，时间戳精确到毫秒，保证每秒有120万条数据生成）   
方式一：每隔一秒生成一个数据文件  
方式二：直接调用数据库接口写入  

测试方案(如何使用该工具完成案例)：
TODO

### 历史数据集1（测点少）
数据生成在一个csv文件中，数据总量约1TB  
测点数：60个指标*20辆车=1,200个测点  
数据生成间隔：N（每个测点每隔N时间产生一条数据，时间戳精确到毫秒，数据周期持续1年） 

测试方案(如何使用该工具完成案例)：
TODO

### 历史数据集2（测点多）
数据生成在一个csv文件中，数据总量约1TB  
测点数：60个指标*20,000辆车=1,200,000个测点  
数据生成间隔：1s（每个测点每隔1s产生一条数据，时间戳精确到毫秒，数据周期持续M时间）  

测试方案(如何使用该工具完成案例)：
TODO


## 自定义数据库
如果你的数据库不是基于InfluxDB、Elasticsearch 、Cassandra 、MongoDB、OpenTSDB中的任何一种，或者数据格式与这些数据库不一致，请自行添加数据库类型。  

方法是：仿照bulk_load、bulk_query_gen、cmd文件夹下的代码，重写一个数据库模型
