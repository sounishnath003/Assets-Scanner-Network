#  IT Network Asset Mgmt Script
An Advanced and versatile tool who scans through your network, and find devices and gives an extensive information about that.

At the initial stage I am focusing upon a CLI based application tool, later will think to build GUI.

![GOLANG](https://miro.medium.com/max/3152/1*Ifpd_HtDiK9u6h68SZgNuA.png)

## Workflows WIP 🤞

- [x] Network CIDR support (eg: 12.168.0.1/24)
- [x] Scans all devices over network.
- [ ] Detect what type of device is it.
- [x] Also check which ports are open.
- [x] Device information
- [x] Extensive information like which software are installed on that machines
- [x] IPParsing and Junk Remote RPC connection check.
- [x] Send RPC Syn-Ping b/w devices and then background script execution started.
- [x] Saving all alive host machines and installed software details in a payload file in jsonStructObject in GO.
- [x] Synchronizing Multiple Go-Routines and Light weight Thead support.
- [x] Optimising the Load of execution.
## Workflows TODOs ⭐

- [ ] Synchronizing with active channels and cancel them.
- [ ] Optimizing load while execution

## Technology Used 🔥

Currently I have not yet figure it out how will I build the idea into production ready application.
Just a thought - GO will be a better choice or the other.