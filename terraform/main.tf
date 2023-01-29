terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "3.5.0"
    }
  }
}

provider "google" {
  credentials = file("key.json")

  project = "devopsteste-375800"
  region  = "us-central1"
  zone    = "us-central1-c"
}
#resource "google_compute_firewall" "default" {
#  name    = "default-firewall"
#  network = "default"

#    allow {
#    protocol = "icmp"
#}

#  allow {
#    protocol = "tcp"
#    ports    = ["80", "8080"]
#  }

#}


resource "google_compute_instance" "vmtestdevops" {
  name         = "vmtestdevops"
  machine_type = "e2-micro"
  zone         = "us-central1-c"


  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2204-lts"
      }
  }

        metadata_startup_script = "sudo apt-get update; sudo apt-get install -yq docker.io; sudo apt-get install -yq golang; sudo cd /; sudo git clone https://github.com/sales0108/devops_teste.git; sudo docker build -t my-golang-app /devops_teste/docker/.; sudo docker run -itd -p 8080:8080 --name my-running-app my-golang-app; sudo docker exec -d my-running-app go run . "

  network_interface {
    network = "default"
    access_config {
      // Ephemeral public IP
    }
  }
}
