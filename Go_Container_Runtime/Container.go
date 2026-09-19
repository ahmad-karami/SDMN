package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

// Usage: Container.go <container_name> <memory_limit_in_MB>
// The program re-runs itself with "child" as the first argument to start the container process.
func main() {
	switch os.Args[1] {
	case "child":
		child(os.Args[2], os.Args[3])
	default:
		extractFile(os.Args[1])

		namespace(os.Args[1])
	}
}

// namespace starts the child process that becomes the container.
func namespace(Container string) {
	fmt.Printf("Running main %v as %d\n", os.Args[1:], os.Getpid())

	// /proc/self/exe is this binary, so the child enters the "child" branch of main.
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[1:]...)...)

	// Attach the child to this terminal.
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// Start the child in new UTS, PID, mount, and network namespaces.
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET,
		// This also makes Go remount / as private, so mounts in the child do not propagate to the host.
		Unshareflags: syscall.CLONE_NEWNS,
	}
	cmd.Run()

}

// child runs inside the new namespaces. It applies the memory limit, sets the hostname,
// switches to the container's root filesystem, and starts a shell.
func child(Container string, M_max string) {
	fmt.Printf("Running child %v as %d\n", os.Args[1:], os.Getpid())

	Cgroup(Container, M_max)
	// The new UTS namespace lets the container have its own hostname.
	syscall.Sethostname([]byte(Container))
	Ip()
	// Switch to the extracted root filesystem and mount a /proc for the new PID namespace.
	syscall.Chroot(Container)
	syscall.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")
	os.Setenv("PS1", "root@\\h:\\w$ ")
	cmd := exec.Command("/bin/bash")

	// Attach the shell to this terminal.
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Wait for the shell to exit, then remove the /proc mount.
	cmd.Run()

	syscall.Unmount("/proc", 0)

}

// Cgroup creates a cgroup named after the container, sets its memory limit to M_max megabytes,
// and moves the current process into it, so the shell started later inherits the limit.
func Cgroup(Container, M_max string) {

	dir := "/sys/fs/cgroup/" + Container

	os.Mkdir(dir, 0777)

	ioutil.WriteFile(filepath.Join(dir, "/memory.max"), []byte(M_max+"M"), 0700)
	ioutil.WriteFile(filepath.Join(dir, "/notify_on_release"), []byte("1"), 0700)
	ioutil.WriteFile(filepath.Join(dir, "/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700)

}

// file_sys downloads the Ubuntu 20.04 base filesystem archive if it is not already present.
func file_sys() {
	path := "/home/mr_king/projects/SDMN/"

	file_path := "/home/mr_king/projects/SDMN/ubuntu-base-20.04.2-base-amd64.tar.gz"
	file_url := "http://cdimage.ubuntu.com/ubuntu-base/releases/20.04/release/ubuntu-base-20.04.2-base-amd64.tar.gz"
	if _, err := os.Stat(file_path); os.IsNotExist(err) {
		fmt.Printf("ubuntu:20.04 Filesys does not exist\n")
		cmd := exec.Command("wget", "-P", path, file_url)
		cmd.Run()
		fmt.Printf("File downloaded\n")
	}
}

// extractFile unpacks the base filesystem into a directory named after the container.
func extractFile(Container string) {
	file_sys()
	os.Mkdir(Container, 0777)
	file_path := "/home/mr_king/projects/SDMN/ubuntu-base-20.04.2-base-amd64.tar.gz"

	cmd := exec.Command("tar", "xzvf", file_path, "-C", Container)
	cmd.Run()

	fmt.Printf("File %s extracted successfully!\n", Container)
}

// Ip prints the interfaces of the current network namespace.
func Ip() {
	cmd := exec.Command("ip", "addr")
	out, _ := cmd.Output()
	fmt.Println(string(out))
}
