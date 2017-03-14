#Lego EV3 REST API @ Go

####Installation:
- Download [Go](http://golang.org/dl/)
- Setup the $GOPATH: (you can setup in the other folder, up to you)
    - cd $HOME
    - mkdir gocode
    - cd gocode
    - mkdir src
    - mkdir bin
    - nano .bash_profile : 
        - export GOPATH=“$HOME/gocode”
        - export PATH=$PATH:$GOPATH/bin 
 - Get the GB package:
    - go get -u github.com/constabulary/gb/...
    
 - Install dependencies:
    - cd [PROJECT_ROOT]/src
    - gb vendor restore
    - (CTRL-C to exit manually, if command don't exit after install)
 - gb build all
 - if all is OK, you will see [bin] folder
    - cd bin
    - ./ev3-handle (to start app)
    
- Good ideas about the package [Design](https://www.goinggo.net/2017/02/package-oriented-design.html) philosophy


####Project targets:
- Deliver fast and easy Lego EV3 module communication API
- Export REST API for integration with 3rd-party applications
- Cross-platform execution
- Language independed integration
- Open source

####TODO:
- ~~Make a DEMO with connection to EV3 via Serial over BT~~
- ~~Define a list of supported operations~~
- ~~Implement Serial wrapper~~
- ~~Implement command builder according to Lego protocol~~
- ~~Implement response parser according to Lego protocol~~
- ~~Implement set of Serial supported operations~~
- ~~Design REST APIs~~
- ~~Implement REST APIs~~
- Cover by UT
- Provide integration documentation

####Supported brick commands:
- Play sound
- Get device list
- Motor control:
    - Start
    - Stop
    - Speed mode
    - Angle mode
    - _Time mode*_
    - _Get angle*_
- Get sensor data:
    - Light
    - Color
    - Touch
    - Distance
    - Gyro
    
_* testing is pending_

####Helpful links:
- https://goo.gl/FeoZe7
- https://goo.gl/Usm7Gr
- https://goo.gl/0LzQqF

####Rest API routes:
	GET /ports

	POST /sound/{volume}/frequency/{frequency}/duration/{duration}

	POST /motors/{port}/start
	POST /motors/{port}/stop/{brake}
	POST /motors/{port}/speed/{speed}
	GET /motors/{port}

	GET /sensors/{port}/color
	GET /sensors/{port}/luminosity
	GET /sensors/{port}/distance
	GET /sensors/{port}/click
	GET /sensors/{port}/clickcount
	GET /sensors/{port}/gyro/angle
	GET /sensors/{port}/gyro/gravity
	GET /sensors/{port}
