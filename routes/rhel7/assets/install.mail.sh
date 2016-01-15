#!/bin/bash

HOSTCONTROL_PASS=$(grep "pass=" settings.cfg | cut -d= -f2)

echo Ensuring creation of mail tables in nimble database

#Create the virtual domains table with the following command:
#mysql nimble -e "CREATE TABLE mail_domains (domain varchar(50) NOT NULL, PRIMARY KEY (domain) );"

#Create a table to handle mail forwarding with the following command:
#mysql nimble -e "CREATE TABLE mail_forwardings (source varchar(80) NOT NULL, destination TEXT NOT NULL, PRIMARY KEY (source) );"

#Create the users table with the following command:
#mysql nimble -e "CREATE TABLE mail_users (email varchar(80) NOT NULL, password varchar(20) NOT NULL, PRIMARY KEY (email) );"

#Create a transports table with the following command:
#mysql nimble -e "CREATE TABLE mail_transport ( domain varchar(128) NOT NULL default '', transport varchar(128) NOT NULL default '', UNIQUE KEY domain (domain) );"


echo Copying some configuration templates into place
cp -fv common/src/rhel7/mail/mysql-virtual_domains.cf /etc/postfix/mysql-virtual_domains.cf
cp -fv common/src/rhel7/mail/mysql-virtual_forwardings.cf /etc/postfix/mysql-virtual_forwardings.cf
cp -fv common/src/rhel7/mail/mysql-virtual_mailboxes.cf /etc/postfix/mysql-virtual_mailboxes.cf
cp -fv common/src/rhel7/mail/mysql-virtual_email2email.cf /etc/postfix/mysql-virtual_email2email.cf

echo Writing passwords to those configuration files...
/usr/bin/replace __PASSWORD__ "$HOSTCONTROL_PASS" -- /etc/postfix/mysql-virtual_domains.cf
/usr/bin/replace __PASSWORD__ "$HOSTCONTROL_PASS" -- /etc/postfix/mysql-virtual_forwardings.cf
/usr/bin/replace __PASSWORD__ "$HOSTCONTROL_PASS" -- /etc/postfix/mysql-virtual_mailboxes.cf
/usr/bin/replace __PASSWORD__ "$HOSTCONTROL_PASS" -- /etc/postfix/mysql-virtual_email2email.cf

echo Securing those files that we just put passwords in
chmod -v o= /etc/postfix/mysql-virtual_*.cf
chgrp -v postfix /etc/postfix/mysql-virtual_*.cf


echo Ensuring vmail user exists as dovecot user alias
grep -q vmail /etc/passwd && echo vmail user already exists || useradd -d /home/vmail -u $(id -u dovecot) -g $(id -g dovecot) -o vmail

echo Ensuring vmail group exists as dovecot alias
grep -q vmail /etc/group && echo Group vmail exists || groupadd -g $(id -g dovecot) -o vmail



echo "Deaddrop postfix configuration... running postconf... pray to  your diety (if you wish dave)..."
postconf -e 'myhostname = '`hostname -f`
postconf -e 'mydestination = $myhostname, localhost, localhost.localdomain'
postconf -e 'mynetworks = 127.0.0.0/8'
postconf -e 'inet_interfaces = all'
postconf -e 'message_size_limit = 30720000'
postconf -e 'virtual_alias_domains ='
postconf -e 'virtual_alias_maps = proxy:mysql:/etc/postfix/mysql-virtual_forwardings.cf, mysql:/etc/postfix/mysql-virtual_email2email.cf'
postconf -e 'virtual_mailbox_domains = proxy:mysql:/etc/postfix/mysql-virtual_domains.cf'
postconf -e 'virtual_mailbox_maps = proxy:mysql:/etc/postfix/mysql-virtual_mailboxes.cf'
postconf -e 'virtual_mailbox_base = /home/vmail'
postconf -e 'virtual_uid_maps = static:'$(id -u vmail)
postconf -e 'virtual_gid_maps = static:'$(id -g vmail)
postconf -e 'smtpd_sasl_type = dovecot'
postconf -e 'smtpd_sasl_path = private/auth'
postconf -e 'smtpd_sasl_auth_enable = yes'
postconf -e 'broken_sasl_auth_clients = yes'
postconf -e 'smtpd_sasl_authenticated_header = yes'
postconf -e 'smtpd_recipient_restrictions = permit_mynetworks, permit_sasl_authenticated, reject_unauth_destination'
postconf -e 'smtpd_use_tls = yes'
postconf -e 'smtpd_tls_cert_file = /etc/pki/dovecot/certs/dovecot.pem'
postconf -e 'smtpd_tls_key_file = /etc/pki/dovecot/private/dovecot.pem'
postconf -e 'virtual_create_maildirsize = yes'
postconf -e 'virtual_maildir_extended = yes'
postconf -e 'proxy_read_maps = $local_recipient_maps $mydestination $virtual_alias_maps $virtual_alias_domains $virtual_mailbox_maps $virtual_mailbox_domains $relay_recipient_maps $relay_domains $canonical_maps $sender_canonical_maps $recipient_canonical_maps $relocated_maps $transport_maps $mynetworks $virtual_mailbox_limit_maps'
postconf -e 'virtual_transport = dovecot'
postconf -e 'dovecot_destination_recipient_limit = 1'



echo Ensuring that postfix has dovecot in the master config
sed -i '/^dovecot/d' /etc/postfix/master.cf
echo -e 'dovecot   unix   -   n   n   -   -   pipe\n   flags=DRhu   user=vmail:vmail   argv=/usr/libexec/dovecot/dovecot-lda -f ${sender} -d ${recipient}' >> /etc/postfix/master.cf


echo Copying dovecot templates
cp -fv common/src/rhel7/mail/dovecot.conf /etc/dovecot/dovecot.conf
cp -fv common/src/rhel7/mail/dovecot-sql.conf /etc/dovecot/dovecot-sql.conf

echo Setting hostname in dovecot.conf
/usr/bin/replace __HOSTNAME__ `hostname -f` -- /etc/dovecot/dovecot.conf

/usr/bin/replace __DOVECOTUID__ $(id -u dovecot) -- /etc/dovecot/dovecot.conf
/usr/bin/replace __DOVECOTGID__ $(id -g dovecot) -- /etc/dovecot/dovecot.conf


echo Setting password in dovecot-sql.conf
/usr/bin/replace __PASSWORD__ $HOSTCONTROL_PASS -- /etc/dovecot/dovecot-sql.conf

echo Securing dovecot-sql.conf, it has a password...
chgrp -v dovecot /etc/dovecot/dovecot-sql.conf
chmod -v o= /etc/dovecot/dovecot-sql.conf


echo Ensuring that dovecot and postfix start on boot
systemctl enable postfix
systemctl enable dovecot


echo Ensuring any changes are applied to postfix with a start or restart
if [ -f /var/spool/postfix/pid/master.pid ]; then
        systemctl restart postfix
else
        systemctl start postfix
fi


echo Ensuring any changes are applied to dovecot with a start or restart
if [ -f /var/run/dovecot/master.pid ]; then
        systemctl restart dovecot
else
        systemctl start dovecot
fi


echo Making sure /var/run/dovecot/auth-master is readable by dovecot
chown -v vmail:vmail /var/run/dovecot/auth-master
