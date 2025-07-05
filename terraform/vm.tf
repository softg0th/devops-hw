resource "yandex_vpc_address" "public-ip-devops-vm" {
  name = "public-ip-devops-vm"
  external_ipv4_address {
    zone_id = "ru-central1-a"
  }
}

resource "yandex_compute_instance" "devops-vm" {
  folder_id          = ""
  name               = "devops-vm"
  hostname           = "devops-vm"
  description        = "Virtual machine"
  zone               = "ru-central1-a"

  resources {
    cores         = 2
    memory        = 4
    core_fraction = 20
  }

  boot_disk {
    initialize_params {
      image_id = ""
      type     = "network-hdd"
      size     = 35
    }
  }

  network_interface {
    subnet_id      = data.yandex_vpc_subnet.a.id
    nat            = true
    nat_ip_address = yandex_vpc_address.public-ip-devops-vm.external_ipv4_address[0].address
  }

  scheduling_policy {
    preemptible = true
  }

  metadata = {
    user-data = file("users.yml")
  }
}
