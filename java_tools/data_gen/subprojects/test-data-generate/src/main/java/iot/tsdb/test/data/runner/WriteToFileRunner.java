package iot.tsdb.test.data.runner;

import java.io.BufferedWriter;
import java.io.FileWriter;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.LinkedBlockingQueue;
import java.util.concurrent.TimeUnit;

import iot.tsdb.test.data.generator.DataGenerator;
import iot.tsdb.test.data.meta.DataSetMeta;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;

@RequiredArgsConstructor
@Slf4j
public class WriteToFileRunner implements Runnable {

    private final int batchSize;
    private final int queueSize;
    private final String fileName;
    private final DataSetMeta meta;
    private final long seed;
    private final int userType;

    private volatile BlockingQueue<List<String>> queue;
    private ExecutorService dataGenerateService;

    @Override
    public void run() {
        queue = new LinkedBlockingQueue<>(queueSize);

        dataGenerateService = Executors.newSingleThreadExecutor();

        dataGenerateService.submit(this::generateData);
        dataGenerateService.shutdown();

        readAndWrite();
    }

    private void generateData() {
        DataGenerator generator = new DataGenerator(meta, seed, userType);
        List<String> list = new ArrayList<>(batchSize);
        while (generator.hasNext()) {
            list.add(generator.next());
            if (list.size() == batchSize) {
                try {
                    queue.put(list);
                } catch (InterruptedException e) {
                    throw new IllegalStateException(e);
                }
                list = new ArrayList<>(batchSize);
            }
        }
        if (list.size() > 0) {
            queue.add(list);
        }
    }

    private void readAndWrite() {
        List<String> list;
        int confirmEnd = 0;
        try (BufferedWriter writer = new BufferedWriter(new FileWriter(fileName))){
            while (true) {
                try {
                    list = queue.poll(10, TimeUnit.MILLISECONDS);
                    if (list != null ) {
                        for (String line : list) {
                            writer.write(line);
                            writer.newLine();
                        }
                    } else if (dataGenerateService.isTerminated()) {
                        confirmEnd++;
                        if (confirmEnd > 1) {
                            break;
                        }
                    }
                } catch (InterruptedException e) {
                    throw new IllegalStateException(e);
                }
            }

        } catch (IOException e) {
            log.error("write to file error. file={}", fileName, e);
        }
    }
}
