# 时序数据库性能基准测试说明
工具基于influxdb-comparisons，以BDC-TS为基准打造的压测工具。支持以下的数据库，或者基于以下数据库改写的数据库。

+ InfluxDB
+ Elasticsearch ([announcement blog here](https://influxdata.com/blog/influxdb-markedly-elasticsearch-in-time-series-data-metrics-benchmark/))
+ Cassandra ([InfluxDB Tops Cassandra in Time-Series Data & Metrics Benchmark](https://www.influxdata.com/influxdb-vs-cassandra-benchmark-time-series-metrics/))
+ MongoDB ([InfluxDB is 27x Faster vs MongoDB for Time-Series Workloads](https://www.influxdata.com/influxdb-is-27x-faster-vs-mongodb-for-time-series-workloads/))
+ OpenTSDB

## 一、准备工作
需要安装golang和安装测试工程

### 1: 安装golang

```powershell
yum install golang
```

### 2: 安装测试工程

测试工程包括生成数据工具、导入工具、生成查询语句工具和查询测试工具

#### 如果没有指定GOPATH，需要指定GOPATH

```powershell
export GOPATH=/root/go  # 指定工程目录，随便指定一个地方
```

#### 安装生成数据工具

```powershell
go get github.com/caict-benchmark/BDC-TS/cmd/bulk_data_gen
```

#### 安装导入数据工具

导入数据需要根据你基于的数据库不同，安装不同的导入工具
```powershell
# influx
go get github.com/caict-benchmark/BDC-TS/cmd/bulk_load_influx

# ES
go get github.com/caict-benchmark/BDC-TS/cmd/bulk_load_es

# OPENTSDB
go get github.com/caict-benchmark/BDC-TS/cmd/bulk_load_opentsdb
```


## 二、测试工具基本使用介绍
在安装完golang和测试工具以后，就可以开始测试了

### 1、启动数据库
测试之前，需要把你的时序数据库启动。测试工具默认去读写的是localhost:8086，如果你的数据库不是在这个地址，请在使用工具的时候指定参数-url

### 2、生成数据

生成数据有两种方式，分别是生成数据到文件和边生成数据边写入数据库

#### 生成数据到文件
以influx为例（其他数据库替换工具名即可）
```powershell
$GOPATH/bin/bulk_data_gen --seed=123 --use-case=vehicle --scale-var=1 --format=influx-bulk | gzip > influx_bulk_records__usecase_vehicle__scalevar_1__seed_123.gz
```
use-case：这里使用的vehicle，也就是BDC-TS标准，请不要修改  
scalevar：定义有多少个设备同时上报，BDC-TS案例中约定20000或者20个车辆  
format： 写es、influx、opentsdb等，根据实际填入  
timestamp-start：数据开始时间 格式诸如 2008-01-01T08:00:01Z  
timestamp-end：数据结束时间 格式诸如 2008-01-01T08:00:01Z  

如，20000个设备产生1秒的数据应该使用以下命令
```powershell
$GOPATH/bin/bulk_data_gen --seed=123 --use-case=vehicle --scale-var=20000 --format=es-bulk --timestamp-start=2008-01-01T08:00:00Z --timestamp-end=2008-01-01T08:00:01Z | gzip > es_bulk_records_usecase_vehicle__scalevar_20000_seed_123.gz
```  


#### 边生成数据边导入数据库
以influx为例（其他数据库替换工具名即可）
```powershell
$GOPATH/bin/bulk_data_gen --seed=123 --use-case=vehicle --scale-var=1 --format=influx-bulk | $GOPATH/bin/bulk_load_influx  -workers 10 
```

### 3、导入数据
以influx为例（其他数据库替换工具名即可）
```powershell
cat influx_bulk_records__usecase_vehicle__scalevar_1__seed_123.gz | gunzip | $GOPATH/bin/bulk_load_influx --batch-size=5000 --workers=2
```

### 4、生成查询语句
TODO

### 5、执行查询
TODO

### 6、测试结束后清理数据
以influx为例，其他的DB的清理方法欢迎补充
```powershell
curl -XPOST 'http://localhost:8086/query?q=drop%20database%20benchmark_db'
```

## 三、时序数据库基准测试(BDC-TS)
我们工程的核心，是实现BDC-TS的测试。BDC-TS测试方案详见(CTSDB最佳实践)：https://github.com/caict-benchmark/BDC-TS/blob/master/practices/CTSDB_Tencent/README.md  
大家可以参考这个最佳实践进行测试

### 实时数据集
参考国标GB／T 32960.3-2016 电动汽车远程服务与管理系统技术规范 第3部分：通讯协议及数据格式，模拟实时生成的新能源汽车的运行数据
60个指标*20,000辆车=1,200,000个测点
数据生成间隔：1s（每个测点每隔1s产生一条数据，时间戳精确到毫秒，每秒共有120万个数据点生成，持续写入1小时，数据量约为73.8GB）

### 历史数据集

数据集1：测点少
数据生成到文件，数据总量约620.5GB
测点数：60个指标*20辆车=1,200个测点，数据跨度为1年

数据集2：测点多
数据生成到文件，数据总量约738GB
测点数：60个指标*20,000辆车=1,200,000个测，数据跨度为10小时 

## 四、自定义数据库
如果你的数据库不是基于InfluxDB、Elasticsearch 、Cassandra 、MongoDB、OpenTSDB中的任何一种，或者数据格式与这些数据库不一致，请自行添加数据库类型。或者联系gdchaochao进行协助  

方法是：仿照bulk_load、bulk_query_gen、cmd文件夹下的代码，重写一个数据库模型
