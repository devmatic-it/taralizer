locals {
  # Empty and trivial collections.
  empty_list = []
  empty_map  = {}

  # Nested and trailing-comma collections.
  nested_list = [
    "a",
    ["b", "c"],
    [],
  ]

  nested_map = {
    a = 1,
    b = {
      c = "d",
    },
    e = [1, 2],
  }

  # Map without commas (HCL2 allows newline-separated entries).
  no_comma_map = {
    x = 1
    y = 2
    z = 3
  }

  # Tuple with mixed types.
  mixed_tuple = [1, "two", true, null]

  # Spread into a list.
  spread = [for t in ["p", "q"] : upper(t)]

  # Set constructor.
  set_expr = { for k, v in { a = 1, b = 2 } : k => v * 10 }
}
