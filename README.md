# Ent Index Migration Test

## Issue

When creating a model with a custom table name (through `Config`) that is the singular form of the model name, generated migration code cycles between dropping and creating the index.

## Running

To generate the migration code, run `migrate.go`, which will produce the `migration.sql` file. The first time it's run, replace the `$1` with the correct table name.

After migrating the database, run `migrate.go` again. This will overwrite `migration.sql` and produce a file that contains a line to drop the index.

After dropping the index, run `migrate.go` again. The new `migration.sql` file will contain a line to create the index again the same as it did initially.

After this the cycle repeats.

## Experimenting

Try removing the `Config` method on the `Employee` model, the migration cycle behavior goes away.

Try using a different table name (other than `"employee"`, e.g. `"emp"`), the migration cycle behavior goes away.