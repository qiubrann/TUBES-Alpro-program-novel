package main

import (
	"fmt"
	"strings"
)

const MAX = 50

type Novel struct {
	ID    int
	Title string
}

type Score struct {
	ID    int
	Value float64
}

type Genre struct {
	ID     int
	Genres []string
}

var novelList = [MAX]Novel{
	{0, "Shadow Slave"},
	{1, "Lord of the Mysteries"},
	{2, "Reverend Insanity"},
	{3, "The Beginning After The End"},
	{4, "Omniscient Reader's Viewpoint"},
	{5, "Trash of the Count's Family"},
	{6, "The Legendary Mechanic"},
	{7, "A Regressor’s Tale of Cultivation"},
	{8, "The Author's POV"},
	{9, "The Villain Wants to Live"},
	{10, "Nano Machine"},
	{11, "My House of Horrors"},
	{12, "Grandson of the Holy Emperor is a Necromancer"},
	{13, "The Second Coming of Gluttony"},
	{14, "The Novel's Extra"},
	{15, "The World After the Fall"},
	{16, "Damn Reincarnation"},
	{17, "Solo Leveling"},
	{18, "The Beginning After the End (Manhwa)"},
	{19, "How to Live as the Enemy Prince"},
	{20, "I’m Really Not The Demon God’s Lackey"},
	{21, "Emperor’s Domination"},
	{22, "Kidnapped Dragons"},
	{23, "The Demon Prince Goes to the Academy"},
	{24, "Little Tyrant Doesn’t Want to Meet with a Bad End"},
	{25, "I Fell into the Game with Instant Kill"},
	{26, "The Return of the Disaster-Class Hero"},
	{27, "The Regressed Demon Lord is Kind"},
	{28, "The Greatest Estate Developer"},
	{29, "I Became a Genius Engineer"},
	{30, "Everyone Else is a Returnee"},
	{31, "Re:Zero kara Hajimeru Isekai Seikatsu"},
	{32, "Kill the Hero"},
	{33, "SSS-Class Suicide Hunter"},
	{34, "Versatile Mage"},
	{35, "Ranker Who Lives A Second Time"},
	{36, "Sword God in a World of Magic"},
	{37, "Villain Retirement"},
	{38, "The Villainess Lives Twice"},
	{39, "Release That Witch"},
	{40, "The Return of the Crazy Demon"},
	{41, "The Extra’s Survival"},
	{42, "My Death Flags Show No Sign of Ending"},
	{43, "Beware of Chicken"},
	{44, "I’m the Evil Lord of an Intergalactic Empire!"},
	{45, "I Quit Being The Male Lead’s Rival"},
	{46, "Ending Maker"},
	{47, "To Be a Power in the Shadows!"},
	{48, "Dungeon Defense"},
	{49, "Supremacy Games"},
}

var scoreList = [MAX]Score{
	{0, 4.9}, {1, 4.9}, {2, 4.9}, {3, 4.8}, {4, 4.8}, {5, 4.8}, {6, 4.7}, {7, 4.7}, {8, 4.7}, {9, 4.7},
	{10, 4.6}, {11, 4.6}, {12, 4.6}, {13, 4.6}, {14, 4.6}, {15, 4.6}, {16, 4.6}, {17, 4.5}, {18, 4.5}, {19, 4.5},
	{20, 4.5}, {21, 4.5}, {22, 4.5}, {23, 4.5}, {24, 4.5}, {25, 4.5}, {26, 4.4}, {27, 4.4}, {28, 4.4}, {29, 4.4},
	{30, 4.4}, {31, 4.4}, {32, 4.4}, {33, 4.4}, {34, 4.4}, {35, 4.4}, {36, 4.3}, {37, 4.3}, {38, 4.3}, {39, 4.3},
	{40, 4.3}, {41, 4.3}, {42, 4.3}, {43, 4.3}, {44, 4.3}, {45, 4.2}, {46, 4.2}, {47, 4.2}, {48, 4.2}, {49, 4.2},
}

var genreList = [MAX]Genre{}

