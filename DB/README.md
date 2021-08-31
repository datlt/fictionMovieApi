# How To Migrate DB

To create DB table and seeds data, take flowing step

1. Download migrate tool from https://github.com/golang-migrate/migrate/releases/tag/v4.14.1
Select binary that matches your OS

For more information about the tool, view https://github.com/golang-migrate/migrate

2. In this folder, run the binary downloaded in step 1, using the flowing command

$ migrate -database ${POSTGRESQL_URL} -path migrations up