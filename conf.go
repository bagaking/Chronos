package main

type configEntry struct {
	path string;
}

type config struct {
	[]configEntry entries
}


func loadConfig() {
	file, _ := os.Open(CONF_SH) 
}