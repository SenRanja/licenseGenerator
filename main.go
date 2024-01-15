package main

import (
	"crypto/md5"
	"encoding/hex"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/SenRanja/licenseGenerator/gcmAesCipher"
	"strings"
)

func formatString(input string) string {
	if len(input) != 16 {
		return "Invalid input length"
	}
	parts := []string{
		input[:4],
		input[4:8],
		input[8:12],
		input[12:],
	}
	result := strings.Join(parts, "-")

	return result
}

func LicenseGenerator(Ip string, Time string) string {
	//	### license生成算法
	//
	//	IV、Key值固定（不写在此处）
	//
	//	1. 字符串拼接：`Ip + "shenyanjian" + Time`，编码为`[]byte`类型
	//	2. 然后`AES-GCM`加密得到密文，直接对byte类型的密文进行`md5`生成hex摘要
	//	3. 取前16字节
	//	4. 四四分割，得到license，样子如:`34a6-0b6d-a1de-e321`

	plainText := Ip + "shenyanjian" + Time
	p_b := []byte(plainText)
	c := gcmAesCipher.EncryptGCM(p_b)
	//c_h := hex.EncodeToString(c)
	m := md5.New()
	m.Write(c)
	digest := hex.EncodeToString(m.Sum(nil))[:16]
	license := formatString(digest)
	return license
}

func makeUI() (*widget.Label, *widget.Entry, *widget.Label, *widget.Entry, *widget.Label, *widget.Entry) {
	IpAddressOut := widget.NewLabel("IP address")
	IpAddressIn := widget.NewEntry()

	TimeOut := widget.NewLabel("Start time")
	TimeIn := widget.NewEntry()

	LicenseLabel := widget.NewLabel("License")
	LicenseOut := widget.NewEntry()

	// 逻辑部分
	IpAddressIn.OnChanged = func(content string) {
		LicenseOut.SetText(LicenseGenerator(IpAddressIn.Text, TimeIn.Text))
	}
	TimeIn.OnChanged = func(content string) {
		LicenseOut.SetText(LicenseGenerator(IpAddressIn.Text, TimeIn.Text))
	}

	return IpAddressOut, IpAddressIn, TimeOut, TimeIn, LicenseLabel, LicenseOut
}

// 暂时图省事儿，逻辑部分直接写makeUI()中
func main() {
	a := app.New()
	w := a.NewWindow("License Generator")
	w.SetContent(container.NewVBox(makeUI()))
	w.Resize(fyne.NewSize(300, 200))
	w.Show()
	a.Run()
}
