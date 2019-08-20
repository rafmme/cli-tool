VERSION ?= 0.0.1  
NAME ?= "PROJECT FOXTROT"  
AUTHOR ?= "Timileyin Farayola"   
  
.PHONY: build move remove

build:  
	chmod u+x ./script.sh; source ./script.sh; build_executable

move:  
	chmod u+x ./script.sh; source ./script.sh; move_to_path

remove:
	chmod u+x ./script.sh; source ./script.sh; remove

DEFAULT: build
