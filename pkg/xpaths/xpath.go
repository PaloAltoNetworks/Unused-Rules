package xpaths

//SysInfo: Returns data containing the system's information
func SysInfo() string {
	return "&type=op&cmd=<show><system><info></info></system></show>"
}

//HitCount: Returns data containing security rules hit count
func HitCount() string {
	return "&type=op&cmd=<show><rule-hit-count><vsys><vsys-name><entry+name='vsys1'><rule-base><entry+name='security'><rules><all/></rules></entry></rule-base></entry></vsys-name></vsys></rule-hit-count></show>"
}
