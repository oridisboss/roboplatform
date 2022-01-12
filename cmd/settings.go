package main

type Setting struct {
	Name  string `gorm:"primaryKey"`
	Value string
}

var gSettings map[string]string

func LoadSettings() {
	gSettings = make(map[string]string)

	var sts []Setting

	dbg.Find(&sts)
	for i := range sts {
		gSettings[sts[i].Name] = sts[i].Value
	}
}

func SaveSetting(name string, value string) {
	s := Setting{Name: name, Value: value}
	dbg.Save(&s)
	gSettings[name] = value
}

func GetSetting(name string) string {
	return gSettings[name]
}
