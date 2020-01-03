package iot.tsdb.test.data.generator;

import static iot.tsdb.test.data.meta.DataConfiguration.CSV_SPLITOR;
import static iot.tsdb.test.data.meta.DataConfiguration.fields;
import static iot.tsdb.test.data.meta.DataConfiguration.getArea;
import static iot.tsdb.test.data.meta.DataConfiguration.getBjlx;
import static iot.tsdb.test.data.meta.DataConfiguration.getDistrict;
import static iot.tsdb.test.data.meta.DataConfiguration.getLine;
import static iot.tsdb.test.data.meta.DataConfiguration.getMpid;
import static iot.tsdb.test.data.meta.DataConfiguration.getProvince;
import static iot.tsdb.test.data.meta.DataConfiguration.getSystem;

import java.text.DecimalFormat;
import java.util.Random;

import com.google.common.collect.AbstractIterator;

import iot.tsdb.test.data.meta.DataSetMeta;

/**
 * timestamp,provice,city,system,mpid,cuserid...
 */
public class DataGenerator extends AbstractIterator<String> {
    private DecimalFormat df = new DecimalFormat("0.00");

    private final Random random;
    protected final DataSetMeta meta;
    private final int userType;

    private int currentUserCount;
    private int currentTimeSeriesIndex;

    public DataGenerator(DataSetMeta meta, long seed, int userType) {
        this.meta = meta;
        this.userType = userType;
        random = new Random(seed);
    }

    @Override
    protected String computeNext() {
        int userId = currentUserId();
        if (userId > meta.getEndUserId()) {
            return endOfData();
        }

        long timestamp = meta.calculateTimestamp(currentTimeSeriesIndex);
        currentTimeSeriesIndex++;

        if (isTimeEnd()) {
            nextUser();
        }

        return toLine(timestamp, userId, userType);
    }


    private boolean isTimeEnd() {
        return currentTimeSeriesIndex >= meta.getLineCountPerUser();
    }

    private int currentUserId() {
        return meta.getStartUserId() + currentUserCount;
    }

    private void nextUser() {
        currentTimeSeriesIndex = 0;
        currentUserCount++;
    }

    private String toLine(long timestamp, int cuserid, int userType) {
        LineBuilder lineBuilder = new LineBuilder();
        lineBuilder.append(timestamp)
                .append(cuserid)
                .append(getProvince(cuserid))
                .append(getDistrict(cuserid))
                .append(getSystem(cuserid))
                .append(getMpid(cuserid))
                .append(getBjlx(cuserid, userType))
                .append(getLine(cuserid))
                .append(getArea(cuserid));

        for (String field : fields) {
            double value = random.nextDouble() * 1000000;
            lineBuilder.append(df.format(value));
        }
        return lineBuilder.build();
    }

    private class LineBuilder {
        private StringBuilder stringBuilder;

        LineBuilder() {
            stringBuilder = new StringBuilder(512);
        }

        LineBuilder append(Object o) {
            stringBuilder.append(o);
            stringBuilder.append(CSV_SPLITOR);
            return this;
        }

        String build() {
            if (stringBuilder.length() > 1) {
                return stringBuilder.substring(0, stringBuilder.length() - 1);
            }
            return "";
        }
    }
}
