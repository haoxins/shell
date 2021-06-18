#!/usr/bin/env zsh

action=$1
param1=$2
param2=$3

fmt() {
  echo $param1 | tr '[:upper:]' '[:lower:]' | tr ' ' '-' | tr '_' '-'
}

os_version() {
  local name=`uname -s`
  local isDarwin=`echo $name | grep -i darwin`
  local isLinux=`echo $name | grep -i linux`

  if [ ! -z $isDarwin ]; then
    # system_profiler SPSoftwareDataType
    sw_vers -productVersion
  elif [ ! -z $isLinux ]; then
    cat /etc/os-release
  fi
}

my_ip() {
  ifconfig | grep inet | grep netmask | grep broadcast | cut -d: -f2 | awk '{ print $2}'
}

pmem() {
  local memusage=`ps -o vsz -p $param1 | grep -v VSZ`
  ((memusage /= 1000))
  echo $memusage
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

dotar() {
  # $param1: pkg, $param2: dir
  tar -zcvf ${param1}.tar.gz $param2
}

untar() {
  tar -zxvf $param1
}

usage() {
  echo Usage
  echo '  - fmt     [str]'
  echo '  - osversion'
  echo '  - pm      [pid]          show process memory'
  echo '  - port    [port]'
  echo '  - ports                  show listening ports'
  echo '  - ps      [text]'
  echo '  - psport  [port]'
  echo '  - serve   [port]'
  echo '  - tar     [pkg]  [dir]'
  echo '  - untar   [pkg]'
}

case $action in
  fmt)
    fmt
  ;;
  ip)
    my_ip
  ;;
  osversion)
    os_version
  ;;
  pm)
    pmem
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
  tar)
    dotar
  ;;
  untar)
    untar
  ;;
  *)
    usage
  ;;
esac
