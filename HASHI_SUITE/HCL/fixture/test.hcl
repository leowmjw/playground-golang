#############################################################################
# Directors will act as the public facing node (put in isolated public subnet 50..)
# with Traefik/Fabio nodes load balancing to internal web services
# Difference is Worker nodes should be in the private subnet (start 100..)
#############################################################################

variable "node_type" {
  description = "Node Type; will use this to refactor for both Directors/Workers"
  default = "director"
}


resource "aws" "mynode" {
  name = "Dudue.."
}
