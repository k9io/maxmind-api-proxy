/*
** Copyright (C) 2023-2024 - Champ Clark III <cclark _AT_ k9.io>
**
** This program is free software; you can redistribute it and/or modify
** it under the terms of the GNU General Public License Version 2 as
** published by the Free Software Foundation.  You may not use, modify or
** distribute this program under any other version of the GNU General
** Public License.
**
** This program is distributed in the hope that it will be useful,
** but WITHOUT ANY WARRANTY; without even the implied warranty of
** MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
** GNU General Public License for more details.
**
** You should have received a copy of the GNU General Public License
** along with this program; if not, write to the Free Software
** Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA 02111-1307, USA.
 */

package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	if len(os.Args) == 1 {
		log.Fatalln("No configuration file specified.")
	}

	log.Printf("Firing up Maxmind-API-Proxy!\n")

	/* Load configuration into "global" memory. */

	Load(os.Args[1])

	Redis_Init()

	log.Printf("Connect to Redis.\n")

	log.Printf("Setting gin to \"%s\" mode.\n", Config.Http_Mode)

	gin.SetMode(Config.Http_Mode)

	router := gin.Default()

	/* Always force API authentication! */

	router.Use(Authenticate_API())

	router.GET("/:ip_address", Maxmind_Query_IP)

	/* Listen for TLS or non-encrypted traffic */

	if Config.Http_TLS == true {

		log.Printf("Listening for TLS traffic on %s.\n", Config.Http_Listen)

		err := router.RunTLS(Config.Http_Listen, Config.Http_Cert, Config.Http_Key)

		if err != nil {
			log.Fatalf("Cannot bind it %s or cannot open %s or %s.\n", Config.Http_Listen, Config.Http_Cert, Config.Http_Key)
		}

	} else {

		log.Printf("Listening for traffic on %s.\n", Config.Http_Listen)

		err := router.Run(Config.Http_Listen)

		if err != nil {
			log.Fatalf("Cannot bind it %s.\n", Config.Http_Listen)
		}

	}

}
