#!/bin/bash
# ensure that innodb is file-per-table
echo Ensuring that innodb is file per table in my.cnf
sed -i -e '/innodb-file-per-table/d' /etc/my.cnf
sed -i -e 's:\[mysqld\]:&\ninnodb-file-per-table=1:g' /etc/my.cnf


# ensure skip-grant-tables is in the /etc/my.cnf
echo Starting mysql with skip-grant-tables
sed -i -e 's:\[mysqld\]:&\nskip-grant-tables:g' /etc/my.cnf

# ensure mysql is on now with those grant tables being ignored
if [ -f /var/run/mariadb/mariadb.pid ]; then
        systemctl restart mariadb
else
        systemctl start mariadb
fi

# Ensure that things aren't borked due to upgrade from mysql55 to mysql56u
mysql_upgrade

# generate a random passowrd
PASS=$(openssl rand -base64 12)
# This copy of the password is mysql query safe
PASS_SAFE=$(echo $PASS | sed 's/\(['"'"'\]\)/\\\1/g')

# reset the password
echo "Setting up the root password"
mysql -e "UPDATE mysql.user SET Password=PASSWORD('$PASS_SAFE') WHERE User='root';"

# remove remote root access
echo "Ensuring root can not remotely access mysql"
mysql -e "DELETE FROM mysql.user WHERE User='root' AND Host NOT IN ('localhost', '127.0.0.1', '::1');"

# remove test db and privileges
echo "Ensuring the test db is dropped"
mysql -e "DROP DATABASE test;"
mysql -e "DELETE FROM mysql.db WHERE Db='test' OR Db='test\\_%'"


# remove anonymous users
echo "Removing any anonymous users"
mysql -e "DELETE FROM mysql.user WHERE User='';"

# flush privileges
echo Flusing privileges
mysql -e 'FLUSH PRIVILEGES;'

# ensure mysql starts on restart
systemctl enable mariadb


# Remove grant table skip
echo Removing skip-grant-tables
sed -i -e '/^skip-grant-tables/d' /etc/my.cnf

# Ensure mysqld is on
if [ -f /var/run/mariadb/mariadb.pid ]; then
        systemctl restart mariadb
else
        systemctl start mariadb
fi


# create /root/.my.cnf
echo "Creating /root/.my.cnf"
echo "[client]" > /root/.my.cnf
echo "user=root" >> /root/.my.cnf
echo "password=\"$PASS_SAFE\"" >> /root/.my.cnf
chown -v root:root /root/.my.cnf
chmod -v 640 /root/.my.cnf


# ensure database is setup
HCDB=$(mysql -e 'show databases' | grep -c ^hostcontrol$)
if [ "$HCDB" == "0" ]; then
        echo Creating database
        mysql -e 'create database hostcontrol'
        mysql hostcontrol < common/hostcontrol.sql
fi

HC_PASS=$(openssl rand -base64 12)
HC_PASS_SAFE=$(echo $HC_PASS | sed 's/\(['"'"'\]\)/\\\1/g')

sed -i -e 's:^dbname=.*:dbname=hostcontrol:g' settings.cfg
sed -i -e 's:^user=.*:user=hostcontrol:g' settings.cfg
sed -i -e 's:^pass=.*:pass='$HC_PASS':g' settings.cfg

if [ -S /var/lib/mysql/mysql.sock ]; then
	sed -i -e 's:^socket=.*:socket=/var/lib/mysql/mysql.sock:g' settings.cfg
fi

echo Creating hostcontrol user
mysql -e "DROP USER 'hostcontrol'@'localhost';"
mysql -e "DROP USER 'hostcontrol'@'127.0.0.1';"
mysql -e "GRANT ALL PRIVILEGES ON *.* TO hostcontrol@localhost IDENTIFIED BY '$HC_PASS_SAFE' WITH GRANT OPTION;"
mysql -e "GRANT ALL PRIVILEGES ON *.* TO hostcontrol@127.0.0.1 IDENTIFIED BY '$HC_PASS_SAFE' WITH GRANT OPTION;"
mysql -e "flush privileges;";
