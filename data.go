package cha

type Seed struct {
	ChineseFirstName []string
	ChineseLastName []string
	FirstName []string
	MiddleName []string
	LastName []string
}
// ref: https://github.com/nuysoft/Mock/blob/refactoring/src/mock/random/name.js
var seed = Seed{
	FirstName: []string{
		// male
		"James", "John", "Robert", "Michael", "William",
		"David", "Richard", "Charles", "Joseph", "Thomas",
		"Christopher", "Daniel", "Paul", "Mark", "Donald",
		"George", "Kenneth", "Steven", "Edward", "Brian",
		"Ronald", "Anthony", "Kevin", "Jason", "Matthew",
		"Gary", "Timothy", "Jose", "Larry", "Jeffrey",
		"Frank", "Scott", "Eric",
		// female
		"Mary", "Patricia", "Linda", "Barbara", "Elizabeth",
		"Jennifer", "Maria", "Susan", "Margaret", "Dorothy",
		"Lisa", "Nancy", "Karen", "Betty", "Helen",
		"Sandra", "Donna", "Carol", "Ruth", "Sharon",
		"Michelle", "Laura", "Sarah", "Kimberly", "Deborah",
		"Jessica", "Shirley", "Cynthia", "Angela", "Melissa",
		"Brenda", "Amy", "Anna",
	},
	MiddleName : []string{
		// male
		"James", "John", "Robert", "Michael", "William",
		"David", "Richard", "Charles", "Joseph", "Thomas",
		"Christopher", "Daniel", "Paul", "Mark", "Donald",
		"George", "Kenneth", "Steven", "Edward", "Brian",
		"Ronald", "Anthony", "Kevin", "Jason", "Matthew",
		"Gary", "Timothy", "Jose", "Larry", "Jeffrey",
		"Frank", "Scott", "Eric",
		// female
		"Mary", "Patricia", "Linda", "Barbara", "Elizabeth",
		"Jennifer", "Maria", "Susan", "Margaret", "Dorothy",
		"Lisa", "Nancy", "Karen", "Betty", "Helen",
		"Sandra", "Donna", "Carol", "Ruth", "Sharon",
		"Michelle", "Laura", "Sarah", "Kimberly", "Deborah",
		"Jessica", "Shirley", "Cynthia", "Angela", "Melissa",
		"Brenda", "Amy", "Anna",
		"J", "R", "C", "H", "G",
	},
	LastName: []string{
		"Smith", "Johnson", "Williams", "Brown", "Jones",
		"Miller", "Davis", "Garcia", "Rodriguez", "Wilson",
		"Martinez", "Anderson", "Taylor", "Thomas", "Hernandez",
		"Moore", "Martin", "Jackson", "Thompson", "White",
		"Lopez", "Lee", "Gonzalez", "Harris", "Clark",
		"Lewis", "Robinson", "Walker", "Perez", "Hall",
		"Young", "Allen",
	},
	ChineseFirstName: []string{
		"王", "李", "张", "刘", "陈", "杨", "赵", "黄", "周", "吴",
		"徐", "孙", "胡", "朱", "高", "林", "何", "郭", "马", "罗",
		"梁", "宋", "郑", "谢", "韩", "唐", "冯", "于", "董", "萧",
		"程", "曹", "袁", "邓", "许", "傅", "沈", "曾", "彭", "吕",
		"苏", "卢", "蒋", "蔡", "贾", "丁", "魏", "薛", "叶", "阎",
		"余", "潘", "杜", "戴", "夏", "锺", "汪", "田", "任", "姜",
		"范", "方", "石", "姚", "谭", "廖", "邹", "熊", "金", "陆",
		"郝", "孔", "白", "崔", "康", "毛", "邱", "秦", "江", "史",
		"侯", "邵", "孟", "龙", "万", "段", "雷", "钱", "汤", "操",
		"尹", "黎", "易", "常", "武", "乔", "贺", "赖", "龚", "文",
	},
	ChineseLastName: []string{
		"思", "扉", "疏", "枫", "君", "月", "清", "歌", "陵", "乌", "怡", "致",
	},
}
