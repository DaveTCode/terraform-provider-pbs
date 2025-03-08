#!/bin/sh

pbs_conf_file=/etc/pbs.conf
mom_conf_file=/var/spool/pbs/mom_priv/config
hostname=$(hostname)

# replace hostname in pbs.conf and mom_priv/config
sed -i "s/PBS_SERVER=.*/PBS_SERVER=$hostname/" $pbs_conf_file
sed -i "s/\$clienthost .*/\$clienthost $hostname/" $mom_conf_file

ssh-keygen -A
/usr/sbin/sshd

/etc/init.d/pbs start

/opt/pbs/bin/qmgr -c "create queue test queue_type=execution"
/opt/pbs/bin/qmgr -c "set queue test started=true"
/opt/pbs/bin/qmgr -c "set queue test enabled=true"
/opt/pbs/bin/qmgr -c "set queue test resources_default.nodes=1"
/opt/pbs/bin/qmgr -c "set queue test resources_default.walltime=3600"

exec "$@"
