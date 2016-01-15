#!/bin/bash

# Add mod_ruid2
cd common/src/rhel7/mod-ruid2

# Compile and enable mod_ruid2 in apache conf
apxs -a -i -l cap -c mod_ruid2.c

# Go back to previous dir
cd -


# Ensure phpMyAdmin is accessible
cat <<EOM >/etc/httpd/conf.d/phpMyAdmin.conf
Alias /phpMyAdmin /usr/share/phpMyAdmin
Alias /phpmyadmin /usr/share/phpMyAdmin
<Directory "/usr/share/phpMyAdmin">
        <IfModule mod_authz_core.c>
                Require all granted
        </IfModule>
        <IfModule !mod_authz_core.c>
                Order allow,deny
                Allow from all
        </IfModule>
</Directory>
<Directory "/usr/share/phpMyAdmin/setup">
        <IfModule mod_authz_core.c>
                Require all denied
        </IfModule>
        <IfModule !mod_authz_core.c>
                Order allow,deny
                Deny from all
        </IfModule>
</Directory>
EOM

# Ensure reverse proxy for hostcontrol
cat <<EOM >/etc/httpd/conf.d/hostcontrol.conf
<VirtualHost _default_:80>
#	ServerName null
#	ServerAlias *
	Include conf.d/hostcontrol.inc
</VirtualHost>
EOM

cat <<EOM >/etc/httpd/conf.d/hostcontrol.inc
<LocationMatch "^/(?!phpmyadmin|roundcubemail|shellinabox)">
	ProxyPass http://localhost:1337
</LocationMatch>

SSLProxyEngine on
SSLProxyVerify off
ProxyPass /shellinabox https://localhost:4200/
EOM

cat <<EOM >/etc/httpd/conf.d/roundcubemail.conf
#
# Round Cube Webmail is a browser-based multilingual IMAP client
#

Alias /roundcubemail /usr/share/roundcubemail

# Define who can access the Webmail
# You can enlarge permissions once configured

<Directory /usr/share/roundcubemail/>
    <IfModule mod_authz_core.c>
        # Apache 2.4
        Require all granted
    </IfModule>
    <IfModule !mod_authz_core.c>
        # Apache 2.2
        Order Allow,Deny
        Allow from all
    </IfModule>
</Directory>

# Define who can access the installer
# keep this secured once configured

<Directory /usr/share/roundcubemail/installer/>
    <IfModule mod_authz_core.c>
        # Apache 2.4
        Require local
    </IfModule>
    <IfModule !mod_authz_core.c>
        # Apache 2.2
        Order Deny,Allow
        Deny from all
        Allow from 127.0.0.1
        Allow from ::1
    </IfModule>
</Directory>

# Those directories should not be viewed by Web clients.
<Directory /usr/share/roundcubemail/bin/>
    Order Allow,Deny
    Deny from all
</Directory>
<Directory /usr/share/roundcubemail/plugins/enigma/home/>
    Order Allow,Deny
    Deny from all
</Directory>
EOM


# setup roundcube
ROUNDCUBEPASSWORD=$(openssl rand -base64 12)
mysql -e 'drop database roundcubedb'
mysql -e 'create database roundcubedb'
mysql -e "grant all privileges on roundcubedb.* to roundcubeuser@localhost identified by '$ROUNDCUBEPASSWORD';"
mysql -e 'FLUSH PRIVILEGES'
mysql roundcubedb < /usr/share/roundcubemail/SQL/mysql.initial.sql
sed -i -e "s/\$config\['db_dsnw'\].*/\$config['db_dsnw'] = 'mysql:\/\/roundcubeuser:$ROUNDCUBEPASSWORD@localhost\/roundcubedb';/g" /etc/roundcubemail/defaults.inc.php
ln -s /etc/roundcubemail/defaults.inc.php  /etc/roundcubemail/config.inc.php
sed -i "s/\$config\['create_default_folders'\] = false;/\$config\['create_default_folders'\] = true;/g" /etc/roundcubemail/defaults.inc.php


# Ensure that SSL has a default vhost for SSL:
sed -i 's/<VirtualHost .*:443>/<VirtualHost _default_:443>/g' /etc/httpd/conf.d/ssl.conf

# Remove any instancest of hostcontrol.inc and re-add just 1 (avoiding duplicates)
sed -i -e '/Include conf.d\/hostcontrol.inc/d' /etc/httpd/conf.d/ssl.conf
sed -i '/<\/VirtualHost>/i Include conf.d\/hostcontrol.inc' /etc/httpd/conf.d/ssl.conf

# Ensure that NameVirtualHost is declared before include of conf.d
sed -i -e '/^NameVirtualHost \*:80/d' /etc/httpd/conf/httpd.conf
sed -i '/Include conf.d/i NameVirtualHost \*:80' /etc/httpd/conf/httpd.conf

# Remove any instancest of vhosts.d include directory (avoiding duplicates)
#sed -i -e 's:Include vhosts.d/\*.conf::g' /etc/httpd/conf/httpd.conf
sed -i -e '/IncludeOptional vhosts.d\/\*.conf/d' /etc/httpd/conf/httpd.conf

# Ensure we include vhosts.d conf files.
echo 'IncludeOptional vhosts.d/*.conf' >> /etc/httpd/conf/httpd.conf

# make sure that vhosts.d directory exists
mkdir -p /etc/httpd/vhosts.d

# Ensure welcome.conf is empty
echo > /etc/httpd/conf.d/welcome.conf

# Ensure apache is enabled on boot
systemctl enable httpd

# start or restart apache
if [ -f /var/run/httpd/httpd.pid ]; then
	systemctl restart httpd
else
	systemctl start httpd
fi
