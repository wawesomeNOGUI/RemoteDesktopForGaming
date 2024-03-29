package main

/*
#include <windows.h>

void MouseMove(int nXMove, int nYMove)
{
	int nX, nY;
	int nScreenWidth = GetSystemMetrics(SM_CXVIRTUALSCREEN);
	int nScreenHeight = GetSystemMetrics(SM_CYVIRTUALSCREEN);
	int nScreenLeft = GetSystemMetrics(SM_XVIRTUALSCREEN);
	int nScreenTop = GetSystemMetrics(SM_YVIRTUALSCREEN);

	INPUT Input = { 0 };
  nX = (int)((((double)(nXMove)-nScreenLeft) * 65536) / nScreenWidth + 65536 / (nScreenWidth));
	nY = (int)((( (double) (nYMove)-nScreenTop) * 65536) / nScreenHeight + 65536 / (nScreenHeight));
	Input.type = INPUT_MOUSE;
	Input.mi.dwFlags = MOUSEEVENTF_ABSOLUTE | MOUSEEVENTF_MOVE;
	Input.mi.dx = nX;
	Input.mi.dy = nY;
	Input.mi.time = GetTickCount();
	Input.mi.dwExtraInfo = GetMessageExtraInfo();
	SendInput(1, &Input, sizeof(INPUT));
}

void MouseMoveFor3DGames(int nXMove, int nYMove)
{
	int nX, nY;
	int nScreenWidth = GetSystemMetrics(SM_CXVIRTUALSCREEN);
	int nScreenHeight = GetSystemMetrics(SM_CYVIRTUALSCREEN);
	int nScreenLeft = GetSystemMetrics(SM_XVIRTUALSCREEN);
	int nScreenTop = GetSystemMetrics(SM_YVIRTUALSCREEN);

	POINT cursor;
	GetCursorPos(&cursor);

	INPUT Input = { 0 };
	nX = (int)((((double)(nXMove + cursor.x)-nScreenLeft) * 65536) / nScreenWidth + 65536 / (nScreenWidth));
	nY = (int)((( (double) (nYMove + cursor.y)-nScreenTop) * 65536) / nScreenHeight + 65536 / (nScreenHeight));
	Input.type = INPUT_MOUSE;
	Input.mi.dwFlags = MOUSEEVENTF_ABSOLUTE | MOUSEEVENTF_MOVE;
	Input.mi.dx = nX;
	Input.mi.dy = nY;
	Input.mi.time = GetTickCount();
	Input.mi.dwExtraInfo = GetMessageExtraInfo();
	SendInput(1, &Input, sizeof(INPUT));
}


void MouseMoveRaw (int x, int y )
{
  //double fScreenWidth    = ::GetSystemMetrics( SM_CXSCREEN )-1;
  //double fScreenHeight  = ::GetSystemMetrics( SM_CYSCREEN )-1;
  //double fx = x*(65535.0f/fScreenWidth);
  //double fy = y*(65535.0f/fScreenHeight);

  INPUT  Input={0};
  Input.type      = INPUT_MOUSE;
  Input.mi.dwFlags  = MOUSEEVENTF_MOVE|MOUSEEVENTF_ABSOLUTE;
  Input.mi.dx = x;
  Input.mi.dy = y;
  SendInput(1,&Input,sizeof(INPUT));
}

void MouseDown ()
{
  INPUT  Input={0};
  Input.type      = INPUT_MOUSE;
  Input.mi.dwFlags  = MOUSEEVENTF_LEFTDOWN;
  SendInput(1,&Input,sizeof(INPUT));
}

void MouseUp ()
{
  INPUT  Input={0};
  Input.type      = INPUT_MOUSE;
  Input.mi.dwFlags  = MOUSEEVENTF_LEFTUP;
  SendInput(1,&Input,sizeof(INPUT));
}

void MouseWheel (int a)
{
	INPUT  Input={0};
	Input.type = INPUT_MOUSE;
	Input.mi.mouseData = a;
	Input.mi.dwFlags = MOUSEEVENTF_WHEEL;
	SendInput(1,&Input,sizeof(INPUT));
}

void RightMouseDown ()
{
  INPUT  Input={0};
  Input.type      = INPUT_MOUSE;
  Input.mi.dwFlags  = MOUSEEVENTF_RIGHTDOWN;
  SendInput(1,&Input,sizeof(INPUT));
}

void RightMouseUp ()
{
  INPUT  Input={0};
  Input.type      = INPUT_MOUSE;
  Input.mi.dwFlags  = MOUSEEVENTF_RIGHTUP;
  SendInput(1,&Input,sizeof(INPUT));
}

void KeySimulate (WORD keyAscii, _Bool down) //https://github.com/octalmage/robotjs/blob/master/src/keypress.c
{
	//Convert the ascii code to key scan code
	//UINT VKCode=LOBYTE(VkKeyScan(keyAscii));
  //UINT ScanCode=MapVirtualKey(VKCode,0);
	UINT ScanCode=MapVirtualKey(keyAscii & 0xff, 0);

  INPUT ip;
  // Set up a generic keyboard event.
  ip.type = INPUT_KEYBOARD;
  ip.ki.wScan = ScanCode; // hardware scan code for key (works for more applications than wVk)
  ip.ki.time = 0;
  ip.ki.dwExtraInfo = 0;
  ip.ki.wVk = 0; // virtual-key code for which key, set to 0 if not using

  if(down){
    ip.ki.dwFlags = KEYEVENTF_SCANCODE;
    SendInput(1, &ip, sizeof(INPUT));
  }else{
    // Release the key
		// Set the scan code for keyup
		ScanCode |= 0x80;
		ip.ki.wScan = ScanCode;
    ip.ki.dwFlags = KEYEVENTF_SCANCODE | KEYEVENTF_KEYUP; // KEYEVENTF_KEYUP for key release
    SendInput(1, &ip, sizeof(INPUT));
  }

}
//https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-keybdinput
//https://stackoverflow.com/questions/5607849/how-to-simulate-a-key-press-in-c
*/
import "C"

