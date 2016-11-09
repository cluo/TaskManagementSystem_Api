mongodump -h 172.16.2.25 -d local -o ~/dbbackup
mongorestore -h 172.16.2.25 -d local --drop ~/dbbackup/localÏ