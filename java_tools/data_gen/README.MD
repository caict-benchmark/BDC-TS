 编译打包：
 
 ./gradlew :test-data-generate:jar
 
 生成的jar文件位于:
 
 subprojects/test-data-generate/build/libs
 
 直接java -jar test-data-generate-1.0-SNAPSHOT.jar 运行可查看帮助

 数据集1的生成:
  对于9000万低压用户数据, userCount设置为90000000, userType设置为0, startUserId设置为0, startDate设置为2019-01-01, lineCountPerUser设置为12, 其他参数自行设置. 如：

 ```
 java -jar test-data-generate-1.0-SNAPSHOT.jar generateDataSet1 --userCount 90000000 --userType 0 --startUserId 0 --startDate 2019-01-01 --lineCountPerUser 12 --maxThreadCount 48 --fileCount 48 --output [输出目录]
 ```
 
 对于200万中高压用户数据, userCount设置为2000000, userType设置为1, startUserId设置为0, startDate设置为2019-01-01, lineCountPerUser设置为60, 其他参数自行设置。
 ```
 java -jar test-data-generate-1.0-SNAPSHOT.jar generateDataSet1 --userCount 2000000 --userType 1 --startUserId 0 --startDate 2019-01-01 --lineCountPerUser 60 --maxThreadCount 48 --fileCount 48 --output [输出目录]
 ```

 数据集2的生成:

 userCount设置为6000000, userType设置为0, startUserId设置为0, startTimestamp设置为1546272000000, lineCountPerUser设置为672, 其他参数自行设置

 ```
 java -jar test-data-generate-1.0-SNAPSHOT.jar generateDataSet2 --userCount 6000000 --userType 0 --startUserId 0 --startTimestamp 1546272000000 --lineCountPerUser 672 --maxThreadCount 48 --fileCount 48 --output [输出目录]
 ```