import (
	"log"
	"net"
	"net/http"
	//"strings"
	"encoding/json"
	"fmt"
	//"strconv"
	"time"
	"sync"
	"flag"

	"github.com/gorilla/websocket"

	"github.com/kbinani/screenshot"

	"github.com/pion/mediadevices"
	"github.com/wawesomeNOGUI/RemoteDesktopForGaming/signal"
	//"github.com/pion/mediadevices/pkg/frame"
	"github.com/pion/mediadevices/pkg/prop"
	"github.com/pion/webrtc/v3"

	// If you don't like x264, you can also use vpx by importing as below
	// "github.com/pion/mediadevices/pkg/codec/vpx" // This is required to use VP8/VP9 video encoder
	// or you can also use openh264 for alternative h264 implementation
	//"github.com/pion/mediadevices/pkg/codec/openh264"
	// or if you use a raspberry pi like, you can use mmal for using its hardware encoder
	//"github.com/pion/mediadevices/pkg/codec/mmal"
	//"github.com/pion/mediadevices/pkg/codec/opus" // This is required to use opus audio encoder
	//"github.com/pion/mediadevices/pkg/codec/x264" // This is required to use h264 video encoder
	"github.com/wawesomeNOGUI/RemoteDesktopForGaming/x264"

	// Note: If you don't have a camera or microphone or your adapters are not supported,
	//       you can always swap your adapters with our dummy adapters below.
	// _ "github.com/pion/mediadevices/pkg/driver/videotest"
	// _ "github.com/pion/mediadevices/pkg/driver/audiotest"
	//_ "github.com/pion/mediadevices/pkg/driver/camera" // This is required to register camera adapter
	//_ "github.com/pion/mediadevices/pkg/driver/microphone" // This is required to register microphone adapter
	_ "github.com/pion/mediadevices/pkg/driver/screen"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	//First Wait for the password:
	_, message, err2 := c.ReadMessage() //ReadMessage blocks until message received
	if err2 != nil {
		log.Println("readPassErr:", err)
	}

	if string(message) != *password {
		return  //if password wrong don't let gamer connect
	}

	//webrtc stuffffffffff
	/*
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
	*/
	// Create a new RTCPeerConnection
	//peerConnection, err := api.NewPeerConnection(config)
	peerConnection, err := api.NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		panic(err)
	}

	// Set the handler for ICE connection state
	// This will notify you when the peer has connected/disconnected
	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		log.Printf("Connection State has changed %s \n", connectionState.String())
		// Disconnected
		if connectionState == 5 || connectionState == 6 || connectionState == 7 {
			err := peerConnection.Close() //deletes all references to this peerconnection in mem and same for ICE agent (ICE agent releases the "closed" status)?
			if err != nil {               //https://www.w3.org/TR/webrtc/#dom-rtcpeerconnection-close
				fmt.Println(err)
			}

			//Stop pressing keys/mouse in case disconnected before keyUp/mouseup messages
			C.MouseUp()
			C.RightMouseUp()

			//Delete pressed keys, and do keyups
			keysDown.Range(func(key, value interface{}) bool {
				C.KeySimulate(C.WORD(key.(float64)), false) //keyup
				keysDown.Delete(key)
				return true
			})
		}
	})


