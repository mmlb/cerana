#!/usr/bin/env bash

set -e

# args:
# $1: relative path to .test file

dir=$(dirname $1)
name=$(basename $1)
out=$1.run.out
exec 2> $out

which consul &>/dev/null

cid=$(docker run -dti \
	--cap-add SYS_ADMIN \
	--device  /dev/zfs:/dev/zfs \
	--env     TMPDIR=$(mktemp -d) \
	--env     KV=${KV:-consul} \
	--name    $name \
	--volume  $PWD:/mistify:ro \
	--volume  /sys/fs/cgroup:/sys/fs/cgroup:ro \
	--volume  /tmp:/tmp \
	mistifyio/mistify-os:zfs-stable-api
)

[[ -n $cid ]]
sleep .25

docker cp $(which consul) $cid:/usr/bin/
docker cp $(which etcd) $cid:/usr/bin/

docker exec $cid sh -c "cd /mistify/$dir; ./$name -test.v" >&2 || ret=$?;

docker kill  $cid > /dev/null || :
docker rm -v $cid > /dev/null || :
flock /dev/stdout -c "echo '### TEST  $name'; cat $out"
exit $ret
