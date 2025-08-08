#!/bin/sh

pbs_conf_file=/etc/pbs.conf
mom_conf_file=/var/spool/pbs/mom_priv/config
hostname=$(hostname)

# If PBS_SERVER environment variable is set, use it; otherwise use current hostname
if [ -n "$PBS_SERVER" ]; then
    pbs_server=$PBS_SERVER
else
    pbs_server=$hostname
fi

# replace hostname in pbs.conf and mom_priv/config
sed -i "s/PBS_SERVER=.*/PBS_SERVER=$pbs_server/" $pbs_conf_file
sed -i "s/\$clienthost .*/\$clienthost $hostname/" $mom_conf_file

ssh-keygen -A
/usr/sbin/sshd

/etc/init.d/pbs start

# Only create queues and setup on the main PBS server
if [ "$hostname" = "$pbs_server" ] || [ -z "$PBS_SERVER" ]; then
    # Wait for PBS to be fully ready - check if we can list queues
    echo "Waiting for PBS to be ready..."
    for i in $(seq 1 30); do
        if /opt/pbs/bin/qstat -Q >/dev/null 2>&1; then
            echo "PBS is ready!"
            break
        fi
        echo "Waiting... ($i/30)"
        sleep 2
    done
    
    /opt/pbs/bin/qmgr -c "create queue test queue_type=execution"
    /opt/pbs/bin/qmgr -c "set queue test started=true"
    /opt/pbs/bin/qmgr -c "set queue test enabled=true"
    /opt/pbs/bin/qmgr -c "set queue test resources_default.nodes=1"
    /opt/pbs/bin/qmgr -c "set queue test resources_default.walltime=3600"
    
    # Set up workq as default queue
    /opt/pbs/bin/qmgr -c "set server default_queue=workq"
    
    # Create one node for import testing using the server's own hostname
    /opt/pbs/bin/qmgr -c "create node pbs"
    /opt/pbs/bin/qmgr -c "set node pbs comment='Pre-existing node for import testing'"
    /opt/pbs/bin/qmgr -c "set node pbs resources_available.ncpus=8"
    /opt/pbs/bin/qmgr -c "set node pbs resources_available.mem=16gb"
    
    # Create one hook for import testing
    /opt/pbs/bin/qmgr -c "create hook test"
    /opt/pbs/bin/qmgr -c "set hook test enabled=true"
    /opt/pbs/bin/qmgr -c "set hook test event=execjob_begin"
    /opt/pbs/bin/qmgr -c "set hook test order=1"
    /opt/pbs/bin/qmgr -c "set hook test type=site"
    /opt/pbs/bin/qmgr -c "set hook test user=pbsadmin"
    /opt/pbs/bin/qmgr -c "set hook test fail_action=none"
    
    # Create one resource for import testing
    /opt/pbs/bin/qmgr -c "create resource test type=size"
    /opt/pbs/bin/qmgr -c "set resource test flag=h"
    
    echo "PBS setup complete!"
fi

exec "$@"
