#!/bin/bash

PWD=$(pwd)

tmux new-window -n "server"
tmux send-keys "cd $PWD && air" "C-m"
tmux split-window -v
tmux send-keys "cd $PWD && podman-compose up" "C-m"
tmux select-pane -t 0
tmux split-window -h
