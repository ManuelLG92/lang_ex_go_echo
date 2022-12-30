package language

import (
	"api.go.com/echo/config"
	"github.com/jinzhu/gorm"
)

type Language struct {
	gorm.Model
	Name    string `json:"name" gorm:"type:varchar(50)"`
	IsoCode string `json:"iso_code" gorm:"type:varchar(10)"`
}

func MigrateLanguagesTable() {
	config.DbGlobal.AutoMigrate(&Language{})
}

/*  VALUES('English', 'en') ,
('Afar', 'aa') ,
('Abkhazian', 'ab') ,
('Afrikaans', 'af') ,
('Amharic', 'am') ,
('Arabic', 'ar') ,
('Assamese', 'as') ,
('Aymara', 'ay') ,
('Azerbaijani', 'az') ,
('Bashkir', 'ba') ,
('Belarusian', 'be') ,
('Bulgarian', 'bg') ,
('Bihari', 'bh') ,
('Bislama', 'bi') ,
('Bengali/Bangla', 'bn') ,
('Tibetan', 'bo') ,
('Breton', 'br') ,
('Catalan', 'ca') ,
('Corsican', 'co') ,
('Czech', 'cs') ,
('Welsh', 'cy') ,
('Danish', 'da') ,
('German', 'de') ,
('Bhutani', 'dz') ,
('Greek', 'el') ,
('Esperanto', 'eo') ,
('Spanish', 'es') ,
('Estonian', 'et') ,
('Basque', 'eu') ,
('Persian', 'fa') ,
('Finnish', 'fi') ,
('Fiji', 'fj') ,
('Faeroese', 'fo') ,
('French', 'fr') ,
('Frisian', 'fy') ,
('Irish', 'ga') ,
('Scots/Gaelic', 'gd') ,
('Galician', 'gl') ,
('Guarani', 'gn') ,
('Gujarati', 'gu') ,
('Hausa', 'ha') ,
('Hindi', 'hi') ,
('Croatian', 'hr') ,
('Hungarian', 'hu') ,
('Armenian', 'hy') ,
('Interlingua', 'ia') ,
('Interlingue', 'ie') ,
('Inupiak', 'ik') ,
('Indonesian', 'in') ,
('Icelandic', 'is') ,
('Italian', 'it') ,
('Hebrew', 'iw') ,
('Japanese', 'ja') ,
('Yiddish', 'ji') ,
('Javanese', 'jw') ,
('Georgian', 'ka') ,
('Kazakh', 'kk') ,
('Greenlandic', 'kl') ,
('Cambodian', 'km') ,
('Kannada', 'kn') ,
('Korean', 'ko') ,
('Kashmiri', 'ks') ,
('Kurdish', 'ku') ,
('Kirghiz', 'ky') ,
('Latin', 'la') ,
('Lingala', 'ln') ,
('Laothian', 'lo') ,
('Lithuanian', 'lt') ,
('Latvian/Lettish', 'lv') ,
('Malagasy', 'mg') ,
('Maori', 'mi') ,
('Macedonian', 'mk') ,
('Malayalam', 'ml') ,
('Mongolian', 'mn') ,
('Moldavian', 'mo') ,
('Marathi', 'mr') ,
('Malay', 'ms') ,
('Maltese', 'mt') ,
('Burmese', 'my') ,
('Nauru', 'na') ,
('Nepali', 'ne') ,
('Dutch', 'nl') ,
('Norwegian', 'no') ,
('Occitan', 'oc') ,
('(Afan)/Oromoor/Oriya', 'om') ,
('Punjabi', 'pa') ,
('Polish', 'pl') ,
('Pashto/Pushto', 'ps') ,
('Portuguese', 'pt') ,
('Quechua', 'qu') ,
('Rhaeto-Romance', 'rm') ,
('Kirundi', 'rn') ,
('Romanian', 'ro') ,
('Russian', 'ru') ,
('Kinyarwanda', 'rw') ,
('Sanskrit', 'sa') ,
('Sindhi', 'sd') ,
('Sangro', 'sg') ,
('Serbo-Croatian', 'sh') ,
('Singhalese', 'si') ,
('Slovak', 'sk') ,
('Slovenian', 'sl') ,
('Samoan', 'sm') ,
('Shona', 'sn') ,
('Somali', 'so') ,
('Albanian', 'sq') ,
('Serbian', 'sr') ,
('Siswati', 'ss') ,
('Sesotho', 'st') ,
('Sundanese', 'su') ,
('Swedish', 'sv') ,
('Swahili', 'sw') ,
('Tamil', 'ta') ,
('Telugu', 'te') ,
('Tajik', 'tg') ,
('Thai', 'th') ,
('Tigrinya', 'ti') ,
('Turkmen', 'tk') ,
('Tagalog', 'tl') ,
('Setswana', 'tn') ,
('Tonga', 'to') ,
('Turkish', 'tr') ,
('Tsonga', 'ts') ,
('Tatar', 'tt') ,
('Twi', 'tw') ,
('Ukrainian', 'uk') ,
('Urdu', 'ur') ,
('Uzbek', 'uz') ,
('Vietnamese', 'vi') ,
('Volapuk', 'vo') ,
('Wolof', 'wo') ,
('Xhosa', 'xh') ,
('Yoruba', 'yo') ,
('Chinese', 'zh') ,
('Zulu', 'zu') ;*/
