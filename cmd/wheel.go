package cmd

import (
	"errors"
	"fmt"
	"rac-tester/tool"
	"strconv"

	"github.com/manifoldco/promptui"
)

func WheelTest(isSim bool) {

	var id int
	var time int

	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("RobotID must be 0-11")
		}
		id, _ = strconv.Atoi(input)
		if id < 12 && id >= 0 {
		} else {
			return errors.New("RobotID must be 0-11")
		}
		return nil
	}

	validate2 := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Time cannnot be string")
		}
		time, _ = strconv.Atoi(input)
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . | bold }} ",
		Valid:   "{{ . | bold | green }} ",
		Invalid: "{{ . | bold |red }} ",
		Success: "{{ . | bold | green }} ",
	}

	prompt := promptui.Prompt{
		Label:     "RobotID",
		Templates: templates,
		Validate:  validate,
	}

	prompt2 := promptui.Prompt{
		Label:     "Time(S)",
		Templates: templates,
		Validate:  validate2,
	}

	_, err := prompt.Run()

	_, err = prompt2.Run()

	//3msごとにコマンドを送っているので3で割っておくことで正しい秒数に変更している
	//そのうち改善するかも

	cmd := tool.Commad{
		Id:         uint32(id),
		Veltangent: 0.3,
		IsSim:      isSim,
		Loop:       time * 333,
	}

	//result2, err := prompt2.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	tool.SendCmd(cmd)

	for {
		prompt := promptui.Select{
			Label:     "Again?",
			Items:     []string{"Yes", "No"},
			CursorPos: 0,
		}
		idx, _, err := prompt.Run() //入力を受け取る

		if idx == 0 {
			tool.SendCmd(cmd)
		} else {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	return //ここで終了
}
