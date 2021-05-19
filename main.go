package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"text/template"
	"time"
)

// structure contenant toutes les données d'un artiste
type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Relations    map[string][]string `json:"DatesLocations"`
	Visible      bool
}

// structure contenant les données envoyées au html
type Result struct {
	Alert       bool
	ListeVilles []string //fixe : contient la liste des villes dans les filtres
	MaxVilles   int      //fixe : contient le nombre maximal de villes pour un artist dans les filtres
	MaxMembers  int      //fixe : contient le nombre maximal de membres pour un artiste dans les filtres
	Artist      []Artist //variable : contient les artistes à afficher
}

// si une variable string est présente dans un tableau
func Si_A_dans_tab_B(element string, tab []string) bool {

	for _, v := range tab {
		if v == element {
			return true
		}
	}
	return false
}

// sert a extraire les données d'une api
func Fun_api(result interface{}, adresse string) {

	vinutile, err1 := http.Get(adresse)
	vinutile2, err2 := ioutil.ReadAll(vinutile.Body)
	if err1 != nil {
		fmt.Println(err1)
	}
	if err2 != nil {
		fmt.Println(err2)
	}
	json.Unmarshal(vinutile2, result)
}

//initialiser le résultat avec une partie fixe
func Initialiser_result(artistes []Artist) Result {

	max_members := 1
	max_villes := 1
	liste_villes := make([]string, 0)
	liste_villes = append(liste_villes, "-----")
	var result Result

	for _, artist := range artistes {

		if max_members < len(artist.Members) {
			max_members = len(artist.Members)
		}

		if max_villes < len(artist.Relations) {
			max_villes = len(artist.Relations)
		}

		for key := range artist.Relations {
			vp := strings.Split(key, "-")

			if !Si_A_dans_tab_B(vp[0]+"-->"+vp[1], liste_villes) {
				liste_villes = append(liste_villes, vp[0]+"-->"+vp[1])
			}
		}
	}
	sort.Strings(liste_villes)

	result.Alert = false
	result.MaxMembers = max_members
	result.MaxVilles = max_villes
	result.ListeVilles = liste_villes
	return result
}

