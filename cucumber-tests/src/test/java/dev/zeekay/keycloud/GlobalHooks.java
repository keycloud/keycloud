package dev.zeekay.keycloud;

import io.cucumber.java.Before;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class GlobalHooks {
    private static boolean dunit = false;
    private static Thread server;
    private static Process process;
    @Before
    public void beforeAll() {
        if(!dunit) {
            Runtime.getRuntime().addShutdownHook(new Thread(() -> {
                process.destroy();
                server.stop();
            }));
            server = new Thread(() -> {
                try {
                    process = Runtime.getRuntime().exec("./../server/server");
                    StringBuilder output = new StringBuilder();

                    BufferedReader reader = new BufferedReader(
                            new InputStreamReader(process.getInputStream()));

                    String line;
                    while ((line = reader.readLine()) != null) {
                        output.append(line + "\n");
                    }
            } catch (IOException e) {
                e.printStackTrace();
            }});
            server.start();
            dunit = true;
        }
    }
}
