package main

import data.OwnersArray as OwnersArray
import input.owners

containsIn(Array, val) {
	Array[_] = val
}

# title: owners variable shall exist
deny[msg] {
	not owners

	key_list := object.keys(input)
	msg = sprintf("owners variable not found in %v\n", [key_list])
}

deny[msg] {
	not containsIn(OwnersArray, owners)
	msg = sprintf("owners value %v is forbidden (expected one of %v)\n", [owners, OwnersArray])
}
