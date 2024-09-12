> [!CAUTION]
> This is now deprecated and merged with [opensearch-migration](https://github.com/opensearch-project/opensearch-migrations/tree/main/dashboardsSanitizer) project. 


# dashboard-sanitizer

If the Kibana objects are exported from 7.10.2 or later version, it may not be loaded into OpenSearch Dashboards successfully. This tool makes it compatible to OpenSearch Dashboards by fixing the version numbers for each kibana object and removes any incompatible objects. 

```
Usage
--source string
        The Elastic dashboard object file in ndjson.
--output string
        The OpenSearch compatible dashboard object file in ndjson. (default "os_dashboard_objects.ndjson")
--version
        Prints the version number
```
example:
```
dashboard-sanitizer --source source.ndjson --output to-os-dashboards.ndjson
```
TODO: 
1. ~~Remove lens from panel json and dashboard references.~~ 
2. Fix the max version number based on actual migration version.
3. Reintroduce the summary at the last line.
