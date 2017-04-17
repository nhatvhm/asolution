Beego user registration example.
Objectives:
help others to bootstrap and get some ideas/examples around user registration micro service written in Go lang and BeeGo framework.
(well, feedback or contribution is always welcome).
Architecture:
Front-end Twitter Bootstrap, application server Go Lang and BeeGo, user DB MySQL and sessions in Redis.
End result Self contained micro service with the following services:
- SignUp.
- Account activation (send activation link email).
- SignIn.
- Captcha
- Logout.
- ResetPassword (Send reset password link email).
- Delete account.
- Profile/Account page.
URI entry points
/
(just an example how other controllers can use sessions from redis)
/accounts/signup
/accounts/signin
/accounts/signout
/accounts/verify/(uuid)
/accounts/delete
/accounts/profile
/accounts/forgotpassword
/accounts/resetpassword/(uuid)
Let’s get started!
Okay so what do you need for this example? VirtualBox with one VM (Tested on Ubuntu 14.04).
Just to make it simple for you to “hit” the VM from your host’s machine browser, You can set the networking to use only one nic in bridged mode. If this option is not possible in your environment you can use two nic’s #1 using nat (for internet access) and nic#2 Host-only Adapter so you can SSH in and access the VM from the host’s web browser.
For this example I’ll assume that you can finish this step but feel free to reach out if you need help.
Installation:
Note: for simplicity execute the following as root
ssh to the vm
Install MySQL
apt-get update --fix-missing
apt-get install mysql-server-5.6
* For this tutorial you can leave the root’s password blank
Create empty database in MySQL
mysql

mysql> create database beego_ureg;
Query OK, 1 row affected (0.00 sec)
Install Redis
apt-get install make gcc tcl
wget http://download.redis.io/releases/redis-3.2.0.tar.gz
tar -C /opt -xzf redis-3.2.0.tar.gz
cd /opt/redis-3.2.0/
make
make test
Start redis in background
/opt/redis-3.2.0/src/redis-server &
Install GoLang
wget https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.6.2.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile
source /etc/profile
Setup go environment variables
mkdir $HOME/gocode
echo 'export GOPATH=$HOME/gocode' >> .bashrc
echo 'export PATH=$PATH:$GOPATH/bin' >> .bashrc
source .bashrc
Installing the bee tool
apt-get install git mercurial

go get github.com/beego/bee
Almost there finally let’s pull the source code and explain some internals
go get github.com/TalLannder/beego-ureg
Note: all dependencies located in the vendor folder
if all went well so far let’s continue to the fun part
first things first, change to the source directory
cd $HOME/gocode/src/github.com/TalLannder/beego-ureg
now there’s only a few small things you need to update in conf/app.conf before you start the application server and play around.
vi conf/app.conf
base_url = dev01.loc:8080
This variable is used to construct the urls on the pages for resources like js,css,img etc. Now for everything to work as expected add an entry to your host's host file to point dev01.loc to the IP of the VM, so for example, if you bridged the VM and it got an IP, say, 192.168.0.158. Then if your host is a Windows7 machine edit the hosts file to add that entry.
C:\Windows\System32\drivers\etc\hosts
...
192.168.0.158 dev01.loc dev01
gmail_account = <YourUser>@gmail.com 
Gmail account that the application server will use to sent emails.
gmail_account_password = <PasswordHere>
Gmail's account password.
NOTE:
Gmail account might be subject to https://support.google.com/accounts/answer/6010255?hl=en
Start it plz do!
bee run
if all goes well the output should be similar to:
root@ubuntu:~/gocode/src/github.com/TalLannder/beego-ureg# bee run
2015/08/12 22:25:35 [INFO] Uses 'beego-ureg' as 'appname'
2015/08/12 22:25:35 [INFO] Initializing watcher...
2015/08/12 22:25:35 [TRAC] Directory(/root/go/src/github.com/TalLannder/beego-ureg/controllers)
2015/08/12 22:25:35 [TRAC] Directory(/root/go/src/github.com/TalLannder/beego-ureg)
2015/08/12 22:25:35 [TRAC] Directory(/root/go/src/github.com/TalLannder/beego-ureg/models)
2015/08/12 22:25:35 [TRAC] Directory(/root/go/src/github.com/TalLannder/beego-ureg/routers)
2015/08/12 22:25:35 [TRAC] Directory(/root/go/src/github.com/TalLannder/beego-ureg/tests)
2015/08/12 22:25:35 [TRAC] Directory(/root/go/src/github.com/TalLannder/beego-ureg/utilities/pbkdf2)
2015/08/12 22:25:35 [INFO] Start building...
2015/08/12 22:25:36 [SUCC] Build was successful
2015/08/12 22:25:36 [INFO] Restarting beego-ureg ...
2015/08/12 22:25:36 [INFO] ./beego-ureg is running...
create table `account`
    -- --------------------------------------------------
    --  Table Structure for `github.com/TalLannder/beego-ureg/models.Account`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `account` (
        `uid` varchar(255) NOT NULL PRIMARY KEY,
        `first` varchar(255) NOT NULL DEFAULT '' ,
        `last` varchar(255) NOT NULL DEFAULT '' ,
        `email` varchar(255) NOT NULL DEFAULT ''  UNIQUE,
        `phone1` varchar(255) NOT NULL DEFAULT '' ,
        `phone2` varchar(255) NOT NULL DEFAULT '' ,
        `address` varchar(255) NOT NULL DEFAULT '' ,
        `password` varchar(255) NOT NULL DEFAULT '' ,
        `registration_uid` varchar(255) NOT NULL DEFAULT '' ,
        `registration_date` datetime NOT NULL,
        `password_reset_uid` varchar(255) NOT NULL DEFAULT ''
    ) ENGINE=InnoDB;
2015/08/12 22:25:36 [asm_amd64.s:2232] [I] http server Running on :8080
Now if you manage to finish the steps above you should be able to open the browser on your host machine and see the index page.