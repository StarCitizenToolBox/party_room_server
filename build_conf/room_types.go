package build_conf

import "party_room_server/protos"

var RoomTypesData = []*protos.RoomType{
	{
		Id:   "bounty",
		Name: "赏金",
		Desc: "赏金猎人、战斗信标、地堡 等战斗为主的赏金任务",
		SubTypes: []*protos.RoomSubtype{
			{Id: "bounty_air_war", Name: "空战赏金"},
			{Id: "bounty_beacon", Name: "信标赏金"},
			{Id: "bounty_fps", Name: "FPS赏金"},
			{Id: "bounty_player", Name: "玩家赏金"},
		},
	},
	{
		Id:   "mining",
		Name: "采矿",
		Desc: "FPS采矿、矿车采矿、矿船采矿等",
	},
	{
		Id:   "salvage",
		Name: "打捞",
		Desc: "以 秃鹫 回收者 为主的打捞回收作业任务",
		SubTypes: []*protos.RoomSubtype{
			{Id: "salvage_traditional", Name: "传统打捞"},
			{Id: "salvage_combat", Name: "先打后捞"},
		},
	},
	{
		Id:   "express",
		Name: "快递",
		Desc: "游戏内各物流公司的快递任务",
	},
	{
		Id:   "medical_rescue",
		Name: "医疗/救援",
		Desc: "以游戏内玩家援助为主的救援任务（医疗，托运，加油 等）",
		SubTypes: []*protos.RoomSubtype{
			{Id: "medical", Name: "医疗"},
			{Id: "consignment", Name: "牵引托运"},
			{Id: "refuel", Name: "加油"},
		},
	},
	{
		Id:   "missions",
		Name: "大型任务",
		Desc: "围攻奥里森、异种威胁、Jumptown、九尾封锁、阿灵顿帮伊德里斯任务等",
		SubTypes: []*protos.RoomSubtype{
			{Id: "missions_orison", Name: "围攻奥里森"},
			{Id: "missions_XenoThreat", Name: "异种威胁"},
			{Id: "missions_nine_tails", Name: "九尾封锁"},
			{Id: "missions_collection", Name: "其他大型任务"},
		},
	},
	{
		Id:   "journey",
		Name: "旅游",
		Desc: "参访著名景点、兴趣点",
	},
	{
		Id:   "ac",
		Name: "竞技场指挥官",
		Desc: "以竞技场指挥官为主的玩法",
		SubTypes: []*protos.RoomSubtype{
			{Id: "ac_dog_fight", Name: "自由飞"},
			{Id: "ac_pirate", Name: "海盗潮"},
			{Id: "ac_van", Name: "弯度潮"},
		},
	},
	{
		Id:   "global",
		Name: "随便玩",
		Desc: "随便玩点啥，无固定分类",
	},
	{
		Id:   "newbie",
		Name: "新手频道",
		Desc: "萌新！快到碗里来！",
	},
}
