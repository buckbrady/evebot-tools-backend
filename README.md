# EVEBot Tools Backend

## Commands

### Create Database migration
```bash
atlas migrate diff <migration_name> \
  --dir "file://migrations" \
  --to "file://schema.hcl" \
  --dev-url "docker://postgres/16/test"
```