#!/bin/sh
process_name="docker"

ps aux | grep -v grep | grep "$process_name"
