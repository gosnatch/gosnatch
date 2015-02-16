package main

import (
    "fmt"
    "github.com/gosnatch/gosnatch/gosnatch"
    // "github.com/gosnatch/gosnatch/selfupdate"
    "github.com/spf13/viper"
)

func updateHandler() {
    fmt.Println("Update Successfull!")
}

func main() {

    gosnatch.InitConfig()

    if viper.GetBool("UpdateOnStart") {
        go gosnatch.UpadeShowExceptions()
        go gosnatch.DailySearch(true)
    }

    go gosnatch.DailySearch(false)

    // if viper.GetBool("AutoUpdate") {
    //     var updater = &selfupdate.Updater{
    //         CurrentVersion: gosnatch.VERSION,
    //         ApiURL:         "http://updates.gosnatch.me/",
    //         BinURL:         "http://updates.gosnatch.me/",
    //         DiffURL:        "http://updates.gosnatch.me/",
    //         Dir:            "update/",
    //         CmdName:        "", // app name
    //         Verbose:        true,
    //         CheckInterval:  120,
    //     }

    //     if updater != nil {
    //         go updater.BackgroundRun()
    //     }
    // }

    gosnatch.StartWebserver()
}
