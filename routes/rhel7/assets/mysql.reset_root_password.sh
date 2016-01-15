#!/bin/bash

# ensure skip-grant-tables is in the /etc/my.cnf
echo Starting mysql with skip-grant-tables
sed -i -e 's:\[mysqld\]:&\nskip-grant-tables:g' /etc/my.cnf

# ensure mysql is on now with those grant tables being ignored
if [ -f /var/run/mariadb/mariadb.pid ]; then
	systemctl restart mariadb
else
	systemctl start mariadb
fi

# generate a random passowrd
PASS=$(openssl rand -base64 12)
# This copy of the password is mysql query safe
PASS_SAFE=$(echo $PASS | sed 's/\(['"'"'\]\)/\\\1/g')

# reset the password
echo "Resetting the root password"
mysql -e "UPDATE mysql.user SET Password=PASSWORD('$PASS_SAFE') WHERE User='root';"




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

