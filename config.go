/*
** Copyright (C) 2023 - Champ Clark III <dabeave _AT_ gmail.com>
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

	"encoding/json"
	"io/ioutil"
)

type Configuration struct {
	API_Key      string
	Cache_Errors bool

	Http_Listen string
	Http_TLS    bool
	Http_Cert   string
	Http_Key    string
	Http_Mode   string

	Maxmind_Username string
	Maxmind_Password string
	Maxmind_Url      string

	Redis_Enabled    bool
	Redis_Host       string
	Redis_Port       int
	Redis_Password   string
	Redis_Database   string
	Redis_Cache_Time int
}

var Config *Configuration

func Load(ConfigFile string) {

	json_file, err := ioutil.ReadFile(ConfigFile)

	if err != nil {
		log.Fatalf("Cannot open %s.\n", ConfigFile)
	}

	err = json.Unmarshal(json_file, &Config)

	if err != nil {
		log.Fatalf("Cannot parse configuration file %s.\n", ConfigFile)
	}

	/* Sanity Check */

	if Config.API_Key == "" {
		log.Fatalf("Cannot find 'api_key' in %s.\n", ConfigFile)
	}

	if Config.Http_Listen == "" {
		log.Fatalf("Cannot find 'http_listen' in %s.\n", ConfigFile)
	}

	if Config.Http_Cert == "" {
		log.Fatalf("Cannot find 'http_cert' in %s.\n", ConfigFile)
	}

	if Config.Http_Key == "" {
		log.Fatalf("Cannot find 'http_key' in %s.\n", ConfigFile)
	}

	if Config.Http_Mode == "" {
		log.Fatalf("Cannot find 'http_mode' in %s.\n", ConfigFile)
	}

	if Config.Http_Mode != "release" && Config.Http_Mode != "debug" && Config.Http_Mode != "test" {
		log.Fatalf("Invalid 'http_mode' :  %s.  Valid 'http_modes' are 'release', 'debug' and 'test'\n", Config.Http_Mode)
	}

	if Config.Maxmind_Username == "" {
		log.Fatalf("Cannot find 'maxmind_username' in %s.\n", ConfigFile)
	}

	if Config.Maxmind_Password == "" {
		log.Fatalf("Cannot find 'maxmind_password' in %s.\n", ConfigFile)
	}

	if Config.Maxmind_Url == "" {
		log.Fatalf("Cannot find 'maxmind_url' in %s.\n", ConfigFile)
	}

	if Config.Redis_Host == "" {
		log.Fatalf("Cannot find 'redis_host' in %s.\n", ConfigFile)
	}

	if Config.Redis_Port == 0 {
		log.Fatalf("Cannot find 'redis_port' in %s.\n", ConfigFile)
	}

	if Config.Redis_Password == "" {
		log.Fatalf("Cannot find 'redis_password' in %s.\n", ConfigFile)
	}

	if Config.Redis_Database == "" {
		log.Fatalf("Cannot find 'redis_database' in %s.\n", ConfigFile)
	}

	/* Cache_Time is in hours */

	if Config.Redis_Cache_Time == 0 {
		log.Fatalf("Cannot find 'redis_cache_time' in %s.\n", ConfigFile)
	}

}
