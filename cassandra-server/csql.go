package main

type csqlh_dt struct {
	ascii   string
	bigint  int64
	boolean bool
	double  float64
	float   float32
}

type KeySpace struct { // object exists in gocql

}

// Check if table exists in KeySpace
func tableExists(string KeySpace) bool {

	return false
}

// Create table in cluster, if exists return an error
func CreateTable(tablename string, columns map[string]csqlh_dt, values []string) error {
	var err error

	return err
}

// Create new keyspace in cassandra client
func CreateKeyspace() error {
	var err error

	return err
}

// Connect to kafka as consumer and subscribe to relevant topics
func ConnectKafka(kafkabrokers []string, topics []string) error {
	var err error
	return err
}

func PrintKeyspaces() {

}
func csqlh_uuid() string {

	return ""
}
