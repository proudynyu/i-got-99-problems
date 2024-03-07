# 99 Problems in Go

Based on [99 OCaml](https://v2.ocaml.org/learn/tutorials/99problems.html)

# 1 - Write a function last : 'a list -> 'a option that returns the last element of a list. (easy)

**DONE**

```ocaml
# last ["a" ; "b" ; "c" ; "d"];;
- : string option = Some "d"
# last [];;
- : 'a option = None
```

# 2. Find the last but one (last and penultimate) elements of a list. (easy)

**DONE**

```ocaml
# last_two ["a"; "b"; "c"; "d"];;
- : (string * string) option = Some ("c", "d")
# last_two ["a"];;
- : (string * string) option = None
```

# 3. Find the K'th element of a list. (easy)

**DONE**

```ocaml
# at 3 ["a"; "b"; "c"; "d"; "e"];;
- : string option = Some "c"
# at 3 ["a"];;
- : string option = None
```

# 4. Find the number of elements of a list. (easy)

**DONE**

```ocaml
# length ["a"; "b"; "c"];;
- : int = 3
# length [];;
- : int = 0
```

# 5. Reverse a list. (easy)

- OCaml standard library has List.rev but we ask that you reimplement it.

**DONE**

```ocaml
# rev ["a"; "b"; "c"];;
- : string list = ["c"; "b"; "a"]
```

# 6. Find out whether a list is a palindrome. (easy)
HINT: a palindrome is its own reverse.

**DONE**

```ocaml
# is_palindrome ["x"; "a"; "m"; "a"; "x"];;
- : bool = true
# not (is_palindrome ["a"; "b"]);;
- : bool = true
```

# 7. Flatten a nested list structure. (medium)

**DONE**

```ocaml
# (* There is no nested list type in OCaml, so we need to define one
     first. A node of a nested list is either an element, or a list of
     nodes. *)
  type 'a node =
    | One of 'a 
    | Many of 'a node list;;
type 'a node = One of 'a | Many of 'a node list
```

# 8. Eliminate consecutive duplicates of list elements. (medium)

```ocaml
# compress ["a"; "a"; "a"; "a"; "b"; "c"; "c"; "a"; "a"; "d"; "e"; "e"; "e"; "e"];;
- : string list = ["a"; "b"; "c"; "a"; "d"; "e"]
```
