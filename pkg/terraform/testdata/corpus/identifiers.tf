variable "project" {
  type    = string
  default = "p"
}

variable "instances" {
  type    = list(object({ name = string, zone = string }))
  default = []
}

data "google_client_config" "current" {}

module "vpc" {
  source  = "./modules/vpc"
  project = var.project
}

locals {
  # Plain and chained identifiers.
  plain        = "x"
  var_chain    = var.project
  local_chain  = local.plain
  data_chain   = data.google_client_config.current.project
  module_chain = module.vpc.output_name

  # Splat and index navigation.
  splat_names   = var.instances[*].name
  first_zone    = var.instances[0].zone
  index_chain   = var.instances[1].name
  star_chain    = var.instances[*].zone

  # Keyword-ish identifiers (in, provider as traversal segments).
  in_list       = [for x in var.instances : x.name]
  provider_ref  = provider.google.project

  # Numeric index into a literal.
  numeric_index = [10, 20, 30][2]

  # Attribute after a splat.
  attr_after_splat = var.instances[*].name[0]
}

output "names" {
  value = local.splat_names
}
