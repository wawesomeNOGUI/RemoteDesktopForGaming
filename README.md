# RemoteDesktopForGaming
Hi! This is a remote desktop program for a Windows PC being used remotely by someone's web browser. 
This remote desktop program is built specially for people wanting to play video games remotely.
The program contains helper functions to simulate first or third person camera control used by most 3D games, and can play any 2D games
great with the remote keyboard events.

The program consists of a Go program you run on your Windows PC that serves as a web server and WebRTC peer.
Any client who wants to connect navigates to the website in their browser, and the Go program streams the PC's
screen to the browser client, and lets the browser send back mouse events and keyboard events to execute remotely on the PC.
