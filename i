fmt() {
  echo $param1 | tr '[:upper:]' '[:lower:]' | tr ' ' '-' | tr '_' '-'
my_ip() {
  ifconfig | grep inet | grep netmask | grep broadcast | cut -d: -f2 | awk '{ print $2}'
port() {
  netstat -vanp tcp | grep $param1 | grep LISTEN | wc -l
  echo
  netstat -vanp tcp | grep $param1 | grep LISTEN
  fmt)
    fmt
  ip)
    my_ip
  port)
    port