//====================DataChannel For User Controls=============================
//Create a reliable datachannel with label "TCP" for all other communications
reliableChannel, err := peerConnection.CreateDataChannel("TCP", nil)
if err != nil {
  panic(err)
}

// Register channel opening handling
reliableChannel.OnOpen(func() {})

var rawInput bool = false
var fpMouse bool = false

reliableChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
	  //First check for bad values
		if msg.Data == nil {
			return
		}

		//Check For Mouse Clicks
		if string(msg.Data) == "mouseDown" {
			C.MouseDown()
			return
		}else if string(msg.Data) == "mouseUp" {
			C.MouseUp()
			return
		}else if string(msg.Data) == "rightMouseDown" {
			C.RightMouseDown()
			return
		}else if string(msg.Data) == "rightMouseUp" {
			C.RightMouseUp()
			return
		}

		//Check for if the message is for raw mouse input
		if string(msg.Data) == "rawOn" {
			rawInput = true
			return
		}else if string(msg.Data) == "rawOff" {
			rawInput = false
			//fmt.Println(rawInput)
			return
		}

		//Check for first person mouse
		if string(msg.Data) == "fpMouseOn" {
			fpMouse = true
			return
		}else if string(msg.Data) == "fpMouseOff" {
			fpMouse = false
			//fmt.Println(rawInput)
			return
		}


    //User Input Map
    controls := make(map[string]interface{})

    if err := json.Unmarshal(msg.Data, &controls); err != nil {
      fmt.Println(err)
    }

    if _, ok := controls["X"]; ok {
      mouseX := controls["X"].(float64)  //Javascript uses float64?
      mouseY := controls["Y"].(float64)

				if rawInput{
					//Relative movement, current cursor x & y + passed x & y ints
					C.MouseMoveRaw(C.int(mouseX), C.int(mouseY))
				}else if fpMouse{
					//Relative movement, current cursor x & y + passed x & y ints
					C.MouseMoveFor3DGames(C.int(mouseX), C.int(mouseY))
				}else{
					//Absolute coords movement
					C.MouseMove(C.int(mouseX), C.int(mouseY))
				}

    }else if _, ok := controls["keyDown"]; ok {
			if controls["keyDown"].(float64) != 17 {  //Extended keys work differnt, check out robot.js keypress.c, 17 is the ctrl key
				//Simulate one press so we don't have to wait for keyLoop for first press
				C.KeySimulate(C.WORD(controls["keyDown"].(float64)), true)
				keysDown.Store(controls["keyDown"].(float64), 0)
			}

		}else if _, ok := controls["keyUp"]; ok {
			if controls["keyUp"].(float64) != 17 {  //Extended keys work differnt, check out robot.js keypress.c
				//delete value from the sync.Map and make key "dirty", to mark for
				//key deletion in garbage collection? https://golang.org/src/sync/map.go?s=9414:9451#L282
				keysDown.Delete(controls["keyUp"].(float64))
				//simulate keyup
				C.KeySimulate(C.WORD(controls["keyUp"].(float64)), false)
			}

		}else if _, ok := controls["wheel"]; ok {
			C.MouseWheel(C.int(controls["wheel"].(float64)))
		}

})
//==============================================================================

