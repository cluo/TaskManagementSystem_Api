mongodump -h 172.16.0.114 -d local -o ~/dbbackup
mongorestore -h 172.16.0.114 -d local --drop ~/dbbackup/local
mongorestore -h 127.0.0.1 -d local --drop ~/dbbackup/local



rm -Rf ~/Library/Application\ Support/PremiumSoft\ CyberTech/Navicat*
rm -Rf ~/Library/Caches/com.prect.NavicatPremium 
rm -Rf ~/Library/Preferences/com.prect.NavicatPremium.plist
