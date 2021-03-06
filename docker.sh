#!/bin/bash

fcn=$1
remain_params=""
for ((i = 2; i <= $#; i++)); do
	j=${!i}
	remain_params="$remain_params $j"
done
function viewContainerPort() {
	CMD="docker container port $1 $2"
	if [ -n "$2" ]; then
		if [ "$3" == "--ip" ]; then
			$CMD | awk '{split($0,a,":"); print a[1]}'
			exit 0
		else
			$CMD | awk '{split($0,a,":"); print a[2]}'
			exit 0
		fi
	fi
	$CMD
}
function imageTrim(){
#    WARNING! This will remove all images without at least one container associated to them.
    docker image prune -a
}
$fcn $remain_params
