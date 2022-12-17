# dashboard-sanitizer
Dashboards ndjson sanitizer
Usage
```
-output string
        The OpenSearch compatible dashboard object file in ndjson. (default "os_dashboard_objects.ndjson")
 -source string
        The Elastic dashboard object file in ndjson.
```
example:
```
dashboard-sanitizer --source source.ndjson --output to-os-dashboards.ndjson
```
TODO: 
1. Remove lens from panel json and dashboard references.
2. Fix the max version number based on actual migration version.
3. Reintroduce the summary at the last line.
