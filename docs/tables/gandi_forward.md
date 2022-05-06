# Table: gandi_forward

The `gandi_forward` table can be used to query information about your email forwards and you must specify which domain in the where or join clause using the `domain` column.

## Examples

### List emails forwards of a domain

```sql
select
  source, destinations
from
  gandi_forward
where
  domain='example.net';
```
