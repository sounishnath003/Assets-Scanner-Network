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
import java.net.NetworkInterface;
import java.util.ArrayList;
import java.util.Enumeration;
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

        for(int i=1; i < 20;i++) {
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
        String c1 = "Get-WmiObject -Class Win32_Product -ComputerName ";
        String c2 = " | select name, vendor, version, InstallDate, caption, IdentifyingNumber, PackageName, ProductID, WarrantyDuration, Description, InstallSource, PackageCode, WarrantyStateDate | ConvertTo-Json -depth 100 | out-file " + ip + ".json" ;
        String cmd = c1 + ip + c2;
        System.out.println(cmd);
        new Thread(() -> {
            try {
                Runtime runner = Runtime.getRuntime();
                Process pr = runner.exec(new String[]{"powershell", cmd});

                BufferedReader buf = new BufferedReader(new InputStreamReader(pr.getInputStream()));
                pr.waitFor();
                String out = "";
                while ((out=buf.readLine()) != null) {
                    System.out.println(out);
                }
            }catch (Exception e) {
                e.printStackTrace();
            }
        }).start();
    }

    public static void getDetailsOfAllNetworkInterfaces() throws Exception {
        Enumeration<NetworkInterface> interfaces = NetworkInterface.getNetworkInterfaces();

        while (interfaces.hasMoreElements()) {
            NetworkInterface networkInterface = interfaces.nextElement();
            // drop inactive
            if (!networkInterface.isUp())
                continue;

            // smth we can explore
            Enumeration<InetAddress> addresses = networkInterface.getInetAddresses();
            while(addresses.hasMoreElements()) {
                InetAddress addr = addresses.nextElement();
                System.out.println(String.format("NetInterface: name [%s], ip [%s]",
                        networkInterface.getDisplayName(), addr.getHostAddress()));
            }
        }
    }

    public static void main(String[] args) throws Exception {
        getDetailsOfAllNetworkInterfaces();
        getNetworkIPsWithThreading();
    }
}