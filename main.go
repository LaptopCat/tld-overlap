package main

import (
	"os"
	"strings"
)

/* do not touch list below, run the make-tld-arr.py script to update it if needed*/
var tlds = [1467]string{"cpa", "vacations", "saarland", "mh", "bet", "creditunion", "ax", "trade", "ice", "pwc", "sb", "zero", "systems", "nz", "fedex", "信息", "lidl", "verisign", "cymru", "as", "software", "亚马逊", "itau", "bharti", "gu", "cool", "اتصالات", "promo", "download", "bt", "rogers", "москва", "fresenius", "lifestyle", "amfam", "virgin", "gallery", "gd", "monster", "nissan", "fund", "rwe", "mp", "realestate", "한국", "dnp", "condos", "flowers", "ga", "motorcycles", "io", "vermögensberater", "woodside", "lk", "archi", "partners", "pm", "house", "ferrero", "organic", "tc", "phone", "cat", "hyatt", "淡马锡", "md", "intuit", "手机", "政务", "భారత్", "ski", "webcam", "ලංකා", "redstone", "playstation", "genting", "lancaster", "通販", "online", "vc", "how", "school", "tech", "trv", "lilly", "energy", "markets", "tui", "futbol", "我爱你", "industries", "airforce", "voting", "airtel", "imamat", "vip", "services", "jmp", "marriott", "norton", "garden", "agency", "pid", "mm", "market", "gifts", "kitchen", "moda", "google", "info", "box", "softbank", "nba", "集团", "bw", "sina", "kred", "موريتانيا", "diy", "bible", "movie", "bcg", "expert", "tr", "chase", "hughes", "alipay", "fashion", "prod", "hiv", "nc", "fitness", "drive", "erni", "gg", "deal", "mlb", "lotte", "bot", "case", "срб", "band", "コム", "ferrari", "sex", "kpmg", "credit", "المغرب", "flights", "tokyo", "dad", "дети", "able", "kerryhotels", "amica", "عمان", "公益", "cbs", "pfizer", "rocher", "ieee", "pink", "tours", "youtube", "hitachi", "mtn", "mn", "olayan", "ricoh", "rsvp", "ollo", "surgery", "eus", "swatch", "next", "rocks", "cafe", "닷컴", "is", "hr", "jo", "shiksha", "kindle", "شبكة", "at", "天主教", "dating", "scot", "ma", "th", "democrat", "орг", "pioneer", "istanbul", "trust", "fi", "電訊盈科", "select", "gw", "country", "hotmail", "bz", "co", "pro", "澳門", "live", "like", "pccw", "ck", "佛山", "lat", "cleaning", "joburg", "bio", "mormon", "press", "reliance", "网址", "dm", "xin", "kids", "walmart", "got", "mk", "nissay", "locus", "bank", "am", "bms", "otsuka", "grocery", "goo", "porn", "kosher", "pramerica", "ελ", "allfinanz", "mobile", "sakura", "ভাৰত", "by", "bo", "ne", "guru", "bingo", "pics", "pet", "philips", "bv", "bg", "招聘", "नेट", "care", "works", "bike", "firmdale", "audio", "ارامكو", "pg", "uno", "aq", "boutique", "mattel", "gratis", "البحرين", "analytics", "schwarz", "accenture", "barcelona", "courses", "vi", "nr", "catholic", "cheap", "swiss", "music", "asda", "discover", "northwesternmutual", "uy", "style", "ভারত", "im", "dish", "mitsubishi", "forex", "ru", "sling", "mw", "taipei", "cz", "mv", "leclerc", "ie", "camera", "city", "food", "net", "football", "link", "healthcare", "recipes", "boo", "fit", "ruhr", "post", "flir", "rip", "auspost", "vegas", "lds", "中信", "quest", "td", "jaguar", "blue", "kp", "guide", "gs", "dvag", "mx", "honda", "solar", "safe", "mt", "gm", "ישראל", "quebec", "limo", "author", "бел", "yamaxun", "landrover", "in", "kg", "tatar", "toyota", "bid", "bb", "audible", "farm", "sandvikcoromant", "gov", "cc", "winners", "reise", "pnc", "hyundai", "photography", "gp", "za", "aetna", "android", "nec", "shop", "tickets", "commbank", "global", "voto", "بارت", "ストア", "afl", "clinic", "foo", "ngo", "dance", "بيتك", "coop", "pt", "degree", "land", "xfinity", "कॉम", "company", "christmas", "org", "cbre", "ikano", "travel", "inc", "wales", "yun", "abb", "سورية", "cf", "ファッション", "hosting", "kerrylogistics", "gay", "닷넷", "mma", "bridgestone", "showtime", "orange", "final", "weir", "ps", "grainger", "melbourne", "商店", "qpon", "mckinsey", "mr", "pru", "zuerich", "sg", "health", "cy", "bn", "family", "med", "be", "save", "watches", "bzh", "ipiranga", "dubai", "abogado", "hu", "ltda", "lv", "scholarships", "bar", "art", "technology", "aws", "consulting", "design", "cx", "express", "cipriani", "nf", "docs", "delta", "kddi", "dell", "gle", "ba", "bentley", "wtf", "international", "xyz", "network", "gripe", "คอม", "ovh", "hk", "sew", "ไทย", "pub", "ls", "men", "na", "мкд", "career", "hotels", "no", "mint", "pictures", "je", "business", "lc", "madrid", "total", "tushu", "tennis", "cw", "storage", "space", "tube", "lexus", "ong", "tatamotors", "ril", "bcn", "ລາວ", "ee", "fans", "hockey", "travelersinsurance", "cfd", "living", "sca", "бг", "onl", "rehab", "cyou", "juegos", "circle", "bloomberg", "juniper", "obi", "guge", "zip", "center", "mit", "გე", "kn", "helsinki", "br", "ею", "insure", "volvo", "家電", "digital", "shoes", "ht", "room", "sas", "team", "praxi", "dz", "call", "chintai", "العليان", "sydney", "tz", "guardian", "bradesco", "film", "sanofi", "bostik", "discount", "sk", "tirol", "zara", "do", "fm", "az", "channel", "goodyear", "golf", "krd", "rent", "rs", "omega", "vodka", "hisamitsu", "latino", "jot", "glass", "hangout", "bs", "taobao", "silk", "بازار", "nrw", "fido", "lundbeck", "tdk", "spot", "thd", "bj", "wiki", "rw", "exchange", "autos", "مصر", "limited", "airbus", "ms", "finance", "dealer", "builders", "koeln", "cv", "lr", "watch", "haus", "foundation", "itv", "physio", "phd", "abudhabi", "holdings", "com", "video", "tm", "racing", "xbox", "data", "blog", "ua", "world", "ch", "stockholm", "bond", "goldpoint", "wed", "香格里拉", "build", "bayern", "gold", "financial", "law", "stream", "tj", "vn", "luxe", "deals", "furniture", "homesense", "events", "deloitte", "kiwi", "lease", "ki", "aquarelle", "realtor", "nhk", "tci", "lpl", "ag", "doctor", "broker", "store", "lego", "engineer", "today", "ir", "natura", "alstom", "nico", "bing", "wtc", "edeka", "pk", "study", "agakhan", "fk", "politie", "codes", "faith", "berlin", "hdfc", "موقع", "eurovision", "km", "ren", "toys", "farmers", "page", "みんな", "fujitsu", "sharp", "hermes", "укр", "epson", "homes", "aol", "bananarepublic", "statebank", "nikon", "lb", "uk", "godaddy", "ink", "ads", "dot", "irish", "reisen", "associates", "imdb", "cooking", "இலங்கை", "yodobashi", "sncf", "pictet", "de", "tw", "vu", "amazon", "maif", "lgbt", "tunes", "jpmorgan", "nokia", "loan", "sx", "et", "li", "vet", "台湾", "top", "moto", "cloud", "nextdirect", "bi", "zm", "ibm", "wme", "mtr", "gn", "kinder", "avianca", "lifeinsurance", "oracle", "pf", "map", "republican", "baseball", "cuisinella", "ph", "香港", "vivo", "diamonds", "sr", "慈善", "ing", "jio", "jll", "free", "ro", "aw", "search", "shell", "whoswho", "party", "athleta", "aramco", "prudential", "mov", "immobilien", "vig", "gy", "xerox", "reviews", "cbn", "trading", "cfa", "chrome", "kerryproperties", "al", "uz", "toray", "java", "menu", "holiday", "theatre", "comsec", "cards", "my", "report", "联通", "sj", "se", "contractors", "jeep", "グーグル", "fly", "compare", "gb", "plumbing", "bosch", "datsun", "properties", "crown", "giving", "direct", "soccer", "sohu", "ninja", "games", "点看", "broadway", "prof", "gallo", "citi", "creditcard", "schmidt", "gmx", "சிங்கப்பூர்", "rugby", "viva", "tjmaxx", "网店", "arte", "frontdoor", "attorney", "ком", "pa", "fan", "ec", "walter", "netflix", "alibaba", "red", "xihuan", "abbvie", "ac", "pin", "best", "jp", "tab", "sbs", "canon", "basketball", "you", "pohl", "baby", "int", "skin", "allstate", "icu", "olayangroup", "capetown", "goog", "marketing", "guitars", "coffee", "globo", "np", "삼성", "poker", "radio", "ml", "ventures", "town", "ڀارت", "hkt", "target", "plus", "academy", "legal", "zone", "secure", "management", "scb", "esq", "soy", "ve", "大拿", "sh", "fox", "navy", "mil", "star", "mg", "vin", "sale", "si", "sport", "tiffany", "rest", "hospital", "af", "kim", "fr", "bnpparibas", "redumbrella", "akdn", "机构", "security", "lamer", "tel", "hair", "club", "st", "accountants", "сайт", "community", "gal", "arab", "pw", "corsica", "קום", "sz", "онлайн", "mo", "ایران", "sbi", "er", "il", "name", "tf", "temasek", "staples", "gea", "actor", "hdfcbank", "casa", "review", "institute", "இந்தியா", "read", "talk", "buzz", "srl", "tkmaxx", "infiniti", "fishing", "luxury", "samsung", "组织机构", "sc", "sa", "unicom", "clothing", "السعودية", "coupons", "rodeo", "immo", "etisalat", "jcb", "ug", "game", "supplies", "商标", "gi", "tvs", "الاردن", "中国", "calvinklein", "ye", "cruise", "bf", "americanfamily", "网站", "中國", "ftr", "desi", "py", "ooo", "education", "run", "us", "merckmsd", "cl", "cu", "museum", "gucci", "jewelry", "fage", "қаз", "new", "boston", "在线", "ευ", "tn", "coach", "rio", "mutual", "企业", "yoga", "wanggou", "tips", "audi", "gives", "surf", "时尚", "cityeats", "微博", "مليسيا", "dental", "day", "娱乐", "now", "bbc", "money", "iq", "anquan", "bm", "catering", "gf", "xxx", "bom", "shaw", "tax", "host", "健康", "amsterdam", "here", "mz", "booking", "banamex", "capital", "accountant", "brother", "mq", "भारत", "bd", "hamburg", "osaka", "kaufen", "pl", "george", "icbc", "gdn", "black", "ltd", "عرب", "teva", "भारोत", "世界", "lasalle", "boats", "wolterskluwer", "tl", "hn", "capitalone", "equipment", "stc", "kyoto", "eat", "kz", "oldnavy", "asia", "lol", "insurance", "همراه", "kia", "dvr", "tk", "training", "alsace", "love", "paris", "su", "ericsson", "tienda", "barefoot", "restaurant", "tg", "tattoo", "ott", "ad", "reit", "jprs", "bestbuy", "casino", "frogans", "ਭਾਰਤ", "schaeffler", "observer", "diet", "ng", "photo", "ally", "mobi", "rich", "ws", "lefrak", "pizza", "me", "dog", "wow", "tt", "dk", "yt", "taxi", "crs", "travelers", "smart", "cd", "salon", "cn", "ni", "cr", "nra", "संगठन", "rentals", "vana", "嘉里大酒店", "jetzt", "versicherung", "sener", "auto", "bargains", "dentist", "horse", "嘉里", "science", "shouji", "tools", "购物", "brussels", "smile", "locker", "nagoya", "भारतम्", "volkswagen", "windows", "loans", "ups", "ge", "spa", "mom", "frl", "lamborghini", "buy", "nyc", "blockbuster", "cricket", "computer", "aco", "click", "safety", "广东", "repair", "clinique", "om", "clubmed", "sy", "solutions", "aero", "stcgroup", "政府", "place", "marshalls", "moe", "قطر", "miami", "nab", "wien", "netbank", "cab", "abc", "book", "католик", "earth", "moi", "ubank", "hiphop", "lanxess", "firestone", "la", "photos", "kr", "рф", "americanexpress", "church", "كاثوليك", "forsale", "アマゾン", "au", "office", "memorial", "amex", "green", "realty", "cg", "sony", "parts", "shopping", "cruises", "nowruz", "msd", "qa", "careers", "nl", "yokohama", "mu", "viajes", "fire", "viking", "engineering", "chanel", "الجزائر", "statefarm", "ubs", "تونس", "date", "group", "cal", "mango", "nexus", "sm", "toshiba", "dunlop", "八卦", "theater", "ಭಾರತ", "gent", "news", "schule", "latrobe", "man", "gift", "pars", "yahoo", "stada", "media", "extraspace", "ae", "dtv", "wine", "investments", "sv", "open", "homedepot", "ss", "florist", "boehringer", "ceo", "gr", "fun", "gallup", "gt", "sucks", "tires", "play", "bbva", "mc", "makeup", "cern", "saxo", "lt", "joy", "谷歌", "ഭാരതം", "বাংলা", "weibo", "villas", "weber", "kuokgroup", "fail", "fo", "weather", "neustar", "cam", "新加坡", "fidelity", "website", "lipsy", "anz", "shia", "lu", "singles", "gh", "aaa", "site", "meet", "forum", "camp", "beer", "origins", "セール", "contact", "exposed", "arpa", "hm", "gl", "bauhaus", "suzuki", "gmbh", "wang", "ford", "university", "gmail", "ابوظبي", "okinawa", "shangrila", "delivery", "hbo", "email", "pn", "vote", "bmw", "work", "dds", "dj", "charity", "graphics", "hsbc", "protection", "bh", "pay", "uol", "social", "sfr", "flickr", "vermögensberatung", "app", "سودان", "bofa", "homegoods", "ntt", "es", "one", "tiaa", "ca", "va", "construction", "directory", "rexroth", "sky", "seek", "fairwinds", "sandvik", "song", "comcast", "gap", "lplfinancial", "nike", "jobs", "to", "seven", "پاکستان", "apple", "abbott", "claims", "cisco", "supply", "ryukyu", "wf", "ଭାରତ", "skype", "voyage", "azure", "maison", "fast", "monash", "beauty", "ar", "tv", "college", "بھارت", "show", "pr", "hot", "adult", "microsoft", "panasonic", "citic", "台灣", "lighting", "рус", "tmall", "london", "prime", "domains", "gq", "فلسطين", "beats", "ポイント", "餐厅", "mba", "win", "pharmacy", "studio", "飞利浦", "հայ", "dev", "sl", "lacaixa", "progressive", "ci", "書籍", "vanguard", "公司", "عراق", "edu", "cars", "移动", "coupon", "fyi", "ly", "support", "vlaanderen", "aig", "vision", "life", "mortgage", "jm", "kfh", "axa", "мон", "car", "enterprises", "ist", "samsclub", "网络", "chat", "army", "citadel", "weatherchannel", "eg", "mls", "feedback", "blackfriday", "商城", "visa", "aarp", "ke", "sexy", "zw", "gop", "auction", "vg", "durban", "ao", "ifm", "sap", "ky", "sd", "ggee", "emerck", "seat", "williamhill", "llp", "fish", "pe", "kw", "aeg", "jnj", "كوم", "barclays", "eco", "komatsu", "dclk", "it", "moscow", "中文网", "eu", "cba", "so", "richardli", "yachts", "dhl", "bbt", "lotto", "meme", "baidu", "gmo", "kpn", "lawyer", "yandex", "dabur", "frontier", "cash", "nfl", "cm", "re", "barclaycard", "ai", "africa", "tjx", "ping", "fj", "kh", "gbiz", "estate", "zappos", "productions", "ismaili", "クラウド", "sarl", "wedding", "cologne", "mini", "nu", "游戏", "biz", "help", "sn", "apartments", "食品", "caravan", "新闻", "lincoln", "nowtv", "ભારત", "امارات", "property", "llc", "dupont", "id"}
/* do not touch list above*/

func main() {
	if len(os.Args) < 2 {
		println("supply me with some arguments dumbass")
		println("example: " + os.Args[0] + " damn")
		os.Exit(1)
	} else {
		word := string(os.Args[1])
		word = strings.Replace(word, " ", "", -1)
		word = strings.ToLower(word)
		println("searching for tld overlaps with " + word)
		for _, tld := range tlds {
			if strings.HasSuffix(word, tld) && len(tld) != len(word) {
				domain := string(word[:len(word)-len(tld)])
				if !strings.HasSuffix(domain, ".") {
					println("found " + domain + "." + tld)
				}
			}
		}
	}
}