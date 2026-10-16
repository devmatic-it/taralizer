variable "pools" {
  type = list(object({
    name     = string
    max_pods = number
  }))
}

variable "policies" {
  type    = map(string)
  default = {}
}

variable "enable_netpol" {
  type    = bool
  default = false
}

resource "google_container_cluster" "cluster" {
  name               = "c"
  location           = "us-central1-a"
  initial_node_count = 1

  dynamic "node_pool" {
    for_each = var.pools
    content {
      name              = node_pool.value.name
      max_pods_per_node = node_pool.value.max_pods
    }
  }

  dynamic "ip_allocation_policy" {
    for_each = var.policies
    content {
      cluster_secondary_range_name = ip_allocation_policy.key
      subnetwork_name              = ip_allocation_policy.value
    }
  }

  dynamic "network_policy" {
    for_each = var.enable_netpol ? [1] : []
    content {
      enabled  = true
      provider = "CALICO"
    }
  }

  node_config {
    labels = {
      pool = "default"
    }
  }
}

resource "google_compute_instance" "plain" {
  name         = "vm"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-12"
    }
  }

  network_interface {
    network = "default"
  }
}
