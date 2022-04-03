# codesamples-go-smart-home-device-simulator
Smart device source code for the https://github.com/giantpanda9/codesamples-nodejs-ts-smart-home-simulator no necessity to run it from here, unless you want to rebuild it, if you run the goSmart device from here, the Node.js part should exchange signals via post requests (it is an emulator after all) with this pseudo and the binary that Node.js system executes will be stopped, because port 8081 already busy.
# Installation
1) go1.18.linux-amd64.tar.gz from https://go.dev/
2) cd /path/to/downloaded/go1.18.linux-amd64.tar.gz
3) sudo tar -C /usr/local -xzf go1.18.linux-amd64.tar.gz
3-1) [optional] do sudo rm -rf /usr/local/go if you have other versions of Go installed - not needed for the first time installation
4) sudo nano ~/.profile
5) add to the end of file the following lines:
```
export GOPATH=$HOME/work
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```
6) cd /path/to/the/folder/of/the/cloned/project
7) go mod init goSmart 
8) go install 'github.com/gorilla/mux@latest' - this should download latest gorilla/mux module to the $HOME/work directory, which was mentioned at point 5)
9) go run goSmart.go - to run at once
OR 
10) go build goSmart.go - to create a binary

# The questions that could be asked on the first place for this code (FAQ based on probability):
 
 ## Q: Why AJAX is used instead of IoT signalling system for an instance? 
 **A:** The answer will be given in three distinct sections, each of those provides a single seprate reason why AJAX is used.

 - As per task decription this is a simulator, a model, which required to trace and debug possible ways of exchanging an ivformation between modules and "the brain", AJAX is the most developed data exchanging system, a mature one, this means that there are popular instruments to debug it - like Postman that has been used as part of this development process - and it is much easier to check, which signal goes where and when, this saved time during development process. In short, to develop and debug quicker and save time.   

 - Besides, as per task description I should keep in mind the architecture of the device/system and the system I am going to simulate in some way - not to 100% - was already described on free code camp - https://www.freecodecamp.org/news/the-most-robust-and-secure-home-automation-system-6d0ddbb39f29/ Though they use MongoDB there to keep a JSON signals and commands, I hardcode them(it is a quick-to-build-simulator, not a real system,afterall). There are many similar systems described via the Internet - the one proposed via freecodecamp is the closest one. In short, analogues are present to keep them in mind.
 
 -  This is not a 100% simulator system, in some way. Given the necessity it could be expanded to be working as a real hardware based Smart House system given some architecture additions and tiny changes would be implemented. Raspberry Pi systems (and similar systems as well) did support the Golang installation and if Golang part, which located in a bios folder will be upgraded to access the hardware ports of the Raspberry Pi via the low-level implementation (Golang is like C/C++ afterall) the bios part will stay on Raspberry Pi (let's say under Raspbian, but any Linux would do) it can still communicate with the central "brain" - hub via AJAX, while actual and real Switches, Air Conditioners and Heaters will connected from there by smartGo binary via hardware ports using driver-boards or hand-made interface boards connected to Raspberry Pi input ports. This way require additional investigation and not requested via task description directly therefore it is postponed and only shallow research is done for now on this matter. In short, even in this state it is theoretically possible to upgrate this system to a real working one - even based on AJAX - or so it seems.
   

  ## Q: Why Golang?
  **A:** The answer again will be given in three distinct sections, each of those provides a single seprate reason why AJAX is used.
  
  - Roleplay reasons, it is like in a theater, the Smart Device system should be considered like something external and even alien to the hub or brain system, because technically "the brain" should be controlling any Smart Device (reasonably speaking), so Golang's purpose is to simulate this somewhat random smart device. And as I have been told during communication about the possible implementation - that would be a great and interesting point to (a quote) "implement matser control using one language and other units in another one". 
     
 - Besides the Go binary is not just a simulator - like it was said in question 1, the Golang generated binary is a basis for the further development of Basic input and output system (BIOS), which will translate AJAX to some hardware port signals to a real Smart Devices using Raspberry Pi, for example. 
     
 - The main difference of Go and JavaScript and Node.js is a different approach to multi-threading/multi-processing parallel systems. Go based on Goroutines a seprate process that could be created from a main one and yet have an easy access to the whole lexical and name spacing of the main program. In some cases it would be easy, like in a current example and in systems where this could become harder a messages that Goroutines can exchange can solve this issue. In particular case, Golang was really helpful, when solving the issue with the Heater, which was and intellectual device in its own turn. The Golang allowed to use an AfterFunc function to create a seprate Goroutine, which is supposed to turn off the heater after 25 minutes passed in a simple and elegant way as shown below:
```
if signal == "cold" && comingBackSoon { // if we received a singal that somebody would come back soon
			if (heaterSwitch == false) { // double check that the heater is turned off
				heaterSwitch = true /// turn it on
				time.AfterFunc(25 * time.Minute, func() { // and schedule a Goroutine to run in 25 minutes
					heaterSwitch = false // to turn it off
				})
			}
		}
```

  ## Q: Is Golang an OOP style language?
  **A:** 
  Though Golang does not have a "classic" OOP implementation (not classic does not mean it is improper), with the "class"-like keywords, it still supports and hierarchy of objects, which are the main characteristic of OOP designed applications, but in more lightweight way. Types in Go can be embedded in other types, which result in something similar to subclassing. It supports encapsulation, interfaces and methods. The only main (and noticeable) difference from classic OOP is that Go is using Composition instead of Ineritance, which allows the creation of complex types by combining objects or components of other types, while inheritance based on a base and/or parent class ancestry. So, summarzing all that said, the Go does not violate the rule to implement the test in proper OOP style, it is still a proper-OOP language, just does it differently.

  ## Q: Why some of code implementations are different from task description and some additional data taken from the Weather providing API?
  **A:** I took a liberty, I hope it will not violate the rules in a significant way, to improve the logic of the Smart Device simultor working in a very tiny way as follows:
  
 - The Switch - it looks like that "hot" and "cold" signals are being send to it at a random pace, no conditions described (at least in my edition of the task description or I missed those) in the document. So, I took additional data from the Weather API proposed and now the switch will turn on and off based on the time of sunrise and sunset. Not complete approach, to improve it, it would also be great to take cloudiness and visibility into account, but just do not wanted to violate the rules too much.
    
- In case of air conditioner unit, it seems like its temperature would increase up to no limit, the same could be said in relation to decrease - to avoid this a comfort temperature parameter is added (its value could be changed in the config file) to limit the temperature that Air Conditioner Unit might produce both for "cold" and "hot".
    
 - And in case of Heater, nothing changed too much, just added a check that it is already turned on, so that no new Gourutines would be created to turn it off - of course those would finish their task and stop working in order of creation (the first one, will exit last), but there is no reason to create too many similar instances of code to execute.
