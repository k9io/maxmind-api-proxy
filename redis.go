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
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()

func Redis_Init() {

	connect_string := fmt.Sprintf("%s:%d", Config.Redis_Host, Config.Redis_Port)

	database, _ := strconv.Atoi(Config.Redis_Database)

	rdb = redis.NewClient(&redis.Options{
		Addr:     connect_string,
		Password: Config.Redis_Password,
		DB:       database,
	})

}

func Redis_Check_Cache(ip_address string) []byte {

	key := fmt.Sprintf("geoip:%s", ip_address)

	val, err := rdb.Get(ctx, key).Result()

	if err != nil {
		return nil
	}

	return []byte(val)
}

func Redis_Store_Cache(json_data string, ip_address string) {

	key := fmt.Sprintf("geoip:%s", ip_address)

	err := rdb.Set(ctx, key, json_data, time.Duration(Config.Redis_Cache_Time)*time.Hour).Err()

	if err != nil {
		log.Printf("ERROR: Can't SET into Redis!")
		return
	}

}
