package main

import data.EnvArray as EnvArray
import input.env

containsIn(Array, val) {
	Array[_] = val
}

# title: env variable shall exist
deny[msg] {
	not env

	key_list := object.keys(input)
	msg = sprintf("env variable not found in %v\n", [key_list])
}

deny[msg] {
	not containsIn(EnvArray, env)
	msg = sprintf("env value %v is forbidden (expected one of %v)\n", [env, EnvArray])
}
