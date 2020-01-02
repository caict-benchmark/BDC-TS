package iot.tsdb.test.data;

import com.beust.jcommander.JCommander;

import iot.tsdb.test.data.command.GenerateDataSet1;
import iot.tsdb.test.data.command.GenerateDataSet2;

public class Main {
    public static void main(String[] args) {
        GenerateDataSet1 generateDataSet1 = new GenerateDataSet1();
        GenerateDataSet2 generateDataSet2 = new GenerateDataSet2();

        JCommander jc = JCommander.newBuilder()
                .addCommand(generateDataSet1)
                .addCommand(generateDataSet2)
                .build();
        jc.parse(args);

        String commandName = jc.getParsedCommand();
        if (commandName != null) {
            Runnable runnable = (Runnable) jc.getCommands().get(commandName).getObjects().get(0);
            runnable.run();
        } else {
            jc.usage();
        }
    }
}
