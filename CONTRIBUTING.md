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

Publish the snap package into the snap store.

```sh
# create account on snapcraft.io , register a snap package listing, login in terminal
snapcraft login

# publish snap package as latest/edge
snapcraft upload --release=edge gobrew_x.y.z_amd64.snap

# publish snap package as latest/beta
snapcraft upload --release=beta gobrew_x.y.z_amd64.snap

# publish snap package as latest/candidate
snapcraft upload --release=candidate gobrew_x.y.z_amd64.snap

# publish snap package as latest/stable
snapcraft upload --release=stable gobrew_x.y.z_amd64.snap
```

Change the state of the release version (promote):

```sh
# from latest/edge to latest/beta
snapcraft promote --from-channel edge --to-channel beta gobrew

# from latest/beta to latest/candidate
snapcraft promote --from-channel beta --to-channel candidate gobrew

# from latest/candidate to latest/stable
snapcraft promote --from-channel candidate --to-channel stable gobrew

# from latest/edge to latest/stable
snapcraft promote --from-channel edge --to-channel stable gobrew

# from latest/beta to latest/stable
snapcraft promote --from-channel beta --to-channel stable gobrew
```

You can just trigger a new build in the [snapcraft management dashboard](https://snapcraft.io/gobrew/builds), so you do not need anything like packaging nor uploading nor promoting.