func listGenre() {
	genreList[0] = Genre{ID: 0, Genres: []string{"Action", "Dark-Fantasy", "Adventure"}}
	genreList[1] = Genre{ID: 1, Genres: []string{"Mystery", "Supernatural", "Dark-Fantasy"}}
	genreList[2] = Genre{ID: 2, Genres: []string{"Dark-Fantasy", "Psychological", "Cultivation"}}
	genreList[3] = Genre{ID: 3, Genres: []string{"Fantasy", "Isekai", "Action"}}
	genreList[4] = Genre{ID: 4, Genres: []string{"Post-Apocalypse", "Action", "Tragedy"}}
	genreList[5] = Genre{ID: 5, Genres: []string{"Fantasy", "Comedy", "Reincarnation"}}
	genreList[6] = Genre{ID: 6, Genres: []string{"Sci-fi", "Action", "Reincarnation"}}
	genreList[7] = Genre{ID: 7, Genres: []string{"Cultivation", "Regression", "Fantasy"}}
	genreList[8] = Genre{ID: 8, Genres: []string{"Fantasy", "Action", "Slice-of-Life"}}
	genreList[9] = Genre{ID: 9, Genres: []string{"Fantasy", "Villain-MC", "Drama"}}
	genreList[10] = Genre{ID: 10, Genres: []string{"Martial-Arts", "Revenge", "Fantasy"}}
	genreList[11] = Genre{ID: 11, Genres: []string{"Horror", "Mystery", "Supernatural"}}
	genreList[12] = Genre{ID: 12, Genres: []string{"Fantasy", "Magic", "Action"}}
	genreList[13] = Genre{ID: 13, Genres: []string{"Fantasy", "Adventure", "Tragedy"}}
	genreList[14] = Genre{ID: 14, Genres: []string{"Fantasy", "Action", "Isekai"}}
	genreList[15] = Genre{ID: 15, Genres: []string{"Fantasy", "Post-Apocalypse", "Adventure"}}
	genreList[16] = Genre{ID: 16, Genres: []string{"Fantasy", "Reincarnation", "Action"}}
	genreList[17] = Genre{ID: 17, Genres: []string{"Action", "Dungeon", "Fantasy"}}
	genreList[18] = Genre{ID: 18, Genres: []string{"Fantasy", "Isekai", "Action"}}
	genreList[19] = Genre{ID: 19, Genres: []string{"Fantasy", "Political", "Drama"}}
	genreList[20] = Genre{ID: 20, Genres: []string{"Mystery", "Supernatural", "Fantasy"}}
	genreList[21] = Genre{ID: 21, Genres: []string{"Cultivation", "Action", "Revenge"}}
	genreList[22] = Genre{ID: 22, Genres: []string{"Fantasy", "Regression", "Action"}}
	genreList[23] = Genre{ID: 23, Genres: []string{"Fantasy", "Drama", "Action"}}
	genreList[24] = Genre{ID: 24, Genres: []string{"Fantasy", "Villain-MC", "Sci-fi"}}
	genreList[25] = Genre{ID: 25, Genres: []string{"Fantasy", "Game-Elements", "Isekai"}}
	genreList[26] = Genre{ID: 26, Genres: []string{"Fantasy", "Action", "Revenge"}}
	genreList[27] = Genre{ID: 27, Genres: []string{"Fantasy", "Cultivation", "Comedy"}}
	genreList[28] = Genre{ID: 28, Genres: []string{"Fantasy", "Comedy", "Kingdom-Building"}}
	genreList[29] = Genre{ID: 29, Genres: []string{"Sci-fi", "Action", "Isekai"}}
	genreList[30] = Genre{ID: 30, Genres: []string{"Fantasy", "Isekai", "Action"}}
	genreList[31] = Genre{ID: 31, Genres: []string{"Fantasy", "Isekai", "Psychological"}}
	genreList[32] = Genre{ID: 32, Genres: []string{"Fantasy", "Dungeon", "Revenge"}}
	genreList[33] = Genre{ID: 33, Genres: []string{"Fantasy", "Regression", "Action"}}
	genreList[34] = Genre{ID: 34, Genres: []string{"Comedy", "Magic", "Action"}}
	genreList[35] = Genre{ID: 35, Genres: []string{"Fantasy", "Dungeon", "Revenge"}}
	genreList[36] = Genre{ID: 36, Genres: []string{"Fantasy", "Magic", "Adventure"}}
	genreList[37] = Genre{ID: 37, Genres: []string{"Comedy", "Villain-MC", "Fantasy", "Urban-Fantasy"}}
	genreList[38] = Genre{ID: 38, Genres: []string{"Fantasy", "Drama", "Political"}}
	genreList[39] = Genre{ID: 39, Genres: []string{"Fantasy", "Magic", "Kingdom-Building"}}
	genreList[40] = Genre{ID: 40, Genres: []string{"Martial-Arts", "Action", "Comedy"}}
	genreList[41] = Genre{ID: 41, Genres: []string{"Fantasy", "Survival", "Action"}}
	genreList[42] = Genre{ID: 42, Genres: []string{"Fantasy", "Villain-MC", "Reincarnation"}}
	genreList[43] = Genre{ID: 43, Genres: []string{"Cultivation", "Slice-of-Life", "Comedy"}}
	genreList[44] = Genre{ID: 44, Genres: []string{"Sci-fi", "Villain-MC", "Comedy"}}
	genreList[45] = Genre{ID: 45, Genres: []string{"Fantasy", "Romance", "Comedy"}}
	genreList[46] = Genre{ID: 46, Genres: []string{"Fantasy", "Adventure", "Romance"}}
	genreList[47] = Genre{ID: 47, Genres: []string{"Fantasy", "Comedy", "Action"}}
	genreList[48] = Genre{ID: 48, Genres: []string{"Dark-Fantasy", "Villain-MC", "Psychological"}}
	genreList[49] = Genre{ID: 49, Genres: []string{"Sci-fi", "Action", "Game-Elements"}}
}

