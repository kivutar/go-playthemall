package menu

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"

	"github.com/libretro/ludo/video"
)

func downloadThumbnail(list *entry, i int, url, folderPath, path string) {
	fmt.Println("Download " + url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		list.children[i].thumbnail = menu.icons["img-broken"]
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		list.children[i].thumbnail = menu.icons["img-broken"]
		return
	}

	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		list.children[i].thumbnail = menu.icons["img-broken"]
		return
	}

	out, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		list.children[i].thumbnail = menu.icons["img-broken"]
		return
	}
	defer out.Close()

	io.Copy(out, resp.Body)
}

func drawThumbnail(list *entry, i int, system, gameName string, x, y, w, h, scale float32) {
	usr, _ := user.Current()
	folderPath := usr.HomeDir + "/.ludo/thumbnails/" + system + "/Named_Snaps/"
	path := folderPath + gameName + ".png"
	url := "http://thumbnails.libretro.com/" + system + "/Named_Snaps/" + gameName + ".png"

	if list.children[i].thumbnail == 0 || list.children[i].thumbnail == menu.icons["img-dl"] {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			list.children[i].thumbnail = video.NewImage(path)
		} else if list.children[i].thumbnail != menu.icons["img-dl"] {
			list.children[i].thumbnail = menu.icons["img-dl"]
			go downloadThumbnail(list, i, url, folderPath, path)
		}
	}

	vid.DrawImage(
		list.children[i].thumbnail,
		x, y, w, h, scale,
		video.Color{R: 1, G: 1, B: 1, A: 1},
	)
}
