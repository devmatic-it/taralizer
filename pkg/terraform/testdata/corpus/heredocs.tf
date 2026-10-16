resource "google_compute_instance" "script_vm" {
  name = "vm"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-12"
    }
  }

  network_interface {
    network = "default"
  }

  metadata = {
    startup-script = <<SCRIPT
#!/bin/bash
echo hello
if [ -f /etc/motd ]; then
  echo "hi"
fi
for i in 1 2 3; do
  echo "$i {braces} and # hashes"
done
SCRIPT
    user-data = <<-EOT
      indented line one
      indented line two with "quotes" and ${interpolation}
    EOT
  }

  description = <<DESCRIPTION
A native DESCRIPTION heredoc.
DESCRIPTION
}

resource "google_compute_instance" "eof_vm" {
  name = "vm2"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-12"
    }
  }

  network_interface {
    network = "default"
  }

  metadata_startup_script = <<EOF
native EOF tag body
EOF
}
