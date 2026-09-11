# Lesson 5: Processes, Containers, and Isolation

## Objective

Understand what container isolation actually provides (and its real limits) at the OS level — namespaces and capabilities — and apply least-privilege thinking (system-engineering Lesson 10) concretely to container configuration.

## Prerequisites

OS Lesson 1 (processes and the user/kernel boundary — containers are built entirely from OS-level primitives, not a separate virtualization technology), system-engineering Lesson 10 (least privilege — this lesson is that principle's specific application to container configuration).

## Learn

**Containers are not virtual machines — a genuinely important distinction, not just terminology.** A VM virtualizes hardware, running a completely separate OS kernel per VM, with strong isolation enforced by the hypervisor. A container is fundamentally a set of Linux kernel features — namespaces and cgroups — applied to an ordinary process, running on the **same kernel** as the host and every other container on that host. This matters directly: a kernel-level vulnerability can potentially be exploited to escape container isolation and affect the host or other containers, in a way that's structurally much harder (though not literally impossible) for a properly configured VM's hypervisor boundary.

**Namespaces: what makes a containerized process see a restricted view of the system.** Linux namespaces (PID, network, mount, user, and others) each restrict what a process can *see* within that namespace category — a process in its own PID namespace sees only its own process tree (its own "PID 1"), not the host's full process list; a process in its own network namespace has its own network interfaces, isolated from the host's; a process in its own mount namespace has its own filesystem view, isolated from the host's actual filesystem outside what's explicitly mounted in. This is what makes a container feel like an isolated system, while it's actually just an ordinary Linux process (OS Lesson 1) with a restricted, namespaced view of shared kernel resources.

**Capabilities: fine-grained privilege, beyond the traditional root/non-root binary.** Traditional Unix permission is coarse: a process either runs as root (with essentially unrestricted privilege) or as a non-root user (with restricted privilege) — Linux capabilities break "root's privilege" into dozens of specific, individually grantable capabilities (e.g. `CAP_NET_BIND_SERVICE` for binding to low-numbered network ports, `CAP_SYS_ADMIN` for a wide range of administrative operations). A well-configured container drops every capability not specifically needed, rather than running as full root (or with the default capability set Docker provides, which is broader than most applications actually need) — this is least privilege (system-engineering Lesson 10), applied at the kernel-capability level specifically.

**The real limit of container isolation, worth understanding rather than assuming away.** Because containers share the host kernel, a container running as root (even within its own namespaced view) still has meaningfully more attack surface than a properly capability-restricted, non-root container — if a kernel vulnerability or a container-escape bug is found, a root-privileged container process has a much easier path to full host compromise than a correctly restricted one. This is exactly why "don't run containers as root, and drop unnecessary capabilities" isn't optional hardening advice for the paranoid, it's a real, meaningful reduction in the actual blast radius of a container-escape scenario.

## Attempt

1. Run a container (Docker or equivalent) and, from inside it, inspect `/proc/1/status` and process listing — confirm you see a restricted process tree (your container's own processes only, with your container's main process as PID 1), directly observing PID namespace isolation in action, distinct from what `ps aux` on the actual host machine would show.

2. Run the same container with `--network host` (disabling network namespace isolation) versus without it, and compare what network interfaces are visible from inside the container in each case (`ip addr` or equivalent) — confirm the isolated version shows only the container's own virtual network interface, while the host-networked version shows the host's actual interfaces directly, demonstrating namespace isolation's real, observable effect.

3. Run a container as root (the Docker default, unless explicitly configured otherwise) and inspect its capability set (`capsh --print`, or check `/proc/self/status`'s `Cap*` fields). Then run the same container with explicit capability dropping (`docker run --cap-drop=ALL --cap-add=<only what's needed>`) and compare the capability sets — confirm the restricted version genuinely has a smaller capability set, and test that an operation requiring a dropped capability (e.g. attempting to bind to a low-numbered port without `CAP_NET_BIND_SERVICE`, if your test app doesn't need it) correctly fails in the restricted container while succeeding in the default one.

4. Configure a container to run as a non-root user explicitly (via a `USER` directive in the Dockerfile, or `--user` flag) rather than the default root, and confirm — using the same capability inspection as step 3 — that running as non-root further reduces the effective privilege even before any explicit capability-dropping, directly demonstrating that "don't run as root" and "drop unnecessary capabilities" are two separate, complementary hardening steps, not the same thing.

## Verify

Present your step 1 process-tree comparison (container's restricted view vs. what the host actually shows), your step 2 network-namespace comparison, and your step 3-4 capability-set comparisons across default-root, capability-dropped, and non-root-user configurations, showing the concrete privilege reduction at each step.

## Failure drill

Take a capability-dropped, non-root container from steps 3-4 and attempt an operation that *should* be prevented by your restricted configuration but that you suspect might still be possible due to some capability or permission you didn't think to drop (e.g. writing to a mounted host directory, if your container has an unnecessarily broad volume mount) — actually test this, don't just assume your restrictions are complete. If you find something still works that shouldn't, given your intended security posture, identify specifically what additional restriction (a more limited volume mount, an additional dropped capability, a read-only root filesystem via `--read-only`) would be needed, and apply it. Explain why this kind of "test what you think is restricted, don't just configure and assume" verification step is necessary — configuring capability drops and namespaces correctly is easy to get subtly wrong, and only actually testing the specific operations you intend to prevent (not just checking that the configuration flags are present) gives real confidence the restriction works as intended.

## Transfer

If TARDOC or Mahall's deployment runs in Docker containers (a reasonable assumption given modern deployment practices, worth confirming for your actual setup), audit their actual container configuration against this lesson's hardening checklist: do they run as non-root, do they drop unnecessary capabilities, is the root filesystem read-only where practical, are volume mounts scoped as narrowly as possible. Report your actual findings, and for any gap found, describe what applying this lesson's specific hardening step would require for that real container.

## Done when

You've directly observed namespace isolation's effect on process visibility and network interfaces, you've compared and tested real capability differences between default-root, capability-dropped, and non-root container configurations, and you've verified — not just configured and assumed — that your hardening actually prevents the specific operations it's meant to prevent, finding and closing at least one gap via the failure drill's testing discipline.
