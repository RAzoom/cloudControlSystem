package utils

import (
	"encoding/json"
	"log"
	"cloudControlSystem/cloudStruct"
)

func LoadFullSchema() cloudStruct.FinalStruct {
	body1 := []byte(`{
  "mainNginxAddr": "192.168.0.1:8080",
  "groupMap": {
    "group1": {
      "nginxAddr": "192.168.0.1:7000",
      "consulAddr": "192.168.0.1:6000",
      "serverMap": {
        "server1": {
          "role": "server1",
          "ip": "192.168.0.1",
          "dockerPort": "3232",
          "ServiceMap": {
            "nginx": {
              "containerConfig": {
                "Image": "nginx:latest"
              }
            }
          }
        },
        "server2": {
          "role": "server2",
          "ip": "192.168.0.2",
          "dockerPort": "3232"
        }
      }
    },
    "group2": {
      "nginxAddr": "192.168.1.1:7000",
      "consulAddr": "192.168.1.1:6000",
      "serverMap": {
        "server3": {
          "role": "server1",
          "ip": "192.168.1.1",
          "dockerPort": "3232"
        },
        "server4": {
          "role": "server2",
          "ip": "192.168.1.2",
          "dockerPort": "3232"
        }
      }
    }
  }
}`)
	var schema = cloudStruct.FinalStruct{}
	err := json.Unmarshal(body1, &schema)
	if err != nil {
		log.Fatal(err)
	}

	return schema

}
