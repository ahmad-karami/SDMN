# Minimal Container Runtime in Go

`Container.go` is a minimal container runtime that starts an interactive shell in an isolated environment. It uses Linux UTS, PID, mount, and network namespaces, a cgroup memory limit, and `chroot` into an Ubuntu 20.04 base filesystem. The program downloads the base filesystem archive with `wget` when it is missing and extracts it into a directory named after the container. It then starts a copy of itself in the new namespaces, and that child process sets up the container and runs `/bin/bash`.

## How it works
Run as root:

```bash
go run Container.go <container_name> <memory_limit_in_MB>
```

`extractFile` calls `file_sys`, which downloads the archive into a hardcoded directory if it is missing, and then extracts it into `<container_name>`. `namespace` re-runs the program through `/proc/self/exe` with the `child` argument and new namespaces. In the child, `Cgroup` creates `/sys/fs/cgroup/<container_name>`, writes the limit to `memory.max`, and adds the process to `cgroup.procs`. `child` then sets the hostname, prints `ip addr`, runs `chroot`, mounts `/proc`, sets a `PS1` prompt with the hostname, and starts `/bin/bash`. It unmounts `/proc` when the shell exits.
