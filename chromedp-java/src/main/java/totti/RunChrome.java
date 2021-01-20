package totti;

import java.io.IOException;

import utils.FlagUtils;

public class RunChrome {
    public static void main(String[] args) throws IOException, InterruptedException {
        String methodName = "Page.navigate";
        // {"id":14,"sessionId":"7DD78405FDAAD20E11D3B02EB0293B50","method":"Page.navigate","params":{"url":"http://www.baidu.com"}}
        RunChrome runChrome = new RunChrome();
        // ws://127.0.0.1:51574/devtools/browser/d3abe4e8-4644-4252-9a8e-d620fc03a0a3
        ProcessBuilder processBuilder = new ProcessBuilder(FlagUtils.getFlags());
        processBuilder.start();

        runChrome.wait();
    }
}
