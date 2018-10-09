package main

import (
    "fmt"
    "github.com/BurntSushi/toml"
)

type Config struct {
    Server ServerConfig
    Db     DbConfig
}

type ServerConfig struct {
    Host  string        `toml:"host"`
    Port  string        `toml:"port"`
    Slave []SlaveServer `toml:"slave"`
}

type DbConfig struct {
    User string `toml:"user"`
    Pass string `toml:"pass"`
}

type SlaveServer struct {
    Weight int    `toml:"weight"`
    Ip     string `toml:"ip"`
}

func main() {
    var config Config
    _, err := toml.DecodeFile("config.tml", &config)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Host: %s\n", config.Server.Host)
    fmt.Printf("Port: %s\n", config.Server.Port)
    for k, v := range config.Server.Slave {
        fmt.Printf("Slave %d\n", k)
        fmt.Printf("  weight: %d\n", v.Weight)
        fmt.Printf("  ip: %s\n", v.Ip)
    }

    fmt.Printf("DB User: %s\n", config.Db.User)
    fmt.Printf("DB Pass: %s\n", config.Db.Pass)
}
