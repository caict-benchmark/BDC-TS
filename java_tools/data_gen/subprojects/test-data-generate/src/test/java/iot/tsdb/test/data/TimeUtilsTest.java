package iot.tsdb.test.data;

import static iot.tsdb.test.data.util.TimeUtils.toDateStr;
import static iot.tsdb.test.data.util.TimeUtils.toTimestamp;

import java.time.LocalDate;

import org.junit.Test;

public class TimeUtilsTest {

    @Test
    public void test() {
        long timestamp = toTimestamp(LocalDate.of(2018, 01, 01));
        System.out.println(timestamp);
        System.out.println(toDateStr(timestamp));
        System.out.println(LocalDate.of(2018, 01, 01).plusMonths(12));
    }
}
