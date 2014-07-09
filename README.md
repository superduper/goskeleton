goskeleton
==========

The project uses few tools for dependency management, so the build process requires few extra steps.

### Prerequisites

As you may have guessed:

* Go 1.3(or 1.2.1)

As the code requires a lot of external libraries some tools are preferred to handle them.

For golang package management the cool but complex way is [gpm](https://github.com/pote/gpm) and [gvp](https://github.com/pote/gvp).

* GPM https://github.com/pote/gpm
* GPM-GIT https://github.com/babo/gpm-git
* GVP https://github.com/pote/gvp (optional, but preferred)

A simple alternative is the included `bin/getdeps.sh` scripts which does the same without any requirement.

### Building from source


#### 1. Prepare environment
Initialize this directory at the first use of gvp here.

```
gvp init
```

#### 2. Get dependencies with GPM
Following will automatically get all dependencies.

```
make deps
```

#### 3. Compile the binary

```
make compile
```
