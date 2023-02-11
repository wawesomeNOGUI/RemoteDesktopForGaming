# RemoteDesktopForGaming
Hi! This is a remote desktop program for a Windows PC being used remotely by someone's web browser. 
This remote desktop program is built specially for people wanting to play video games remotely.
The program contains helper functions to simulate first or third person camera control used by most 3D games, and can play any 2D games
great with the remote keyboard events.

The program consists of a Go program you run on your Windows PC that serves as a web server and WebRTC peer.
Any client who wants to connect navigates to the website in their browser, and the Go program streams the PC's
screen to the browser client, and lets the browser send back mouse events and keyboard events to execute remotely on the PC.

# Downloading
You need to download the source code and exe from the github releases tab on the right.

# Running
After you have downloaded the source code and screenGrabWebRTC.exe:

 - Extract the source code wherever you'd like, and place screenGrabWebRTC.exe inside that folder.
 - Now inside that extracted folder you should have the public folder, screenGrabWebRTC.exe, and screenGrabWebRTC.go.
 - Next navigate into the public folder and open index.html with a text editor and change the address the websocket request is made to to the computer you want to remote into [here](https://github.com/wawesomeNOGUI/RemoteDesktopForGaming/blob/948e7983cd68dee31316d20a1ad2485d74ddf1ef/public/index.html#L126), and edit the resolution (the two 1600s are width and two 900s are height) in the mouse events [here](https://github.com/wawesomeNOGUI/RemoteDesktopForGaming/blob/948e7983cd68dee31316d20a1ad2485d74ddf1ef/public/index.html#L236).
 - Next port forward 80 on your router and map it to the local computer address you want to remote into.
 - Finally open command prompt and execute `screenGrabWebRTC.exe -h` to see all the flag options you have. (e.g screenGrabWebRTC.exe -ip 127.0.0.1 -fps 30)
 
 Thanks for checking out this project!
