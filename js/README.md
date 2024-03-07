# 99 Problems in JS

Based on [99 OCaml](https://v2.ocaml.org/learn/tutorials/99problems.html)

# 1 - Write a function last : 'a list -> 'a option that returns the last element of a list. (easy)

```ocaml
# last ["a" ; "b" ; "c" ; "d"];;
- : string option = Some "d"
# last [];;
- : 'a option = None
```
