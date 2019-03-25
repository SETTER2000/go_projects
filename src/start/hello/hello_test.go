package hello

/**
https://www.litres.ru/static/or3/view/or.html?art_type=4&file=31991950&bname=Go%2520%25D0%25BD%25D0%25B0%2520%25D0%25BF%25D1%2580%25D0%25B0%25D0%25BA%25D1%2582%25D0%25B8%25D0%25BA%25D0%25B5&art=27072032&user=13451207&uuid=054e7800-c1a8-11e7-aff2-0cc47a5453d6&cover=/static/bookimages/31/99/19/31991942.bin.dir/31991942.cover.jpg&uilang=ru
 */
import "testing"

func TestName(t *testing.T) {
	name := GetName()
	if name != "World!"{
		t.Error("Функция getName отдала неожиданное значение!")
	}
}
