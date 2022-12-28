package cmd

import (
	"errors"
	"fmt"
	"rac-tester/tool"
	"strconv"

	"github.com/manifoldco/promptui"
)

func KickTest(isSim bool) {

	var id int

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

	result, err := prompt.Run()

	cmd := tool.Commad{
		Id:         uint32(id),
		Kickspeedx: 1,
		IsSim:      isSim,
		Loop:       100,
	}

	//result2, err := prompt2.Run()
	if err != nil {
		fmt.Println(result)
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
