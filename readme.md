[![baby-gopher](https://raw2.github.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)

#LogFingJira

A simple Jira work-logging CLI tool written in go-lang.

##Installation 
+ Install go, get the detailed documentation for installation [here](http://golang.org/doc/install).
+ make sure GOPATH and GOROOT environment variables are set correctly.
+ type "go install github.com/mantishK/logfingjira" on your terminal.

##Documentation
Set the environment variables as follows - 

+ JUNAME - Your Jira username.
+ JCOMPNAME - Name of your company as it appears in the jira url (eg - If http://www.XXXXX.atlassian.com is the url you use for Jira website, you need to set XXXXX part only as the value of this environment variable).

### Flags
+ m - The jira log message flag.
+ d - The Jira interval/duration (e.g 1h, 1d, etc) flag.
+ i - The Jira issue key (e.g AAP-420) flag.

###Example Usage
```  logfingjira -m "pretended to work on this issue and resigned" -i "AAP-420" -d "1h"```

##Contributing
Suggestions and pull requests are welcome.

##Credits

+ [crypto/ssh](https://code.google.com/p/go/source/browse?repo=crypto#hg%2Fssh)
+ [Jira API](https://www.atlassian.com/software/jira)
