# Contributing to gochlog #

This is a short guide on how to contribute things to gochlog.

## Reporting a bug ##

Bug reports are welcome.  Check your issue exists with the latest
version first. Please add when submitting:

  * gochlog version (eg output from `gochlog version`)
  * Which OS you are using and how many bits (eg Windows 7, 64 bit)
  * The command you were trying to run (eg `gochlog`)
  * A log of the command (eg output from `GOCHLOG_LOG_LEVEL=debug gochlog`)
    * if the log contains secrets then edit the file with a text editor first to obscure them

## Submitting a pull request ##

If you find a bug that you'd like to fix, or a new feature that you'd
like to implement then please submit a pull request via Github.

If it is a big feature then make an issue first so it can be discussed.

You'll need a Go environment set up with GOPATH set.  See [the Go
getting started docs](https://golang.org/doc/install) for more info.

First in your web browser press the fork button on [gochlog's Github
page](https://github.com/shad7/gochlog).

Now in your terminal

    go get github.com/shad7/gochlog
    cd $GOPATH/src/github.com/shad7/gochlog
    git remote rename origin upstream
    git remote add origin git@github.com:YOURUSER/gochlog.git

Make a branch to add your new feature

    git checkout -b my-new-feature

And get hacking.

When ready - run the unit tests for the code you changed

    make test

The same way tests are run by Travis when you make a pull request but
you can do this yourself locally too.

Make sure you

  * Add documentation for a new feature
  * Add unit tests for a new feature
  * squash commits down to one per feature
  * rebase to master `git rebase master`

When you are done with that

    git push origin my-new-feature

Go to the Github website and click [Create pull
request](https://help.github.com/articles/creating-a-pull-request/).

You patch will get reviewed and you might get asked to fix some stuff.

If so, then make the changes in the same branch, squash the commits,
rebase it to master then push it to Github with `--force`.

## Test / Build / Cross-Compile

gochlog's tests are run from the go testing framework, so at the top
level you can run this to run all the tests.

**Assumes that you have Docker Toolbox / Docker for Mac / Docker for Windows installed.**

```bash
# make docker available locally
eval "$(docker-machine env)"
# runs all tests (includes formatting and linting)
make test
# run all tests and generates code coverage (includes formatting and linting)
make cover
# builds the default binary (linux amd64); (includes formatting and linting)
make build
# generates binaries for multiple different OS and Archicture combinations; (test run first)
make xcompile
```

#### Setting up Docker

If you are using Docker Toolbox and need to create a mount for the code follow this:
http://stackoverflow.com/questions/30864466/whats-the-best-way-to-share-files-from-windows-to-boot2docker-vm/30865500#30865500

To make it easy the name of the mount should be called `gochlog` and the directory should be created and mounted at `/gochlog`.

```bash
VBoxManage sharedfolder add default --name gochlog --hostpath <localpath> --automount
docker-machine ssh
sudo mkdir -p /gochlog
sudo mount -t vboxsf gochlog /gochlog
```

## Adding New Dependency

**If prompted to use specific versions, always select "Yes"**

```bash
DEP=<package> make add-dep
```

#### Example

```bash
DEP=github.com/Sirupsen/logrus make add-dep
```

## Making a release ##

[Release](RELEASE.md)

## Writing a new style ##

**Coming Soon**

