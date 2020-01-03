package iot.tsdb.test.data.command;

import java.io.File;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

import com.beust.jcommander.Parameter;

import iot.tsdb.test.data.meta.DataSetMeta;
import iot.tsdb.test.data.runner.WriteToFileRunner;

public abstract class AbstractGenerateCommand implements Runnable {

    private String fileSuffix = ".csv";

    @Parameter(names = "--seed", description = "seed for generate random value")
    private long seed = 123445432123L;
    @Parameter(names = "--maxThreadCount")
    protected int maxThreadCount = 30;
    @Parameter(names = "--startTimestamp", description = "default is 2019.01.01 00:00:00")
    protected long startTimestamp = 1546272000000L;
    @Parameter(names = "--output", description = "output dir", required = true)
    protected String output;
    @Parameter(names = "--fileCount", description = "generate file currentTimeSeriesIndex")
    protected int fileCount = 1;
    @Parameter(names = "--metric", description = "metric")
    protected String metric = "datapoints";

    @Parameter(names = "--queueSize", description = "queue size")
    protected int queueSize = 10000;
    @Parameter(names = "--batchSize", description = "batch size")
    protected int batchSize = 100;
    @Parameter(names = "--lineCountPerUser", description = "time range")
    protected int lineCountPerUser = 1;
    @Parameter(names = "--userCount", description = "user currentTimeSeriesIndex")
    protected int userCount = 90000000;
    @Parameter(names = "--startUserId")
    protected int startUserId = 0;
    @Parameter(names = "--userType", description = "0 for low voltage, 1 for medium and high voltage")
    protected int userType;


    private ExecutorService executorService;

    @Override
    public void run() {
        prepareOutputDir();

        int threadCount = Math.min(fileCount, maxThreadCount);
        executorService = Executors.newFixedThreadPool(threadCount);

        int countPerFile = userCount/fileCount;
        for (int i = 0; i < fileCount; i++) {
            int startId = countPerFile * i + startUserId;
            int count = i == (fileCount - 1) ? userCount - countPerFile * i : countPerFile;
            int endId = startId + count - 1;
            DataSetMeta meta = buildDataSetMeta(startId, endId);

            WriteToFileRunner runner = new WriteToFileRunner(batchSize, queueSize, getFileName(i), meta, seed, userType);
            executorService.submit(runner);
        }
        executorService.shutdown();
    }

    abstract protected DataSetMeta buildDataSetMeta(int startId, int totalCount);

    private String getFileName(int index) {
        return output + "/" + metric + "_" + index + fileSuffix;
    }

    private void prepareOutputDir() {
        File file = new File(output);
        if (file.exists() && !file.isDirectory()) {
            throw new IllegalArgumentException("file: "+ output + " is not a directory.");
        }
        if (!file.exists()) {
            if (!file.mkdirs()) {
                throw new IllegalStateException("Create dir fail. dir=" + output);
            }
        }
    }
}
