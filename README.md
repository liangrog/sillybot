# SillyBot
A silly bot only can move around 5X5 spaces

## Bot Requirements
[see here](requirement.txt)

## Installation
Download the [binary](https://github.com/liangrog/sillybot/releases) to your local machine. 

Note only linux and mac versions are provided

Depends on your version (linux or mac), run below command to change name and make it executable:
```
$ mv sillybot-[linux/mac] sillybot && chmod +x sillybot
```
## How to run
Run below command follow the instruction in [requirement](requirement.txt)
```
$ ./sillybot
```

Alternatively you can put all instructions in a file following the format of `example.yaml`, then just run
```
$ ./sillybot -f example.yaml
```

To get command help, run
```
$ ./sillybot -h
```
