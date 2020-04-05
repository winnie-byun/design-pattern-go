package common

import "log"

func Build(service string) string {
	log.Println("building...")
	return "successfully built " + service
}

func Start(service string) string {
	log.Println("starting...")
	return "successfully  started " + service
}

func Stop(service string) string {
	log.Println("stoping...")
	return "successfully  stopped " + service
}

func Clean(service string) string {
	log.Println("cleaning...")
	return "successfully  cleaned " + service
}
