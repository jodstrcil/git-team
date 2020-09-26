# Git Team :rainbow::handshake::rainbow:

![Build Status](https://img.shields.io/github/workflow/status/jodstrcil/git-team/Go?style=flat-square)

### CLI for pairing


Small cli tool to format your commit messages including your collaborators, and the ticket you are working on.

## How to install it

Clone the project locally. Then build the code.
```
go build
```

You can install it to make it accessible from anywhere. 
```
go install 
```

That's it! :tada:

**GOPATH**

Make sure your `PATH` includes the `$GOPATH/bin` directory, so your commands can
be easily used:
```
export PATH=$PATH:$GOPATH/bin
```
 
 ## How to use it
 
 In the *config.yml* set your teammates names, you can optionally add the prefix of the jira ticket for your team. 
 Once this is set: 
 
 ```
 gitteam commit -m "some message" -j <ticket_number> -p <collaborator_shortname> 
 ```
### Acknoledgement 
This is built using *urfave's* cli [github.com/urfave/cli](https://github.com/urfave/cli)