package data

import "github.com/mukuvi/Tribes/internal/models"

var Tribes = []models.Tribe{
	// BANTU - Central Region
	{Name: "Kikuyu", Region: "Central Kenya", Tags: []string{"agriculture", "business", "entrepreneurship", "largest", "central"}, Population: 8148000, Counties: 5, Languages: 1},
	{Name: "Meru", Region: "Eastern Kenya", Tags: []string{"farming", "miraa", "education", "mt kenya", "ameru"}, Population: 1975000, Counties: 2, Languages: 1},
	{Name: "Embu", Region: "Eastern Kenya", Tags: []string{"agriculture", "tea", "coffee", "mt kenya"}, Population: 770000, Counties: 1, Languages: 1},
	{Name: "Mbeere", Region: "Eastern Kenya", Tags: []string{"millet", "sorghum", "dryland", "farming"}, Population: 250000, Counties: 1, Languages: 1},
	{Name: "Tharaka", Region: "Eastern Kenya", Tags: []string{"farming", "pastoral", "tana river", "dryland"}, Population: 175000, Counties: 1, Languages: 1},

	// BANTU - Eastern Region
	{Name: "Kamba", Region: "Eastern Kenya", Tags: []string{"woodcarving", "trading", "military", "sculpture", "eastern"}, Population: 4663000, Counties: 3, Languages: 1},

	// BANTU - Coastal Region
	{Name: "Mijikenda", Region: "Coast", Tags: []string{"nine tribes", "kaya forests", "coastal", "fishing", "traditions"}, Population: 2489000, Counties: 5, Languages: 9},
	{Name: "Taita", Region: "Coast", Tags: []string{"mining", "agriculture", "gemstones", "taitahills"}, Population: 350000, Counties: 1, Languages: 1},
	{Name: "Taveta", Region: "Coast", Tags: []string{"agriculture", "cross-border", "tanzania", "bananas"}, Population: 30000, Counties: 1, Languages: 1},
	{Name: "Pokomo", Region: "Coast", Tags: []string{"fishing", "farming", "tana river", "floodplains"}, Population: 115000, Counties: 1, Languages: 1},
	{Name: "Swahili", Region: "Coast", Tags: []string{"islam", "trade", "architecture", "coastal", "lamu"}, Population: 50000, Counties: 3, Languages: 1},
	{Name: "Bajuni", Region: "Coast", Tags: []string{"fishing", "islands", "lamu", "coastal", "swahili"}, Population: 15000, Counties: 1, Languages: 1},
	{Name: "Boni", Region: "Coast", Tags: []string{"hunter-gatherers", "forest", "lamu", "indigenous"}, Population: 12000, Counties: 1, Languages: 1},

	// BANTU - Western Region
	{Name: "Luhya", Region: "Western Kenya", Tags: []string{"agriculture", "hospitality", "farming", "community", "western"}, Population: 6823000, Counties: 6, Languages: 18},
	{Name: "Kisii", Region: "Nyanza", Tags: []string{"bananas", "soapstone", "carving", "highlands"}, Population: 2700000, Counties: 1, Languages: 1},
	{Name: "Kuria", Region: "Nyanza", Tags: []string{"agriculture", "pastoral", "migori", "cross-border"}, Population: 315000, Counties: 1, Languages: 1},
	{Name: "Suba", Region: "Nyanza", Tags: []string{"fishing", "lakeside", "victoria", "islands"}, Population: 160000, Counties: 1, Languages: 1},

	// NILOTES - Highland Nilotes
	{Name: "Kalenjin", Region: "Rift Valley", Tags: []string{"athletics", "running", "highlands", "farming", "champions"}, Population: 6358000, Counties: 7, Languages: 1},
	{Name: "Sengwer", Region: "Rift Valley", Tags: []string{"forest", "hunting", "cherangani", "indigenous"}, Population: 35000, Counties: 1, Languages: 1},
	{Name: "Sabaot", Region: "Rift Valley", Tags: []string{"mt elgon", "agriculture", "forest", "highlands"}, Population: 290000, Counties: 1, Languages: 1},
	{Name: "Terik", Region: "Rift Valley", Tags: []string{"farming", "nandi", "tea", "highlands"}, Population: 120000, Counties: 1, Languages: 1},
	{Name: "Nandi", Region: "Rift Valley", Tags: []string{"athletics", "tea", "warriors", "highlands"}, Population: 950000, Counties: 1, Languages: 1},
	{Name: "Kipsigis", Region: "Rift Valley", Tags: []string{"tea", "athletics", "agriculture", "highlands"}, Population: 2100000, Counties: 2, Languages: 1},
	{Name: "Tugen", Region: "Rift Valley", Tags: []string{"agriculture", "baringo", "pastoral"}, Population: 850000, Counties: 1, Languages: 1},
	{Name: "Marakwet", Region: "Rift Valley", Tags: []string{"irrigation", "escarpment", "farming"}, Population: 200000, Counties: 1, Languages: 1},
	{Name: "Pokot", Region: "Rift Valley", Tags: []string{"pastoral", "farming", "highlands", "warriors"}, Population: 750000, Counties: 2, Languages: 1},

	// NILOTES - River-Lake Nilotes
	{Name: "Luo", Region: "Nyanza", Tags: []string{"fishing", "music", "culture", "lakeside", "politics"}, Population: 5260000, Counties: 4, Languages: 1},

	// NILOTES - Plains Nilotes
	{Name: "Maasai", Region: "Rift Valley", Tags: []string{"pastoral", "warriors", "tourism", "culture", "heritage"}, Population: 1190000, Counties: 2, Languages: 1},
	{Name: "Samburu", Region: "Rift Valley", Tags: []string{"pastoral", "nomadic", "warriors", "northern", "culture"}, Population: 320000, Counties: 1, Languages: 1},
	{Name: "Turkana", Region: "North Western", Tags: []string{"nomadic", "pastoral", "arid", "resilient", "heritage"}, Population: 1020000, Counties: 1, Languages: 1},
	{Name: "Njemps", Region: "Rift Valley", Tags: []string{"fishing", "lake baringo", "pastoral", "smallest"}, Population: 40000, Counties: 1, Languages: 1},
	{Name: "Teso", Region: "Western Kenya", Tags: []string{"agriculture", "busia", "cross-border", "uganda"}, Population: 410000, Counties: 1, Languages: 1},
	{Name: "El Molo", Region: "North Eastern", Tags: []string{"fishing", "endangered", "lake turkana", "smallest"}, Population: 1100, Counties: 1, Languages: 1},

	// CUSHITIC - North Eastern
	{Name: "Somali", Region: "North Eastern", Tags: []string{"pastoral", "islam", "trade", "nomadic", "arid"}, Population: 2738000, Counties: 3, Languages: 1},
	{Name: "Borana", Region: "North Eastern", Tags: []string{"pastoral", "gada system", "oromo"}, Population: 355000, Counties: 1, Languages: 1},
	{Name: "Gabra", Region: "North Eastern", Tags: []string{"pastoral", "chalbi desert", "camel", "nomadic"}, Population: 125000, Counties: 1, Languages: 1},
	{Name: "Rendille", Region: "North Eastern", Tags: []string{"camel herding", "pastoral", "marsabit", "nomadic"}, Population: 90000, Counties: 1, Languages: 1},
	{Name: "Orma", Region: "Coast/North Eastern", Tags: []string{"pastoral", "tana river", "cattle", "nomadic"}, Population: 80000, Counties: 1, Languages: 1},
	{Name: "Watta", Region: "Coast", Tags: []string{"hunter-gatherers", "archery", "indigenous"}, Population: 15000, Counties: 1, Languages: 1},
	{Name: "Burji", Region: "Northern Kenya", Tags: []string{"trading", "agriculture", "ethiopia"}, Population: 40000, Counties: 1, Languages: 1},
	{Name: "Dasenach", Region: "North Western", Tags: []string{"pastoral", "lake turkana", "ethiopia"}, Population: 25000, Counties: 1, Languages: 1},
	{Name: "Sakuye", Region: "North Eastern", Tags: []string{"pastoral", "marsabit", "nomadic"}, Population: 25000, Counties: 1, Languages: 1},
	{Name: "Wayyu", Region: "North Eastern", Tags: []string{"pastoral", "oromo", "nomadic"}, Population: 20000, Counties: 1, Languages: 1},
	{Name: "Mukogodo", Region: "Laikipia", Tags: []string{"maasai", "yaaku", "hunter-gatherers"}, Population: 5000, Counties: 1, Languages: 1},
}
