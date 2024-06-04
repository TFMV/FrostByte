#!/bin/bash

set -e

# Wait for the services to be up and running
sleep 30

# Set up Iceberg tables
spark-sql --master spark://spark:7077 -e "
CREATE DATABASE IF NOT EXISTS iceberg_db;
USE iceberg_db;
CREATE TABLE IF NOT EXISTS sample_table (id INT, name STRING) USING iceberg;
INSERT INTO sample_table VALUES (1, 'Alice'), (2, 'Bob');
"

# Set up Polaris
psql -h polaris -U polaris -d polaris -c "
CREATE TABLE IF NOT EXISTS sample_table (id INT PRIMARY KEY, name TEXT);
INSERT INTO sample_table VALUES (1, 'Alice'), (2, 'Bob');
"
