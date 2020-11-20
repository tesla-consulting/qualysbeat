# qualysbeat

A TeslaConsulting Beat

### Configure

Put the qualysbeat root folder in your linux path system 
```
/usr/share/qualysbeat
```
take a `qualysbeat.yml` file from the root path and put him into
```
/etc/qualysbeat/
```
Configure the `qualysbeat.yml` with your favorite parameters

-----

Put the beater/qualys.conf file into `/etc/qualysbeat/` and set the file with client's credentials

```
{
	"qualys":{
		"user":"...";
		"password":"...";
		"cliente":"..."
	}
}
```
cliente is optional parameter

### Execute

execute command:

```
/usr/share/qualysbeat/qualysbeat -c /etc/qualysbeat/qualysbeat.yml -e
```
