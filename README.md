# LiteFS on Fly

This is a working example of a web app that can depend on many locally accessed SQLite on disk.

These databases are replicated automatically across all fly nodes with [LiteFS](https://fly.io/docs/litefs/).

All the databases can be backed up easily with a cron job or LiteFS Cloud.

## Benefits

This is essentially a globally replicated application with incremental rollouts automatic fallover for both web app and database, and very quick disaster recovery.

It can be ran for $0/mo on Fly under its free plan. It includes:
 - 3 app instances in Dallas, Ashburn, and Seattle (1 core, 256Mb RAM)
 - Each app instance runs a Go web app and has a volume mounted in each region.
 - Each region contains a copy all the databases (hundreds to thousands) on disk.
 - Each of the databases are replicated automatically to the other 2 regions
 - For another $5/mo, LiteFS Cloud will backup each database every 5 minutes

## Development

Run the webapp locally

```
air
```

## Deployment

Follow the guide: https://fly.io/docs/litefs/getting-started-fly/


```
flyctl launch
```

Make sure Consul is attached to the app so that all nodes can be linked together

```
flyctl consul attach
```


