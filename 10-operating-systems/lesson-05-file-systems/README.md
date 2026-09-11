# Lesson 5: File Systems

## Objective

Understand how a filesystem represents files and directories on disk — inodes, allocation strategies, metadata — and reason about crash consistency: what happens if power is cut mid-write.

## Prerequisites

Computer architecture Lesson 7 (storage performance — this lesson covers what's stored and how, building on that lesson's coverage of raw storage latency).

## Learn

**The inode: separating a file's identity from its name.** In Unix-style filesystems, a file's actual metadata (size, permissions, timestamps, and pointers to the data blocks holding its content) lives in a structure called an inode, identified by a number — not by its filename. A directory is just a special file mapping names to inode numbers. This separation is *why* a hard link works: two different filenames in possibly different directories can point to the exact same inode, meaning the exact same underlying data and metadata — genuinely the same file with two names, not two copies. `ls -i` shows a file's inode number directly, and `stat <file>` shows the full inode metadata.

**Allocation strategies.** A file's data typically doesn't sit in one contiguous disk region — the filesystem allocates it in blocks (fixed-size chunks, e.g. 4KB) that may be scattered across the disk, tracked via pointers stored in (or reachable from) the inode. Older/simpler designs used direct block pointers plus indirect pointer blocks (a pointer to a block that itself contains more pointers, extending how large a file's block list can grow without bloating every inode); modern filesystems (ext4, XFS) commonly use extents (a start block plus a length, describing a contiguous run in one entry) which are more efficient for large, mostly-contiguous files, directly connecting back to computer architecture Lesson 4's point about sequential access being cheaper — extents are partly a data structure optimized around that same physical reality.

**Crash consistency: the central hard problem.** A file write is rarely a single atomic operation from the filesystem's perspective — updating a file's content, its size metadata, and possibly directory entries can require multiple separate disk writes. If power is lost between some of these writes completing and others not, the filesystem can be left in an inconsistent state (e.g. a directory entry pointing at an inode whose data blocks were never actually written, or a file whose size metadata doesn't match its actual content). This is the exact problem journaling filesystems (ext4, most modern filesystems) solve: before making the real changes, the filesystem writes a compact log (the journal) describing the intended changes; after a crash, replaying the journal on next mount can bring the filesystem back to a consistent state without needing to fully rescan the entire disk.

**Why this matters for your actual work.** Database engines (PostgreSQL, which the postgresql track covers) build their own crash-consistency mechanisms (write-ahead logging, WAL) *on top of* the filesystem's own guarantees, precisely because a database transaction's atomicity requirements are stronger and more specific than what a general-purpose filesystem's journal alone guarantees — understanding filesystem-level crash consistency here is the conceptual foundation the database-internals track's recovery lesson builds on directly.

## Attempt

1. Create a file, then create a hard link to it (`ln original.txt hardlink.txt`, not `ln -s` which is a symlink, a genuinely different mechanism). Use `ls -i` to confirm both filenames report the identical inode number. Modify the content via one filename and confirm the change is visible through the other filename — direct proof they're the same underlying file, not a copy.

2. Use `stat <file>` on a file you create and modify, and identify the metadata fields that change after you edit its content (`mtime` — modification time) versus fields that only change on a metadata-only operation like `chmod` (`ctime` — inode change time, distinct from mtime, a commonly confused pair) versus a field that changes only when the file is read (`atime` — access time, though many modern systems mount filesystems with `noatime` for performance, meaning this may not update at all — check whether yours does).

3. Delete the original filename from step 1 (`rm original.txt`) while the hard link (`hardlink.txt`) still exists, and confirm the file's content is still fully accessible via the remaining link — demonstrating that a file's actual data persists until its **link count** (also visible via `stat`, or `ls -l`'s second column) drops to zero, not until any *specific* name is removed.

4. Research (via `man` pages or documentation, not required to implement) how your specific filesystem's journal mode works — for ext4 specifically, the three journaling modes (`journal`, `ordered`, `writeback`) offer different tradeoffs between consistency guarantees and performance. Write a short explanation, in your own words, of what specifically `ordered` mode (the common default) guarantees and does not guarantee about the relationship between metadata and data writes after a crash.

## Verify

For step 3, report the exact link count (from `stat` or `ls -l`) before deletion (should be 2, since two names point to the same inode) and after deleting one name (should drop to 1), confirming the data remained accessible the whole time via the surviving link.

## Failure drill

Attempt the equivalent experiment with a **symlink** instead of a hard link: create a symlink to a file, then delete the original file the symlink points to, and confirm accessing the symlink now fails (a "dangling" symlink) — unlike the hard link case, where the data survived because the inode itself was still referenced. Explain, using the inode-vs-filename distinction from Learn, exactly why these two link types behave so differently: a hard link is another name for the *same inode*, directly incrementing its link count, while a symlink is a separate small file whose content is just a path string, with no connection to the target's inode or link count at all — deleting the target leaves the symlink pointing at a name that no longer resolves to anything.

## Transfer

If TARDOC or Mahall ever performs multi-step file operations (e.g. writing a file, then updating a database record referencing it, or writing to a temp file then renaming it into place — a common atomic-update pattern, since a rename within the same filesystem is typically a single atomic metadata operation, unlike a direct overwrite), describe, using this lesson's crash-consistency concepts, what could go wrong if the process crashes between the two steps, and why the temp-file-then-rename pattern specifically is used to minimize (though not eliminate entirely, depending on journal mode) that window of inconsistency.

## Done when

You can explain the inode/filename separation and have directly demonstrated it with a hard link surviving the original filename's deletion, you can distinguish mtime/ctime/atime and state what specifically changes each one, and you can explain in your own words what a filesystem journal protects against and why database engines still need their own additional crash-consistency mechanism on top of it.
