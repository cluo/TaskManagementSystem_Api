mongodump -h 127.0.0.1 -o ~/dbbackup
mongorestore -h 172.16.2.25 -d local --drop ~/dbbackup/local