package iot.tsdb.test.data.command;

import java.time.LocalDate;
import java.time.format.DateTimeFormatter;

import com.beust.jcommander.Parameter;
import com.beust.jcommander.Parameters;

import iot.tsdb.test.data.meta.DataSetMeta;

@Parameters(commandNames = "generateDataSet1", commandDescription = "Generate data set 1")
public class GenerateDataSet1 extends AbstractGenerateCommand {

    @Parameter(names = "--startDate", description = "Format yyyy-MM-dd”")
    private String startDate = "2018-01-01";

    @Override
    protected DataSetMeta buildDataSetMeta(int startId, int endId) {

        return new DataSetMeta.DataSet1(metric, startId, endId, lineCountPerUser,
                LocalDate.parse(startDate, DateTimeFormatter.ofPattern("yyyy-MM-dd")));
    }
}
