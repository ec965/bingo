#!/bin/bash

PWD=$(pwd)
CMD="docker-compose"

if command -v "podman-compose" &>/dev/null
then
  CMD="podman-compose"
fi

tmux new-window -n "server"
tmux send-keys "cd $PWD && air" "C-m"
tmux split-window -v
tmux send-keys "cd $PWD && $CMD up" "C-m"
tmux select-pane -t 0
tmux split-window -h
