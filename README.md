# tagext
Tag Extractor<br>
install<br>
```➜  ~ go get -u -v github.com/kawakatz/tagext```

usage<br>
- Extract the contents of the Anchor tag<br>
```➜  ~ curl -s https://github.com/ | tagext -tag a```
- Extract the contents of the Anchor tag as html (You can give the output to this tool again)<br> 
```➜  ~ curl -s https://github.com/ | tagext -tag a -html true```
- Extract the href attributes of the Anchor tag<br>
```➜  ~ curl -s https://github.com/ | tagext -tag a -attr href```
