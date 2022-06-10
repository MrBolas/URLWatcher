# URLWatcher
URL watcher built in Go

## What is URL Watcher
Go CLI program that allows monitorization of last modifications of an URL.
The program will launch individual Watchers (Go routines) and periodically sends a HEAD request to the target url.
The last date where the website was changed is on the response header "last-modified". 

## Available Commands
| Command      | Description |
| ----------- | ----------- |
| load \<arg\>  | Takes as an argument a path to a file with urls in each line and loads it into the url watcher       |
| add \<args\>  | Takes as argument one or more urls and adds them to the url watcher, alternative to load        |
| ls   | lists all the created watchers with their id, status, urls and last modified date        |
| stop \<arg\> | Takes as an argument the id of the watcher and stops it |
| start all   | starts all the watchers  |
| start \<args\> | Takes as an argument the ids of the watchers and starts them |

### Examples
#### Startup
```
----------URL Watcher Shell-----------
-> 
```

#### Load file
```
----------URL Watcher Shell-----------
-> load testfolder/links
-> 
```

#### Add Watcher to url
```
----------URL Watcher Shell-----------
-> add http://www.google.com
-> 
```

#### List unactive Watchers 
```
-> ls
-------------------------------------------------- Watchers --------------------------------------------------
ID                                      status          url                     last modified
4ff4beda-b166-48f8-b047-2f6816cfb25a  Not Watching  http://www.google.com 
6f3df36b-cd1d-486b-b4a8-78a3867b1c85  Not Watching  https://manpages.debian.org/ 
69cc19e1-511e-4b08-89e1-ecf3cab275e3  Not Watching  http://www.google.com 
098e1216-3f58-4dfa-b1e6-deb64fd4acd1  Not Watching  http://example.com/ 
fafca455-ec32-48b6-9e50-ebbf89a6e766  Not Watching  http://zealwebtech.com/ 
b1dd8095-eaf0-4a72-aede-796e5a0d1cd2  Not Watching  http://techvynsys.com/v2/ 
-> 
```

#### List active Watchers 
```
-> start all
-> ls
-------------------------------------------------- Watchers --------------------------------------------------
ID                                      status          url                     last modified
23d1b057-949a-472b-91e1-ed56be4b26a0  Watching  https://manpages.debian.org/ Thu, 09 Jun 2022 21:38:57 GMT
db2a001c-c343-40d2-9059-fee7e63c2d19  Watching  http://www.google.com 
dbf3cc3b-35b9-4b52-9c5d-5d740817068b  Watching  http://example.com/ Thu, 17 Oct 2019 07:18:26 GMT
a0581acb-a8a4-468e-983c-9fc5bf819ed4  Watching  http://zealwebtech.com/ Thu, 15 Sep 2016 11:22:10 GMT
2f589fd2-be81-46fc-87ff-5c6ab43fa052  Watching  http://techvynsys.com/v2/ Sat, 09 Jul 2016 18:43:44 GMT
-> 
```

#### Start/Stop Watcher by Id 
```
----------URL Watcher Shell-----------
-> add http://www.google.com
-> ls
-------------------------------------------------- Watchers --------------------------------------------------
ID                                      status          url                     last modified
9fe48c02-4f08-4da3-a0c6-16698777b048  Not Watching  http://www.google.com 
-> start 9fe48c02-4f08-4da3-a0c6-16698777b048
-> ls
-------------------------------------------------- Watchers --------------------------------------------------
ID                                      status          url                     last modified
9fe48c02-4f08-4da3-a0c6-16698777b048  Watching  http://www.google.com 
-> stop 9fe48c02-4f08-4da3-a0c6-16698777b048
ls
2022/06/10 02:07:41 Watcher 9fe48c02-4f08-4da3-a0c6-16698777b048 was terminated
-> -------------------------------------------------- Watchers --------------------------------------------------
ID                                      status          url                     last modified
9fe48c02-4f08-4da3-a0c6-16698777b048  Not Watching  http://www.google.com 
-> 
```