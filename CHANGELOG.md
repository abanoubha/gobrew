# CHANGE LOG

This page shows the roadmap with versioned tasks along with releases.

## 24.09.07

- get all Homebrew Core formulas
- save core_formulas as a file onto the disk
- get each package JSON file
- get count of packages which are written/built in Go language
- ability to set the language or build system
- include **dependencies** in calculation
- include **build dependencies** in calculation
- include **test dependencies** in calculation
- include **recommended dependencies** in calculation
- include **optional dependencies** in calculation
- count all versions of the language by default ([commit](https://github.com/abanoubha/gobrew/commit/7de9e76c03401ce70568417db550eda590bff919))
- re-download Homebrew/Core formulae index JSON file if the local one is older than 7 days ([commit](https://github.com/abanoubha/gobrew/commit/2a9713b90dd319203ec7692df81fb6c8e5759277))
- show all languages and count of their packages depend on them
- show all packages depends on specific language

## 25.03.03

- gobrew snap package support x64/amd64, arm64, and riscv64 architectures
- release all archs for stable channel

## 25.03.05

- fix: use latest codebase in the snap package

## 25.03.06

- set the minimum Go version to 1.24
- upgrade Cobra lib/dep

## 25.03.07

- support all platforms & archs by letting snapcraft build dashboard handle the build process

## next

- output SVG chart of specified programming languages
- get statistics from APT package manager too
