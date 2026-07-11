locals { service_tags = concat(["${var.router}.enable=true"], var.entrypoints == "" ? [] : ["${var.router}.http.routers.${var.service_name}.entrypoints=${var.entrypoints}"], var.extra_tags) }
