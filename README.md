# Git Team 
![Build Status](https://img.shields.io/github/workflow/status/jodstrcil/git-team/Go?style=flat-square)

### CLI for collaboration

<img src="https://cdn-images-1.medium.com/max/1600/0*gwlwlr7qMuwOYvnS" alt="Go team" width="200"/>

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
 
 In the *config.yml* set your teammates names, you can optionally add the prefix of the jira ticket for your team: 
 For example: 

```
---
users:
  - shortname: emarley
    fullname: Elaine Marley
    email: emarley@monkeyisland.com
  - shortname: gthreepwood
    fullname: Guybrush Threepwood
    email: gthreepwood@monkeyisland.com

jiratag: MYTEAM
```

 If you want to use this tool globally you need to export the path of you config file to environment variable GIT_TEAM_CONFIG
 
 ```
 echo 'export GIT_TEAM_CONFIG=<path_to_config_file' >> ~/.zshrc
 ```
 Once this is set you are ready to go: 

 ```
 gitteam commit -m "some message" -j <ticket_number> -p <collaborator_shortname> 
 ```
### Acknoledgement 
This is built using *urfave's* cli [github.com/urfave/cli](https://github.com/urfave/cli)

The gopher picture is taken from [configuring-googles-pixelbook](https://hackernoon.com/configuring-googles-pixelbook-to-contribute-to-go-2be955c21936)
