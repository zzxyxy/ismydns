echo off

[ -z "$DNS" ] && echo "DNS is not set" && exit 1

/ismydns $DNS
