# CsvToDB
Convert any CSV file to a DB

## Supported Databases:

- Sqlite

More sql and nosql databases will be supported later.

## Build

### With docker
````shell
docker build --output . .
````
This will build a docker image, copy source files, build the project and copy the compiled binary file to your local folder.

## Run
````shell
csvToDB --csv file.csv [--primarykey column]
````
- file.csv: The CSV file to parse.
- column: optional, the column to set as primary key.