package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/zackmacharia/Unused-Rules/datastruct"
	"github.com/zackmacharia/Unused-Rules/pkg/client"
	"github.com/zackmacharia/Unused-Rules/pkg/key"
	"github.com/zackmacharia/Unused-Rules/pkg/xpaths"
)

func main() {
	var (
		hostname = os.Getenv("SITEAFW")
		user     = os.Getenv("PANOUSER")
		password = os.Getenv("PANOPWD")
		months = 10
	)

	apiKey := key.ApiKey(hostname, user, password)

	cmd := xpaths.HitCount()

	url := "https://" + hostname + "/api/?&key=" + apiKey + cmd

	resp := client.FwClient(url)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var data datastruct.Response

	err = xml.Unmarshal(body, &data)

	rules := data.Result.RuleHitCount.Vsys.VsysNumber.Rulebase.RuleType.Rules.RuleDetails
	n := len(rules)

	fmt.Println("*****Rules with Zero Hit Count*****")
	for i := 0; i < n; i++ {
		if data.Result.RuleHitCount.Vsys.VsysNumber.Rulebase.RuleType.Rules.RuleDetails[i].HitCountInt == 0 {
			a := fmt.Sprint("Name:", data.Result.RuleHitCount.Vsys.VsysNumber.Rulebase.RuleType.Rules.RuleDetails[i].Name) + " " +
				fmt.Sprintln("HitCount:0")
			fmt.Println(a)
		}
	}

	fmt.Println("***** Rules with no hit count for the last", months, "Months *****")
	for i := 0; i < n; i++ {
		if data.Result.RuleHitCount.Vsys.VsysNumber.Rulebase.RuleType.Rules.RuleDetails[i].HitCountInt > 0 {
			t := time.Now().AddDate(0, -months, 0) // represents the last number of months from now
			l := time.Unix(data.Result.RuleHitCount.Vsys.VsysNumber.Rulebase.RuleType.Rules.RuleDetails[0].LastHitTS, 0)
			if l.Before(t) {
				r := fmt.Sprint("Name:", data.Result.RuleHitCount.Vsys.VsysNumber.Rulebase.RuleType.Rules.RuleDetails[i].Name) + " " +
					fmt.Sprint("Last Hit:", l, " Hit Count:", data.Result.RuleHitCount.Vsys.VsysNumber.Rulebase.RuleType.Rules.RuleDetails[i].HitCountInt)
				fmt.Println(r)
			}
		}
	}
}
