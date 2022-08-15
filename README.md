# ipinfo_go
Tiny program that fetches ipinfo.io information and prints it in a table

Example output:
```
$> go run . 1.1.1.1
+----------+--------------------------+
| IP       | 1.1.1.1                  |
| Hostname | one.one.one.one          |
| City     | Los Angeles              |
| Region   | California               |
| Country  | US                       |
| Loc      | 34.0522,-118.2437        |
| Org      | AS13335 Cloudflare, Inc. |
| Postal   | 90076                    |
| Timezone | America/Los_Angeles      |
+----------+--------------------------+
```
