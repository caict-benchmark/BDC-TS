# BDC-TS基准测试最佳实践 (CTSDB)
CTSDB是腾讯云上主推的时序数据库。接下来介绍如何使用CTSDB实现此BDC-TS的基准测试的最佳实践。

## 案例一：实时数据集测试

### 测试要求
测点数：60个指标*20,000辆车=1,200,000个测点  
数据生成间隔：1s（每个测点每隔1s产生一条数据，时间戳精确到毫秒，保证每秒有120万条数据生成）   
方式一：每隔一秒生成一个数据文件  
方式二：直接调用数据库接口写入 

### 测试实现步骤
#### 1、产生数据
use-case：这里使用的vehicle，也就是BDC-TS标准，请不要修改  
scalevar：定义有多少个设备同时上报，这个案例中约定20000个vehicle同时上报数据，所以是20000，请不要修改  
format： 写es、influx、opentsdb等，根据实际填入  
timestamp-start：数据开始时间 格式诸如 2008-01-01T08:00:01Z  
timestamp-end：数据结束时间 格式诸如 2008-01-01T08:00:01Z  
  
如，20000个设备产生1秒的数据应该使用以下命令  
```powershell
$GOPATH/bin/bulk_data_gen --seed=123 --use-case=vehicle --scale-var=20000 --format=es-bulk --timestamp-start=2008-01-01T08:00:00Z --timestamp-end=2008-01-01T08:00:01Z | gzip > es_bulk_records_usecase_vehicle__scalevar_20000_seed_123.gz
```  
  
但是，由于测试需要持续一段时间，比如1小时，需要产生3600个1秒的数据。所以使用上面的命令逐秒生成数据效率太低，需要借助脚本批量生成。脚本地址./practices/CTSDB_TENCENT/gen_data.sh，调用脚本命令：
```powershell
sh gen_data.sh --sec 3600 --output ./data_scale_20000_hour_1 --scale 20000 --interval 1 --format es --start-time 1546272001
```
--sec 是发送多长时间的数据  
--format 是你使用的数据库类型，如es、influx
其余参数不改即可

#### 2、导入数据到数据库
接上面的命令，在./data生成了一批文件（每个文件产生了1秒20万个数据），于是调用下面命令，写入数据到数据库
```powershell
sh load_data.sh --input ./data_scale_20000_hour_1 --format es --sleep 1 --start-time 1514736001
```
发送导入数据请求后sleep一秒，再发送下一个，所以这里的--sleep参数为1，这个参数请不要修改  

导入速度结果会如下显示：
```powershell
loaded 20160 items in 4.125944sec with 2 workers (mean point rate 4886.153966 items/sec, mean value rate 298055.391939/s, 9.34MB/sec from stdin)
```

#### 3、结果汇总
汇总上面的结果日志 TODO

  
  
## 案例二：历史数据集1（测点少）
### 测试要求
数据生成在一个csv文件中，数据总量约1TB  
测点数：60个指标*20辆车=1,200个测点  
数据生成间隔：N（每个测点每隔N时间产生一条数据，时间戳精确到毫秒，数据周期持续1年）
### 测试实现步骤
#### 1、产生数据
```powershell
sh gen_data.sh --sec 31536000 --output ./data_scale_20_year_1 --scale 20 --interval 86400 --format es --start-time 1546358401
``` 
--format 是你使用的数据库类型，如es、influx
其余参数不改即可
#### 2、导入数据到数据库
```powershell
sh load_data.sh --input ./data_scale_20_year_1 --format es --sleep 0
```
--format 是你使用的数据库类型，如es、influx
其余参数不改即可
#### 3、结果汇总
汇总上面的结果日志 TODO
  
  
## 案例三：历史数据集2（测点多）
### 测试要求
数据生成在一个csv文件中，数据总量约1TB  
测点数：60个指标*20,000辆车=1,200,000个测点  
数据生成间隔：1s（每个测点每隔1s产生一条数据，时间戳精确到毫秒，数据周期持续M时间）
### 测试实现步骤
#### 1、产生数据
```powershell
sh gen_data.sh --sec 36000 --output ./data_scale_20000_hour_10 --scale 20000 --interval 1 --format es
```
#### 2、导入数据到数据库
跟实时数据导入类似，这里sleep时间调整为0即可
```powershell
sh load_data.sh --input ./data_scale_20000_hour_10 --format es --sleep 0
```
导入速度结果会如下显示：
```powershell
loaded 20160 items in 4.125944sec with 2 workers (mean point rate 4886.153966 items/sec, mean value rate 298055.391939/s, 9.34MB/sec from stdin)
```
#### 3、结果汇总
汇总上面的结果日志 TODO