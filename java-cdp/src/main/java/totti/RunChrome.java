package totti;

import utils.FlagUtils;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Random;
import java.util.concurrent.TimeUnit;
import java.util.stream.Collectors;

public class RunChrome {
    public static void main(String[] args) throws IOException, InterruptedException {
        String tempDir = "/home/C5311429/temp/chromedb-runner" + new Random().nextInt(1000);
        File file = new File(tempDir);
        if (!file.exists()) {
            file.mkdir();
        }
        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            file.delete();
        }));

        String methodName = "Page.navigate";
        String path = "/usr/bin/google-chrome";
        // String macPath = "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome";

        // {"id":14,"sessionId":"7DD78405FDAAD20E11D3B02EB0293B50","method":"Page.navigate","params":{"url":"http://www.baidu.com"}}
        // ws://127.0.0.1:51574/devtools/browser/d3abe4e8-4644-4252-9a8e-d620fc03a0a3

//        File.createTempFile("chromedb-runner123", "");

        List<String> cmd = new ArrayList<>();
        cmd.add(path);
        cmd.addAll(Arrays.stream(FlagUtils.getFlags(tempDir)).collect(Collectors.toList()));
        ProcessBuilder processBuilder = new ProcessBuilder(cmd);

        processBuilder.start();

        new Thread(() -> {
            try {
                FileInputStream fileInputStream = new FileInputStream(file.getPath() + File.separator + "DevToolsActivePort");
                BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(fileInputStream));
                String line = "";
                int count = 0;
                while ((++count < 20) && (line = bufferedReader.readLine()) != null) {
                    System.out.println("line is ======" + line);
                }

            } catch (Exception e) {
                e.printStackTrace();
            }
        });

        TimeUnit.SECONDS.sleep(10);
    }
}
