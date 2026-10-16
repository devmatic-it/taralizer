terraform {
  required_version = ">= 1.3"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.50"
    }
  }
}

provider "google" {
  project     = var.project_id
  region      = var.region
  credentials = file(var.credentials_file)
}

variable "project_id" {
  type    = string
  default = "example-project"
}

variable "region" {
  type    = string
  default = "us-central1"
}

variable "credentials_file" {
  type    = string
  default = ""
}

resource "google_compute_network" "vpc" {
  name                    = "app-vpc"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "app_subnet" {
  name          = "app-subnet"
  ip_cidr_range = "10.0.0.0/24"
  region        = var.region
  network       = google_compute_network.vpc.id
}

resource "google_compute_instance" "app" {
  name         = "app-server"
  machine_type = "e2-medium"
  zone         = "${var.region}-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-12"
    }
  }

  network_interface {
    subnetwork = google_compute_subnetwork.app_subnet.name

    access_config {
      network_tier = "PREMIUM"
    }
  }

  metadata_startup_script = <<SCRIPT
#!/bin/bash
apt-get update
apt-get install -y nginx
systemctl enable nginx
SCRIPT

  metadata = {
    env = var.region
  }

  tags = ["http-server", "https-server"]
}

resource "google_sql_database_instance" "app_db" {
  name             = "app-db"
  database_version = "POSTGRES_15"

  settings {
    tier = "db-custom-2-7680"

    ip_configuration {
      ipv4_enabled    = false
      private_network = google_compute_network.vpc.self_link
    }
  }
}

resource "google_sql_database" "app" {
  name     = "appdb"
  instance = google_sql_database_instance.app_db.name
}

resource "google_storage_bucket" "assets" {
  name          = "${var.project_id}-assets"
  location      = var.region
  force_destroy = true

  lifecycle_rule {
    condition {
      age = 30
    }

    action {
      type = "Delete"
    }
  }
}

output "instance_ip" {
  value = google_compute_instance.app.network_interface[0].access_config[0].nat_ip
}