//Add Screen Capture
	for _, track := range s.GetTracks() {
		track.OnEnded(func(err error) {
			log.Printf("Track (ID: %s) ended with error: %v\n",
				track.ID(), err)
		})

		_, err = peerConnection.AddTransceiverFromTrack(track,
			webrtc.RtpTransceiverInit{
				Direction: webrtc.RTPTransceiverDirectionSendonly,
			},
		)
		if err != nil {
			panic(err)
		}

	}

	// Create an offer
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		panic(err)
	}

	//Create a channel that is blocked until ICE Gathering is complete
	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	// Sets the LocalDescription, looks for ICE candidates
	err = peerConnection.SetLocalDescription(offer)
	if err != nil {
		panic(err)
	}

	//Wait for ICE gathering to complete (non-trickle ICE)
	<-gatherComplete

	//dt = time.Now()
	//log.Print(dt.String())

	//Output the SDP with the final ICE candidate
	log.Println(*peerConnection.LocalDescription())

	//Send the SDP with the final ICE candidate to the browser as our offer
	err = c.WriteMessage(1, []byte(signal.Encode(*peerConnection.LocalDescription()))) //write message back to browser, 1 means message in byte format
	if err != nil {
		log.Println("write:", err)
	}

	//Wait for the browser to return an answer (its SDP)
	msgType, message, err2 := c.ReadMessage() //ReadMessage blocks until message received
	if err2 != nil {
		log.Println("readSDPErr:", err)
	}

	answer := webrtc.SessionDescription{}

	signal.Decode(string(message), &answer) //set offer to the decoded SDP
	log.Print(answer, msgType)

	// Set the remote SessionDescription
	err = peerConnection.SetRemoteDescription(answer)
	if err != nil {
		panic(err)
	}

	//=====================Trickle ICE==============================================
	//Make a new struct to use for trickle ICE candidates
	var trickleCandidate webrtc.ICECandidateInit
	var leftBracket uint8 = 123 //123 = ascii value of "{"

	for {
		_, message, err2 := c.ReadMessage() //ReadMessage blocks until message received
		if err2 != nil {
			fmt.Println("readTrickleICEErr:", err)
		}

		//If staement to make sure we aren't adding websocket error messages to ICE
		if message[0] == leftBracket {
			//Take []byte and turn it into a struct of type webrtc.ICECandidateInit
			//(declared above as trickleCandidate)
			err := json.Unmarshal(message, &trickleCandidate)
			if err != nil {
				fmt.Println("errorUnmarshal:", err)
			}

			fmt.Println(trickleCandidate)

			err = peerConnection.AddICECandidate(trickleCandidate)
			if err != nil {
				fmt.Println("errorAddICE:", err)
			}
		}

	}

}

//===================Key Simulation Stuff=======================================
//repeatedly press keys that are down
func keyLoop(){
	for {
		keysDown.Range(func(key, value interface{}) bool {
			//simulate keyDown
			if value.(int) == 1 {
				C.KeySimulate(C.WORD(key.(float64)), true)
			}else{
				 //the 1 means has been through the keyLoop once, so start doing more presses
				keysDown.Store(key, 1)
			}

			return true
		})
		time.Sleep(time.Millisecond*150)
	}
}

var keysDown sync.Map
//==============================================================================

//==================Global WebRTC Vars==========================================
//var peerConnection PeerConnection
var s mediadevices.MediaStream
var x264Params x264.Params
var codecSelector *mediadevices.CodecSelector
var mediaEngine = webrtc.MediaEngine{}
//var api = webrtc.NewAPI(webrtc.WithMediaEngine(&mediaEngine))
var api *webrtc.API
var password *string
//==============================================================================

