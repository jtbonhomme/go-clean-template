package main

import data.EnvArray as EnvArray

containsIn(Array, val) {
	Array[_] = val
}

# title: env variable shall exist
deny[msg] {
	not input.env

	key_list := object.keys(input)
	msg = sprintf("env variable not found in %v\n", [key_list])
}

deny[msg] {
	not containsIn(EnvArray, input.env)
	msg = sprintf("env value %v is forbidden (expected one of %v)\n", [input.env, EnvArray])
}
