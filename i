#!/usr/bin/env zsh

action=$1
param1=$2
param2=$3

fmt() {
  echo $param1 | tr '[:upper:]' '[:lower:]' | tr ' ' '-' | tr '_' '-'
}

my_ip() {
  ifconfig | grep inet | grep netmask | grep broadcast | cut -d: -f2 | awk '{ print $2}'
}

ports() {
  netstat -an | grep LISTEN
}

port() {
  netstat -vanp tcp | grep $param1 | grep LISTEN | wc -l
  echo
  netstat -vanp tcp | grep $param1 | grep LISTEN
}

ps_find_by_port() {
  lsof -i :$param1
}

ps_find() {
  ps -ef | grep $param1
}

serve() {
  python3 -m http.server $param1
}

usage() {
  echo Usage
  echo '  - fmt     [str]'
  echo '  - port    [port]'
  echo '  - ports                  show listening ports'
  echo '  - ps      [text]'
  echo '  - psport  [port]'
  echo '  - serve   [port]'
}

case $action in
  fmt)
    fmt
  ;;
  ip)
    my_ip
  ;;
  port)
    port
  ;;
  ps)
    ps_find
  ;;
  psport)
    ps_find_by_port
  ;;
  serve)
    serve
  ;;
  *)
    usage
  ;;
esac
