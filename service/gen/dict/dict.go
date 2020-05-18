package dict

// prefix
const (
    // 天朝单姓
    CHINESE_PREFIX = iota + 1000
    // 天朝复姓
    CHINESE_PREFIX_TWO
    // 日本姓氏
    JAPANESE_PREFIX
    // 西方姓氏
    WEST_PREFIX
    // 功夫名字
    KONGFU_PREFIX
    // 天材地宝
    TREASURE_PREFIX
)

// suffix
const (
    // 天朝男性专用名
    CHINESE_MAN_SUFFIX = iota + 2000
    // 天朝女性专用名
    CHINESE_WOMAN_SUFFIX
    // 日本男性专用名
    JAPANESE_MAN_SUFFIX
    // 日本女性专用名
    JAPANESE_WOMAN_SUFFIX
    // 西方男性专用名
    WEST_MAN_SUFFIX
    // 西方女性专用名
    WEST_WOWAN_SUFFIX
    // 组织
    ORG_SUFFIX
    // 地点
    PLACE_SUFFIX
    // 功法
    KONGFU_SUFFIX
    // 天材地宝
    TREASURE_SUFFIX
    // 通用文字
    WORD
)

var (
    // 天朝单姓
    chinese_prefix = []string{"赵","钱","孙","李","周","吴","郑","王","冯","陈","诸","卫","蒋","沈","韩","杨","朱","秦","尤","许","何","吕","施","张","孔","曹","严","华","金","魏","陶","姜","戚","谢","邹","喻","柏","水","窦","章","云","苏","潘","葛","奚","范","彭","郎","鲁","韦","昌","马","苗","凤","花","方","俞","任","袁","柳","酆","鲍","史","唐","费","廉","岑","薛","雷","贺","倪","汤","滕","殷","罗","毕","郝","邬","安","常","乐","于","时","傅","皮","卡","齐","康","伍","余","元","卜","顾","孟","平","黄","和","穆","萧","尹","姚","邵","堪","汪","祁","毛","禹","狄","米","贝","明","臧","计","伏","成","戴","谈","宋","茅","庞","熊","纪","舒","屈","项","祝","董","粱","杜","阮","蓝","闵","席","季","麻","强","贾","路","娄","危","江","童","颜","郭","梅","盛","林","刁","钟","徐","邱","骆","高","夏","蔡","田","樊","胡","凌","霍","虞","万","支","柯","咎","管","卢","莫","经","房","裘","缪","干","解","应","宗","丁","宣","贲","邓","郁","单","杭","洪","包","诸","左","石","崔","吉","钮","龚","程","嵇","邢","滑","裴","陆","荣","翁","荀","羊","於","惠","甄","魏","家","封","芮","羿","储","靳","汲","邴","糜","松","井","段","富","巫","乌","焦","巴","弓","牧","隗","山","谷","车","侯","宓","蓬","全","郗","班","仰","秋","仲","伊","宫","宁","仇","栾","暴","甘","钭","厉","戎","祖","武","符","刘","景","詹","束","龙","叶","幸","司","韶","郜","黎","蓟","薄","印","宿","白","怀","蒲","台","从","鄂","索","咸","籍","赖","卓","蔺","屠","蒙","池","乔","阴","郁","胥","能","苍","双","闻","莘","党","翟","谭","贡","劳","逄","姬","申","扶","堵","冉","宰","郦","雍","却","璩","桑","桂","濮","牛","寿","通","边","扈","燕","冀","郏","浦","尚","农","温","别","庄","晏","柴","翟","阎","充","慕","连","茹","习","宦","艾","鱼","容","向","古","易","慎","戈","廖","庚","终","暨","居","衡","步","都","耿","满","弘","匡","国","文","寇","广","禄","阙","东","殴","殳","沃","利","蔚","越","夔","隆","师","巩","厍","聂","晁","勾","敖","融","冷","訾","辛","阚","那","简","饶","空","曾","毋","沙","乜","养","鞠","须","丰","巢","关","蒯","相","查","后","荆","红","游","竺","权","逯","盖","后","桓","公",}
    // 天朝复姓
    chinese_prefix_two = []string{"欧阳","太史","端木","上官","司马","东方","独孤","南宫","万俟","闻人","夏侯","诸葛","尉迟","公羊","赫连","澹台","皇甫","宗政","濮阳","公冶","太叔","申屠","公孙","慕容","仲孙","钟离","长孙","宇文","司徒","鲜于","司空","闾丘","子车","亓官","司寇","巫马","公西","颛孙","壤驷","公良","漆雕","乐正","宰父","谷梁","拓跋","夹谷","轩辕","令狐","段干","百里","呼延","东郭","南门","羊舌","微生","公户","公玉","公仪","梁丘","公仲","公上","公门","公山","公坚","左丘","公伯","西门","公祖","第五","公乘","贯丘","公皙","南荣","东里","东宫","仲长","子书","子桑","即墨","达奚","褚师","吴铭",}
    // 日本姓氏
    japanese_prefix = []string{"鹤田","香取","野泽","麻生","小田切","草翦","稻垣","木村","中居","濑户","山下","酒井","松本","石田","柴崎","藤原","福山","江口","唐泽","长泽","椎名","松岛","白石","铃木","堂本","仲间","织田","泷泽","妻夫木","药师丸","余贵","石黑","丰川","平宫","工藤","赤西","生田","高岛","松山","井之原","锦户","城田","竹野内","广末","二宫","石垣","小松","小栗","田中","滨崎","滨田","幸田","志田","香椎","山本","原田","永山","栗山","前田","冈部","忍成","寺岛","黑木","水野","伊势谷","野口","土屋","北乃","绫濑","泽尻","荣仓","加藤","宫崎","风间","户田","山口","井川","深田","米仓","佐藤","小池","上野","伊东","须藤","长濑","倍赏","岸谷","赤坂","中村","相叶","今井","黑川","伊藤","五十岚","冈田","野际","岛谷","堤","加濑","吉田","观月","深津","洼冢","役所","山田","吹石","吉冈","内田","阿部","吉泽","松田","长谷川","国仲","上川","北村","宝生","京野","天海","中山","中谷","香川","吉永","冈本","相武","向井","稻森","成海","市川","玉山","龟梨","松下","高桥","仲代","井上","吉川","手冢","友坂","宫泽","樱井","大野","多部未","上户","平冈","能濑","手越","宇多田","仓木","安室奈","美木","小野","中岛","竹中","中井","吉高","安藤","川岛","菊川","管野","安倍","市原","小泉","苍井","加藤","浅野","冢本","筱原","白川","村川","矢田","三浦","入江","管谷","小仓","水岛","大政","上原","蛯原","津川","阵内","内山","江角","柳叶","西田","常盘","树木","高冈","泽口","南野","田口","相田","相马","押尾","佐佐木","秋山","北川","松坂","高仓","三船","栗原","松雪","横山","武田","岩佐","丹波","行定","渡部","本木","桃井","储形","乙羽","大冢","泽村","中越","夏川","森田","三宅","坂本","华原","细川","小林","渥美","泽田","北野","黑川","小室","寺尾","今村","小津","深作","大岛","玉置","筱田","寺山","若松","黑泽","沟口","押井","岩井","谷村","宇津","西村","矢泽","稻山","吉武","八尾","古尾谷","贯地谷","不破","若月","高村","伊佐","牛岛","杉山","神木","松川","要","堺","本乡","水川","释由","石原","藤木","平山","笕利","饭田","饭岛","堀北","广濑","藤井","片濑","谷原","金子","江户川","福田","津岛","横沟","佐野","丸山","平井","柳井","有坂","水桥","铜谷","草野","内博","南泽","樱庭","新垣","末永","伊崎","森村","高木","川端","沟端","横光","芥川","矢井田","藤田","森山","持田","一青","松尾","尾崎","小川","大江","三岛","清少","夏目","清水","爱内","伴都","黑石","古谷","松浦","清浦","后藤","远藤","增田","小山","满","野间","村上","森","三枝","竹井","坂井","新居","石川","藤本","大仓","安部","池田","岸本","岩田","北原","宇德","上木","近江","水树","小出","冈崎","加护","玉木","奥井","中原","植田","植草","东山","梶浦","绀野","原田","原纱","金田","锦织","能登","牧野","堀江","石松","堀内","广桥","池泽","南里","千叶","小西","近藤","三木","折笠","河原木","神田","野中","野岛","川澄","种村","桑岛","宫小路","福井","丰崎","藤堂","西门","花泽","桧山","新谷","高泽见","大谷","久川","早见","纪野","茅原","中森","大原","神谷","藤村","户松","阪口","桑谷","小林","小野","上杉","源","饭冢","菊地","生天目","名冢","武内","新井","横手","越智","松谷","岭","樱内","齐藤","斋藤","太田","木下","福永","千野","鸠山","渡边","菊池","美部浓","末弘","平山","石桥","大久保","秋月","竹内","武见","松冈","岸","犬神","金田一","竹下","内藤","柏原","泉谷","大泉","森高","森下","牛尾","安西","正田","小和田","黑田","森嘉","松崎","森永","加纳","野田","荒船","近卫","细川护","千","江崎","叶山","濑名","龟山","杉尾","臼井","久保田","奥泽","小石川","冰室","朝仓","杉崎","星野","矢吹","真壁","生野","沟口","冲岛","町田","田村","西川","小泽","池内","大淹","梅田","山崎","北田","小岩井","片桐","内野","水原","纯名","黑崎","森口","吉本","井筒","筒井","井之上","长岭","浅见","野村","品川","生濑","黑谷","游川","八木","土井","难波","片山","北井","别所","五代","田渊","小椋","高丸","市村","长冢","秋吉","吹越","日向志","藏原","长井","杉村","奥贯","望月","井田","桥爪","神尾","道明寺","美作","三条","大河原","青池","日向","中岛","重村","堀口","楠田","周防","宅间","小牧","重冈","星谷","佐伯","江黑","坂上","笹峰","浅井","利根川","山野","宫下","赤井","家富","飞松","樱田","山室","水黑","彩田","大卫","栗卷","佐田","石野","富浦","加贺","坪井","三城","武藤","佐竹","织部","鹤见","水月","桥田","田岛","岩本","西浦","叶野","泷村","日比","野弥","小柳","北岛","宫林","胜亦","大森","美山","大杉","中江","平野","堂岛","大泽","田山",}
    // 西方姓氏
    west_prefix = []string{"亚伦","亚伯","亚伯拉罕","亚当","艾德里安","阿尔瓦","亚历克斯","亚历山大","艾伦","艾伯特","阿尔弗雷德","安德鲁","安迪","安格斯","安东尼","亚瑟","奥斯汀","本","本森","比尔","鲍伯","布兰登","布兰特","布伦特","布莱恩","布鲁斯","卡尔","凯里","卡斯帕","查尔斯","采尼","克里斯","克里斯蒂安","克里斯多夫","科林","科兹莫","丹尼尔","丹尼斯","德里克","唐纳德","道格拉斯","大卫","丹尼","埃德加","爱德华","艾德文","艾略特","埃尔维斯","埃里克","埃文","弗朗西斯","弗兰克","富兰克林","弗瑞德","加百利","加比","加菲尔德","加里","加文","乔治","基诺","格林","格林顿","哈里森","雨果","汉克","霍华德","亨利","伊格纳缇伍兹","伊凡","艾萨克","杰克","杰克逊","雅各布","詹姆士","詹森","杰弗瑞","杰罗姆","杰瑞","杰西","吉姆","吉米","乔","约翰","约翰尼","约瑟夫","约书亚","贾斯汀","凯斯","肯","肯尼斯","肯尼","凯文","兰斯","拉里","劳伦特","劳伦斯","利安德尔","李","雷欧","雷纳德","利奥波特","劳伦","劳瑞","劳瑞恩","卢克","马库斯","马西","马克","马科斯","马尔斯","马丁","马修","迈克尔","麦克","尼尔","尼古拉斯","奥利弗","奥斯卡","保罗","帕特里克","彼得","菲利普","菲比","昆廷","兰德尔","伦道夫","兰迪","列得","雷克斯","理查德","里奇","罗伯特","罗宾","罗宾逊","洛克","罗杰","罗伊","赖安","萨姆","萨米","塞缪尔","斯考特","肖恩","肖恩","西德尼","西蒙","所罗门","斯帕克","斯宾塞","斯派克","斯坦利","史蒂文","斯图亚特","特伦斯","特里","蒂莫西","汤米","汤姆","托马斯","托尼","泰勒","范","弗恩","弗农","文森特","沃伦","卫斯理","威廉",}
    // 功夫名字
    kongfu_prefix = []string{"准提","穿日","断肠","极光","无间","龙虎","玄雷","虎皮","悲问","老阳","凌霄","崩山","活杀","白虎","自然","地煞","摩诃","邪风","斩龙","蛤蟆","雪花","灭神","屠仙","盘古","摧日","陨星","飞霜","封神","赤铁","回春","金刚","赤星","离魂","凤舞","鱼翔","一刹","悟禅","紫虹","飞鱼","天网","罗刹","斩月","绝神","盖世","雷火","梁尘","蓮花","破地","赤虹","龙圣","无上","孔雀","拾柒","明心","血月","苍羽","续断","十方","天元","阿含","拂柳","至高","无华","天浪","断龙","仙佛","血魔","赤炎","无踪","翔龙","烈日","白月","巨鲸","灭魔","雷音","惊虹","托月","鸿爪","五色","圣光","戒冥","空蝉","飞鹤","兑泽","虎爪","飞雪","药叉","三泰","震天","六道","飞仙","虎狼","神鸟","龙脑","灭妖","水月","灵龙","双皇","鱼龙","天悲","凄煌","雨霖","银月","大蛇","菩提","法性","玄武","太极","辟心","临枫","朱雀","飞蝎","释迦","霓霞","哭魂","蚀精","水仙","六门","童子","翔鹰","碎肉","罗汉","玄鬼","白雾","咆哮","钢骨","缘生","陨日","吸血","翠羽","天武","闪电","血焰","金钟","渡缘","金蚕","涡卷","天虎","金焰","紫羽","火凤","捕风","百炼","凤鸟","托天","倒海","灵光","傲世","暗鸦","真元","梅花","豹爪","追魂","灵鹤","碧羽","极火","风火","血虹","星象","玄云","无限","参妖","斩海","闇炎","青霞","虎豹","钢翼","菩萨","浑天","天道","千鹤","银光","炼神","无音","生死","十八","七色","妖灵","狼牙","炼狱","碧炎","踏云","圣兽","销魂","紫阳","大衍","天魔","赤凤","黑龙","风刃","断空","阳明","凤凰","邪影","绝命","玄冥","断魂","紫雾","玄光","锁穴","封魔","真武","飞燕","绝对","辟邪","紫霄","乘风","贯清","金蝉","焚魂","终极","飞狼","血阳","重生","金乌","闇影","飘雪","鳞光","暗黑","烈光","闼婆","地狱","暴虎","水鸟","仁王","北斗","般若","易筋","惊天","倚月","射手","逍遥","拾捌","风魔","镰鼬","飞狐","圆融","残阳","定禅","至尊","南斗","夜叉","醉梦","旋锋","魔灵","佛光","狼王","三清","银环","天极","夺魄","狂风","灵猴","丹霞","流离","断筋","玄阳","龙凰","神鹰","丹鼎","霸鲸","灵霄","灭世","赤月","青雷","九宫","紫凰","无痕","紫霜","孤灯","千蛛","映月","滴血","天蚕","追星","斩铁","游鱼","灵心","狮子","无双","天鹤","螳螂","天火","白光","大日","无影","花蝶","飞虎","太初","枯禅","破元","无敌","飞星","震雷","斩钢","迴旋","日扬","溷元","皇极","狂浪","九天","八卦","清风","寒月","碧蛇","虎牙","焚野","九阳","奔雷","皇龙","真龙","擎天","赤羽","七彩","奇木","惊魂","玛瑙","凶神","腾龙","天罡","霸王","赤龙","天龙","霓光","千里","凤爪","白焰","灵狐","龙凤","暴雨","血海","灵蛇","渡罪","森罗","冰锋","日月","十绝","无悔","真传","六星","邪光","三阴","凶灵","陆龟","天心","摧魂","旋风","封脉","闇雷","蛊风","风霜","鬼神","荼毒","射日","万丈","昇龙","玄月","迦蓝","离恨","疾风","罗王","如意","琉璃","炼佛","穿山","白虹","龙胆","无妄","霸皇","除仙","巨鲨","掣雷","天鬼","抱残","惊龙","碎骨","金虎","缘灭","大易","赤霄","济生","夺魂","精铜","白阳","梨花","玄阴","狂涛","战龙","杀霸","雷光","暴风","霸鬼","绞肉","金翼","飞豹","丧门","净心","圣火","毒蛛","乾元","问世","凶冥","绝地","极乐","杀芒","白云","连环","无量","纯阳","琼浆","封咒","兽王","焰佛","毒龙","阴风","金胎","无定","问缘","飞鹰","缺月","渡业","烛龙","丽炎","血雨","黑虎","梦鬼","弥勒","炼气","大海","毒圣","星月","轮迴","血网","金狮","大千","飞花","血手","火云","毒蝎","摧心","紫焰","极阳","蝉翼","绝世","迅蛇","七曜","白骨","招邪","飞云","血龙","虎尾","水镜","巽风","伏魔","莲花","八神",}
    // 天材地宝
    treasure_prefix = []string{"霸封","皇术","魔斧","魔手","鬼指","捌破","圣手","万咒","妖破","神指","龙变","龙舞","龙咒","霸隐","魔舞","凤刀","神杀","妖爆","斧","鬼咒","妖封","圣解","圣变","法","拳","佛解","皇刀","龙雷","霸咒","皇斩","吟","神爪","龙劈","妖步","佛劈","妖拳","佛法","冥刀","妖解","轰","霸斩","皇枪","圣爪","鬼腿","皇瘴","皇手","龙指","魔雷","神法","妖雷","魔杀","枪","佛破","龙刀","佛变","术","圣掌","霸吟","双闪","佛咒","鬼爆","壹咒","皇法","佛掌","鬼斩","神腿","魔剑","神手","皇步","妖术","霸手","霸劈","斩","霸枪","龙剑","刀","鬼轰","魔术","龙拳","魔斩","皇劲","佛剑","皇斧","皇闪","劲","皇爆","神舞","猴瘴","魔瘴","龙爪","神斩","魔封","龙手","圣斧","圣拳","杀","破","皇杀","龙轰","虎腿","霸斧","神破","步","瘴","印","霸爆","神封","圣术","妖隐","霸印","霸破","冻法","腿","魔指","咒","佛闪","神轰","舞","亿杀","妖掌","皇雷","魔拳","爪","妖咒","霸变","爆","龙闪","皇封","亿轰","霸瘴","圣隐","手","魅剑","霸闪","鬼掌","圣雷","指","鬼封","佛指","皇劈","劈","妖闪","变","霸轰","妖印","鬼瘴","妖枪","神解","魔枪","鬼斧","佛舞","鬼闪","龙解","龙术","鬼劈","鬼术","神术","圣斩","梵指","妖吟","鬼枪","解","妖轰","剑","皇拳","佛步","圣步","佛轰","掌","皇隐","魔解","妖斧","神刀","神印","妖斩","妖舞","皇咒","隐","龙法","封","霸刀","佛瘴","妖刀","神咒","佛雷","雷","魔爆","圣印","鬼杀","圣封","鬼变","皇变","神劲","佛拳","佛术","闪","魔爪","皇剑","佛斧",}

    // -----------------------------

    // 天朝男性专用名
    chinese_man_suffix = []string{"佑","义","庐","寅","启","楷","君","辉","楠","澄","少","志","诚","远","材","烈","枫","鹏","博","祜","杰","谨","勋","信","暄","英","胜","锟","奋","恒","湖","基","杞","钊","涛","译","中","如","恩","善","睿","鑫","翘","海","祥","伦","刚","祯","俊","继","泽","望","轩","南","铭","日","培","福","骞","昀","骏","伟","才","勇","嘉","烁","锦","翰","运","本","承","栋","斌","丘","阳","奇","浩","濡","鹤","民","加","化","洋","德","凡","初","乘","帝","彬","光","厚","汝","瑞","佳","建","洁","沛","帆","延","维","圣","邦","天","辰","良","逸","玉","锋","源","槐","仕","钦","宇","梁","琛","峻","星","超","侠","然","贤","波","实","哲","人","璇","起","晖","业","礼","振","喆","涉","升","锐","允","皓","润","捷","钧","玲","载","益","晓","翱","腾","峰","落","驰","晨","济","凯","坤","辞","炳","原","谛","震","宾","子","年","斯","绍","鸿","尧","泰","裕","胤",}

    // 天朝女性专用名
    chinese_woman_suffix = []string{"竹","苑","鑫","媛","铃","依","琪","佳","小","娅","颖","亭","宜","曼","昕","曦","彬","莎","晞","尔","羽","香","婧","倩","春","瑶","昊","筠","澜","璐","宛","妍","丹","聪","斯","婵","蕊","克","炜","咏","瑗","栀","如","芳","言","洁","静","萱","敏","采","慧","之","晨","莲","美","桐","姝","漫","煦","怡","蓉","瑾","阑","琴","黛","楚","懿","祥","枫","珍","霞","玉","兰","寒","冰","凝","雪","思","青","珊","菁","琇","琰","飘","婉","桃","涵","蕾","丽","昭","梦","纯","仪","茵","若","柔","雯","雁","优","莹","雨","艺","希","莉","彩","凡","岚","函","珠","笙","舜","晗","菡","紫","月","晓","爽","琬","阳","书","芙","可","霭","窈","婕","鸽","立","滢","秀","棠","清","佩","纨","芸","淑","帛","橘","南","仁","熙","芝","亦","楠","梓","琳","钰","兆","诗","音","徽","婷","玲","玖","悦","眉","冬","毓","旭","晶","蓓","霏","馨","未","洵","璇","君","荷","娜","女","新","卉","灵","茜","馥","英","恩","璟","露","欣","彤","沛","初","育","园","锦","菲","心","妮","玥","琼","欢","雅","韵","琛","嘉","薇","碧","杉","子","磬","歆","畅","茗","风","令","卿","玟","亚","萌",}

    // 日本男性专用名
    japanese_man_suffix = []string{"兹","海","瞬","东","比","部","和","隼","也","结","心","晟","树","易","续","一","斗","人","太","茸","拓","航","宏","凑","向","新","平","大","翔","北","泰","白","奏","生","优","亮","天","治","野","之","花","智","直","煌","飒","汰","雄","渉","郎","三","実","石","光","十","悠","凉","介","枫","木","慕","遥","明","黑","星","陆","马","丈","雨","泽","空","庆","尔","虎","服","朔","文","龙","乐","少","圣","银","勇","成","五","邪","健","哲","阳","希","葵","佑","真","干","苍","咏","春",}

    // 日本女性专用名
    japanese_woman_suffix = []string{"弥","津","好","舞","纱","由","瑶","椿","萌","彩","希","晶","幸","纯","郁","朝","纪","百","惠","七","里","见","凛","泉","麻","秋","桃","奈","亚","杏","香","瞳","启","音","合","早","凌","和","华","夏","琴","静","露","熏","温","馨","绫","丽","光","映","织","衣","歌","久","佳","美","真","冬","小","代","芳","菜","诗","结","千","花","爱","日","理","穗","绘","樱","江","乃","玲","孝","蕾","梨","嘉","央","春","绪","文","珠","耶","子","芽","穂",}

    // 西方男性专用名
    west_man_suffix = []string{"艾伦","艾布特","埃布尔","艾布纳","亚伯拉罕","亚岱尔","亚当","艾狄生","阿道夫","亚度尼斯","亚德里恩","亚恒","艾伦","艾伯特","奥德里奇","亚历山大","亚尔弗列得","阿尔杰","阿尔杰农","艾伦","奥斯顿","阿尔瓦","阿尔文","亚尔维斯","亚摩斯","安得烈","安德鲁","安迪","安其罗","安格斯","安斯艾尔","安东尼","安托万","安东尼奥","阿奇尔","阿奇柏德","亚力士","亚尔林","亚尔曼","阿姆斯特朗","阿诺","阿诺德","阿瑟","艾文","亚撒","亚希伯恩","亚特伍德","奥布里","奥格斯格","奥古斯汀","艾富里","柏得温","班克罗夫特","巴德","巴洛","巴奈特","巴伦","巴里特","巴里","巴萨罗穆","巴特","巴顿","巴特莱","巴泽尔","比其尔","宝儿","贝克","班","班尼迪克","班杰明","班奈特","班森","博格","格吉尔","格纳","伯尼","伯特","伯顿","柏特莱姆","毕维斯","比尔","宾","毕夏普","布莱尔","布莱克","布莱兹","鲍伯","布兹","博格","鲍里斯","波文","柏宜斯","布德","布拉德利","布雷迪","布兰登","布莱恩","布拉得里克","布鲁克","布鲁斯","布鲁诺","巴克","伯骑士","巴尔克","布尼尔","波顿","拜伦","西泽","卡尔文","凯里","卡尔","凯尔","卡特","凯希","塞西尔","锡德里克","查德","钱宁","契布曼","查尔斯","夏佐","切斯特","克莱斯特","克里斯汀","克里斯托弗","克莱尔","克拉伦斯","克拉克","克劳德","克雷孟特","克利夫兰","柯利福","克利福德","克莱得","考伯特","考尔比","科林","康拉德","科里","康那理惟士","康奈尔","克莱凸","柯蒂斯","西里尔","戴纳","丹尼尔","达尔西","达内尔","达伦","迪夫","戴维","迪恩","邓普斯","丹尼斯","戴里克","得文","狄克","多米尼克","唐","唐纳修","唐纳德","道格拉斯","杜鲁","杜克","邓肯","唐恩","德维特","迪伦","额尔","艾德","伊登","埃德加","埃德蒙","爱迪生","爱德华","埃德温","爱格伯特","伊莱","易莱哲","埃利奥特","埃利斯","爱尔玛","埃尔罗伊","埃尔顿","埃尔维斯","爱曼纽","伊诺克","艾利克","欧内斯特","尤金","尔文","伊夫力","富宾恩","费利克斯","费迪南德","费奇","费兹捷勒","福特","弗朗西斯","法兰克","法兰克林","弗雷得力克","加布力尔","加尔","盖理","盖文","吉恩","杰佛理","杰夫","乔治","杰拉尔德","吉尔伯特","贾艾斯","葛兰","哥达","高德佛里","戈登","格雷格","格雷戈里","葛里菲兹","格罗佛","古斯塔夫","盖","霍尔","哈利","汉米敦","哈帝","哈伦","哈利","哈罗德","哈丽雅特","哈里","哈维","海登","海拾兹","享利","赫伯特","赫尔曼","希拉里","海勒","霍伯特","霍根","哈瑞斯","好尔德","休伯特","修","雨果","汉弗莱","汉特","海曼","伊恩","英格兰姆","艾勒","艾萨克","伊西多","艾凡","埃尔维斯","杰克","雅各布","贾里德","杰森","杰","杰夫","杰弗里","杰勒米","哲罗姆","杰理","杰西","吉姆","乔","约翰","琼纳斯","强纳生","约瑟夫","乔休尔","乔伊斯","朱利安","朱利尔斯","贾斯丁","基思","凯利","肯恩","肯尼迪","肯尼思","肯特","科尔","凯文","金姆","金","科克","凯尔","蓝伯特","蓝斯","劳瑞","劳伦斯","列夫","伦恩","伦农","利奥","伦纳德","利奥波德","勒斯","里斯特","利瓦伊","刘易斯","赖昂内尔","路","路易斯","陆斯恩","路德","莱尔","林顿","林恩","麦基","麦尔肯","曼德尔","马卡斯","马里奥","马克","马伦","玛希","马歇尔","马丁","马文","马特","马休","摩里斯","马克斯","马克西米兰","马克斯韦尔","马勒第兹","穆尔","麦克","米契尔","密克","麦克","麦尔斯","米路","门罗","曼特裘","穆尔","摩尔根","摩帝马","摩顿","摩西","摩菲","莫雷","麦伦","纳特","奈登","奈宝尼尔","尼尔","尼尔森","纽曼","尼克勒斯","尼克","奈哲尔","诺亚","诺尔","诺曼","诺顿","欧格登","奥利佛","奥玛","奥利尔","奥斯本","奥斯卡","奥斯蒙","奥斯维得","奥狄斯","奥特","欧恩","裴吉","帕克","培迪","帕特里克","保罗","派恩","斐瑞","皮特","彼得","菲尔","菲利普","波特","普雷斯科特","普利莫","昆特","昆尼尔","昆西","昆","昆顿","雷契尔","雷尔夫","伦道夫","雷蒙德","雷哲","里根","雷吉诺德","鲁宾","雷克斯","理查德","罗伯特","罗宾","洛克","罗德","罗得里克","罗德尼","罗恩","罗纳德","罗里","罗伊","鲁道夫","鲁伯特","莱安","山姆","辛普森","撒姆尔","山迪","撒克逊","史考特","肖恩","夕巴斯汀","锡德","锡得尼","席尔维斯特","赛门","所罗门","史宾杜","史丹","史丹佛","史丹尼","史帝文","史帝夫","史都华德","塔伯","泰勒","泰德","泰伦斯","希尔保特","西奥多","托玛士","帝福尼","堤姆","帝摩斯","托拜厄斯","托比","托德","汤姆","汤尼","特瑞西","特洛伊","杜鲁门","泰勒","泰伦","尤里西斯","阿普顿","尤莱亚",}

    // 西方女性专用名
    west_wowan_suffix = []string{"布兰妮","艾薇儿","希尔顿","梅根福克斯","莫妮卡","贝鲁奇","维多利亚","叶莲娜","科里科娃","辛迪","克劳馥","安吉丽娜","乔丹","凯特","卡梅隆","迪亚兹","妮可","基德曼","斯坦福妮","森莫","卡斯特","娜奥米","沃茨","珍妮佛","安妮斯顿","瓦妮莎","帕拉迪斯","亚米拉","莉苔希娅","考斯特","安吉丽娜","朱丽","娜塔丽","波特曼","凯普瑞丝","布瑞特","哈莉","贝瑞","凯莉","布鲁克","珍妮佛","洛佩兹","纳奥米","坎贝尔","凯瑟琳","泽塔琼斯","麦当娜","凯莉","米洛","莎朗","斯通","胡凯莉","米歇尔","菲佛","伊莉莎白","赫莉","丽贝卡","伊丽莎","库斯伯特","碧昂斯","克里斯蒂特","林顿","克劳蒂娅","希弗","凯瑞奥特斯","海蒂","克鲁穆","丹妮拉","玛丽娅","凯莉","米莲妮","茱丽娅","罗伯茨","库尔尼科娃","梅格","瑞恩","莎拉","布莱曼","蜜雪儿","凯特","温丝莱特","朱迪","福斯特","克里斯","蒂娜","安妮","海瑟薇","席琳","迪翁","西尔维","斯通","赛尔玛","海耶克","惠特尼","休斯顿","凯特","贝金塞尔","蒂朵","艾丽雅","辛吉斯","珍妮","杰克逊","安伯","瓦莉塔","艾蕾","莎摩儿","乔齐娜","格琳薇尔","帕梅拉","安德森","杰西卡","薇诺娜","赖德","夏洛蒂","菲丝希尔","苏菲玛索","奥黛丽","赫本","英格丽","褒曼","玛丽莲","梦露","费雯丽","波姬","小丝","秀兰","邓波儿","朱丽叶特","比诺什","戴安娜","凯瑟琳","赫本","伊丽莎白","泰勒","葛丽泰","嘉宝","丽芙","泰勒","珍妮弗","康奈利","格丽蕾丝凯莉","碧姬","芭铎","朱莉","安德鲁斯","简","方达","罗密","施奈德","索菲娅","罗兰","芭芭拉","史翠姗","娜塔莎","金斯基","梅丽尔","斯特里普","黛咪","摩儿",}

    // 组织
    org_suffix = []string{"帮","教","宗","联盟","宫","党","王朝","门","会","楼","阁","殿","团","家族","派","山庄","庵","坊","族","馆","谷","塔","寺","书院",}

    // 地点
    place_suffix = []string{"江","河","湖","海","山","峰","城","郡","州","洲","都","台","岛","国","林","谷","泽","园","村","乡","镇","县","港","堡","湾","洋","镇","观","府","庄","崖","阜","道","坊","山","坡","坝","陂","洞","区",}

    // 功法
    kongfu_suffix = []string{"霸封","皇术","魔斧","魔手","鬼指","捌破","圣手","万咒","妖破","神指","龙变","龙舞","龙咒","霸隐","魔舞","凤刀","神杀","妖爆","斧","鬼咒","妖封","圣解","圣变","法","拳","佛解","皇刀","龙雷","霸咒","皇斩","吟","神爪","龙劈","妖步","佛劈","妖拳","佛法","冥刀","妖解","轰","霸斩","皇枪","圣爪","鬼腿","皇瘴","皇手","龙指","魔雷","神法","妖雷","魔杀","枪","佛破","龙刀","佛变","术","圣掌","霸吟","双闪","佛咒","鬼爆","壹咒","皇法","佛掌","鬼斩","神腿","魔剑","神手","皇步","妖术","霸手","霸劈","斩","霸枪","龙剑","刀","鬼轰","魔术","龙拳","魔斩","皇劲","佛剑","皇斧","皇闪","劲","皇爆","神舞","猴瘴","魔瘴","龙爪","神斩","魔封","龙手","圣斧","圣拳","杀","破","皇杀","龙轰","虎腿","霸斧","神破","步","瘴","印","霸爆","神封","圣术","妖隐","霸印","霸破","冻法","腿","魔指","咒","佛闪","神轰","舞","亿杀","妖掌","皇雷","魔拳","爪","妖咒","霸变","爆","龙闪","皇封","亿轰","霸瘴","圣隐","手","魅剑","霸闪","鬼掌","圣雷","指","鬼封","佛指","皇劈","劈","妖闪","变","霸轰","妖印","鬼瘴","妖枪","神解","魔枪","鬼斧","佛舞","鬼闪","龙解","龙术","鬼劈","鬼术","神术","圣斩","梵指","妖吟","鬼枪","解","妖轰","剑","皇拳","佛步","圣步","佛轰","掌","皇隐","魔解","妖斧","神刀","神印","妖斩","妖舞","皇咒","隐","龙法","封","霸刀","佛瘴","妖刀","神咒","佛雷","雷","魔爆","圣印","鬼杀","圣封","鬼变","皇变","神劲","佛拳","佛术","闪","魔爪","皇剑","佛斧",}

    // 天材地宝
    treasure_suffix = []string{"毒骨","妖榴","鬼砂","毒玉","丝","仙岩","铁","神丹","妖血","爪","钢","天魂","圣棕","幻砂","冥兰","秘草","鬼棕","宝焰","毒甲","宝魄","佛霖","魂","神浪","蛊草","幻火","血","奇眼","玄金","灵枝","冥魄","真牙","玄皮","真丹","妖浪","琉","神叶","妖牙","真藤","仙波","鳞","毒珠","天丝","妖丝","冥浪","蛊丝","天草","鬼琉","玉","毒瘤","灵泉","真露","秘翼","砂","棕","奇鳞","蛊霞","宝鳞","幻土","佛绸","毒焰","蛊露","魔蕊","天肠","宝干","龙竹","灵瘴","仙钢","妖棕","芽","宝牙","佛丝","秘睛","毒肠","神精","蛊眼","宝枝","圣血","冥露","毛","冥血","魔血","神卵","天硫","佛睛","龙乳","幻丝","佛铁","仙瘴","宝水","妖璃","毒毛","油","灵油","龙玉","圣油","秘沙","天皮","幻金","仙炎","波","冥砂","仙焰","真睛","佛爪","秘水","宝瘤","龙爪","奇榴","幻芝","奇焱","土","圣实","龙牙","妖果","幻藤","魔瘤","冥雾","毒露","蛊莲","宝灵","妖莲","奇芽","神焰","晶","筋","仙丝","蛊芝","仙莲","毒眼","龙砂","魔参","真石","龙睛","灵精","仙火","竹","龙筋","牙","毒牙","魔冰","幻焰","神瘴","宝琉","仙叶","睛","玄土","天革","魔翼","佛泪","幻玉","灵","魔珠","冥花","毒皮","龙骨","天瘴","龙棕","仙羽","甲","宝缕","硫","仙焱","毒璃","秘缎","佛缕","天参","鬼血","宝精","圣翼","精","玄骨","天琉","圣牙","毒丝","妖油","灵灵","玄藤","圣心","实","丹","真土","奇铁","奇牙","骨","奇丝","真珠","灵焱","丹气","灵花","真骨","皮","冥睛","妖卵","毒冰","宝珠","秘干","玄牙","璃","天芝","玄筋","仙甲","干","仙灵","眼","肉","仙兰","泪","炎","冥心","灵布","秘精","灵干","魔玉","虎睛","岩","佛肠","妖兰","蛊玉","毒气","天枝","佛莲","珠","魔目","魄","目","革","仙珠","天砂","神硫","宝浪","灵眼","乳","蛊泉","龙实","佛魂","缎","仙根","霖","羽","卵","仙花","魔实","宝芽","神眼","仙魂","幻革","神钢","龙珠","火","秘霖","秘脑","神丝","根","蛊筋","金","宝露","宝花","宝波","天浪","心","鬼干","蛊岩","圣参","宝丹","毒脑","灵土","真翼","蕊","龙油","毒石","玄波","缕","花","枝","天石","雾","佛血","秘莲","芝","翅","蛊兰","毒油","宝玉","玄肠","妖魂","宝铁","幻兰","冥铁","佛炎","神泉","霞","魔铁","冰","龙花","毒霖","宝莲","兰","神炎","仙魄","真绸","奇玉","仙卵","秘琉","毒灵","佛冰","冥毛","奇泉","秘兰","玄甲","沙","天焱","佛牙","焰","魔芽","灵血","玄铁","宝土","蛊骨","莲","蛊璃","神土","妖花","妖根","真炎","真草","幻花","宝布","露","宝丝","蛊根","蛊缎","蛊爪","天精","内丹","幻草","鬼兰","真筋","灵火","秘瘴","叶","毒果","毒缕","仙砂","天露","泉","妖霖","翼","真璃","草","奇根","天丹","佛砂","灵参","肠","宝革","鬼叶","宝血","灵肉","宝根","神翅","石","冥干","神筋","玄血","瘤","焱","灵翼","瘴","天芽","魔花","灵脑","仙牙","鬼竹","玄晶","毒砂","冥筋","魔油","毒铁","榴","布","圣魂","宝石","鬼翼","幻牙","龙叶","水","毒绸","妖晶","天金","妖露","灵魂","秘竹","天脑","圣目","妖沙","宝肠","秘泉","佛浪","玄肉",}

    // 通用文字
    word = []string{"榆","邯","肇","莆","同","枣","邳","横","门","石","乌","嵩","无","都","随","仁","环","曲","泉","磐","承","三","洛","贡","辰","桓","虹","上","兰","登","郑","利","水","寨","闵","洋","衢","六","宁","堰","灯","民","集","红","洱","湖","北","象","尉","徽","赵","族","洪","溪","竹","源","岱","璧","港","甸","玉","嘉","沭","尼","诸","柘","井","翼","静","白","唐","裕","贵","阿","远","磁","坊","玛","塔","温","阴","崇","壁","溧","化","棱","沧","岳","西","椒","海","益","营","泗","盱","泾","卢","曹","坛","鸡","孟","金","恩","合","汇","庐","台","墨","盖","阳","获","修","铜","丰","广","岷","梁","莒","安","桦","廊","靖","家","陕","始","川","余","府","聊","岗","沛","高","东","虎","功","树","丹","义","宜","滨","伊","嫩","沁","彦","旗","寿","柳","万","张","碑","云","梦","彭","钢","梅","鄱","如","吉","塞","渭","施","依","姜","盘","吴","锦","来","汀","黑","彬","骅","监","克","界","绍","诏","济","莱","钟","宣","宿","许","内","宝","票","岭","鲁","乐","仓","邱","兴","友","英","漳","左","萨","氏","孝","原","柏","元","州","淳","首","濉","栗","敦","铁","肃","志","宕","射","良","衡","霍","子","埠","香","坪","凌","婺","宾","善","汪","保","陇","浦","庄","公","亭","华","林","祥","野","政","娄","荆","昭","春","岚","舒","谷","汉","陵","连","迁","田","征","邑","资","武","游","仪","镇","应","鹿","通","边","侯","栾","郧","洞","蔡","洮","青","范","巨","仙","四","扬","蓟","荔","泸","罗","顺","农","运","任","明","易","城","临","黟","池","汾","灵","信","封","昌","峰","常","泰","肥","桐","芜","浪","力","泊","商","微","中","荣","休","干","呈","十","漠","和","惠","文","清","冈","闽","浮","秭","德","抚","津","瑞","头","泽","库","陶","容","眙","滦","舟","自","徐","繁","定","汝","关","缙","格","光","赤","开","鱼","涟","丘","平","沽","沟","鄂","间","拜","庆","赣","淮","句","朔","晋","杭","凤","山","慈","双","灌","长","大","天","辉","延","胶","富","江","楼","厦","龙","敖","濮","康","邓","嵊","密","雄","莲","国","叶","姚","成","昆","峡","略","怀","留","巧","汤","河","奉","贤","行","辽","蒙","固","黄","莘","五","潭","口","福","蚌","沂","风","碌","浑","遂","南","乡","花","归","夏","九","松","新","沙","柔","交","盈","古","朝","永","建","太","银","渡","图",}
)

var Dict map[int][]string


func init(){
    Dict = map[int][]string{
        CHINESE_PREFIX: chinese_prefix,
        CHINESE_PREFIX_TWO: chinese_prefix_two,
        JAPANESE_PREFIX: japanese_prefix,
        WEST_PREFIX: west_prefix,
        KONGFU_PREFIX: kongfu_prefix,
        TREASURE_PREFIX: treasure_prefix,

        CHINESE_MAN_SUFFIX: chinese_man_suffix,
        CHINESE_WOMAN_SUFFIX: chinese_woman_suffix,
        JAPANESE_MAN_SUFFIX: japanese_man_suffix,
        JAPANESE_WOMAN_SUFFIX: japanese_woman_suffix,
        WEST_MAN_SUFFIX: west_man_suffix,
        WEST_WOWAN_SUFFIX: west_wowan_suffix,
        ORG_SUFFIX: org_suffix,
        PLACE_SUFFIX: place_suffix,
        KONGFU_SUFFIX: kongfu_suffix,
        TREASURE_SUFFIX: treasure_suffix,
        WORD: word,
    }
}
