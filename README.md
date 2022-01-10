# MGR "Must Gather Reader" 
MGR "_not the final name_" is a simple TUI interface to navigate and view OpenShift 4 must-gather files.

# How to run it:
- Download and run the executable file "mgr":
~~~
# ./mgr <MG path>
~~~
- Or, you can clone the project and build your own executable files as follows [Recommended]:
~~~
# git clone https://github.com/AmrGanz/MGR.git
# cd MGR
# go mod init mgr
# go mod tidy
# go build . -o mgr
# ./mgr <MG path>
~~~
- The main interface is devided into multiple areas as shown in the following screenshot:
![Alt text](https://github.com/AmrGanz/MGR/blob/main/InterfaceAreas.png?raw=true)
- You will always start with area "1" or List number 1.
- Most of the flow will start from List 1 then you have new options to choose from in List 2 and so on till you finally get an output in the Text Output area.
- Sometimes you will ge the output in the Text Area from selecting options in List 2, so you don't have to always get it from the last List.
- `CopyMode` button will bring focus to the contents in the Text Output area by hiding other areas.
- While in the `CopyMode`, you can use your mouse to highlight and Select/Copy lines.
- To go back to the main interfacr, press the `Esc` key in your keyboard.
- The `SearchBox` is used to do a simple filtering of the `Text Output` and it is case sensitive
- To retrive the original content, clean the `SearchBox` then press on the `Search Button`

# A screenshot of the main interface showing an example output:
![Alt text](https://github.com/AmrGanz/MGR/blob/main/MainInterface.png?raw=true)

# Added Features:
- Read Projects, Operators, Nodes, MCP, MC, and PersistentVolume resources details.
- Text output is similar to what you get with the oc client
- Quick navigation between cluster resources
- Filter output text using the SearchBox
- Copy from the output text using the CopyMode button
- Ability to use the mouse
- Read a MG that is generated from multiple images
- Descrypt Secrets, and MachineConfigs

# To be added Features "TBA":

- More cluster reources to be added
- Color coding
- Dynamically set the initial size of the interface areas according to the initial terminal size
- Display and search containers full logs
- Keyboard shourtcuts
- Some selections will give an output as "TBA" which means that this feature will be added later

# Notes:
- This tool is using [tview](https://github.com/rivo/tview) go library
- In my lab I am using go version go1.16.8
- This tool is in a PoC phase and more features will be added to it soon
- Please feel free to report any bugs or a required feature/enhancement
