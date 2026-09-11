# Lesson 1: Pages, Records, and Storage

## Objective

Understand how a database turns rows into bytes on disk — fixed-size pages, record layout, and the overhead that comes with making data both queryable and durable — by designing and implementing your own minimal page format.

## Prerequisites

Computer architecture Lesson 7 (storage performance — a database's page size choice is directly informed by that lesson's I/O latency numbers), OS Lesson 5 (filesystems — a database page is the database engine's own layer built on top of filesystem blocks).

## Learn

**Why pages, not variable-length reads.** A database doesn't read individual rows from disk — it reads and writes fixed-size **pages** (commonly 4KB, 8KB, or 16KB, chosen to align with or be a multiple of the underlying filesystem/storage block size from computer architecture Lesson 7, so a single page read/write maps cleanly onto the storage layer's own I/O granularity rather than crossing block boundaries wastefully). Every row lives inside some page, and the database's buffer pool (Lesson 2) caches whole pages in memory, not individual rows — this fixed-granularity design is what makes buffer pool management tractable: uniform-size units are much simpler to allocate, evict, and track than arbitrary-size ones.

**Record layout within a page.** A page needs to store multiple variable-length records (rows can have different actual byte sizes, e.g. due to variable-length text fields) plus enough metadata to find and interpret them. A common design — the **slotted page** — stores a small, fixed-size "slot array" at the start of the page (each slot holding an offset and length pointing to where a record actually lives), with the records themselves packed from the *end* of the page growing backward, and free space in the middle. This lets records be added, removed, or resized without shifting every other record's position — only the slot array entry needs updating, which is exactly the kind of design tradeoff (a layer of indirection in exchange for cheaper mutation) that shows up repeatedly in systems design.

**Fixed-width vs. variable-width fields.** A record layout typically separates fixed-width fields (integers, fixed-length values — directly readable at a known byte offset within the record) from variable-width fields (text, blobs — requiring either an embedded length prefix, matching Lesson 1 of the networking track's framing concept applied to on-disk data instead of network data, or a pointer to variable-length storage elsewhere). This distinction is why adding a `VARCHAR` column to a table has historically been a heavier schema-migration operation than adding a fixed-width `INT` column in some database engines — the record layout itself has different implications for each.

**Storage overhead, concretely.** Every record needs some bookkeeping beyond its raw data: a slot array entry, possibly a null-value bitmap (tracking which nullable fields are actually null, since a NULL isn't stored as data at all), and per-page metadata (free-space tracking, a page header). This overhead is real and measurable — a naive record format can spend a surprisingly large fraction of a page's bytes on bookkeeping rather than actual data, which matters directly for storage cost and for how many rows fit per page (fewer rows per page means more pages to scan for the same logical dataset, directly costing more I/O per query).

## Attempt

1. Design a slotted page format on paper first: a page header (page size, number of slots, free space offset), a slot array (each slot: record offset + record length), and record storage growing from the end of the page. Decide your page size (e.g. 4096 bytes) and write out, with actual byte offsets, what a page containing 3 small records would look like.

2. Implement your page format in Go or C: functions to insert a record into a page (finding free space, adding a slot array entry), read a record given a slot number, and delete a record (marking its slot as free, without necessarily compacting immediately — note explicitly whether your delete implementation reclaims the freed space right away or defers it, and why).

3. Serialize a small table's worth of records (e.g. 50 records with a mix of a fixed-width integer field and a variable-width text field) into your page format, writing multiple pages as needed once a page fills up. Then deserialize them back and confirm every record round-trips correctly (matches the original data exactly).

4. Measure actual storage overhead: for your 50 test records, compute the total raw data size (sum of just the meaningful field bytes) versus the total bytes actually written to disk (including all slot array entries, headers, and any padding). Report the overhead as a percentage, and identify which specific component (slot array, header, padding) contributes the most.

## Verify

Report the exact overhead percentage computed in step 4, and for at least 3 of your 50 test records, show the raw bytes as actually stored (slot entry plus record bytes) alongside the original in-memory value, confirming byte-for-byte your serialization format is deterministic and correctly reversible.

## Failure drill

Insert records into a page until it's nearly full, then delete several records from the *middle* of the page (not the most recently added ones) without compacting the freed space, and then attempt to insert a new record that's larger than any single individual freed slot but smaller than the *total* freed space across all the deleted slots combined. Confirm your naive implementation either fails to insert it (even though enough total space technically exists, just fragmented across non-contiguous freed slots) or, if you implemented compaction, succeeds only because compaction consolidated the fragmented free space first. Explain, in your own words, why this is genuinely analogous to external memory fragmentation (a classic OS/allocator concept) applied to on-disk page layout, and why real database storage engines periodically perform page compaction (or vacuum, in PostgreSQL's specific terminology) for exactly this reason.

## Transfer

If TARDOC or Mahall's PostgreSQL tables (covered practically in the postgresql-engineering track) use `VARCHAR`/`TEXT` columns extensively, describe, using this lesson's fixed-vs-variable-width discussion, what PostgreSQL is likely doing differently internally for a fixed-width `INTEGER` column versus a variable-length `TEXT` column at the page/record level — you don't need to look up PostgreSQL's exact internal format, just reason from this lesson's general principles about what a variable-length field's storage must additionally account for that a fixed-width field's doesn't.

## Done when

Your slotted page implementation correctly serializes and deserializes a real batch of mixed fixed/variable-width records with zero data loss, you've measured and can explain your format's actual storage overhead percentage, and you've directly reproduced a fragmentation failure (a large-enough-in-aggregate but non-contiguous free space scenario) and can explain why it's the same underlying phenomenon as OS-level external fragmentation, just at the page-layout level instead of the memory-allocator level.
