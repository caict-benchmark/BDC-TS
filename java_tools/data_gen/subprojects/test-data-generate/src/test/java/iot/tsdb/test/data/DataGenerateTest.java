package iot.tsdb.test.data;

import iot.tsdb.test.data.generator.DataGenerator;
import iot.tsdb.test.data.meta.DataSetMeta;
import org.junit.Test;

import java.time.Duration;
import java.time.LocalDate;

public class DataGenerateTest {

    @Test
    public void testLowVoltage() {
        DataSetMeta meta = new DataSetMeta.DataSet1("aaa", 0, 10, 2, LocalDate.of(2018, 1, 1));
        DataGenerator dataGenerator = new DataGenerator(meta, 0, 0);
        while (dataGenerator.hasNext()) {
            System.out.println(dataGenerator.next());
        }
    }

    @Test
    public void testMediumAndHighVoltage() {
        DataSetMeta meta = new DataSetMeta.DataSet1("bbb", 0, 10, 2, LocalDate.of(2018, 1, 1));
        DataGenerator dataGenerator = new DataGenerator(meta, 0, 1);
        while (dataGenerator.hasNext()) {
            System.out.println(dataGenerator.next());
        }
    }

    @Test
    public void testFeeder() {
        DataSetMeta meta = new DataSetMeta.DataSet2("ccc", 0, 10, 2, 1514736000000L, Duration.ofMinutes(15).toMillis());
        DataGenerator dataGenerator = new DataGenerator(meta, 0, 0);
        while (dataGenerator.hasNext()) {
            System.out.println(dataGenerator.next());
        }
    }
}
