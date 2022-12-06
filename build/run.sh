JWT="fuck"
ADDR=":8080"

DBNAME="db"
DBPORT="5432"
DBUSER="alexdenkk"
DBPSWD="12345678"

./server -jwt=$JWT -address=$ADDR -db-name=$DBNAME -db-port=$DBPORT -db-password=$DBPSWD -db-user=$DBUSER
