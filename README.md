# Unused-Rules

This utility queries the firewall and out provides information on Unused rules.

# Prerequisites
Requires GO installed in your system. Use this [link](https://go.dev/doc/install) to download GO.

# How to use the tool
Set up environmental variables on your system for the following:
* Firewall IP or FQDN
* Firewall Usernane
* Firewall Password

Update the variables in the main package
`var (
		hostname = os.Getenv("SITEAFW") -- Change "SITEAFW" to match your enviromental setting for the firewall fqdn or IP address
		user     = os.Getenv("PANOUSER") -- Change "PANOUSER" to match your enviromental setting for the firewall's username.
		password = os.Getenv("PANOPWD") -- Change "PANOPWD" to match your enviromental setting for the firewall's password.
		months = 10 -- Change to the number of past months you want to check for last hits. In this example the script will looks for rules with no hits for the last 10 months from today.
	)`