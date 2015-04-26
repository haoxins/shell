#!/usr/bin/env bash

DIR=`pwd`
CMD="node"
DEST="${DIR}/master.js"

ACTION=$1

usage() {
  echo "usage: {start|stop|restart|status}"
  exit 1;
}

get_pid() {
  # fuck:
  #   - awk '{print $1}' works
  #   - awk "{print $1}" not works
  PID=`ps ax | grep ${CMD} | grep ${DEST} | awk '{print $1}'`
}

start() {
  get_pid

  if [ -z $PID ]; then
    echo "starting"
    $CMD $DEST 2>&1 &
    get_pid
    echo "start success. pid=${PID}"
  else
    echo "already running. pid=${PID}"
  fi
}

stop() {
  get_pid

  if [ -z $PID ]; then
    echo "not running"
  else
    kill -15 $PID
    echo "stopped"
  fi
}

restart() {
  stop
  sleep 0.5
  echo =====
  start
}

status() {
  get_pid

  if [ ! -z $PID ]; then
    echo "process pid: ${PID}"
    all_processes=`ps -ef | grep $PID | grep -v grep`
    echo "${all_processes}"
    total=`echo "${all_processes}" | wc -l`
    echo "all processes count: ${total}"
  else
    echo "not running"
  fi
  exit 0;
}

case "$ACTION" in
  start)
    start
  ;;
  stop)
    stop
  ;;
  restart)
    restart
  ;;
  status)
    status
  ;;
  *)
    usage
  ;;
esac
