package iot.tsdb.test.data.command;

import java.time.Duration;

import com.beust.jcommander.Parameter;
import com.beust.jcommander.Parameters;

import iot.tsdb.test.data.meta.DataSetMeta;

@Parameters(commandNames = "generateDataSet2", commandDescription = "Generate data set 2")
public class GenerateDataSet2 extends AbstractGenerateCommand {

    @Parameter(names = "--interval")
    private long interval = Duration.ofMinutes(15).toMillis();

    @Override
    protected DataSetMeta buildDataSetMeta(int startId, int endId) {
        return new DataSetMeta.DataSet2(metric, startId, endId,
                lineCountPerUser, startTimestamp, interval);
    }
}
