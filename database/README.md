### DB Migrations

Using package Goose: https://github.com/pressly/goose

##### Common Commands:
* `goose create <migration_name> sql` - generate migration file in .sql format
* `goose up` - apply all pending migrations
* `goose down` - roll back latest migration
* `goose status` - shows migration status

