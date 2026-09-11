# Lesson 3: Pointers, Arrays, and Strings

## Objective

Make pointer arithmetic and array-to-pointer decay mechanically predictable, and implement string handling from scratch to understand exactly what the standard library functions you normally take for granted are actually doing.

## Prerequisites

Lesson 2 (memory model — pointer arithmetic operates on the addresses just discussed).

## Learn

**Array decay.** In most expressions, an array name "decays" into a pointer to its first element. `int arr[5]; int *p = arr;` is valid precisely because of this decay — `arr` becomes `&arr[0]`. This is why `sizeof(arr)` inside the same scope gives the full array size, but `sizeof(p)` (or `sizeof` on an array parameter passed to a function, which decays at the function boundary) gives only the pointer's size — a famous, genuinely easy-to-hit trap: `sizeof` on a decayed array parameter silently returns the wrong number, with no compiler warning by default.

**Pointer arithmetic is scaled by type size.** `p + 1` doesn't add 1 byte, it adds `sizeof(*p)` bytes — moving to the next element, not the next byte. This is precisely why `int *p; p+1` and `char *p; p+1` advance by different actual byte amounts even though the arithmetic expression looks identical. `arr[i]` is defined as exactly `*(arr + i)` — array indexing is pointer arithmetic with syntax sugar, which is also why `i[arr]` (reversed) compiles and works identically to `arr[i]`, a fact worth knowing exists even though writing it that way in real code would be needlessly confusing.

**C strings have no built-in length.** A C string is a `char*` (or `char[]`) terminated by a `\0` byte — there is no length field anywhere. Every string function (`strlen`, `strcpy`, `strcat`) must scan forward from the start until it finds the terminator. This is the direct mechanical reason `strlen` is O(n) rather than O(1) (unlike Go's `len(s)`, which reads a stored length field), and it's the direct root cause of buffer overflows: functions like `strcpy` have no way to know the destination buffer's capacity, since C strings carry no length information for the function to check against.

## Attempt

1. Write a function that traverses an array using pointer arithmetic only (no `[]` indexing) — increment a pointer through the array and dereference it, rather than using an index variable. Confirm it produces identical output to an index-based traversal of the same array.

2. Demonstrate the `sizeof` decay trap directly: write a function `void printSize(int arr[])` that prints `sizeof(arr)` inside the function, and call it while also printing `sizeof` on the same array from the caller's scope (where it hasn't decayed). Confirm the two numbers differ, and explain in your own words why the function parameter's `sizeof` cannot report the array's true length — a `int arr[]` function parameter is really just `int *arr` in disguise, the compiler will not warn you.

3. Implement `my_strlen`, `my_strcpy`, and `my_strcat` from scratch (no `<string.h>`), using only pointer arithmetic and the `\0` terminator convention. Test each against the real `strlen`/`strcpy`/`strcat` on several inputs, including an empty string, to confirm matching behavior.

4. Deliberately trigger a buffer overflow: allocate a fixed-size `char buf[8]`, and use your `my_strcpy` (or the real `strcpy`) to copy a string longer than 8 characters into it. Compile and run with `-fsanitize=address` and observe the sanitizer catch the overflow with a specific error and stack trace, rather than silently corrupting adjacent memory the way an unsanitized build would (possibly without any visible symptom at all, depending on what's adjacent in memory).

## Verify

For step 3, produce a table comparing your implementation's output against the standard library's for at least 4 test strings (including an empty string and a string exactly matching the destination buffer's capacity), confirming byte-for-byte agreement.

## Failure drill

Run step 4's overflow *without* the sanitizer flag, on a build with optimizations enabled (`-O2`). Observe that the program frequently doesn't crash and may print output that looks entirely normal, depending on what memory happens to sit adjacent to `buf` on the stack. Explain why "it didn't crash" is not evidence the code is correct — undefined behavior's defining property is that the language makes no guarantee about what happens, including the possibility of nothing visibly wrong happening at all, which is exactly why buffer overflow bugs can sit undetected in shipped code for years until a specific memory layout or input triggers a visible symptom (or a security exploit).

## Transfer

Go strings carry an explicit length internally (a Go string is effectively a pointer plus a length, not a null-terminated byte sequence) and Go slices carry both a length and capacity. Explain, using this lesson's `sizeof`-decay and buffer-overflow examples as the contrast, what specific class of bug this design choice in Go eliminates by construction, compared to C's convention.

## Done when

You can traverse and manipulate arrays using raw pointer arithmetic without indexing syntax, your from-scratch string functions match the standard library's behavior across edge cases including empty strings, and you've directly observed the same buffer overflow both silently succeed (unsanitized) and get caught with a precise diagnostic (sanitized) — and can explain why both outcomes are consistent with "undefined behavior," not just the caught one.
