 编译打包：
 
 ./gradlew :test-data-generate:jar
 
 生成的jar文件位于:
 
 subprojects/test-data-generate/build/libs
 
 直接java -jar test-data-generate-1.0-SNAPSHOT.jar 运行可查看帮助

 数据集1的生成:

 对于9000万低压用户数据, userCount设置为90000000, startUserId设置为0, startDate设置为2019-01-01, metric设置为low_voltage, lineCountPerUser设置为12, 其他参数自行设置
 
 对于200万中高压用户数据, userCount设置为2000000, startUserId设置为0, startDate设置为2019-01-01, metric设置为medium_high_voltage, lineCountPerUser设置为60, 其他参数自行设置

 数据集2的生成:

 userCount设置为6000000, startUserId设置为0, metric设置为feeder, startTimestamp设置为1546272000000, lineCountPerUser设置为672, 其他参数自行设置
