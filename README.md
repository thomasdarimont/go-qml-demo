Demo App with Go and go-qamel (QT/QML)
--- 

Simple demo app for [go-qamel](https://github.com/go-qamel/qamel)

# Setup
Follow the [instructions from the gamel project](https://github.com/go-qamel/qamel/wiki/Installation)

Example qamel profile:  
```
$ qamel profile setup
Thanks for using qamel, QML's binding for Go.

Please specify the target OS for this profile. Possible values are "windows", "linux" and "darwin". 
Keep it empty to use your system OS.

Target OS (default linux) : 

Please specify the target architecture for this profile. Possible values are "386" and "amd64".
Keep it empty to use your system architecture.

Target arch (default amd64) : 

Please specify whether this profile used to build static or shared app.

Use static mode (y/n, default n) : n

Please specify the *full path* to your Qt's tools directory.
This might be different depending on your platform or your target. For example, in Linux with Qt 5.11.1, the tools are located in $HOME/Qt5.11.1/5.11.1/gcc_64/bin/

Qt tools dir : /home/tom/dev/tools/Qt/5.14.1/gcc_64/bin
qmake        : found
moc          : found
rcc          : found

Please specify the *full path* to your compiler.
Keep it empty to use the default compiler on your system.

C compiler (default gcc)   : 
C++ compiler (default g++) : 
Saving profile default...done

Setup finished.
Now you can get started on your own QML app.
```

## Install go dependencies
```
go mod vendor
```

# Build the app
```
$ qamel build -p default
Starting build process.

Load config file...done
Generating vendor files...done
Removing old build files...done
Generating binding files...done
Generating Qt resource file...done
Generating code for QML objects...done
Building app...done

Build finished succesfully.
```

# Run
```
./go-qml-demo
```