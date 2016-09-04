pick
====
A minimal password manager for OS X and Linux.

![demo](https://github.com/bndw/pick/raw/master/demo.gif)

## Install
```sh
$ go get github.com/bndw/pick
```

or

```sh
$ git clone https://github.com/bndw/pick && cd pick
$ ./build
$ ./install
```


## Usage
```
Usage:
  pick [command]

Available Commands:
  add         Add a credential
  cat         Cat a credential
  cp          Copy a credential to the clipboard
  export      Export decrypted credentials in JSON format
  ls          List all credentials
  rm          Remove a credential

Flags:
  -h, --help   help for pick

Use "pick [command] --help" for more information about a command.
```

## Similar software
* [pwd.sh: Unix shell, GPG-based password manager](https://github.com/drduh/pwd.sh)
* [Pass: the standard unix password manager](http://www.passwordstore.org/)
