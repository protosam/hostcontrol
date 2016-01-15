# hostcontrol
This is a webhosting control panel written in Go Lang.
  
#### Get the source
mkdir -p $GOPATH/src/github.com/protosam
cd $GOPATH/src/github.com/protosam
git clone git@gitlab.just.ninja:samuelp/hostcontrol.git
cd hostcontrol
go get
```
  
#### Development run
```
bash dev_run.sh
```
  
#### Build the installer
```
cd $GOPATH/src/github.com/protosam/hostcontrol
bash build.sh
ls -lah build
```

#### Install it after build
```
cd $GOPATH/src/github.com/protosam/hostcontrol/build
bash latest.sh
ls -lah /opt/hostcontrol
```
