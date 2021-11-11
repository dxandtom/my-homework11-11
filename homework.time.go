package main

import "fmt"
//lv0, 输入skill1，2，3 分别输出不同的技能

func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}
func inputskill(a string)string{
	return a
}
func main() {

	//var output []string
	//var ans string
	//var inputskill string
	//
	//output = make([]string, 0)
	//for i := 0; i <= 5; i++ {
	//
	//	for true {
	//		fmt.Scanln(&inputskill)
	//		switch inputskill {
	//		case "skill1":
	//			ans = "龙卷风摧毁停车场！"
	//		case "skill2":
	//			ans = "乌鸦坐飞机！！"
	//		case "skill3":
	//			ans = "黑虎掏心！！！"
	//		}
	//		output = append(output, ans)
	//		if i <= 1 {
	//			fmt.Print(ans, " ")
	//			fmt.Println()
	//		} else {
	//			fmt.Println("游玩次数到达上限")
	//		}
	//	}
	//}



//lv1



	ReleaseSkill("龙卷风摧毁停车场！", func(skillName string) {
		fmt.Println("尝尝我的厉害吧！", skillName)
	})
var ans string
var inputskill string

	for true {
		fmt.Scanln(&inputskill)
		switch inputskill {
		case "skill1":
			ans = "龙卷风再次摧毁停车场！！"
		case "skill2":
			ans = "乌鸦也坐飞机！！"
		case "skill3":
			ans = "黑虎也掏心！！"
		}
fmt.Print(ans, "")
		fmt.Println( )
	}


}

