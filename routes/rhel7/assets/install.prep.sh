#!/bin/bash
# ensure we have epel-release
cat <<EOM >/etc/yum.repos.d/epel-bootstrap.repo
[epel]
name=Bootstrap EPEL
mirrorlist=http://mirrors.fedoraproject.org/mirrorlist?repo=epel-\$releasever&arch=\$basearch
failovermethod=priority
enabled=0
gpgcheck=0
EOM

yum --enablerepo=epel -y install epel-release
rm -f /etc/yum.repos.d/epel-bootstrap.repo

cat <<EOM >/etc/yum.repos.d/mariadb-10.1.repo
# MariaDB 10.1 CentOS repository list - created 2015-12-22 07:21 UTC
# http://mariadb.org/mariadb/repositories/
[mariadb]
name = MariaDB
baseurl = http://yum.mariadb.org/10.1/centos7-amd64
gpgkey=https://yum.mariadb.org/RPM-GPG-KEY-MariaDB
gpgcheck=1
EOM


# Install all of our packages below
yum -y install \
vim gcc gcc-c++ make autoconf yum-plugin-replace pam-devel \
httpd httpd-devel libcap-devel mod_ssl \
mariadb-server phpmyadmin \
php php-devel php-fpm php-mysql php-gd php-mcrypt php-xmlrpc php-xml php-mbstring php-pear \
pdns pdns-backend-mysql bind-utils \
postfix dovecot dovecot-mysql roundcubemail \
vsftpd \
unzip cyrus-sasl openssh-server initscripts cronie sudo libcgroup \
shellinabox


# We still don't like selinux... apache runs as multiple users... this is a multi-user system... and I'm lazy
setenforce 0
sed -i 's/^SELINUX=.*/SELINUX=disabled/g' /etc/sysconfig/selinux


if firewall-cmd --state; then
        echo firewalld is in use, configuring firewalld
        firewall-cmd --zone=public --add-port=80/tcp --permanent
        firewall-cmd --zone=public --add-port=443/tcp --permanent
        firewall-cmd --zone=public --add-port=1337/tcp --permanent
        firewall-cmd --zone=public --add-port=1338/tcp --permanent
        firewall-cmd --zone=public --add-port=3306/tcp --permanent

        firewall-cmd --zone=public --add-port=21/tcp --permanent
        firewall-cmd --zone=public --add-port=22/tcp --permanent
        firewall-cmd --zone=public --add-port=25/tcp --permanent
        firewall-cmd --zone=public --add-port=110/tcp --permanent
        firewall-cmd --zone=public --add-port=143/tcp --permanent
        firewall-cmd --zone=public --add-port=487/tcp --permanent
        firewall-cmd --zone=public --add-port=993/tcp --permanent
        firewall-cmd --zone=public --add-port=995/tcp --permanent


        firewall-cmd --reload
elif iptables --list; then
        echo iptables is in use, configuring iptables
        # Allow our services to be accessible through the firewall.
        iptables -I INPUT -p tcp --dport 80 -j ACCEPT
        iptables -I INPUT -p tcp --dport 443 -j ACCEPT
        iptables -I INPUT -p tcp --dport 1337 -j ACCEPT
        iptables -I INPUT -p tcp --dport 1338 -j ACCEPT
        iptables -I INPUT -p tcp --dport 3306 -j ACCEPT

        iptables -I INPUT -p tcp --dport 21 -j ACCEPT
        iptables -I INPUT -p tcp --dport 22 -j ACCEPT
        iptables -I INPUT -p tcp --dport 25 -j ACCEPT
        iptables -I INPUT -p tcp --dport 110 -j ACCEPT
        iptables -I INPUT -p tcp --dport 143 -j ACCEPT
        iptables -I INPUT -p tcp --dport 487 -j ACCEPT
        iptables -I INPUT -p tcp --dport 993 -j ACCEPT
        iptables -I INPUT -p tcp --dport 995 -j ACCEPT

        iptables-save

else
        echo Was unable to determine firewall in use.
fi

# start or restart sasl... we fucking want this to run.
systemctl restart saslauthd
systemctl enable saslauthd


# Setup shellinabox
sed -i '/^pts/d' /etc/securetty; for i in {0..40}; do echo pts/$i >> /etc/securetty; done
sed -i '/^OPTS=/d' /etc/sysconfig/shellinaboxd
echo 'OPTS="--user-css Normal:+/usr/share/shellinabox/white-on-black.css --disable-ssl-menu -s /:LOGIN"' >> /etc/sysconfig/shellinaboxd
systemctl restart shellinaboxd
systemctl enable shellinaboxd

cp -rfvp common/src/rhel7/startup_scripts/hostcontrol.service /etc/systemd/system/hostcontrol.service
systemctl daemon-reload
systemctl enable hostcontrol
systemctl restart hostcontrol
