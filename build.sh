#!/bin/bash
rm -rf build
mkdir -p build/hostcontrol
cp -rf www build/hostcontrol/
cp -rf template build/hostcontrol/
cp -rf common build/hostcontrol/
cp -rf settings.example.cfg build/hostcontrol/
. pack_assets.sh
go get -u
go build main.go
mv main build/hostcontrol/hostcontrol
cd build/
tar -czvf hostcontrol.tar.gz hostcontrol/
cat <<EOM > latest.sh
sed -e '1,/^exit/d' \$0 | tar -xz
mkdir -p /opt
mv hostcontrol /opt/
cp /opt/hostcontrol/settings.example.cfg /opt/hostcontrol/settings.cfg
/opt/hostcontrol/hostcontrol /api/install
exit
EOM
cat hostcontrol.tar.gz >> latest.sh
rm -rf hostcontrol hostcontrol.tar.gz
chmod +x latest.sh
