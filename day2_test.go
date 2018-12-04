package main

import (
	"bufio"
	"log"
	"strings"
	"testing"
)

func TestDay2(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		/*
		   abcdef contains no letters that appear exactly two or three times.
		   bababc contains two a and three b, so it counts for both.
		   abbcde contains two b, but no letter appears exactly three times.
		   abcccd contains three c, but no letter appears exactly two times.
		   aabcdd contains two a and two d, but it only counts once.
		   abcdee contains two e.
		   ababab co*/
		sampleInput := `abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab`
		checksum := day2checksum(sampleInput)
		if checksum != 12 {
			t.Errorf("got %d expected 12", checksum)
		}
	})
	t.Run("part 1", func(t *testing.T) {
		checksum := day2checksum(day2input)
		if checksum != 7936 {
			t.Errorf("got %d expected 7936", checksum)
		}
	})
	t.Run("part 2 example", func(t *testing.T) {
		common := day2common(`abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz`)
		if common != "fgij" {
			t.Errorf("got %s expected fgij", common)
		}
	})
	t.Run("part 2 ", func(t *testing.T) {
		common := day2common(day2input)
		if common != "lnfqdscwjyteorambzuchrgpx" {
			t.Errorf("got %s expected lnfqdscwjyteorambzuchrgpx", common)
		}
	})
}

func day2checksum(input string) int {
	r := strings.NewReader(input)
	s := bufio.NewScanner(r)
	twos := 0
	threes := 0
	for s.Scan() {
		w := s.Text()
		letterCount := make(map[rune]int)
		for _, l := range w {
			letterCount[l] = letterCount[l] + 1
		}
		for _, v := range letterCount {
			if v == 2 {
				twos++
				break
			}
		}
		for _, v := range letterCount {
			if v == 3 {
				threes++
				break
			}
		}
	}
	return twos * threes
}

func day2common(input string) string {
	r := strings.NewReader(input)
	s := bufio.NewScanner(r)
	words := []string{}
	for s.Scan() {
		w := s.Text()
		if w == "" {
			continue
		}
		words = append(words, w)
	}
	log.Print(words)
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			d := distanceOne(words[i], words[j])
			if d {
				k := diffpos(words[i], words[j])
				log.Print(words[i], " ", words[j], " ", k)
				return words[i][:k] + words[i][k+1:]
			}
		}
	}
	return "FAIL"
}

func distanceOne(a, b string) bool {
	d := 0
	for i := range a {
		if a[i] != b[i] {
			d++
			//			log.Print(a, " ", i, " ", b)
			if d > 1 {
				return false
			}
		}
	}
	return d == 1
}

func diffpos(a, b string) int {
	for i := range a {
		if a[i] != b[i] {
			return i
		}
	}
	return -1
}

