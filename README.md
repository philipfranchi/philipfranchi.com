# Source code for philipfranchi.net

This is the source code for my personal website. It's really a work in progress

## Steps:
1. Create Route53 Domain
2. Spin up EC2
3. Make sure security opens 80, 443, and 22 to public
4. Create and assosciate an Elastic IP (static IP)
5. SSH into machine and install apache (httpd)
6. Create virtual host:80 (philipfranchi.net) in /etc/httpd/conf.d
7. Install mod_ssl, mod_proxy if it isnt on there
8. Turn the virtual host into a reverse proxy, specifying /api 127.0.0.1:8000 as the route
9. Install certbot
10. Stop httpd
11. Run certbot certonly, filling everything out as normal. certbot uses port 80 and 443, so make sure to stop the server
12. Add SSL to the website conf, under virtualhost 443
13. Restart httpd
14. Clone my repo onto the machine and build the files
15. Symlink my build output dir to the dir listed in the conf file for the site (I needed to change the group of the files to be apache `chown -R :apache <files>`) to get this to work
16. Pause to appreciate the front-end running :)
17. Create a cron job (1 01,13 * * * root /usr/bin/certbot renew --quiet) using crontab -e to auto renew the cert twice a day
18. Turn the backend into a service by creating a file under /etc/systemd/system with these contents
```
[Unit]
Description=philipfranchi backend service
After=network.target
StartLimitIntervalSec=0

[Service]
Type=simple
Restart=always
RestartSec=1
User=ec2-user
ExecStart=/usr/bin/env /bin/bash /home/ec2-user/philipfranchi.net/backend/scripts/service.sh

[Install]
WantedBy=multi-user.target\
```