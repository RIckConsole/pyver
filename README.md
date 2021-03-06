# pyver

pyver makes use of update-alternatives to make python versions simple!

pyver will help you easliy set the `python` binary to whichever python version you have installed

## Installation

If you have your go environment setup correctly you can use this one-liner:

``` sh
go get github.com/RIckConsole/pyver && cd ~/go/src/github.com/RIckConsole/pyvar && go install
```

## Usage

pyvar has 4 functions, `set`, `check`, `add`, and `list`

### Check

Check allows you to see the enabled python version

```sh
$ pyver check
/usr/bin/python2.7
```

### Add

Before you can switch to a different python version, you have to `add` it as an alternative

``` sh
$ pyver add python3
python version added successfully!
```

Note: The python version you are adding must be properly installed (in `/usr/bin/`)

### Set

Once you have your python versions added as alternatives, you can switch to them at any time by using `set` (requires root)

``` sh
$ sudo pyver set python3
Python version switched to: python3
```

### List 

If you are unsure which python versions you have added as alternatives, you can use the list function to view them:

``` sh
$ pyver list
```
