# Line comment with a fake arrow <<NOTAHeredoc inside it.
// Another line-comment style (HCL2 accepts both).

/* Block comment spanning lines,
   with <<ALSO_NOT and braces { } inside. */

resource "null_resource" "comments" {
  # trailing comment after an argument
  triggers = {
    a = "1" # hash comment inside a map
    b = "2" // slash-slash comment inside a map
  }
  /* block comment between arguments */
  count = 1
}
