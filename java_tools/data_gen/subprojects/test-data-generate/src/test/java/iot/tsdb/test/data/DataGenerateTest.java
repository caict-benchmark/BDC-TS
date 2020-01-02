package iot.tsdb.test.data;

import iot.tsdb.test.data.generator.DataGenerator;
import iot.tsdb.test.data.meta.DataSetMeta;
import org.junit.Test;

import java.time.LocalDate;

public class DataGenerateTest {

    @Test
    public void test() {
        DataSetMeta meta = new DataSetMeta.DataSet1("aaa", 0, 10, 12, LocalDate.of(2018, 1, 1));
        DataGenerator dataGenerator = new DataGenerator(meta, 0);
        while (dataGenerator.hasNext()) {
            System.out.println(dataGenerator.next());
        }
    }
}
