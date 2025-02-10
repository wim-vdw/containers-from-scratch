# Containers from scratch in Go

## container1 (Simple Linux UTS Namespace Runner)

This program is a minimal container-like process runner for Linux.  
It executes commands in a new UTS namespace, isolating the hostname.

Requirements:

- Must be run on Linux.
- Requires root privileges.

## container2 (Simple Linux UTS Namespace Runner - spawns a new child process)

This program runs commands in a new UTS namespace, simulating a lightweight container by changing the hostname.  
The program spawns a new process and sets the hostname using the `sethostname` system call.  
When a process is created inside a namespace, all child processes it spawns will automatically belong to the same
namespace unless explicitly configured otherwise.

Requirements:

- Must be run on Linux.
- Requires root privileges.

## container3 (Minimal Linux Container Runner)

This program runs commands in a lightweight, isolated container using Linux namespaces.  
It creates a new UTS, PID, and mount namespace while using a chroot environment.

Requirements:

- Must be run on Linux.
- Requires root privileges.
- Needs a valid root filesystem (e.g., Alpine Linux) at `/home/wim/alpinefs`.

## container4 (Lightweight Container with Cgroup Limits)

This program runs commands in an isolated container-like environment using Linux namespaces and cgroups.  
It creates a new UTS, PID, and mount namespace while restricting process count via cgroups.

Requirements:

- Must be run on Linux.
- Requires root privileges.
- Needs a valid root filesystem (e.g., Alpine Linux) at `/home/wim/alpinefs`.
- Requires cgroups to be enabled (`/sys/fs/cgroup`).

Features:

- Isolated hostname (`container`).
- Separate PID and mount namespaces.
- Chrooted environment (`/home/wim/alpinefs`).
- `/proc` remounted inside the container.
- Cgroup limit: Max 10 processes.
