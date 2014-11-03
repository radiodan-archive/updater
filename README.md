Updater
===

A tiny update manager.

## Purpose

Updater checks whether an app (specified by a small JSON file) has any pending updates (specified by a remote API). If there are updates pending, `updater` will download those updates, validate the download and install them into place.

## Filesystem layout

### Apps

`updater` expects apps to exist in the filesystem in a specific way.

A single invocation of `updater check` will check a specific directory on the filesystem, given by the `--target` argument.

Apps must exist in a subdirectory directoy below the target

An app  must contain a `releases` directory containing a subdirectory for each release. A symlink named `current` points to the current release directory.

For example:

    /opt/apps
      |
      +— magic-button
      |     |
      |     +— current -> ../releases/123451234
      |     |
      |     +— releases
      |         |
      |         +— 123451234
      |         |    |
      |         |    +- .deploy
      |         |
      |         +— 123451231
      |             |
      |             +- .deploy
      |
      +— not-a-valid-app
            |
            +— src
            +— bin

`updater` will ignore directories that do not match this layout.

### .deploy file

Each release must contain a `.deploy` file containing the following JSON:

    {
      "name": "radiodan/magic-button",
      "ref": "master",
      "commit": "d0ab003502b32510f8acd1c94b4b41664be6884"
    }

### Workspace

`updater` manages downloads in a workspace directory that you specify. This workspace should be the same on each invocation.

The workspace contains:

  - `downloads` downloaded packages
  - `manifests` JSON files specifying downloaded files
  - `status.json` info about `updater`'s workings

## Remote API

`updater` compares the installed list of releases with a remote API.

The API should contain a single endpoint that returns JSON about the latest releases.

For example:

    {
      "radiodan/magic-button": {
        "master": {
          "commit": "68200ec154d6dbdcc648e5706bb860c57b7fb621",
          "file": "http://example.com/radiodan/magic-button/magic-button-master-68200e.tar.gz",
          "sha1": "0e8ff998e905a2cfa5a4d15afe68622e3db24b72",
          "updated": "2014-10-29T19:30:54.095Z"
        },
        "radiotag": {
          "commit": "307c9c98dd229ba770ee3ad025ffa069bc860e61",
          "file": "http://example.com/radiodan/magic-button/magic-button-radiotag-307c9c.tar.gz",
          "sha1": "7f6566b8708a5f966cf6bfc9d125bd8365314bdb",
          "updated": "2014-10-09T14:14:17.322Z"
        }
      },
      "radiodan/physical-ui": {
        "master": {
          "commit": "d377b9a030c61ef803bc8ff8a5c8732b1d655427",
          "file": "http://example.com/radiodan/physical-ui/physical-ui-master-d377b9.tar.gz",
          "sha1": "6e149e6a098e30ce87adf87e36279637ccd42e2e",
          "updated": "2014-10-20T12:39:02.394Z"
        }
      }
    }

This JSON represents two applications called `radiodan/magic-button` and `radiodan/physical-ui`. `radiodan/magic-button` contains two refs called `master` and `radiotag`.

## Checking for updates

    $ updater check --target=/opt/apps --workspace=/opt/updates

1. `updater` scans target directory for a list of apps with `.deploy` files
2. Requests the latest versions of apps from the remote API
3. Finds apps that have matching names and refs
4. Compares latest commits against deployed commits
5. If different, downloads the archive specified by `file` to the workspace
6. Terminates

## Installing downloaded updates

    $ updater install --target=/opt/apps --workspace=/opt/updates

1. `updater` scans workspace directory for downloaded updates
2. Validates each archive against it's SHA1
3. Installs as a new release in the app's directory
4. Symlinks the new release as `current`

## Installation

`updater` is written in Go and must be compiled for each target platform. If you have a binary, just copy the single `updater` binary into /usr/local/bin.

## License

See [LICENSE](LICENSE)

## Copyright

Copyright 2014 British Broadcasting Corporation
