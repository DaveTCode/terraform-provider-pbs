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

# Generate SSH host keys with stronger algorithms
rm -f /etc/ssh/ssh_host*
ssh-keygen -t ed25519 -f /etc/ssh/ssh_host_ed25519_key -N ""
ssh-keygen -t rsa -b 4096 -f /etc/ssh/ssh_host_rsa_key -N ""
/usr/sbin/sshd

/etc/init.d/pbs start

# Wait for PBS to be fully ready on the main PBS server
if [ "$hostname" = "$pbs_server" ] || [ -z "$PBS_SERVER" ]; then
    echo "Waiting for PBS to be ready..."
    for i in $(seq 1 30); do
        if /opt/pbs/bin/qstat -Q >/dev/null 2>&1; then
            echo "PBS is ready!"
            break
        fi
        echo "Waiting... ($i/30)"
        sleep 2
    done
    
    echo "PBS services started successfully!"
fi

exec "$@"