func main() {
	screenshot.Setup()
	defer screenshot.TearDown()

	//First Get Command Line Arguments (Flags)
	bitRate := flag.Int("bitrate", 5_000_000, "Integer Value For Video BitRate")
	webRTCIP := flag.String("ip", "", "IP for this computer for the browser webRTC peer to connect to")
	password = flag.String("password", "itGameTime", "The Password For the Browser Peer to Type and Send to Connect")
	frameRate := flag.Float64("fps", 60, "frameRate for screen capture (can also be decimal values)")
	keyFrameInterval := flag.Int("keyFrameInterval", 1, "Integer Value How Many Frames Between A KeyFrame (Keyframes can be decoded on their own without other reference frames)")
	flag.Parse()

	fmt.Println("Password to connect (can be changed with -password flag): " + *password)

	if *webRTCIP == "" {
		fmt.Println("Usage Example: screenGrabWebRTC.exe -ip 127.0.0.0")
		fmt.Println("You Can Optionally Change Video Params: screenGrabWebRTC.exe -bitrate 10000000 -ip 127.0.0.0")
		fmt.Println("Or List All Avaible Flags: screenGrabWebRTC.exe -help")
		return
	}

	//Setup Video Stream
	x264Params, err := x264.NewParams()
	if err != nil {
		panic(err)
	}
	x264Params.BitRate = *bitRate    // default 5_000_000 bps (5mbps)
	x264Params.KeyFrameInterval = *keyFrameInterval  //default 60
	x264Params.Preset = x264.PresetVeryfast
	//openh264Params.BitRate = 4_000_000 // 4000kbps
	//openh264Params.BitRate = 0

	codecSelector = mediadevices.NewCodecSelector(
		mediadevices.WithVideoEncoders(&x264Params),
		//mediadevices.WithAudioEncoders(&opusParams),
	)

	codecSelector.Populate(&mediaEngine)

	s, err = mediadevices.GetDisplayMedia(mediadevices.MediaStreamConstraints{
		Video: func(c *mediadevices.MediaTrackConstraints) {
			//c.FrameFormat = prop.FrameFormat(frame.FormatYUY2)
			//c.FrameFormat = prop.FrameFormatExact(frame.FormatI420)
			//c.Width = prop.Int(640)
			//c.Height = prop.Int(480)
			c.FrameRate = prop.Float(*frameRate)
		},
		//Audio: func(c *mediadevices.MediaTrackConstraints) {
		//},
		Codec: codecSelector,
	})
	if err != nil {
		panic(err)
	}

	// Listen on UDP Port 80, will be used for all WebRTC traffic
	udpListener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IP{0, 0, 0, 0},
		Port: 80,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Listening for WebRTC traffic at %s\n", udpListener.LocalAddr())

	// Create a SettingEngine, this allows non-standard WebRTC behavior
	settingEngine := webrtc.SettingEngine{}

	//Our Public Candidate is declared here cause were not using a STUN server for discovery
	//and just hardcoding the open port, and port forwarding webrtc traffic on the router
	settingEngine.SetNAT1To1IPs([]string{*webRTCIP}, webrtc.ICECandidateTypeHost)

	// Configure our SettingEngine to use our UDPMux. By default a PeerConnection has
	// no global state. The API+SettingEngine allows the user to share state between them.
	// In this case we are sharing our listening port across many.
	settingEngine.SetICEUDPMux(webrtc.NewICEUDPMux(nil, udpListener))

	// Create a new API using our SettingEngine & MediaEngine
	api = webrtc.NewAPI(webrtc.WithSettingEngine(settingEngine), webrtc.WithMediaEngine(&mediaEngine))

	//Start keyLoop
	go keyLoop()

	fileServer := http.FileServer(http.Dir("./public"))
	http.HandleFunc("/echo", echo) //this request comes from webrtc.html
	http.Handle("/", fileServer)

	err = http.ListenAndServe(":80", nil) //Http server blocks
	if err != nil {
		log.Fatal(err)
	}

}
