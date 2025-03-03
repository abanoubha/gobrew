# Contributing to gobrew

Make sure you installed [snapcraft](https://snapcraft.io/) first.

Clean the cache(s) of build process in snapcraft.

```sh
snapcraft clean
```

Make sure to created a tag version release in git system and publish it, as snapcraft gets the latest git tag and package it.

```sh
git tag -a 25.03.05 -m "25.03.05 release title, with short highlight"
git push --tags --all
```

Build snap packages for architectures you specified inside the snapcraft.yaml file.

```sh
snapcraft
```

Uninstall snap package.

```sh
sudo snap remove --purge gobrew
```

Install the generated (packaged) snap on your computer to test it.

```sh
sudo snap install gobrew_<version>_<architecture>.snap --devmode --dangerous

# for example
sudo snap install gobrew_25.03.05_amd64.snap --devmode --dangerous
```

Normal installation of upstream gobrew snap package.

```sh
sudo snap install gobrew
```
