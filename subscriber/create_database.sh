HOST=127.0.0.1
PORT=27017
USER=mongo-root
PASSWORD=Test*123

if(($#==1)); then
	$HOST=$1
fi

if(($#==2)); then
	$PORT=$2
fi

if(($#==3)); then
	$USER=$3
fi

if(($#==4)); then
	$PASSWORD=$4
fi

mongo --host $HOST --port $PORT -u $USER -p $PASSWORD --authenticationDatabase admin --eval "db.getSiblingDB('mqtt_database').createUser( { user:'mqtt-owner', pwd:'Test*123', roles:[{role: 'dbOwner', db: 'mqtt_database'}] } );"
#mongo --host $HOST --port $PORT -u $USER -p $PASSWORD --authenticationDatabase admin --eval 'db.new_collection({some_key: "some_value" })'
#mongo --host $HOST --port $PORT -u $USER -p $PASSWORD --authenticationDatabase admin --shell ./create_db.js
