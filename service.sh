#!/bin/bash

case $1 in
    start)
        echo "Starting bot app."
        cd /home/skylark/RemoteLockBot/bin/ && /home/skylark/RemoteLockBot/bin/rlbot &
        ;;
    stop)
        echo "Stopping bot app."
        sudo kill $(sudo lsof -t -i:65000)
        ;;
    *)
        echo "RemoteLock bot app service."
        echo $"Usage $0 {start|stop}"
        exit 1
esac
exit 0