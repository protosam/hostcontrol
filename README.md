# THIS REPO IS NO LONGER MAINTAINED
Sorry about that, I've not had time to work on this in a very long time. If you need a cPanel alternative that's free, I'd recommend checking out CyberPanel: https://cyberpanel.net/

# hostcontrol
This is a webhosting control panel written in Go Lang. The project page with screenshots, compiled binary, and some more info can be found here: http://just.ninja/hostcontrol
  
#### Prepare your RedHat 7/CentOS 7 server before hand
We need these packages to get the source.
```
yum install golang git -y
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
```


#### Get the source
```
mkdir -p $GOPATH/src/github.com/protosam
cd $GOPATH/src/github.com/protosam
git clone git@github.com:protosam/hostcontrol.git
cd hostcontrol
go get
```
  
#### Development run
For the first dev run, we need to do a setup
```
bash dev_run.sh first
```
Aftersetup is done, further runs will just be
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
