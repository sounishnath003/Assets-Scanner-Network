/*
 * Copyright (c) 2021.
 * @author: SOUNISH NATH
 * @Github: https://www.github.com/sounishnath003
 * Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package com.network.assets;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.InetAddress;
import java.util.ArrayList;
import java.util.List;

public class Main {

    public static List<String> activeIps = new ArrayList<>();
    public static List<String> deactivedIps = new ArrayList<>();

    public synchronized static void getNetworkIPsWithThreading() throws InterruptedException {
        final byte[] localhost;
        try {
            localhost = InetAddress.getLocalHost().getAddress();
        } catch (Exception e) {
            return;
        }

        for(int i=1; i < 255;i++) {
            final int j = i;
            // new thread for parallel execution
            new Thread(() -> {
                try {
                    localhost[3] = (byte) j;
                    InetAddress address = InetAddress.getByAddress(localhost);
                    String output = address.toString().substring(1);
                    if (address.isReachable(5000)) {
                        System.out.println(output + " is on the network");
                        activeIps.add(output);

                        execCommands(output);


                    } else {
                        deactivedIps.add(output);
                        System.out.println("Not Reachable: "+output);
                    }
                } catch (Exception e) {
                    e.printStackTrace();
                }
            }).start();     // don't forget to start the thread
        }
    }

    public synchronized static void execCommands(String ip) throws IOException {
        System.out.println("Running for IP: " + ip);
        String cmd = "ls";
        try {
            Runtime runner = Runtime.getRuntime();
            Process pr = runner.exec(new String[]{"powershell", cmd});

            BufferedReader buf = new BufferedReader(new InputStreamReader(pr.getInputStream()));
            pr.waitFor();
            String out = "";
            while ((out=buf.readLine())!=null) {
                System.out.println(out);
            }
        }catch (Exception e) {
            e.printStackTrace();
        }
    }

    public static void main(String[] args) throws Exception {
        getNetworkIPsWithThreading();
        activeIps.stream().forEach(System.out::println);
        deactivedIps.stream().forEach(System.out::println);
    }
}