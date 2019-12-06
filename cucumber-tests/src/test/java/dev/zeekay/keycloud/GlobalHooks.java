package dev.zeekay.keycloud;

import io.cucumber.java.Before;

import java.io.BufferedReader;
import java.io.File;
import java.io.IOException;
import java.io.InputStreamReader;
import java.sql.*;
import org.ini4j.*;

public class GlobalHooks {
    private static boolean dunit = false;
    private static Thread server;
    private static Process process;
    @Before
    public void beforeAll() throws IOException {
        if(!dunit) {
            Runtime.getRuntime().addShutdownHook(new Thread(() -> {
                process.destroy();
                server.stop();
            }));
            Ini ini = new Wini(new File("database.ini"));
            try (Connection con = DriverManager.getConnection(ini.get("postgresql", "url"),
                    ini.get("postgresql", "user") ,
                    ini.get("postgresql", "password"));
                 Statement st = con.createStatement();
                 )
            {
                st.execute("DELETE FROM sessions;DELETE FROM users;DELETE FROM passwds;");
            } catch (SQLException ex) {
                System.out.println(ex);
            }
            //
            server = new Thread(() -> {
                try {
                    process = Runtime.getRuntime().exec("./../server/server.exe");
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
