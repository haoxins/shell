port() {
  netstat -vanp tcp | grep $param1 | grep LISTEN | wc -l
  echo
  netstat -vanp tcp | grep $param1 | grep LISTEN
  port)
    port
