# FGD-Alterra-29
Capstone Project for Alterra Academy.

Contribute by: 
1. [Ilham Prasetya](https://github.com/ilse31)
2. [Wahyu Hauzan Rafi](https://github.com/whauzan)
3. [Muhammad Ghifari](https://github.com/Ghynmo)
172.17.0.2
server {
    listen 80;
    server_name example.com;

    location / {
        proxy_set_header   X-Forwarded-For $remote_addr;
        proxy_set_header   Host $http_host;
        proxy_pass         "http://127.0.0.1:8080";
    }
}
[ec2-user@ip-172-31-85-199 nginx]$ dpkg -L ca-certificates
-bash: dpkg: command not found
[ec2-user@ip-172-31-85-199 nginx]$ sudo find / |grep "\.pem"
/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem
/etc/pki/ca-trust/extracted/pem/email-ca-bundle.pem
/etc/pki/ca-trust/extracted/pem/objsign-ca-bundle.pem
/etc/pki/tls/cert.pem
/etc/pki/tls/certs/cert.pem
/etc/pki/tls/certs/dhparam.pem
/etc/ssl/private/key.pem
/var/lib/docker/overlay2/68e0fd3530922b391b6118c0b87b5a4f6486c6283756a37561393ff042ee21d5/diff/etc/ssl/cert.pem
/var/lib/docker/overlay2/4c87fdf0394d43c0b64d7536eb057089a9033cdcafdd66d198e237fe6c935756/merged/etc/ssl/cert.pem
/usr/lib/python2.7/site-packages/botocore/cacert.pem
/usr/lib/python3.7/site-packages/cfnbootstrap/packages/requests/cacert.pem
[ec2-user@ip-172-31-85-199 nginx]$ 
server {
  listen 443;
  ssl on;
  ssl_certificate /etc/nginx/conf.d/cert.pem;
  ssl_certificate_key /etc/nginx/conf.d/key.pem;
  location / {
  proxy_pass http://123.12.2.1:8080;
  proxy_set_header Host $host;
  proxy_set_header X-Forwarded-For $remote_addr;
 }
}
----------
server {
    listen 443;
    ssl on;
    server_name be.dkku.online;
    ssl_certificate /etc/pki/tls/certs/cert.pem;
    ssl_certificate_key /etc/ssl/private/key.pem;
    location / {
        proxy_set_header   X-Forwarded-For $remote_addr;
        proxy_set_header   Host $http_host;
        proxy_pass         "http://172.17.0.2:8080";
    }
}
