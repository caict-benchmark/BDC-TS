package iot.tsdb.test.data.meta;

import java.time.LocalDate;

import iot.tsdb.test.data.util.TimeUtils;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.Getter;

@Data
@AllArgsConstructor
public abstract class DataSetMeta {
    protected String name;
    protected int startUserId;
    protected int endUserId;
    protected int lineCountPerUser;

    abstract public long calculateTimestamp(int index);

    @Getter
    public static class DataSet1 extends DataSetMeta {
        private LocalDate startDate;

        public DataSet1(String name, int startUserId, int endUserId, int lineCountPerUser, LocalDate startDate) {
            super(name, startUserId, endUserId, lineCountPerUser);
            this.startDate = startDate;
        }

        public long calculateTimestamp(int index) {
            return TimeUtils.toTimestamp(startDate.plusMonths(index));
        }
    }

    @Getter
    public static class DataSet2 extends DataSetMeta {
        private long startTimestamp;
        private long interval;

        public DataSet2(String name, int startUserId, int endUserId, int lineCountPerUser,
                long startTimestamp, long interval) {
            super(name, startUserId, endUserId, lineCountPerUser);
            this.startTimestamp = startTimestamp;
            this.interval = interval;
        }

        public long calculateTimestamp(int index) {
            return startTimestamp + interval * index;
        }
    }
}
