package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	// "github.com/nasOS-official/gfb"
)

// var resX, resY int = gfb.GetResolution("fb0")
// var fb []uint8 = gfb.InitFb()

// func logo() {
// 	if (resX >= 640) && (resY >= 480) {
// 		// drawTestRainbow(fb, (resX-resY)/2, resY+((resX-resY)/2), 0, resY)
// 		// gfb.DrawCircle(fb, resY/2, int(float64(resX)*float64(80)/100.0), 128, 255, 255, 255)
// 		x_center := int(float64(resX) * float64(80) / 100.0)
// 		y_center := resY / 2
// 		radius := 128
// 		gfb.DrawCircle(fb, y_center, x_center, radius, 255, 255, 255)

// 		gfb.DrawRectangle(fb, x_center-61, x_center-35, y_center-37, y_center+67, 27, 173, 255)
// 		gfb.DrawRectangle(fb, x_center-13, x_center+13, y_center-64, y_center+67, 27, 173, 255)
// 		gfb.DrawRectangle(fb, x_center+35, x_center+61, y_center-6, y_center+67, 27, 173, 255)

//			gfb.UpdateScreen(fb)
//		}
//	}
const welcomeru = "Добро пожаловать в программу установки NurOS! \nВыберите язык системы используя стрелки. Нажмите Enter для продолжения."
const welcomeen = "Welcome to the NurOS Installer! \nselect the system language using the arrows. Press Enter to continue."
const welcomekz = "NurOS орнату бағдарламасына қош келдіңіз! \пжүйе тілін таңдаңыз көрсеткілерді қолдану. Жалғастыру үшін Enter пернесін басыңыз."
const Licenceru = `Лицензионное соглашение!

Здравствуйте, благодарим за выбор NurOS.
Данный продукт распространяется по лицензии
GNU General Public License 3.0
Более подробно вы можете прочитать здесь:
https://www.gnu.org/licenses/gpl-3.0.html#license-text

Лицензионное соглашение для продуктов DeltaUI:
(Калькулятор, beta-версия Браузера OpenSerfing)
Пожалуйста, ознакомьтесь с условиями настоящего лицензионного соглашения.
Пользуясь программными продуктами AnmiTali DeltaUI,
Вы соглашаетесь с тем, что:
А) Программа имеет открытый исходный код и вы имеете право изменять его или 
как-либо модифицировать.
Б) Программа предоставляется "как есть", без гарантийных обязательств, явных или
подразумеваемых, либо предусмотренных законодательством, 
включая, но не ограничиваясь этим, гарантии качества, производительности,
пригодности для продажи или для определенной цели.
В) Также не предоставляется никаких гарантий, созданных в результате
заключения сделки, использования или продаж. AnmiTali не 
гарантирует бесперебойную, своевременную и безошибочную работу 
программного обеспечения. Ни при каких условиях AnmiTali не несут 
ответственность за ущерб или убытки, вызванные использованием или 
невозможностью использования данного продукта. 
Г) ПО по данному соглашению предоставляется без явных или 
подразумеваемых гарантий о не нарушениях, и разработчик не дает 
гарантий о не нарушениях любых патентов, авторских прав, торговых
секретов или других прав собственности.
Если вы не согласны с условиями лицензии,
немедленно прекратите использование данного ПО!

Дополнение: Программа установки .apg приложений еще в разработке, также 
распространяется по лицензии GNU General Public License 3.0.

Разработчики:
Игнатьев Илья; Савин Ярослав; Чакилев Арсений (В nasOS-Installer и за идею) и группа разработчиков AnmiTali
const Licenceen = `License Agreement!

Hello, thank you for choosing NurOS.
This product is distributed under license
GNU General Public License 3.0
You can read more details here:
https://www.gnu.org/licenses/gpl-3.0.html#license-text

