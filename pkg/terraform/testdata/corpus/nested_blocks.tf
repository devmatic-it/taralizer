resource "google_compute_instance" "deep" {
  name         = "vm"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-12"

      labels = {
        env = "test"
      }
    }
  }

  network_interface {
    network = "default"

    access_config {
      nat_ip       = null
      network_tier = "PREMIUM"
    }

    alias_ip_range {
      ip_cidr_range         = "10.0.0.0/28"
      subnetwork_range_name = "gke-pods"
    }
  }

  scratch_disk {
    interface = "SCSI"
  }

  guest_accelerator {
    type  = "nvidia-tesla-t4"
    count = 1
  }

  advanced_machine_features {
    threads_per_core = 2
  }

  lifecycle {
    ignore_changes = [metadata]
  }
}

resource "google_sql_database_instance" "nested_db" {
  name             = "db"
  database_version = "POSTGRES_15"

  settings {
    tier              = "db-f1-micro"
    availability_type = "REGIONAL"

    ip_configuration {
      ipv4_enabled    = true
      private_network = google_compute_network.vpc.self_link

      authorized_networks {
        name  = "office"
        value = "203.0.113.0/24"
      }
    }

    backup_configuration {
      enabled            = true
      binary_log_enabled = true
      start_time         = "06:00"
    }

    database_flags {
      name  = "max_connections"
      value = "200"
    }
  }
}

resource "google_compute_network" "vpc" {
  name                    = "vpc"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnet_a" {
  name          = "subnet-a"
  ip_cidr_range = "10.1.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.vpc.id

  secondary_ip_range {
    range_name    = "gke-pods"
    ip_cidr_range = "10.2.0.0/16"
  }
}