// SELECTION SORT skor descending
func urutDataSkor() {
	for i := 0; i < MAX-1; i++ {
		maxIdx := i
		for j := i + 1; j < MAX; j++ {
			if scoreList[j].Value > scoreList[maxIdx].Value {
				maxIdx = j
			}
		}
		scoreList[i], scoreList[maxIdx] = scoreList[maxIdx], scoreList[i]
	}
}

// INSERTION SORT judul ascending
func urutAbjad() {
	for i := 1; i < MAX; i++ {
		temp := novelList[i]
		j := i - 1
		for j >= 0 && strings.ToLower(novelList[j].Title) > strings.ToLower(temp.Title) {
			novelList[j+1] = novelList[j]
			j--
		}
		novelList[j+1] = temp
	}
}

// SEQUENTIAL SEARCH berdasarkan judul
func cariJudul(title string) {
	found := false
	for i := 0; i < MAX; i++ {
		// kalau input kosong, skip
		if title == "" {
		}
		// pakai Contains biar bisa cari yang mengandung kata input
		if strings.Contains(strings.ToLower(novelList[i].Title), strings.ToLower(title)) {
			fmt.Printf("%d. %s (Skor: %.1f)\n", novelList[i].ID, novelList[i].Title, scoreList[novelList[i].ID].Value)

			found = true
		}
	}
	if !found {
		fmt.Println("Novel tidak ditemukan.")
	}
}

func bantuCariSkor() {
	for i := 1; i < MAX; i++ {
		temp := scoreList[i]
		j := i - 1
		for j >= 0 && scoreList[j].Value > temp.Value {
			scoreList[j+1] = scoreList[j]
			j--
		}
		scoreList[j+1] = temp
	}
}

func cariSkor(target float64) {
	bantuCariSkor()

	left, right := 0, MAX-1
	found := false

	for left <= right && !found {
		mid := (left + right) / 2

		if scoreList[mid].Value == target {
			// Ketemu, langsung scan kiri-kanan dari mid
			i := mid
			for i >= 0 && scoreList[i].Value == target {
				fmt.Printf("%d. %s (Skor: %.1f)\n", scoreList[i].ID, novelList[scoreList[i].ID].Title, scoreList[i].Value)

				i--
				found = true
			}
			j := mid + 1
			for j < MAX && scoreList[j].Value == target {
				fmt.Printf("%d. %s (Skor: %.1f)\n", scoreList[j].ID, novelList[scoreList[j].ID].Title, scoreList[j].Value)

				j++
				found = true
			}
			
			left = right + 1
		} else if scoreList[mid].Value < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if !found {
		fmt.Println("Skor tidak ditemukan.")
	}
}

func tampilNovelList() {
	for i := 0; i < MAX; i++ {
		fmt.Printf("%d. %s (%.1f)\n", novelList[i].ID, novelList[i].Title, scoreList[i].Value)
	}
}

func main() {
	listGenre()

	var pilihan int
	for {
		fmt.Println("\n===== MENU =====")
		fmt.Println("1. Tampilkan Data")
		fmt.Println("2. Novel terbaik berdasarkan rating")
		fmt.Println("3. Novel berdasarkan abjad")
		fmt.Println("4. Cari Judul ")
		fmt.Println("5. Cari Skor ")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilNovelList()
		case 2:
			urutDataSkor()
			fmt.Println("Novel terbaik berdasarkan rating: ")
			tampilNovelList()
		case 3:
			urutAbjad()
			fmt.Println("List Novel:")
			tampilNovelList()
		case 4:
			var inputTitle string
			fmt.Print("Masukkan keyword judul: ")
			fmt.Scan(&inputTitle)
			cariJudul(inputTitle)
		case 5:
			var targetScore float64
			fmt.Print("Masukkan skor yang dicari: ")
			fmt.Scan(&targetScore)
			cariSkor(targetScore)
		case 0:
			fmt.Println("Terima kasih.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