License agreement for DeltaUI products:
(Calculator, beta version of OpenSerfing Browser)
Please read the terms of this license agreement.
Using AnmiTali DeltaUI software products,
You agree that:
A) The Program is open source and you have the right to modify or
modify it in any way.
B) The Program is provided "as is", without warranty, express or
implied, or statutory,
including, but not limited to, warranties of quality, performance,
merchantability or fitness for a particular purpose.
C) There are also no guarantees created as a result
of the transaction, use or sales. AnmiTali does not
guarantee the smooth, timely and error-free operation
of the software. Under no circumstances will
AnmiTali be liable for damages or losses caused by the use or 
the inability to use this product. 
D) The software under this agreement is provided without express or
implied warranties of non-infringement, and the developer does not
guarantee non-infringement of any patents, copyrights, trade
secrets or other proprietary rights.
If you do not agree to the terms of the license,
immediately stop using this software!

Addition: The .apg application installer is still in development, and is also
distributed under the GNU General Public License 3.0.

Developers:
Ignatiev Ilya; Savin Yaroslav; Chakilev Arseny (In nasOS-Installer and for the idea) and the AnmiTali development team 
const Licencekz = `Лицензиялық келісім!

Сәлеметсіз бе, NurOS таңдағаныңыз үшін рахмет.
Бұл өнім лицензия бойынша таратылады
GNU General Public License 3.0
Толығырақ мына жерден оқи аласыз:
https://www.gnu.org/licenses/gpl-3.0.html#license-text

Deltaui өнімдеріне арналған лицензиялық келісім:
(Калькулятор, openserfing браузерінің бета нұсқасы)
Осы лицензиялық келісімнің шарттарымен танысыңыз.
AnmiTali deltaui бағдарламалық өнімдерін пайдалану,
Сіз бұған келісесіз:
А) бағдарлама ашық көзі болып табылады және сіз оны өзгертуге құқығыңыз бар немесе 
қалай болса да өзгерту керек.
Б) бағдарлама "сол күйінде", кепілдік міндеттемелерінсіз, анық немесе
көзделген немесе заңнамада көзделген, 
оның ішінде, бірақ онымен шектелмей, сапа кепілдігі, өнімділік,
сатуға немесе белгілі бір мақсатқа жарамдылық.
В) нәтижесінде жасалған кепілдіктер де берілмейді
мәміле жасау, пайдалану немесе сату. AnmiTali емес 
үздіксіз, уақтылы және қатесіз жұмыс істеуге кепілдік береді 
бағдарламалық қамтамасыз ету. Ешқандай жағдайда AnmiTali көтермейді 
пайдаланудан туындаған залал немесе залал үшін жауапкершілік немесе 
бұл өнімді пайдалану мүмкін .стігі. 
Г) осы Келісім бойынша айқын немесе 
бұзушылықтар туралы болжамды кепілдіктер жоқ және әзірлеуші бермейді 
кез келген патенттерді, авторлық құқықтарды, сауда құқықтарын бұзбау туралы кепілдіктер
құпиялар немесе басқа меншік құқықтары.
Егер сіз лицензия шарттарымен келіспесеңіз,
осы бағдарламалық жасақтаманы пайдалануды дереу тоқтатыңыз!

Қосымша: орнату бағдарламасы .APG қосымшалары әлі әзірленуде, сонымен қатар 
GNU General Public License 3.0 лицензиясы бойынша таратылады.

Әзірлеушілер:
Игнатьев Илья; Савин Ярослав; Чакилев Арсений (nasOS-Installer-де және идея үшін) және AnmiTali әзірлеушілер тобы

func showmenu(elem int, menu []string, title string) {
	fmt.Printf("\033c")
	// logo()
	fmt.Printf("\x1b[35m" + title + "\x1b[0m\n")
	for i := 0; i < len(menu); i++ {
		if elem == i {
			fmt.Printf("\x1b[47;30m" + menu[i] + "\x1b[0m\n")
		} else {
			fmt.Println(menu[i])
		}
	}
}

func selectlang(language string) string {
	// logo()
	title := welcomeru
	item := 0
	menu := []string{"Русский", "English", "Қазақша"}
	showmenu(item, menu, title)
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {

		switch key.Code {
		case keys.Up:
			if item != 0 {
				item--
			} else {
				item = len(menu) - 1
			}
			if item == 0 {
				title = welcomeru
				language = "ru"

			} else {
				title = welcomeen
				language = "en"
			
   			} else {
      				title = welcomekz
	  			language = "kz"

			showmenu(item, menu, title)

		case keys.Down:

			if item != len(menu)-1 {
				item++
			} else {
				item = 0
			}
			if item == 0 {
				title = welcomeru
				language = "ru"
			} else {
				title = welcomeen
				language = "en"
			
			} else {
      				title = welcomekz
	  			language = "kz"
			showmenu(item, menu, title)
		case keys.Enter:
			switch item {
			case 0:
				return true, nil
			case 1:
				return true, nil
			}
		case keys.CtrlC:
			return true, nil // Stop listener by returning true on Ctrl+C
		}

		return false, nil // Return false to continue listening
	})
	return language
}
func showLicense(language string) {
	fmt.Printf("\033c")
	// logo()
	title := ""
	exit := 0
	menu := []string{"Do not accept", "Accept"}
	if language == "ru" {
		title = Licenceru
		menu = []string{"Не принимаю", "Принимаю"}

	} else {
		title = Licenceen
		menu = []string{"Do not accept", "Accept"}
	} else {
		title = Licencekz
		menu = []string{"Бас тарту", "Қабылдау"}
	item := 0

	showmenu(item, menu, title)
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {

		switch key.Code {
		case keys.Up:
			if item != 0 {
				item--
			} else {
				item = len(menu) - 1
			}

			showmenu(item, menu, title)

		case keys.Down:

			if item != len(menu)-1 {
				item++
			} else {
				item = 0
			}

			showmenu(item, menu, title)
		case keys.Enter:
			switch item {
			case 0:
				exit = 1
				return true, nil
			case 1:
				exit = 0
				return true, nil
			}
		case keys.CtrlC:
			return true, nil // Stop listener by returning true on Ctrl+C
		}

		return false, nil // Return false to continue listening
	})
	if exit == 1 {
		cmd := exec.Command("shutdown", "now")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// Запускаем команду
		_ = cmd.Start()
		_ = cmd.Wait()
		os.Exit(0)
	}
}

func partiton(language string) string {
	title := ""
	drive := ""
	waiter := ""

	if language == "ru" {
		title = "Выберите диск для разметки"
		waiter = "Пожалуйста подождите 20 секунд."

	} else {
		title = "Select disk drive for partitioning"
		waiter = "Please wait 20 seconds."

	} else {
 		if language == "kz" {
		title = "Белгілеу үшін дискіні таңдаңыз"
		waiter = "20 секунд күтіңіз."

	_ = title
	devices, _ := filepath.Glob("/dev/[sS]d[a-zA-Z]")

	nvmeDevices, _ := filepath.Glob("/dev/nvme[0-9]n[0-9]")

	mmcDevices, _ := filepath.Glob("/dev/mmcblk[0-9]")

	allDevices := append(devices, nvmeDevices...)
	allDevices = append(allDevices, mmcDevices...)
	menu := allDevices
	item := 0

	showmenu(item, menu, title)
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {

		switch key.Code {
		case keys.Up:
			if item != 0 {
				item--
			} else {
				item = len(menu) - 1
			}

			showmenu(item, menu, title)

		case keys.Down:

			if item != len(menu)-1 {
				item++
			} else {
				item = 0
			}

			showmenu(item, menu, title)
		case keys.Enter:
			drive = allDevices[item]
			args := strings.Split("cfdisk"+" "+drive, " ")
			cmd := exec.Command(args[0], args[1])

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin

			// Запускаем команду
			err = cmd.Start()
			if err != nil {
				log.Fatal(err)
			}
			_ = cmd.Wait()
			return true, nil
		case keys.CtrlC:
			return true, nil // Stop listener by returning true on Ctrl+C
		}

		return false, nil // Return false to continue listening
	})
	fmt.Printf("\033c")
	fmt.Print(waiter)
	time.Sleep(20 * time.Second)
	return drive

}

func partboot(language string, instdrive string) string {
	devices := []string{}
	drive := ""
	if _, err := os.Stat("/sys/firmware/efi"); err == nil {
		// Объявляем пустой срез перед условным оператором
		re := regexp.MustCompile(`/dev/[sS]d[a-zA-Z]`)

		if re.MatchString(instdrive) {

			devices, _ = filepath.Glob(instdrive + "[0-9]") // Присваиваем значение внутри блока кода if
		} else {
			devices, _ = filepath.Glob(instdrive + "p[0-9]") // Присваиваем значение внутри блока кода else
		}
		title := ""
		if language == "ru" {
			title = "Выберите раздел для загрузчика (fat32)"

		} else {
			title = "Select partition for boot loader (fat32)"
		}
		menu := devices
		item := 0

		showmenu(item, menu, title)
		keyboard.Listen(func(key keys.Key) (stop bool, err error) {

			switch key.Code {
			case keys.Up:
				if item != 0 {
					item--
				} else {
					item = len(menu) - 1
				}

				showmenu(item, menu, title)

			case keys.Down:

				if item != len(menu)-1 {
					item++
				} else {
					item = 0
				}

				showmenu(item, menu, title)
			case keys.Enter:
				drive = menu[item]
				return true, nil
			case keys.CtrlC:
				return true, nil // Stop listener by returning true on Ctrl+C
			}

			return false, nil // Return false to continue listening
		})

	}
	return drive
}

func partsel(language string, instdrive string) string {
	devices := []string{} // Объявляем пустой срез перед условным оператором
	re := regexp.MustCompile(`/dev/[sS]d[a-zA-Z]`)

	if re.MatchString(instdrive) {

		devices, _ = filepath.Glob(instdrive + "[0-9]") // Присваиваем значение внутри блока кода if
	} else {
		devices, _ = filepath.Glob(instdrive + "p[0-9]") // Присваиваем значение внутри блока кода else
	}

	title := ""

	if language == "ru" {
		title = "Выберите раздел для установки системы"

	} else {

		title = "Select partition for installation system"
	} else {

		title = "Орнату жүйесі үшін бөлімді таңдаңыз"
	}
	drive := ""
	menu := devices
	item := 0

	showmenu(item, menu, title)
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {

		switch key.Code {
		case keys.Up:
			if item != 0 {
				item--
			} else {
				item = len(menu) - 1
			}

			showmenu(item, menu, title)

		case keys.Down:

			if item != len(menu)-1 {
				item++
			} else {
				item = 0
			}

			showmenu(item, menu, title)
		case keys.Enter:
			drive = menu[item]
			return true, nil
		case keys.CtrlC:
			return true, nil // Stop listener by returning true on Ctrl+C
		}

		return false, nil // Return false to continue listening
	})
	return drive

}
func sysinstall(language string, instpart string, bootpart string, instdrive string) {
	title := ""
	if language == "ru" {
		title = "Идет установка, пожалуйста, подождите. Это может занять несколько минут."
	} else {
		title = "Installing, please wait. This may take a few minutes."
	}
	args := strings.Split("mkfs.ext4"+" "+"-q"+" "+instpart, " ")
	cmd := exec.Command(args[0], args[1], args[2])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Запускаем команду
	_ = cmd.Start()
	_ = cmd.Wait()

	args = strings.Split("tune2fs"+" "+"-O"+" "+"^metadata_csum_seed"+" "+instpart, " ")
	cmd = exec.Command(args[0], args[1], args[2], args[3])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Запускаем команду
	_ = cmd.Start()
	_ = cmd.Wait()

	args = strings.Split("tune2fs"+" "+"-O"+" "+"^orphan_file"+" "+instpart, " ")
	cmd = exec.Command(args[0], args[1], args[2], args[3])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Запускаем команду
	_ = cmd.Start()
	_ = cmd.Wait()

	args = strings.Split("e2fsck"+" "+"-f"+" "+instpart, " ")
	cmd = exec.Command(args[0], args[1], args[2])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Запускаем команду
	_ = cmd.Start()
	_ = cmd.Wait()

	args = strings.Split("mount"+" "+instpart+" "+"/mnt", " ")
	cmd = exec.Command(args[0], args[1], args[2])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Запускаем команду
	_ = cmd.Start()
	_ = cmd.Wait()
	args = strings.Split("tar xvzf system.tar.gz -C /mnt", " ")
	cmd = exec.Command(args[0], args[1], args[2], args[3], args[4])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Запускаем команду
	_ = cmd.Start()
	_ = cmd.Wait()

	if bootpart != "" {

		fmt.Println(title)

		args := strings.Split("mkdir -p /mnt/boot/efi", " ")
		cmd := exec.Command(args[0], args[1], args[2])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// Запускаем команду
		_ = cmd.Start()
		_ = cmd.Wait()
		args = strings.Split("mkfs.fat -F32 "+bootpart, " ")
		cmd = exec.Command(args[0], args[1], args[2])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// Запускаем команду
		_ = cmd.Start()
		_ = cmd.Wait()
		args = strings.Split("mount"+" "+bootpart+" "+"/mnt/boot/efi", " ")
		cmd = exec.Command(args[0], args[1], args[2])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// Запускаем команду
		_ = cmd.Start()
		_ = cmd.Wait()
		cmd = exec.Command("bash", "./tab.sh")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// Запускаем команду
		_ = cmd.Start()
		_ = cmd.Wait()

		cmd = exec.Command("grub-install", "--boot-directory=/mnt/boot", instdrive, "--bootloader-id=elyzionos")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// Запускаем команду
		_ = cmd.Start()
		_ = cmd.Wait()

	} else {
		cmd = exec.Command("bash", "./tab.sh")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// Запускаем команду
		_ = cmd.Start()
		_ = cmd.Wait()
		args = strings.Split("arch-chroot /mnt grub-install "+instdrive, " ")
		cmd = exec.Command(args[0], args[1], args[2], args[3])
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, "PATH=\"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/bin:/usr/games:/sbin\"")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// Запускаем команду
		_ = cmd.Start()
		_ = cmd.Wait()

	}
	args = strings.Split("arch-chroot /mnt update-grub ", " ")
	cmd = exec.Command(args[0], args[1], args[2])
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "PATH=\"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/bin:/usr/games:/sbin\"")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Запускаем команду
	_ = cmd.Start()
	_ = cmd.Wait()

	fmt.Printf("\033c")
	if language == "ru" {
		fmt.Println("Пожалуйста, введите пароль суперпользователя. (Пароль не отображается)")
	} else {
		fmt.Println("Please enter the superuser password. (Password is not displayed)")
	}
	args = strings.Split("arch-chroot /mnt passwd", " ")
	cmd = exec.Command(args[0], args[1], args[2])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Запускаем команду
	_ = cmd.Start()
	_ = cmd.Wait()
	if language == "ru" {
		fmt.Println("Пожалуйста, введите пароль пользователя. (Пароль не отображается)")
	} else {
		fmt.Println("Please enter the user password. (Password is not displayed)")
	}
	args = strings.Split("arch-chroot /mnt passwd live", " ")
	cmd = exec.Command(args[0], args[1], args[2], args[3])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Запускаем команду
	_ = cmd.Start()
	_ = cmd.Wait()
	cmd = exec.Command("reboot")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Запускаем команду
	_ = cmd.Start()
	_ = cmd.Wait()

}
func main() {
	language := "ru"
	language = selectlang(language)
	showLicense(language)
	instdrive := partiton(language)
	fmt.Println(instdrive)
	bootpart := partboot(language, instdrive)
	instpart := partsel(language, instdrive)
	sysinstall(language, instpart, bootpart, instdrive)
	// selectpart(language)
	os.Exit(0)
}

//юзеры
//установка
