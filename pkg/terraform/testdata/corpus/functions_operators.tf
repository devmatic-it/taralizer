variable "env" {
  type    = string
  default = "dev"
}

variable "replicas" {
  type    = number
  default = 3
}

variable "tags" {
  type    = list(string)
  default = ["a", "b"]
}

variable "attrs" {
  type    = map(string)
  default = { team = "platform", tier = "gold" }
}

locals {
  # String functions and operators.
  upper_env   = upper(var.env)
  lower_env   = lower(var.env)
  trimmed     = trimspace("  padded  ")
  replaced    = replace(var.env, "d", "D")
  formatted   = format("env=%s replicas=%d", var.env, var.replicas)
  concat_tags = concat(var.tags, ["c"])
  join_str    = join("-", var.tags)
  split_bits  = split(",", "x,y,z")

  # Collection functions.
  tag_count   = length(var.tags)
  attr_keys   = keys(var.attrs)
  attr_vals   = values(var.attrs)
  merged      = merge(var.attrs, { extra = "1" })
  set_union   = union(["a"], ["b"])
  set_inter   = intersect(var.tags, ["a", "z"])
  set_sub     = subtract(var.tags, ["a"])
  chunked     = chunklist(var.tags, 1)
  flattened   = flatten([[var.tags], ["c"]])

  # Numeric and boolean logic.
  doubled     = var.replicas * 2
  remainder   = var.replicas % 2
  abs_val     = abs(-var.replicas)
  is_even     = var.replicas % 2 == 0
  and_expr    = var.replicas > 1 && var.env != ""
  or_expr     = var.replicas < 2 || var.env == "prod"
  not_expr    = !(var.replicas > 10)

  # Conditionals and coalescing.
  display     = var.replicas > 1 ? "many" : "few"
  fallback    = coalesce(null, var.env)
  one_or_null = coalescelist(["x"], ["y"])

  # JSON and encoding.
  encoded     = jsonencode({ env = var.env, tags = var.tags })
  decoded     = jsondecode(local.encoded)
  base64      = base64encode(var.env)

  # Regex.
  has_digit   = can(regex("\\d+", "a1b2"))
}
