terraform {
  required_version = ">= 1.3"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.50"
    }
    google-beta = {
      source  = "hashicorp/google-beta"
      version = "~> 4.50"
    }
  }

  backend "local" {
    path = "state.tfstate"
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}

variable "project_id" {
  type    = string
  default = "example-project"
}

variable "region" {
  type    = string
  default = "us-central1"
}

locals {
  common_tags = {
    project = var.project_id
  }
}

module "network" {
  source  = "./modules/network"
  project = var.project_id
}

output "network_id" {
  value       = module.network.id
  description = "The network id."
}

data "google_compute_image" "debian" {
  family  = "debian-12"
  project = "debian-cloud"
}

resource "google_storage_bucket" "logs" {
  name          = "${var.project_id}-logs"
  location      = var.region
  force_destroy = true

  versioning {
    enabled = true
  }
}

resource "google_compute_instance" "web" {
  name         = "web"
  machine_type = "e2-small"
  zone         = "${var.region}-a"

  boot_disk {
    initialize_params {
      image = data.google_compute_image.debian.self_link
    }
  }

  network_interface {
    network = "default"

    access_config {
      network_tier = "PREMIUM"
    }
  }

  metadata = local.common_tags
}
