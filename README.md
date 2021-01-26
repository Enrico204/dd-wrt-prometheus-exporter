# DD-WRT Prometheus exporter

This prometheus exporter uses some web APIs for DD-WRT in order to collect some data.
It can be used in place of SNMP in routers/releases where SNMP is not available anymore.

Tested with DD-WRT v3.0-r46527 std in WRT54GL.

**Note**: this is a work in progress. Data parsing is not stable at the moment.

# Usage

```sh
./dd-wrt-exporter -url http://192.168.0.1 -username admin -password admin -interfaces eth0,eth1,br0
```

Customize all parameters according your needs. You can use also these environment variables:
* `DDWRT_URL` for router URL
* `DDWRT_USERNAME` for router username
* `DDWRT_PASSWORD` for router password
* `DDWRT_INTERFACES` for the interface list (see above)

The code supports loading those variables from a `.env` file.

# License

See `LICENSE` file.
