package data

import (
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/alexcoder04/meme"
	"github.com/shirou/gopsutil/cpu"
)

func Sysstat() map[string]any {
	c, err := cpu.Percent(time.Second, false)
	if err != nil {
		c = []float64{0}
	}

	var m int
	mo, err := meme.GetMemInfo()
	if err != nil {
		m = 15
	} else {
		m = int((mo.Used + mo.Shared + mo.Buffers) / mo.MemTotal * 100)
	}

	t := 0
	zones, err := os.ReadDir("/sys/class/thermal")
	if err != nil {
		t = 50
	} else {
		for _, z := range zones {
			if !strings.HasPrefix(z.Name(), "thermal_zone") {
				continue
			}
			data, err := os.ReadFile("/sys/class/thermal/" + z.Name())
			if err != nil {
				continue
			}
			if strings.TrimSpace(string(data)) == "cpu-thermal" || strings.TrimSpace(string(data)) == "x86_pkg_temp" {
				data, err := os.ReadFile("/sys/class/thermal/" + z.Name() + "/temp")
				if err != nil {
					continue
				}
				i, err := strconv.Atoi(string(data))
				if err != nil {
					continue
				}
				t = int(i / 1000)
			}
		}
	}

	if t == 0 {
		t = 50
	}

	fs := syscall.Statfs_t{}

	d := 0
	err = syscall.Statfs("/", &fs)
	if err != nil {
		d = 20
	} else {
		totalSize := float64(fs.Blocks) * float64(fs.Bsize)
		usedSize := float64(fs.Blocks-fs.Bfree) * float64(fs.Bsize)
		d = int((usedSize / totalSize) * 100)
	}

	return map[string]any{
		"cpu":  int(c[0]),
		"mem":  m,
		"temp": t,
		"disk": d,
	}
}
