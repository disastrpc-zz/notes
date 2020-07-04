# Simple Network Management Protocol (SNMP)

SNMP is based on UDP, a simple, stateless protocol, and is therefore susceptible to IP spoofing and replay attacks. In addition, the commonly used SNMP protocols 1, 2, and 2c offer no traffic encryption, meaning that SNMP information and credentials can be easily intercepted over a local network. Traditional SNMP protocols also have weak authentication schemes and are commonly
left configured with default public and private community strings.

## The SNMP MIB Tree

The SNMP Management Information Base (MIB) is a database containing information usually related to network management. The database is organized like a tree, where branches represent different organizations or network functions. The leaves of the tree (final endpoints) correspond to specific variable values that can then be accessed, and probed, by an external user. 

The following MIB values correspond to specific Microsoft Windows SNMP parameters and contains much more than network-based information:

![snmp_mib_tree](../img/smnp_mib_tree.png)

## Enumerating SNMP MIB Tree

This command will dump the contents of the tree, provided you know the SNMP read-only community string. 

**SNMPWalk**:

> snmpwalk -c public -v1 -t 10 10.10.10.155

-c <string>: specifies the name of the community string
-v1        : specifies SNMP version
-t <int>   : increase timeout period to 10 seconds

Enumerate all users on a Windows system:

> snmpwalk -c public -v1 10.10.10.155 1.3.6.1.4.1.77.1.2.25

Enumerate all running Windows processes: 

> snmpwalk -c public -v1 10.10.10.155 1.3.6.1.2.1.25.4.2.1.2

Enumerate open TCP ports:

> snmpwalk -c public -v1 10.10.10.155 1.3.6.1.2.1.6.13.1.3

Enumerate installed software:

> snmpwalk -c public -v1 10.10.10.155 1.3.6.1.2.1.25.6.3.1.2