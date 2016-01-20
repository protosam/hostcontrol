#!/bin/bash

sed -i '/chroot_local_user=/c\chroot_local_user=YES' /etc/vsftpd/vsftpd.conf
if ! grep ^chroot_local_user /etc/vsftpd/vsftpd.conf | grep YES; then
	echo Ensuring VSFTPD is chrooted.
	echo 'chroot_local_user=YES' >> /etc/vsftpd/vsftpd.conf
fi

sed -i '/allow_writeable_chroot/d' /etc/vsftpd/vsftpd.conf
if ! grep ^allow_writeable_chroot /etc/vsftpd/vsftpd.conf | grep YES; then
	echo Ensuring VSFTP allow writeable chroot.
	echo 'allow_writeable_chroot=YES' >> /etc/vsftpd/vsftpd.conf
fi

<<<<<<< HEAD
=======

>>>>>>> 9d53418ae2d62b3d958bb6ac7612d652fafe70f6
systemctl enable vsftpd

if [ -f /var/lock/subsys/vsftpd ]; then
	systemctl restart vsftpd
else
	systemctl start vsftpd
fi
