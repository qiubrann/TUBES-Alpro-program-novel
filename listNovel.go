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
	Genres string
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

var genreList = [MAX]Genre{
	{0, "Action, Dark-Fantasy, Adventure"},
	{1, "Mystery, Supernatural, Dark-Fantasy"},
	{2, "Dark-Fantasy, Psychological, Cultivation"},
	{3, "Fantasy, Isekai, Action"},
	{4, "Post-Apocalypse, Action, Tragedy"},
	{5, "Fantasy, Comedy, Reincarnation"},
	{6, "Sci-fi, Action, Reincarnation"},
	{7, "Cultivation, Regression, Fantasy"},
	{8, "Fantasy, Action, Slice-of-Life"},
	{9, "Fantasy, Villain-MC, Drama"},
	{10, "Martial-Arts, Revenge, Fantasy"},
	{11, "Horror, Mystery, Supernatural"},
	{12, "Fantasy, Magic, Action"},
	{13, "Fantasy, Adventure, Tragedy"},
	{14, "Fantasy, Action, Isekai"},
	{15, "Fantasy, Post-Apocalypse, Adventure"},
	{16, "Fantasy, Reincarnation, Action"},
	{17, "Action, Dungeon, Fantasy"},
	{18, "Fantasy, Isekai, Action"},
	{19, "Fantasy, Political, Drama"},
	{20, "Mystery, Supernatural, Fantasy"},
	{21, "Cultivation, Action, Revenge"},
	{22, "Fantasy, Regression, Action"},
	{23, "Fantasy, Drama, Action"},
	{24, "Fantasy, Villain-MC, Sci-fi"},
	{25, "Fantasy, Game-Elements, Isekai"},
	{26, "Fantasy, Action, Revenge"},
	{27, "Fantasy, Cultivation, Comedy"},
	{28, "Fantasy, Comedy, Kingdom-Building"},
	{29, "Sci-fi, Action, Isekai"},
	{30, "Fantasy, Isekai, Action"},
	{31, "Fantasy, Isekai, Psychological"},
	{32, "Fantasy, Dungeon, Revenge"},
	{33, "Fantasy, Regression, Action"},
	{34, "Comedy, Magic, Action"},
	{35, "Fantasy, Dungeon, Revenge"},
	{36, "Fantasy, Magic, Adventure"},
	{37, "Comedy, Villain-MC, Fantasy, Urban-Fantasy"},
	{38, "Fantasy, Drama, Political"},
	{39, "Fantasy, Magic, Kingdom-Building"},
	{40, "Martial-Arts, Action, Comedy"},
	{41, "Fantasy, Survival, Action"},
	{42, "Fantasy, Villain-MC, Reincarnation"},
	{43, "Cultivation, Slice-of-Life, Comedy"},
	{44, "Sci-fi, Villain-MC, Comedy"},
	{45, "Fantasy, Romance, Comedy"},
	{46, "Fantasy, Adventure, Romance"},
	{47, "Fantasy, Comedy, Action"},
	{48, "Dark-Fantasy, Villain-MC, Psychological"},
	{49, "Sci-fi, Action, Game-Elements"},
}

// Sort skor descending (selection sort)
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

// Sort judul ascending (insertion sort)
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

// Cari judul berdasarkan keyword (sequential search)
func cariJudul(title string) {
	found := false
	for i := 0; i < MAX; i++ {
		if strings.Contains(strings.ToLower(novelList[i].Title), strings.ToLower(title)) {
			fmt.Printf("%d. %s (Skor: %.1f)\n", novelList[i].ID, novelList[i].Title, scoreList[novelList[i].ID].Value)
			found = true
		}
	}
	if !found {
		fmt.Println("Novel tidak ditemukan.")
	}
}

// Bantu sort skor ascending untuk binary search
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

// Binary search skor
func cariSkor(target float64) {
	bantuCariSkor()
	left, right := 0, MAX-1
	found := false

	for left <= right && !found {
		mid := (left + right) / 2
		if scoreList[mid].Value == target {
			for i := mid; i >= 0 && scoreList[i].Value == target; i-- {
				fmt.Printf("%d. %s (Skor: %.1f)\n", scoreList[i].ID, novelList[scoreList[i].ID].Title, scoreList[i].Value)
				found = true
			}
			for j := mid + 1; j < MAX && scoreList[j].Value == target; j++ {
				fmt.Printf("%d. %s (Skor: %.1f)\n", scoreList[j].ID, novelList[scoreList[j].ID].Title, scoreList[j].Value)
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

// Tampilkan semua novel
func tampilNovelList() {
	for i := 0; i < MAX; i++ {
		fmt.Printf("%d. %s (%.1f) | Genre: %v\n", novelList[i].ID, novelList[i].Title, scoreList[i].Value, genreList[i].Genres)
	}
}

//cari berdasarkan genre
func cariGenre(keyword string) {
	keyword = strings.ToLower(keyword)
	found := false

	for i := 0; i < MAX; i++ {
		genres := strings.Split(genreList[i].Genres, ", ")
		foundInThisNovel := false

		for _, g := range genres {
			if strings.Contains(strings.ToLower(g), keyword) {
				if !foundInThisNovel {
					fmt.Printf("%d. %s (Skor: %.1f) | Genre: %v\n", novelList[i].ID, novelList[i].Title, scoreList[i].Value, genreList[i].Genres)
					foundInThisNovel = true
					found = true
				}
			}
		}
	}

	if !found {
		fmt.Println("Tidak ditemukan novel dengan genre tersebut.")
	}
}

// Main menu
func main() {

	var pilihan int
	for {
		fmt.Println("\n===== MENU =====")
		fmt.Println("1. Tampilkan Data")
		fmt.Println("2. Novel terbaik berdasarkan rating")
		fmt.Println("3. Novel berdasarkan abjad")
		fmt.Println("4. Cari Judul ")
		fmt.Println("5. Cari Skor ")
		fmt.Println("6. Cari berdasarkan genre")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilNovelList()
		case 2:
			urutDataSkor()
			fmt.Println("Novel terbaik berdasarkan rating:")
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
		case 6:
			var inputGenre string
			fmt.Print("Masukkan keyword genre: ")
			fmt.Scan(&inputGenre)
			cariGenre(inputGenre)
		case 0:
			fmt.Println("Terima kasih.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
