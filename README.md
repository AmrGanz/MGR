# MGR "Must Gather Reader"
A TUI interface to navigate and view OpenShift 4 must-gather files.

# How to run it:
- Download and run the executable file "mgr":
~~~
# ./mgr <MG path>
~~~
- Or, you can clone the project and build your own executable files as follows [Recommended]:
~~~
# git clone https://github.com/AmrGanz/MGR.git
# cd MGR
# go build . -o mgr
# ./mgr <MG path>
~~~

# A screenshot of the main interface:
![Alt text](https://github.com/AmrGanz/MGR/blob/main/MainInterface.png?raw=true)

# Completed Features:
- Read Projects, Operators, Nodes, MCP, MC, and PersistentVolume resources details.
- Filter output text using the SearchBox
- Copy from the output text using the CopyMode button
- Ability to use the mouse

# To be added Features "TBA":
- Color coding
- Display and search containers full logs
- Keyboard shourtcuts
- Descrypt Secrets, and MachineConfigs
- Some options will give an output as "TBA" which means that this option will be added later
- ...

# Notes:
- This software is using [tview](https://github.com/rivo/tview) go library
- In my lab I am using go version go1.16.8
- This tool is in a PoC phase and more features will be added to it soon
