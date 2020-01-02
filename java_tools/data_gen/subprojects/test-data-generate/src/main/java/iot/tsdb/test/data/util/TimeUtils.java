package iot.tsdb.test.data.util;

import java.time.Instant;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.time.format.DateTimeFormatter;

public class TimeUtils {
    private static DateTimeFormatter ftf = DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss");
    private static ZoneId zoneId = ZoneId.of("Asia/Shanghai");

    public static long toTimestamp(String dateStr) {
        LocalDateTime parse = LocalDateTime.parse(dateStr, ftf);
        return LocalDateTime.from(parse).atZone(zoneId).toInstant().toEpochMilli();
    }

    public static String toDateStr(long timestamp) {
        return ftf.format(LocalDateTime.ofInstant(Instant.ofEpochMilli(timestamp), zoneId));
    }

    public static long toTimestamp(LocalDate localDate) {
        return toTimestamp(localDate.toString() + " 00:00:00");
    }

}
