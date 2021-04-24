package main

import (
	"github.com/sampgo/sampgo"
)

var (
	gPlayerHasCitySelected       map[sampgo.Player]bool = make(map[sampgo.Player]bool)
	gPlayerCitySelection         map[sampgo.Player]int  = make(map[sampgo.Player]int)
	gPlayerLastCitySelectionTick map[sampgo.Player]int  = make(map[sampgo.Player]int)

	txtLosSantos   int
	txtSanFierro   int
	txtLasVenturas int

	CITY_LOS_SANTOS   int = 0
	CITY_SAN_FIERRO   int = 1
	CITY_LAS_VENTURAS int = 2
)

func ClassSel_InitCityNameText(txt int) {
	sampgo.TextDrawUseBox(txt, false)
	sampgo.TextDrawLetterSize(txt, 1.25, 3.0)
	sampgo.TextDrawFont(txt, 0)
	sampgo.TextDrawSetShadow(txt, 0)
	sampgo.TextDrawSetOutline(txt, 1)
	sampgo.TextDrawColor(txt, 0xEEEEEE)
	sampgo.TextDrawBackgroundColor(txt, 0x000000FF)
}

func ClassSel_InitTextDraws() {
	txtLosSantos = sampgo.TextDrawCreate(10.0, 380.0, "Los Santos")
	ClassSel_InitCityNameText(txtLosSantos)

	txtSanFierro = sampgo.TextDrawCreate(10.0, 380.0, "San Fierro")
	ClassSel_InitCityNameText(txtSanFierro)

	txtLasVenturas = sampgo.TextDrawCreate(10.0, 380.0, "Las Venturas")
	ClassSel_InitCityNameText(txtLasVenturas)
}

func ClassSel_SetupSelectedCity(p sampgo.Player) {
	if gPlayerCitySelection[p] == -1 {
		gPlayerCitySelection[p] = CITY_LOS_SANTOS
	}

	if gPlayerCitySelection[p] == CITY_LOS_SANTOS {
		sampgo.SetPlayerInterior(p.ID, 0)
		sampgo.SetPlayerCameraPos(p.ID, 1630.6136, -2286.0298, 110.0)
		sampgo.SetPlayerCameraLookAt(p.ID, 1887.6034, -1682.1442, 47.6167, 2)

		sampgo.TextDrawShowForPlayer(p.ID, txtLosSantos)
		sampgo.TextDrawHideForPlayer(p.ID, txtSanFierro)
		sampgo.TextDrawHideForPlayer(p.ID, txtLasVenturas)
	} else if gPlayerCitySelection[p] == CITY_SAN_FIERRO {
		sampgo.SetPlayerInterior(p.ID, 0)
		sampgo.SetPlayerCameraPos(p.ID, -1300.8754, 68.0546, 129.4823)
		sampgo.SetPlayerCameraLookAt(p.ID, -1817.9412, 769.3878, 132.6589, 2)

		sampgo.TextDrawHideForPlayer(p.ID, txtLosSantos)
		sampgo.TextDrawShowForPlayer(p.ID, txtSanFierro)
		sampgo.TextDrawHideForPlayer(p.ID, txtLasVenturas)
	} else if gPlayerCitySelection[p] == CITY_LAS_VENTURAS {
		sampgo.SetPlayerInterior(p.ID, 0)
		sampgo.SetPlayerCameraPos(p.ID, 1310.6155, 1675.9182, 110.7390)
		sampgo.SetPlayerCameraLookAt(p.ID, 2285.2944, 1919.3756, 68.2275, 2)

		sampgo.TextDrawHideForPlayer(p.ID, txtLosSantos)
		sampgo.TextDrawHideForPlayer(p.ID, txtSanFierro)
		sampgo.TextDrawShowForPlayer(p.ID, txtLasVenturas)
	}
}

func init() {
	sampgo.Print("go init() called")

	sampgo.On("goModeInit", func() bool {
		sampgo.Print("Hello from Go!")

		// Functions to be called
		ClassSel_InitTextDraws()

		// Player Class
		sampgo.AddPlayerClass(298, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(299, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(300, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(301, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(302, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(303, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(304, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(305, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(280, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(281, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(282, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(283, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(284, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(285, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(286, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(287, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(288, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(289, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(265, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(266, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(267, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(268, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(269, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(270, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(1, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(2, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(3, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(4, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(5, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(6, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(8, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(42, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(65, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)

		sampgo.AddPlayerClass(86, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(119, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(149, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(208, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(273, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(289, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)

		sampgo.AddPlayerClass(47, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(48, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(49, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(50, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(51, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(52, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(53, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(54, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(55, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(56, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(57, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(58, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(68, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(69, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(70, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(71, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(72, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(73, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(75, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(76, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(78, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(79, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(80, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(81, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(82, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(83, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(84, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(85, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(87, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(88, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(89, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(91, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(92, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(93, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(95, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(96, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(97, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(98, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		sampgo.AddPlayerClass(99, 1759.0189, -1898.1260, 13.5622, 266.4503, -1, -1, -1, -1, -1, -1)
		return true
	})

	sampgo.On("goModeExit", func() bool {
		sampgo.Print("goModeExit!")
		return true
	})

	sampgo.On("playerRequestClass", func(playerid, classid int) bool {

		return true
	})

	sampgo.On("playerConnect", func(p sampgo.Player) bool {
		sampgo.GameTextForPlayer(p.ID, "~w~Grand Larceny", 3000, 4)

		p.SendMessage(-1, "Welcome to {88AA88}G{FFFFFF}rand {88AA88}L{FFFFFF}arceny")

		gPlayerCitySelection[p] = -1
		gPlayerHasCitySelected[p] = false
		gPlayerLastCitySelectionTick[p] = sampgo.GetTickCount()

		return true
	})
}

func main() {}
