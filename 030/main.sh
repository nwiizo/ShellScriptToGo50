#!/bin/bash

psql -d dbname -U username <<EOF
SELECT * FROM users;
EOF
