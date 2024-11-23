package pools

var verbs = []string{"go:hogalu", "come:baralu", "eat:tindi", "drink:piDiyu", "sleep:nidre", "run:odalu",
	"walk:hogalu", "jump:koTTu", "sit:koLLu", "stand:iralu", "write:barodu", "read:occu",
	"speak:maataadalu", "listen:keLuvudu", "see:kaanu", "watch:nodalu", "hear:keLuvudu",
	"touch:thadeyalu", "feel:anubhavisu", "laugh:hasi", "cry:kaayalu", "sing:haadu", "dance:nritya",
	"play:krida", "work:kaaryamadu", "study:adhyayana", "teach:paDeyalu", "learn:kalikodu", "ask:prashne mADu",
	"answer:uttara mADu", "help:saahaya mADu", "like:ishtapaTru", "love:prema mADu", "hate:odhyama mADu",
	"understand:teredu", "forget:tilidu", "remember:tiLidu", "believe:vishwasa mADu", "think:chintane mADu",
	"decide:nirNaya mADu", "choose:thAri mADu", "hope:aashe", "wish:kaamane", "wait:kaalu",
	"start:arambhisu", "finish:kaTi", "open:kaDeyalu", "close:thuralu", "buy:kondu", "sell:beLasi",
	"pay:bhara", "borrow:tedu", "lend:koDalu", "save:kaDeyalu", "spend:khaarchu mADu", "give:koDalu",
	"take:koLLu", "take off:koDu", "put:thagalu", "bring:barisi", "take care:vaati", "build:barisu",
	"destroy:nAshamADu", "create:srishti mADu", "make:mADu", "fix:thiruga", "change:badalu",
	"repair:maramADu", "improve:uttara mADu", "increase:heccu", "decrease:kogalu", "cut:kattalu",
	"add:haakalu", "subtract:kaasu", "multiply:gunisu", "divide:vibhajisu", "measure:mELe mADu",
	"weigh:thoru", "count:gANi", "win:jayisu", "lose:kaLedu", "fight:yudhdha mADu", "argue:avaNa mADu",
	"agree:sahakara mADu", "disagree:virodhi mADu", "invite:aalisi", "accept:svIkArisu", "reject:naPratyAkhyAna",
	"join:joina mADu", "leave:biTTu", "return:hogalu", "arrive:baralu", "depart:hogalu", "wait:naale",
	"cook:agi", "clean:tholake", "wash:kudiyalu", "dry:kaayi", "fold:gaTTi", "open:thodi", "close:kaDisu",
	"cut:kattalu", "slice:thiruvudu", "peel:ogalu", "grind:kalu", "boil:kura", "fry:talusiru",
	"bake:bhakshana", "roast:kadalu", "stir:kelisiddu", "stir fry:hasi", "arrange:sajje mADu", "organize:sajjeyalu",
	"clean:wash", "change:badalu", "decorate:sajja", "repair:marama", "fix:thiruga", "build:barisu",
	"plant:kudumi", "water:paani", "cut:kattalu", "harvest:kaayalu", "pick:koDu", "eat:tindi", "drink:pinu",
	"arrive:baradu", "leave:neevu", "depart:hogi", "join:joina", "combine:samasthite", "separate:bheda",
	"combine:mishra", "empty:kola", "fill:haaki", "check:saMshodhana", "use:upayogi", "reuse:puna:upayogi",
	"borrow:tedu", "lend:koDalu", "pay:kaudi", "receive:padi", "take:thagolu", "offer:thayi",
	"provide:oputal", "look:nodalu", "stare:kannaa padedu", "glance:nirudi", "watch:kanaLige", "sight:olage",
	"return:muttu", "repay:parimara", "act:karama", "react:prathikaara", "respond:uttara",
	"reflect:drishti", "comment:pratibandha", "observe:kaanu", "conclude:nishkarshisu",
	"communicate:maataadalu", "express:vyakti mADu", "imagine:kalpana", "dream:sapna",
	"describe:vishleshana", "represent:pratinidhi", "inform:soochane", "report:jaana", "discuss:samsaara",
	"debate:prashna", "speak:maataadu", "state:nirdharana", "declare:prakashana", "reveal:velavane",
	"announcing:prakashana", "read:padhe", "listen:keLuvu", "hear:keLuvike", "feel:anubhava",
	"understand:teredu", "misunderstand:thappu", "confuse:sandesha", "question:prashna",
	"respond:uttara", "interpret:bhavisu", "analyze:vishleshane", "guess:anumanisi",
	"hypothesize:soochi", "test:pareekshana", "examine:pariksha", "observe:nodalu", "remember:tiLidu",
	"forget:alavade", "realize:thirugu", "acknowledge:anugraha", "recognize:arivu", "apply:upayogi",
	"ask:prashne", "request:vinnana", "insist:niroopaka", "agree:sahakara", "disagree:virodhisalu",
	"accept:sathya", "deny:naPratyaksha", "suggest:sujana", "propose:prastav", "confirm:nishchayisu",
	"invite:aaramayi", "reject:pratibandha", "deny:naPratyaksha", "question:prashna",
	"check:kattalu", "verify:naGalu", "validate:siddhanta", "debate:vivada", "interact:melu",
	"discuss:samsarane", "negotiate:samjhauta", "argue:yudhdha", "assist:saahaya", "help:shaakara",
	"empathize:ahimsa", "sympathize:sympathy", "respond:pratisthana", "offer:oparati",
	"realize:niyama", "understand:sakara", "live:jiivana", "survive:jeevithantraka", "thrive:sampannate",
	"suffer:dukkha", "enjoy:aananda", "celebrate:utsava", "rejoice:anandisu", "appreciate:kaagidi",
	"belong:kaala", "agree:sahakara", "disagree:virodhisalu", "obey:anumati", "command:adhikaara",
	"direct:samvidhaana", "lead:netrutva", "follow:pashchima", "teach:paDeyalu", "warn:alart",
	"advise:sujana", "prohibit:nirodha", "permit:paTravide", "consult:samvaadana", "ask:prashna", "inform:suchane"}