func main() {
	start := time.Now()
	api_artists := "https://groupietrackers.herokuapp.com/api/artists"
	api_relation := "https://groupietrackers.herokuapp.com/api/relation/"

	artist_0 := []Artist{
		{Id: 0, Name: "TARHOUNI Mohamed amine", CreationDate: 0, Visible: false,
			Image: "https://figurhouse.com/wp-content/uploads/2020/10/crash-bandicoot-jet-board-neca-300x300.jpg",
			Members: []string{"<a href=https://www.linkedin.com/in/mohamed-amine-tarhouni-26bb81211/ target=_blank> <i class='fab fa-linkedin fa-3x' style='margin-left:43%;'></i> </a>",
				"<a href=https://github.com/mohamedamine-tarhouni/ target=_blank><i class='fab fa-github-square fa-3x' style='margin-left:43%;'></i></a>",
				"<a href='mailto: mohamedamine.tarhouni@ynov.com'><i class='fas fa-at fa-3x' style='margin-left:43%;' ></i></a>"},
		},
		{Id: 0, Name: "LEICHNIG Coriane", CreationDate: 0, Visible: false,
			Image: "https://intrld.com/wp-content/uploads/2019/01/logo-qalf.png",
			Members: []string{"<a href=https://www.linkedin.com/in/coriane-leichnig-8386b61b6/ target=_blank><i class='fab fa-linkedin fa-3x' style='margin-left:43%;'></i></a>",
				"<a href=https://github.com/VeXs13 target=_blank><i class='fab fa-github-square fa-3x'style='margin-left:43%;'></i></a>",
				"<a href='mailto: coriane.leichnig@ynov.com'> <i class='fas fa-at fa-3x' style='margin-left:43%;'></i></a>"},
		},
		{Id: 0, Name: "FEKAIER Seif", CreationDate: 0, Visible: false,
			Image: "https://www.sortiesdvd.com/itunesimages/tv/saison/7411.jpg",
			Members: []string{"<a href=https://www.linkedin.com/in/seif-fekaier-1439081b6/ target=_blank><i class='fab fa-linkedin fa-3x' style='margin-left:43%;'></i></a>",
				"<a href=https://github.com/H4CK3R5-FS target=_blank><i class='fab fa-github-square fa-3x'style='margin-left:43%;'></i></a>",
				"<a href='mailto: seif.fekaier@ynov.com'> <i class='fas fa-at fa-3x'style='margin-left:43%;'></i></a>"},
		},
		{Id: 0, Name: "FOURNIE William", CreationDate: 0, Visible: false,
			Image: "https://e7.pngegg.com/pngimages/336/734/png-clipart-minion-oscar-wearing-blanket-minions-halloween-ghost-haunted-house-humour-boo-s-adventures-at-home-minions-halloween-thumbnail.png",
			Members: []string{"<i class='fas fa-times fa-3x' style='margin-left:43%;'></i>",
				"<i class='fas fa-times fa-3x'style='margin-left:43%;'></i>",
				"<a href='mailto: William.fournie@ynov.com '> <i class='fas fa-at fa-3x'style='margin-left:43%;'></i></a>"},
		},
	}

	var artist []Artist
	//initialiser la iste des artistes
	Fun_api(&artist, api_artists)

	for i := range artist {
		//ajouter les relations aux artistes
		Fun_api(&artist[i], api_relation+fmt.Sprint(artist[i].Id))
		artist[i].Visible = false
	}
	//initialiser le résultat avec une partie fixe
	result := Initialiser_result(artist)

	fmt.Println("le chargement de l'api a mis", time.Now().Sub(start), " Secondes")
	fmt.Println("-------------------")

	tmpl := template.Must(template.ParseFiles("template/Maquette.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(time.Now())
		//result reviens au même état que la ligne 106 a chaque passage
		result.Artist = Filtres(r, artist)
		result.Alert = false
		//la fonction filtres lui donnes les artistes à ajouter
		if len(result.Artist) == 0 {
			result.Artist = artist_0
			result.Alert = true
		}
		tmpl.Execute(w, result)
		//et l'envoie au html
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.ListenAndServe(":"+port, nil)

}

func Filtres(r *http.Request, artist []Artist) []Artist {
	var result []Artist

	for _, candidat := range artist {
		inter := r.PostFormValue("date-a")
		if inter != "" {
			if inter != strings.Split(candidat.FirstAlbum, "-")[2] {
				continue
			}
		}
		inter = r.PostFormValue("date-c")
		if inter != "" {
			if inter != fmt.Sprint(candidat.CreationDate) {
				continue
			}
		}
		inter = r.PostFormValue("members")
		if !(inter == "" || inter == "0") {
			if inter != fmt.Sprint(len(candidat.Members)) {
				continue
			}
		}
		inter = r.PostFormValue("villes")
		if !(inter == "" || inter == "0") {
			if inter != fmt.Sprint(len(candidat.Relations)) {
				continue
			}
		}
		not_pass := 0
		for i := 1; i <= 3; i++ {
			inter = r.PostFormValue("ville-" + fmt.Sprint(i))
			save := false
			if !(inter == "-----" || inter == "") {
				for key := range candidat.Relations {
					if inter == strings.Replace(key, "-", "-->", 1) {
						save = true
					}
				}
			} else {
				save = true
			}
			if save {
				not_pass++
			}
		}
		if not_pass != 3 {
			continue
		}
		//si je passe tous les filtres avec succés je suis ajouté sur la liste a envoyer au html
		result = append(result, candidat)
		if r.PostFormValue("artist-groupe") != "" {
			result[len(result)-1].Visible = true
		}
	}

	return result
}

//[{"id":1,
//"image":"https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
//"name":"Queen",
//"members":["Freddie Mercury","Brian May","John Daecon","Roger Meddows-Taylor","Mike Grose","Barry Mitchell","Doug Fogie"],
//"creationDate":1970,
//"firstAlbum":"14-12-1973",
//"locations":"https://groupietrackers.herokuapp.com/api/locations/1",
//"concertDates":"https://groupietrackers.herokuapp.com/api/dates/1",
//"relations":"https://groupietrackers.herokuapp.com/api/relation/1"},
