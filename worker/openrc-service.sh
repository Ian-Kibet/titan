#!/sbin/openrc-run

name=$RC_SVCNAME
description="Primelab CPU worker"
supervisor="supervise-daemon"
command="/usr/local/bin/worker"
pidfile="/run/worker.pid"
command_user="primelab:primelab"

depend() {
	after net
}
