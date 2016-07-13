
# Installation

## Compiled version

*mergepath* is a command line program run from a shell like Bash. If you download the
repository a compiled version is in the dist directory. The compiled binary matching
your computer type and operating system can be copied to a bin directory in your PATH.

Compiled versions are available for Mac OS X (amd64 processor), Linux (amd64), Windows
(amd64) and Rapsberry Pi (both ARM6 and ARM7)

### Mac OS X

1. Download **mergepath-binary-release.zip** from [https://github.com/rsdoiel/mergepath/releases/latest](https://github.com/rsdoiel/mergepath/releases/latest)
2. Open a finder window, find and unzip **mergepath-binary-release.zip**
3. Look in the unziped folder and find *dist/macosx-amd64/mergepath*
4. Drag (or copy) *mergepath* to a "bin" directory in your path
5. Open and "Terminal" and run `mergepath -h` to confirm you were successful

### Windows

1. Download **mergepath-binary-release.zip** from [https://github.com/rsdoiel/mergepath/releases/latest](https://github.com/rsdoiel/mergepath/releases/latest)
2. Open the file manager find and unzip **mergepath-binary-release.zip**
3. Look in the unziped folder and find *dist/windows-amd64/mergepath.exe*
4. Drag (or copy) *mergepath.exe* to a "bin" directory in your path
5. Open Bash and and run `mergepath -h` to confirm you were successful

### Linux

1. Download **mergepath-binary-release.zip** from [https://github.com/rsdoiel/mergepath/releases/latest](https://github.com/rsdoiel/mergepath/releases/latest)
2. Find and unzip **mergepath-binary-release.zip**
3. In the unziped directory and find for *dist/linux-amd64/mergepath*
4. Copy *mergepath* to a "bin" directory (e.g. cp ~/Downloads/mergepath-binary-release/dist/linux-amd64/mergepath ~/bin/)
5. From the shell prompt run `mergepath -h` to confirm you were successful

### Raspberry Pi

If you are using a Raspberry Pi 2 or later use the ARM7 binary, ARM6 is only for the first generaiton Raspberry Pi.

1. Download **mergepath-binary-release.zip** from [https://github.com/rsdoiel/mergepath/releases/latest](https://github.com/rsdoiel/mergepath/releases/latest)
2. Find and unzip **mergepath-binary-release.zip**
3. In the unziped directory and find for *dist/raspberrypi-arm7/mergepath*
4. Copy *mergepath* to a "bin" directory (e.g. cp ~/Downloads/mergepath-binary-release/dist/raspberrypi-arm7/mergepath ~/bin/)
    + if you are using an original Raspberry Pi you should copy the ARM6 version instead
5. From the shell prompt run `mergepath -h` to confirm you were successful


## Compiling from source

If you have go v1.6.2 or better installed then should be able to "go get" to install all the **mergepath** utilities and
package. You will need the GOBIN environment variable set. In this example I've set it to $HOME/bin.

```
    GOBIN=$HOME/bin
    go get github.com/rsdoiel/mergepath/...
```

or

```
    git clone https://github.com/rsdoiel/mergepath src/github.com/rsdoiel/mergepath
    cd src/github.com/rsdoiel/mergepath
    make
    make test
    make install
```

