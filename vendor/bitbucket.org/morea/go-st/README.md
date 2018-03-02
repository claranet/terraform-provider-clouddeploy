# Ghost client #

Running tests:

```
$ go test
Testing Ghost client get all apps
All apps retrieved:
{{0xc420498180 0xc4203a7360} [{{55cf9ce5fde8dd0521358a19 0xc42038c950 0xc42038c860 0xc42038c4d0 0xc42038c730 0xc42038c4f0 0xc42034e600} adrien demo-symfony2 prod webfront eu-west-1 t2.micro vpc-3f1eb65a [ghost-devops@morea.fr] {ami-03ce4474 admin subnet-a7e849fe} {iam.ec2.demo ghost-demo [] {0 } [sg-6814f60c] [subnet-a7e849fe]} [{php5 5.4} {apache2 2.2} {php5-mysqlnd 0} {php5-curl 0} {php5-dev 0} {php5-xmlrpc 0} {php5-gd 0} {php5-memcache 0} {php5-redis 0} {pkg libapache2-mod-php5} {php5-zendopcache 0} {php5-intl 0}] [{0xc42038c718 symfony2 https://github.com/KnpLabs/KnpIpsum.git code /var/www 0 0  ZXhpdCAx  }]} {{56740addfde8dd01b2401ea6 0xc42038cb60 0xc42038caa0 0xc42038c990 0xc42038c9f0 0xc42038c9b0 0xc42034eb10} demo demo-symfony2 dev webfront eu-west-1 t2.micro vpc-804a1fe5 [jrespaut@morea.fr] {ami-03ce4474 admin subnet-5c8fb339} {iam.ec2.demo.ghost ghost-demo [] {0 } [sg-4db12629] [subnet-5c8fb339]} [{php5 5.4} {apache2 2.2} {php5-mysqlnd 0} {php5-curl 0} {php5-dev 0} {php5-xmlrpc 0} {php5-gd 0} {php5-memcache 0} {php5-redis 0} {pkg libapache2-mod-php5} {php5-zendopcache 0} {php5-intl 0} {composer 0}] [{0xc42038c9e8 symfony2 https://github.com/symfony/symfony-demo code /var/www 0 0  ZXhpdCAx Y29tcG9zZXIgaW5zdGFsbCAtLW5vLWludGVyYWN0aW9u }]} {{565efc5afde8dd1465d1f649 0xc42038cda0 0xc42038cca0 0xc42038cbb0 0xc42038cc20 0xc42038cbf0 0xc42034ec00} demo demo-symfony2 test webfront eu-west-1 t2.micro vpc-3f1eb65a [jrespaut@morea.fr] {ami-03ce4474 admin subnet-17d7b24e} {iam.ec2.demo.ghost ghost-demo [] {0 } [sg-6814f60c] [subnet-a7e849fe]} [{php5 5.4} {apache2 2.2} {php5-mysqlnd 0} {php5-curl 0} {php5-dev 0} {php5-xmlrpc 0} {php5-gd 0} {php5-memcache 0} {php5-redis 0} {pkg libapache2-mod-php5} {php5-zendopcache 0} {php5-intl 0} {composer 0}] [{0xc42038cbf9 symfony2 https://github.com/symfony/symfony-demo code /var/www 0 0    }]} {{562f87c9fde8dd1ff61370fd 0xc42038d090 0xc42038cf80 0xc42038cf90 0xc42038d030 0xc42038cf70 0xc42034ecc0} demo demo-symfony2 preprod webfront eu-west-1 t2.nano vpc-3f1eb65a [jrespaut@morea.fr] {ami-03ce4474 admin subnet-a7e849fe} {iam.ec2.demo.ghost ghost-demo [] {0 } [sg-6814f60c] [subnet-a7e849fe]} [{php5 5.4} {apache2 2.2} {php5-mysqlnd 0} {php5-curl 0} {php5-dev 0} {php5-xmlrpc 0} {php5-gd 0} {php5-memcache 0} {php5-redis 0} {pkg libapache2-mod-php5_openssh-server/wheezy-backports_openssh-client/wheezy-backports} {php5-zendopcache 0} {php5-intl 0}] [{0xc42038cfd0 symfony2 https://github.com/symfony/symfony-demo code /var/www 0 0  ZWNobyAiaSdtIG9uIHByZXByb2QiCmV4aXQgMA== IyEvYmluL2Jhc2gKCmV4aXQgMA== } {0xc42038cff0 test_utf8 git@github.com:smartholiday/weekendesk-mobile.git system /tmp/test_utf8 0 0    } {0xc42038d01c dummy_3 git@bitbucket.org:morea/ghost.dummy.git system /tmp/dummy3 0 0   ZXhpdCAx }]} {{5773a61efde8dd1cf763f52b 0xc42038d200 0xc42038d1f0 0xc42038d150 0xc42038d198 0xc42038d130 0xc42034eea0} demo build-prezto test bastion eu-west-1 t2.nano vpc-3f1eb65a [jrespaut@morea.fr] {ami-2d0db25e admin subnet-a7e849fe} {iam.ec2.demo.ghost ghost-demo [] {0 } [sg-6814f60c] [subnet-a7e849fe]} [{pkg vim_git}] [{0xc42038d190 symfony2 https://github.com/symfony/symfony-demo code /var/www 0 0  ZWNobyAiaSdtIG9uIHByZXByb2QiCmV4aXQgMA== Y29tcG9zZXIgaW5zdGFsbCAtLW5vLWludGVyYWN0aW9u }]}]}

Testing Ghost client create app
App created: 5839c643fde8dd2de0cf6bf9

Testing Ghost client get single app
App retrieved: 5839c643fde8dd2de0cf6bf9
{{5839c643fde8dd2de0cf6bf9 0xc420360280 0xc420360160 0xc420360190 0xc4203601e0 0xc420360108 0xc4202fc150} demo test-C631CB6B-58FA-9799-08A1-8F8A29F8C9E9 test webfront eu-west-1 t2.nano vpc-123456 [ghost-demo@domain.com] {ami-123456 admin subnet-123456} {test-instance-profile test-key-name [] {0 /dev/xvda} [sg-123456] [subnet-123456]} [{nginx 1.10}] [{0xc4203601a8 testmod git@bitbucket.org/morea/testmod system /tmp/path 0 0    }]}

Testing Ghost client update app
App updated: 5839c643fde8dd2de0cf6bf9

Testing Ghost client delete app
App deleted: 5839c643fde8dd2de0cf6bf9

PASS
ok      bitbucket.org/morea/go-st       2.598s
```
