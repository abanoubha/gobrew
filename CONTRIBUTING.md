# Contributing to gobrew

Clean the cache(s) of build process in snapcraft.

```sh
snapcraft clean
```

Build snap packages for architectures you specified inside the snapcraft.yaml file.

```sh
snapcraft
```

Install the generated (packaged) snap on your computer to test it.

```sh
sudo snap install gobrew_<version>_<architecture>.snap --devmode --dangerous

# for example
sudo snap install gobrew_25.03.05_amd64.snap --devmode --dangerous
```
