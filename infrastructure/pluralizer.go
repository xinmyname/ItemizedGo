package infrastructure

import (
	"regexp"
	"sort"
)

type _RuleTemplate struct {
	rule     string
	template string
}

var _uncountables = [...]string{
	"access", "accommodation", "adulthood", "advertising", "advice",
	"aggression", "aid", "air", "alcohol", "anger", "applause", "arithmetic",
	"art", "assistance", "athletics", "attention", "bacon", "baggage", "ballet",
	"beauty", "beef", "beer", "biology", "bison", "botany", "bread", "butter", "carbon",
	"cash", "chaos", "cheese", "chess", "childhood", "clothing", "coal", "coffee",
	"commerce", "compassion", "comprehension", "content", "corps", "corruption",
	"cotton", "courage", "cream", "currency", "dancing", "danger", "data", "deer",
	"delight", "dignity", "dirt", "distribution", "dust", "economics", "education",
	"electricity", "employment", "engineering", "envy", "equipment", "ethics", "evidence",
	"evolution", "faith", "fame", "fish", "flour", "flu", "food", "freedom", "fuel", "fun",
	"furniture", "garbage", "garlic", "genetics", "gold", "golf", "gossip", "grammar",
	"gratitude", "grief", "ground", "guilt", "gymnastics", "hair", "happiness", "hardware",
	"harm", "hate", "hatred", "health", "heat", "height", "help", "homework", "honesty",
	"honey", "hospitality", "housework", "humour", "hunger", "hydrogen", "ice",
	"importance", "inflation", "information", "injustice", "innocence", "iron", "irony",
	"jealousy", "jelly", "judo", "karate", "kindness", "knowledge", "labour", "lack",
	"laughter", "lava", "leather", "leisure", "lightning", "linguistics", "litter",
	"livestock", "logic", "loneliness", "luck", "luggage", "machinery", "magic",
	"management", "mankind", "marble", "mathematics", "mayonnaise", "means", "measles",
	"meat", "methane", "milk", "money", "moose", "mud", "music", "nature", "news",
	"nitrogen", "nonsense", "nurture", "nutrition", "obedience", "obesity", "oil",
	"oxygen", "passion", "pasta", "patience", "permission", "physics", "poetry",
	"pollution", "poverty", "power", "pronunciation", "psychology", "publicity",
	"quartz", "racism", "rain", "relaxation", "reliability", "research", "respect",
	"revenge", "rice", "rubbish", "rum", "salad", "satire", "scissors", "seaside",
	"series", "shame", "sheep", "shopping", "silence", "sleep", "smoke", "smoking", "snow",
	"soap", "software", "soil", "sorrow", "soup", "species", "speed", "spelling", "steam",
	"stuff", "stupidity", "sunshine", "swine", "symmetry", "tennis", "thirst", "thunder",
	"toast", "tolerance", "toys", "traffic", "transporation", "travel", "trust",
	"understanding", "unemployment", "unity", "validity", "veal", "vengeance", "violence",
}

var _ruleTemplates = [...]_RuleTemplate{
	_RuleTemplate{rule: "(th)is$", template: "$1ese"},
	_RuleTemplate{rule: "(th)at$", template: "$1ose"},
	_RuleTemplate{rule: "(millen)ium$", template: "$1ia"},
	_RuleTemplate{rule: "(l)eaf$", template: "$1eaves"},
	_RuleTemplate{rule: "(r)oof$", template: "$1oofs"},
	_RuleTemplate{rule: "(gen)us$", template: "$1era"},
	_RuleTemplate{rule: "(embarg)o$", template: "$1oes"},
	_RuleTemplate{rule: "arf$", template: "arves"},
	_RuleTemplate{rule: "^(b|tabl)eau$", template: "$1eaux"},
	_RuleTemplate{rule: "^(append|matr)ix$", template: "$1ices"},
	_RuleTemplate{rule: "^(ind)ex$", template: "$1ices"},
	_RuleTemplate{rule: "^(a)pparatus$", template: "$1pparatuses"},
	_RuleTemplate{rule: "^(a)lumna$", template: "$1lumnae"},
	_RuleTemplate{rule: "^(alg|vertebr|vit)a$", template: "$1ae"},
	_RuleTemplate{rule: "^(d)ie$", template: "$1ice"},
	_RuleTemplate{rule: "(m|l)ouse$", template: "$1ice"},
	_RuleTemplate{rule: "^(p)erson$", template: "$1eople"},
	_RuleTemplate{rule: "^(o)x$", template: "$1xen"},
	_RuleTemplate{rule: "^(c)hild$", template: "$1hildren"},
	_RuleTemplate{rule: "(g)oose$", template: "$1eese"},
	_RuleTemplate{rule: "(t)ooth$", template: "$1eeth"},
	_RuleTemplate{rule: "lf$", template: "lves"},
	_RuleTemplate{rule: "(f)oot$", template: "$1eet"},
	_RuleTemplate{rule: "^(wo|work|fire)man$", template: "$1men"},
	_RuleTemplate{rule: "(potat|tomat|volcan)o$", template: "$1oes"},
	_RuleTemplate{rule: "(criteri|phenomen)on$", template: "$1a"},
	_RuleTemplate{rule: "(nebul)a", template: "$1ae"},
	_RuleTemplate{rule: "oof$", template: "ooves"},
	_RuleTemplate{rule: "ium$", template: "ia"},
	_RuleTemplate{rule: "um$", template: "a"},
	_RuleTemplate{rule: "oaf$", template: "oaves"},
	_RuleTemplate{rule: "(thie)f$", template: "$1ves"},
	_RuleTemplate{rule: "fe$", template: "ves"},
	_RuleTemplate{rule: "(buffal|carg|mosquit|torped|zer|vet|her|ech)o$", template: "$1oes"},
	_RuleTemplate{rule: "o$", template: "os"},
	_RuleTemplate{rule: "ch$", template: "ches"},
	_RuleTemplate{rule: "sis$", template: "ses"},
	_RuleTemplate{rule: "(corp)us$", template: "$1ora"},
	_RuleTemplate{rule: "(cact|nucle|alumn|bacill|fung|radi|stimul|syllab)us$", template: "$1i"},
	_RuleTemplate{rule: "(ax)is", template: "$1es"},
	_RuleTemplate{rule: "(sh|zz|ss)$", template: "$1es"},
	_RuleTemplate{rule: "x$", template: "xes"},
	_RuleTemplate{rule: "(t|sp|r|l|b)y$", template: "$1ies"},
	_RuleTemplate{rule: "s$", template: "ses"},
	_RuleTemplate{rule: "$", template: "s"},
}

// PluralOf returns the plural of the specified word if the count is not equal to 1
func PluralOf(word string, count int) string {

	idx := sort.SearchStrings(_uncountables[:], word)
	found := idx < len(_uncountables) && _uncountables[idx] == word

	if count == 1 || len(word) == 0 || found {
		return word
	}

	for _, ruleTemplate := range _ruleTemplates[:] {
		re := regexp.MustCompile(ruleTemplate.rule)

		if re.MatchString(word) {
			return re.ReplaceAllString(word, ruleTemplate.template)
		}
	}

	return word
}
