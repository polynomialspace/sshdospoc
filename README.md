# Usage
`./sshdos -t <host>[:port]`

Options:
```
  -c int
    	number of connections to initiate (default 100)
  -t string
    	target "host" or "host:port" (required)
  -u string
    	user to ssh as; probably doesn't matter (default "root")
```

# About

Very basic proof of concept for SSH denial of service.

Related OpenSSH Settings: `MaxStartups`, `LoginGraceTime`, `MaxAuthTries`. (`PerSourceMaxStartups` as of [8.5](https://www.openssh.com/txt/release-8.5))


# Notes
setting `MaxStartups` to larger values (eg. >1000) may cause high memory usage (>1GB) due to `fork()`ing in the connection handler.

[fail2ban](https://www.fail2ban.org/wiki/index.php/Main_Page) probably helps too.