var day2input = `lnfgdsywjyleogambzuchirkpx
nnfqdskfjyteogambzuchirkpx
lnfqdvvwjyteofambzuchirkpf
lnfqdsvwjyteogvmbzuthirkpn
ltfqdsvwjyoeogambxuchirkpx
lnfqcsvwjytzogacbzuchirkpx
lpfpdsvwjyteogambyuchirkpx
pnfqdsvwjyteogqmbzuchinkpx
lnfqdsvwjytyopambzpchirkpx
lnfqisswjyteogadbzuchirkpx
lnfqdsuwjcteogambzuchirepx
lnfqdovwjnteigambzuchirkpx
lnfbdsvwjxteogambzuchirkax
lnfqdsawjyteogamxzuchiwkpx
lncqdsvwjoteogambzuchirfpx
lnfadsrwjyteogambzuchirktx
lnfqdsvhjyteorazbzuchirkpx
lwfqdsvwjytdogambzuchirkhx
lnfqdhvwjyteogambzuhairkpx
lnfqdsvwjytlogambzgchyrkpx
lnfqdsvwjyteogamnzjwhirkpx
lnfodsvwjyteogahuzuchirkpx
lnfqdsvtjyteogamvzwchirkpx
lnfqdsvwjzueogambzuxhirkpx
lnfqxsvljytxogambzuchirkpx
lnfqdsvwjyteogambpyvhirkpx
lqzqdsvwjnteogambzuchirkpx
lnfqdsvwjyteogumbzichirapx
lnfqdbvwjytedgaubzuchirkpx
lnfqdsvwpyteogabbhuchirkpx
nnfqdsvwryteogambzuchiwkpx
lrfqdsvwjeteogambzuchhrkpx
lnfqdsvwxyteogamtzucyirkpx
lnfsdsvwjyteogambzulhirknx
lnfqdsvwjyreoyambzmchirkpx
ltfqdsvwjytdogkmbzuchirkpx
lnfqwbvcjyteogambzuchirkpx
lnfqdsvwjyteogamrzfchirmpx
lnfqdsvwjqteogambzucwirkpy
lnfqdslwjyfeogambzuchikkpx
lnfqdsvwjybeogambzuchikjpx
lofqysvwjyteogasbzuchirkpx
lnfqusvwjyteogambzucverkpx
lnfqdsvwjyteogaibzuchfrkzx
lnfqdsvwjyleogabbzuchirkcx
lnfqdsvqjyteogambzuchdqkpx
lnfqdsvwjwtewgambzuciirkpx
lnfqisvwjatwogambzuchirkpx
lnfqdgvwjyteogambzuchizkix
lnfqdsxwjyteogambyuehirkpx
lpffdsvwjyteogamrzuchirkpx
lnirdsvwjyteogambzuchirkbx
lnfqdsxdjyteogazbzuchirkpx
lnfgdgvwyyteogambzuchirkpx
lnfqxsvwjyteogambzmcwirkpx
lnxqjsvwjyteogambzuchirkqx
lnrqdsvwjpteogkmbzuchirkpx
lnfqdstwjyteoeahbzuchirkpx
lnfqdsvwtuteogambzuchixkpx
lwfqvsvwjyteogambzughirkpx
lnkqdsvwjyfeogambzuuhirkpx
lvvqdsvwjyteogambzuchirkpn
jndqdsvwjyteogzmbzuchirkpx
enfqdszwjyteogambcuchirkpx
lnfqdsvwiyteogakbauchirkpx
lnfqdsvwjyteogtmbzxcdirkpx
fnfqdswwjyteogawbzuchirkpx
lnfqdsvwjydejqambzuchirkpx
lnqqdsvwjyteogambzucbdrkpx
lnfqdsvwjyteogadbzuchirxcx
lnfqdslwjyyeogambzulhirkpx
lnfqdsvwjytecgambzucgirkpb
lbmqdsvwjyteogamkzuchirkpx
lbfqdsvrjyteogambzuchirapx
lnfqdsmwjyteogambzucfarkpx
lnfqasvwoyteofambzuchirkpx
bnfudsvwjyteogambzucharkpx
lnfrdsvwjytxogambzuchirkpg
lbfqdsvwjyteagambzucdirkpx
lxfqdsvwjytuogambzucjirkpx
lnfqdsvwjcteogamyzuchiikpx
lnfodsvwjyteognmbzuchirapx
ltfqdsvwjytedgaxbzuchirkpx
lnfqdshwjyteogambzucsilkpx
lnfqdsvwpyteohambzuchitkpx
wnzqdsvwjyteogambiuchirkpx
lnfqdsvwayteogambzhchirkpw
ltfqcsvwjrteogambzuchirkpx
lnfqdsvwaytekgamvzuchirkpx
lnfqdsvwjyteogambzokpirkpx
lnfqysbwjyeeogambzuchirkpx
lnsqdsvwjyteogambzuchikkpd
lrfqdsvwjyteogahbzochirkpx
lnfqdsvwjyreoggmbzuchjrkpx
lxfqdsvwjyteogkmbzuchirkpp
enhqdbvwjyteogambzuchirkpx
jnfqdsvwjyteogamczuuhirkpx
lnfqdsvwuyteogadbzuchirkpw
lnfqdsvjjytergambznchirkpx
lnfqdsvwjyteoglmbzuceijkpx
lwfqdsvwjyteogamieuchirkpx
lnfqdsvwjfaeogambzqchirkpx
lfbqdjvwjyteogambzuchirkpx
lnfqdsvwjxteoaambzuchirkpp
lnfqdsvwjyheogjmbzgchirkpx
lnfqdskwjyteonambzuchiikpx
lnfqdwvwjyteogambxuchirkph
pnfqdsvwdyteogambzuchihkpx
lnoqdsvwjyteogaybznchirkpx
lnfqdsvwfyxefgambzuchirkpx
lnfqdsvwjyteotamxzmchirkpx
lnfqdsvwjyteigwmbzuchivkpx
lnfqdsvwjytekgambcuchirkwx
lnfqdsvwjuteogamrzulhirkpx
lnfqdsvwjyteogambzucczrgpx
wnfqzsvwjyteogambduchirkpx
lnfqdsowjyteogambuuthirkpx
lnfqdsvrjyteogcmbzuclirkpx
knfqdsvwgyteogambzuchorkpx
lnaqdsvwjytuogdmbzuchirkpx
lnfrdsvwjyteogambluchigkpx
lnfqqzvwjyteogambzkchirkpx
lnfqdsvwjyteogamuzuchgrkux
lnfqdsvnjyteogxmbznchirkpx
lnfqdsvwjyteolajbzuchdrkpx
lnfqdsvwjypeoiagbzuchirkpx
lnrqdsvwjyteozamuzuchirkpx
lnfqdsvwjytkogaubzucqirkpx
lnkbdsvwjyteogacbzuchirkpx
unfqdsvwjybeogambwuchirkpx
lnfqfsvzjyteogambzuchiikpx
lnfqdsvgjyreogahbzuchirkpx
lnfqdsewjyteogavbeuchirkpx
lnfqdsvwjdteogambbuchidkpx
lnfqdsvwjythogambzcchirkvx
lnfqdscwjyteorambzuchirgpx
cnfqdzvwjyteogambzushirkpx
lnfgdsgwjytedgambzuchirkpx
lnfqdsvwbyteogimbzuchdrkpx
lnfqdsswjyteogambzuohbrkpx
lnfqdsvwjytqogabbzachirkpx
lnfqdsvwjyteogmmbzucqiukpx
lnfxdsrwjyteogambzuchnrkpx
lnfqnqvwjyteogambzuchiwkpx
lffqisvwjyteogambzulhirkpx
lnfqdsxwjydeogambzucfirkpx
lnfqdsvwjyteogambzucjirkrp
lnfqdsnqjyteogambduchirkpx
fnfqdmvwjyteogamlzuchirkpx
lnfqvsvwjyteooamdzuchirkpx
lnfqdsvcyyteogambzuchickpx
onfqdsvwjyqeogambzuchirqpx
znfqdcvwjyteoaambzuchirkpx
lnfqdsvwjzteogambzuchidklx
lnfqjsvwjyteogjmbzuchirkpv
lnfqdsvwjytgorambzuchirppx
lzfqdsvwpfteogambzuchirkpx
lnfidsfwjyteogapbzuchirkpx
lnfodsvwbyteobambzuchirkpx
lnlqdsvwjytefgambzuchfrkpx
lnkqdsvwjyteogambzkchgrkpx
tnfqdsvwjyteoiamhzuchirkpx
lnfqdsvwjyteogamllschirkpx
lnfqdsvwjmthogamizuchirkpx
lnfqdbvwjyteogafbzuchirkpb
lnfxosvwjyteogahbzuchirkpx
lnmqdsvwjyzeogambzuchirkcx
lnfqdevbjytxogambzuchirkpx
lnfqdsvwjyteogamzzudhipkpx
lnfqdszwjyteoqambzuchirkpp
lffqdsvwjyteogamtouchirkpx
lnfqdsvhjytfogambzucharkpx
hnfqdsvwjyteogembzschirkpx
lnfqdsvwjateogambzuchirmpa
lnfqdsvcjyteogambzocairkpx
lnfqdsvwjyteogamwzmchirkpd
lnfqzsvwjyteogdmbzuyhirkpx
lnfqdsvwjytfyglmbzuchirkpx
lnfndsvwjyteogambzuchirktf
gnfqdnvwjytevgambzuchirkpx
lnfqdsvwjyteoganbpuchorkpx
lnfpdsvwnyteogambzucqirkpx
fnfqdstejyteogambzuchirkpx
lnfqlsvwjyteowambzuchirkmx
lnfqdsvwjyteogmmdzuchtrkpx
lnfqdsvwcyteogaqbzuchirkqx
lnfqdsvwjytlogtmbzuchiwkpx
lnfqdsvwoyteogambzuczirkwx
lnfqdsvwjyteogzybzucdirkpx
lnfqdvvwjyteogumbzuchiukpx
lnfqbwvwjyteogambzuchjrkpx
lnfgdsvwjyteogambzvchirkzx
lnfqdsvwjvtjogambzuchiokpx
lnfedsvwjyteogambzuchivkph
lhfqusvwjytaogambzuchirkpx
lnfqdsvwjyteogacbzuihirkpv
lnfwdsvwjyteogambzucokrkpx
lnfqtsvwjpteognmbzuchirkpx
anfqdswwjyteogambzucairkpx
lnfqdsvwjyteorambzuchirlsx
lnfqdsvwjytgogambzychirkpc
lnfqdhvwjyteogambzachirklx
lnfwdsvwjyteogaobquchirkpx
rnfqdsvwjiteogambzuhhirkpx
lnfqdsuwjyemogambzuchirkpx
hnfqdsvwjyteogambzuchprfpx
anfqssvwjyteogambzumhirkpx
lnfkdsvwjyteogafbzqchirkpx
lnfqdsvwjyteogacqzuchirspx
lnfqdskwjyteggambzuchiakpx
lnnqdsvwjyteooambzuchihkpx
lnlqdsvjjyteogambzuchgrkpx
lnfqdsvwjyteogamszochirkex
lnfqbsvwjyteogambzqchirepx
lnfqdsbwjcteogambzhchirkpx
lnfqdwvzjyteogambzechirkpx
ynfadsvwdyteogambzuchirkpx
tnfqdsvwjytuogambzuohirkpx
lnfqdsvwjyteogambzaohivkpx
mnfqisvwjyteogagbzuchirkpx
lnfqbsvwjyueogambzuchirkhx
ynfqdsvwjyteogdmbzuchinkpx
lnfqdwhwjyteogambzuchirqpx
mnfqdsvwjyteogambzfchkrkpx
lnfqdsnwjyteogambzgchiqkpx
lnfqdsvwjytergambzuchiuklx
lnfqdqvjjyteogamtzuchirkpx
lnfqdscwjyteorambzuchzrgpx
enfqdevwjyteogaabzuchirkpx
gnfqdsvbjyteogambzuchirkph
lnfqdxvwjyteogambzubhixkpx
lnfqdsvwjyteogambojchihkpx
lnfqdsvwjytdogambzuzhilkpx
lnfqdsvwjyteogamezuqhirtpx
tnfhdsvwjyteogambzuvhirkpx
lnfzdsvwjnteogahbzuchirkpx
lnfqdsvwjyteogambzfzhirkvx
lnfqqsvwjyteogambzuchirgpo
lufqpsvwjythogambzuchirkpx
lnfqdsvwjyteogzmbzuchimkix
lnwqdspwjyteogambzcchirkpx
lnfqdsowjyteogambzuchigypx
lnfqdnvvjyteogambzucjirkpx
lnfjdsvwryteogambzuchirkcx
lnfqdsvwbyteogambzuchirfpb
lnfqdsvwjyheogambzxchprkpx
lnfqmsvwjytezgambzuchirlpx
lnaqdsvwjyteogamdzuzhirkpx
lnoqdsvwjytebgambfuchirkpx
lnfqdtvwjytvogambzuchirkpv`
