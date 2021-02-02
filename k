#!/usr/bin/env zsh

action=$1
param1=$2
param2=$3

ctx() {
  kubectl config get-contexts
}

use_ctx() {
  kubectl config use-context $param1
}

ns() {
  kubectl get namespace
}

use_ns() {
  kubectl config set-context --current --namespace $param1
}

pods() {
  kubectl get pods -o wide --namespace $param1
}

logs() {
  kubectl logs --namespace $param1 $param2
}

exec() {
  kubectl exec --stdin --tty --namespace $param1 $param2 -- /bin/bash
}

desc() {
  kubectl describe pods $param2 --namespace $param1
}

usage() {
  echo Usage
  echo '  - pods namespace'
  echo '  - logs namespace pod'
  echo '  - exec namespace pod'
  echo '  - desc namespace pod'
}

case $action in
  pods)
    pods
  ;;
  logs)
    logs
  ;;
  exec)
    exec
  ;;
  desc)
    desc
  ;;
  *)
    usage
  ;;
esac
