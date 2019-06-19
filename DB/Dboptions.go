package DB

var DBmap map[int]string

func InitDBoptions(){
	DBmap = make(map[int]string)
	DBmap[0] = "MySql"
	DBmap[1] = "Postgres"
}
