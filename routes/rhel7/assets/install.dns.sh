#!/bin/bash
echo Getting nible user mysql password
HOSTCONTROL_PASS=$(grep "pass=" settings.cfg | cut -d= -f2)

echo Cofiguring DNS...

# For the future.
# wget http://geolite.maxmind.com/download/geoip/database/GeoLiteCity.dat.gz
# gunzip GeoLiteCity.dat.gz
# mv GeoLiteCity.dat /usr/share/GeoIP/GeoIPCity.dat

echo Setting up configuration files
sed -i '/^launch=/d' /etc/pdns/pdns.conf
sed -i '/^gmysql-host=/d' /etc/pdns/pdns.conf
sed -i '/^gmysql-dbname=/d' /etc/pdns/pdns.conf
sed -i '/^gmysql-user=/d' /etc/pdns/pdns.conf
sed -i '/^gmysql-password=/d' /etc/pdns/pdns.conf
sed -i '/^gmysql-socket=/d' /etc/pdns/pdns.conf


sed -i '/# launch=/a gmysql-socket=/var/lib/mysql/mysql.sock' /etc/pdns/pdns.conf
sed -i '/# launch=/a gmysql-password='"$HOSTCONTROL_PASS" /etc/pdns/pdns.conf
sed -i '/# launch=/a gmysql-user=hostcontrol' /etc/pdns/pdns.conf
sed -i '/# launch=/a gmysql-dbname=hostcontrol' /etc/pdns/pdns.conf
sed -i '/# launch=/a gmysql-host=localhost' /etc/pdns/pdns.conf
sed -i '/# launch=/a launch=gmysql' /etc/pdns/pdns.conf

echo Ensuring DNS service auto starts and is on
systemctl enable pdns

if [ -f /var/run/pdns.pid ]; then
	systemctl restart pdns
else
	systemctl start pdns
fi
