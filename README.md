
# Docker Machine Fusion Driver

Create Docker machines locally on VMware Fusion(https://www.vmware.com/products/fusion).

This driver requires VMware VMware Fusion 10 (MacOS)
to be installed on your host. Earlier versions of Fusion might still work
with this driver, but it's not officially supported.

> **The main difference is:**

> Support for macOS Big Sur and fusion 12.1

> Extracted from docker-machine project

> Added dual-network card function , separates public and private networks

## Installation

### From a Release

The latest version of the `docker-machine-driver-fusion` binary is available on the
[GithHub Releases](https://github.com/zhaog918/docker-machine-driver-fusion/releases) page.
Download the the binary that corresponds to your OS into a directory residing in your PATH.

### From Source

Make sure you have installed [Go](http://www.golang.org) and configured [GOPATH](http://golang.org/doc/code.html#GOPATH)
properly. For MacOS and Linux, make sure `$GOPATH/bin` is part of your `$PATH` for MacOS and Linux.
For Windows, make sure `%GOPATH%\bin` is included in `%PATH%`.

Run the following command:

```shell
go get -u github.com/zhaog918/docker-machine-driver-fusion
```


## Usage

```shell
$ docker-machine create --driver=fusion default
```


## Options

- `--fusion-boot2docker-url`: URL for boot2docker image
- `--fusion-configdrive-url`: URL for cloud-init configdrive
- `--fusion-cpu-count`: Number of CPUs for the machine (-1 to use the number of CPUs available)
- `--fusion-disk-size`: Size of disk for host VM (in MB)
- `--fusion-memory-size`: Size of memory for host VM (in MB)
- `--fusion-no-share`: Disable the mount of your home directory
- `--fusion-ssh-password`: SSH password
- `--fusion-ssh-user`: SSH user

#### Environment variables and default values

| CLI option                 | Environment variable   | Default                  |
|----------------------------|------------------------|--------------------------|
| `--fusion-boot2docker-url` | VMFUSION_BOOT2DOCKER_URL | *Latest boot2docker url* |
| `--fusion-configdrive-url` | VMFUSION_CONFIGDRIVE_URL | -                        |
| `--fusion-cpu-count`       | VMFUSION_CPU_COUNT       | `1`                      |
| `--fusion-disk-size`       | VMFUSION_DISK_SIZE       | `20000`                  |
| `--fusion-memory-size`     | VMFUSION_MEMORY_SIZE     | `1024`                   |
| `--fusion-no-share`        | VMFUSION_NO_SHARE        | -                        |
| `--fusion-ssh-password`    | VMFUSION_SSH_PASSWORD    | `tcuser`                 |
| `--fusion-ssh-user`        | VMFUSION_SSH_USER        | `docker`                 |


## License

It's under the Apache 2 license.
