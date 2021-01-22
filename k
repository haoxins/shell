#!/usr/bin/env zsh

action=$1
param1=$2
param2=$3

pods() {
  kubectl get pods -o wide --namespace $param1
}

logs() {
  kubectl logs --namespace $param1 $param2
}

desc() {
  kubectl describe pods $param2 --namespace $param1
}

usage() {
  echo Usage
  echo '  - pods namespace'
  echo '  - logs namespace pod'
  echo '  - desc namespace pod'
}

case $action in
  pods)
    pods
  ;;
  logs)
    logs
  ;;
  desc)
    desc
  ;;
  *)
    usage
  ;;
esac
